package httpapi

import (
	"context"
	"crypto/rand"
	"log/slog"
	"net/http"
	"runtime/debug"
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
