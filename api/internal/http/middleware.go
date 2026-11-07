package httpapi

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"log/slog"
	"net"
	"net/http"
	"runtime/debug"
	"strings"

	"foodies/api/internal/auth"
	"foodies/api/internal/http/apigen"
	"foodies/api/internal/service"
)

type traceIDKey struct{}

// requestID gives every request a trace id, returned as X-Request-Id.
func requestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := rand.Text()
		w.Header().Set("X-Request-Id", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), traceIDKey{}, id)))
	})
}

func traceID(ctx context.Context) string {
	id, _ := ctx.Value(traceIDKey{}).(string)
	return id
}

// recoverPanics turns a panic into a Problem and logs the stack.
func recoverPanics(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if v := recover(); v != nil {
					logger.ErrorContext(r.Context(), "panic",
						"panic", v, "path", r.URL.Path, "trace_id", traceID(r.Context()),
						"stack", string(debug.Stack()))
					writeInternalError(w, r)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func bearer(r *http.Request) string {
	token, ok := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
	if !ok {
		return ""
	}
	return strings.TrimSpace(token)
}

func clientIP(r *http.Request) string {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

type ipKey struct{}

// requestIP returns the address the request came from.
func requestIP(ctx context.Context) string {
	ip, _ := ctx.Value(ipKey{}).(string)
	return ip
}

// authorize is the single gate every generated operation passes: it resolves
// the principal, checks the permission and narrows the scope to it.
func (s *Server) authorize(f apigen.StrictHandlerFunc, operationID string) apigen.StrictHandlerFunc {
	op := apigen.Operations[operationID]
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request any) (any, error) {
		ctx = context.WithValue(ctx, ipKey{}, clientIP(r))
		if op.Public {
			return f(ctx, w, r, request)
		}
		p, err := s.service.Authenticate(ctx, bearer(r), clientIP(r))
		if err != nil {
			return nil, err
		}
		if p.Limit != auth.LimitNone && !limitedAllows(p.Limit, operationID) {
			// The contract has one code for an unfinished sign-in. The web app
			// tells the two apart from credential_status on GET /me.
			detail := "Set up two-factor sign-in to continue."
			if p.Limit == auth.LimitPasswordChange {
				detail = "Choose a new password to continue."
			}
			return nil, service.Fail(http.StatusUnauthorized, "mfa_required", "Finish signing in", detail)
		}
		if op.Permission != "" {
			scope, ok := p.ScopeFor(op.Permission)
			if !ok {
				return nil, service.Fail(http.StatusForbidden, "forbidden", "Not allowed", "Your role does not allow this.")
			}
			p.Active = scope
		}
		ctx = auth.With(ctx, p)
		if op.Restricted && r.Method == http.MethodGet {
			if err := s.service.LogRead(ctx, operationID, r.URL.Path); err != nil {
				return nil, err
			}
		}
		return f(ctx, w, r, request)
	}
}

// limitedAllows lists what a token that has not finished sign-in may call.
func limitedAllows(limit auth.Limit, operationID string) bool {
	switch operationID {
	case "GetMe", "Logout", "RefreshToken":
		return true
	case "ChangePassword", "SetPin":
		return limit == auth.LimitPasswordChange
	case "EnrolMfa", "ConfirmMfa":
		return limit == auth.LimitMfaEnrolment
	}
	return false
}

// unitOfWork runs each write in one transaction, committing only on success.
// Sign-in operations manage their own writes, because a failed attempt has to
// persist its counter even though the request fails.
func (s *Server) unitOfWork(f apigen.StrictHandlerFunc, operationID string) apigen.StrictHandlerFunc {
	if apigen.Operations[operationID].Module == "Auth" {
		return f
	}
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request any) (any, error) {
		if r.Method == http.MethodGet {
			return f(ctx, w, r, request)
		}
		var response any
		err := s.service.InTx(ctx, func(ctx context.Context) error {
			var err error
			response, err = f(ctx, w, r, request)
			return err
		})
		return response, err
	}
}

// capture keeps a copy of a response for the idempotency store.
type capture struct {
	http.ResponseWriter
	status int
	body   bytes.Buffer
}

func (c *capture) WriteHeader(status int) {
	c.status = status
	c.ResponseWriter.WriteHeader(status)
}

func (c *capture) Write(b []byte) (int, error) {
	if c.status == 0 {
		c.status = http.StatusOK
	}
	c.body.Write(b)
	return c.ResponseWriter.Write(b)
}

// idempotent replays the first response to a repeated Idempotency-Key. Only
// successes are kept, so a client can fix a rejected request and resend it. The
// key belongs to the caller, so it runs outside the strict handlers, where the
// raw body is still available.
func (s *Server) idempotent(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("Idempotency-Key")
		token := bearer(r)
		if r.Method != http.MethodPost || key == "" || token == "" {
			next.ServeHTTP(w, r)
			return
		}
		actor, err := s.service.TokenUser(token)
		if err != nil {
			next.ServeHTTP(w, r) // authorize answers with the right error.
			return
		}
		body, err := io.ReadAll(r.Body)
		if err != nil {
			requestError(w, r, err)
			return
		}
		r.Body = io.NopCloser(bytes.NewReader(body))
		sum := sha256.New()
		sum.Write([]byte(r.Method + " " + r.URL.Path + "\n"))
		sum.Write(body)
		hash := hex.EncodeToString(sum.Sum(nil))

		replay, err := s.service.ClaimIdempotency(r.Context(), actor, key, hash)
		if err != nil {
			var fail *service.Error
			if errors.As(err, &fail) {
				responseError(s.logger)(w, r, err)
				return
			}
			s.logger.ErrorContext(r.Context(), "idempotency claim failed", "error", err)
			writeInternalError(w, r)
			return
		}
		if replay != nil {
			for k, v := range replay.Headers {
				w.Header().Set(k, v)
			}
			w.Header().Set("Idempotent-Replayed", "true")
			w.WriteHeader(replay.Status)
			_, _ = w.Write(replay.Body) //nolint:gosec // our own JSON response, stored earlier
			return
		}

		rec := &capture{ResponseWriter: w}
		ctx := context.WithoutCancel(r.Context())
		finished := false
		defer func() {
			if finished {
				return
			}
			// A panic, or a handler that wrote nothing: free the key.
			if err := s.service.ReleaseIdempotency(ctx, actor, key); err != nil {
				s.logger.ErrorContext(ctx, "release idempotency key", "error", err)
			}
		}()
		next.ServeHTTP(rec, r)
		if rec.status >= 200 && rec.status < 300 {
			headers := map[string]string{}
			for _, name := range []string{"Content-Type", "ETag"} {
				if v := w.Header().Get(name); v != "" {
					headers[name] = v
				}
			}
			if err := s.service.FinishIdempotency(ctx, actor, key, rec.status, headers, rec.body.Bytes()); err != nil {
				s.logger.ErrorContext(ctx, "store idempotent response", "error", err)
				return
			}
			finished = true
		}
	})
}
