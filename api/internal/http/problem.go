package httpapi

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"foodies/api/internal/http/apigen"
)

const problemContentType = "application/problem+json"

// problem is what a response needs beyond the fields every Problem shares.
type problem struct {
	status int
	code   apigen.ErrorCode
	// typeName is the last segment of the Problem type URI.
	typeName string
	title    string
	detail   string
}

func writeProblem(w http.ResponseWriter, r *http.Request, p problem) {
	instance := r.URL.Path
	body := apigen.Problem{
		Type:     "/problems/" + p.typeName,
		Title:    p.title,
		Status:   p.status,
		Code:     p.code,
		Instance: &instance,
		TraceId:  traceID(r.Context()),
	}
	if p.detail != "" {
		body.Detail = &p.detail
	}

	w.Header().Set("Content-Type", problemContentType)
	w.WriteHeader(p.status)
	// The status is already sent, so a failed write has nowhere left to go.
	_ = json.NewEncoder(w).Encode(body)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	writeProblem(w, r, problem{
		status:   http.StatusNotFound,
		code:     apigen.ErrorCodeNotFound,
		typeName: "not-found",
		title:    "Not found",
	})
}

// requestError handles requests the generated code could not parse.
func requestError(w http.ResponseWriter, r *http.Request, err error) {
	writeProblem(w, r, problem{
		status:   http.StatusBadRequest,
		code:     apigen.ErrorCodeValidationFailed,
		typeName: "validation-failed",
		title:    "Malformed request",
		detail:   err.Error(),
	})
}

// responseError maps errors from handlers. The contract has no error code for
// an operation that is not built, so 501 carries internal_error.
func responseError(logger *slog.Logger) func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		if errors.Is(err, apigen.ErrNotImplemented) {
			writeProblem(w, r, problem{
				status:   http.StatusNotImplemented,
				code:     apigen.ErrorCodeInternalError,
				typeName: "not-implemented",
				title:    "Not implemented",
			})
			return
		}

		logger.ErrorContext(r.Context(), "request failed",
			"error", err, "path", r.URL.Path, "trace_id", traceID(r.Context()))
		writeInternalError(w, r)
	}
}

func writeInternalError(w http.ResponseWriter, r *http.Request) {
	writeProblem(w, r, problem{
		status:   http.StatusInternalServerError,
		code:     apigen.ErrorCodeInternalError,
		typeName: "internal-error",
		title:    "Internal error",
	})
}
