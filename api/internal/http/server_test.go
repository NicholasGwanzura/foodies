package httpapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"foodies/api/internal/http/apigen"
	"foodies/api/internal/service"
)

type stubQuerier struct{ err error }

func (s stubQuerier) Ping(context.Context) (int32, error) { return 1, s.err }

func TestHealth(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		pingErr    error
		wantStatus int
	}{
		{name: "database reachable", wantStatus: http.StatusNoContent},
		{name: "database down", pingErr: errors.New("connection refused"), wantStatus: http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			svc := service.New(stubQuerier{err: tt.pingErr})
			handler := NewHandler(svc, slog.New(slog.DiscardHandler))
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/api/v1/health", nil))

			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if tt.pingErr == nil {
				if rec.Body.Len() != 0 {
					t.Errorf("body = %q, want empty", rec.Body)
				}
				return
			}
			if got := decodeProblem(t, rec); got.Code != apigen.ErrorCodeInternalError {
				t.Errorf("code = %q, want internal_error", got.Code)
			}
			if bytes.Contains(rec.Body.Bytes(), []byte("refused")) {
				t.Errorf("internal error reached the client: %s", rec.Body)
			}
		})
	}
}

func TestHandlerErrorsAreProblems(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		method     string
		path       string
		query      string
		wantStatus int
		wantCode   apigen.ErrorCode
		wantType   string
	}{
		{
			name:       "operation not built yet",
			method:     http.MethodGet,
			path:       "/api/v1/me",
			wantStatus: http.StatusNotImplemented,
			wantCode:   apigen.ErrorCodeInternalError,
			wantType:   "/problems/not-implemented",
		},
		{
			name:       "unknown path",
			method:     http.MethodGet,
			path:       "/api/v1/nope",
			wantStatus: http.StatusNotFound,
			wantCode:   apigen.ErrorCodeNotFound,
			wantType:   "/problems/not-found",
		},
		{
			name:       "path outside the base path",
			method:     http.MethodGet,
			path:       "/me",
			wantStatus: http.StatusNotFound,
			wantCode:   apigen.ErrorCodeNotFound,
			wantType:   "/problems/not-found",
		},
		{
			name:       "malformed query parameter",
			method:     http.MethodGet,
			path:       "/api/v1/tasks",
			query:      "?limit=many",
			wantStatus: http.StatusBadRequest,
			wantCode:   apigen.ErrorCodeValidationFailed,
			wantType:   "/problems/validation-failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			handler := NewHandler(nil, slog.New(slog.DiscardHandler))
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, httptest.NewRequest(tt.method, tt.path+tt.query, nil))

			got := decodeProblem(t, rec)
			if got.Status != tt.wantStatus || rec.Code != tt.wantStatus {
				t.Errorf("status: body %d, header %d, want %d", got.Status, rec.Code, tt.wantStatus)
			}
			if got.Code != tt.wantCode {
				t.Errorf("code = %q, want %q", got.Code, tt.wantCode)
			}
			if got.Type != tt.wantType {
				t.Errorf("type = %q, want %q", got.Type, tt.wantType)
			}
			if got.Instance == nil || *got.Instance != tt.path {
				t.Errorf("instance = %v, want %q", got.Instance, tt.path)
			}
		})
	}
}

func TestProblemDoesNotLeakInternalErrors(t *testing.T) {
	t.Parallel()

	var logs bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&logs, nil))
	req := httptest.NewRequest(http.MethodGet, "/api/v1/me", nil)
	rec := httptest.NewRecorder()

	requestID(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseError(logger)(w, r, errors.New("pq: password leaked"))
	})).ServeHTTP(rec, req)

	if bytes.Contains(rec.Body.Bytes(), []byte("leaked")) {
		t.Errorf("internal error reached the client: %s", rec.Body)
	}
	if !bytes.Contains(logs.Bytes(), []byte("leaked")) {
		t.Errorf("internal error was not logged: %s", logs.String())
	}
	if got := decodeProblem(t, rec); got.Status != http.StatusInternalServerError {
		t.Errorf("status = %d, want 500", got.Status)
	}
}

func TestRecoverReturnsProblem(t *testing.T) {
	t.Parallel()

	handler := requestID(recoverPanics(slog.New(slog.DiscardHandler))(
		http.HandlerFunc(func(http.ResponseWriter, *http.Request) { panic("boom") }),
	))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/api/v1/me", nil))

	if got := decodeProblem(t, rec); got.Code != apigen.ErrorCodeInternalError {
		t.Errorf("code = %q, want internal_error", got.Code)
	}
}

// decodeProblem checks what every error response shares under spec section 12.4.
func decodeProblem(t *testing.T, rec *httptest.ResponseRecorder) apigen.Problem {
	t.Helper()

	if got := rec.Header().Get("Content-Type"); got != "application/problem+json" {
		t.Errorf("Content-Type = %q", got)
	}
	var p apigen.Problem
	if err := json.Unmarshal(rec.Body.Bytes(), &p); err != nil {
		t.Fatalf("decode problem: %v", err)
	}
	if p.TraceId == "" {
		t.Error("trace_id is empty")
	}
	if got := rec.Header().Get("X-Request-Id"); got != p.TraceId {
		t.Errorf("X-Request-Id = %q, want trace_id %q", got, p.TraceId)
	}
	return p
}
