// Package httpapi serves the API defined by openapi.yaml. Handlers stay thin:
// the generated code parses, the service decides, and errors map to Problems.
package httpapi

import (
	"log/slog"
	"net/http"

	"foodies/api/internal/http/apigen"
	"foodies/api/internal/service"
)

const basePath = "/api/v1"

// Server implements the generated interface. Operations that no slice has
// built yet fall through to the embedded Unimplemented.
type Server struct {
	apigen.Unimplemented

	service *service.Service
}

// NewHandler returns the API with its middleware applied.
func NewHandler(svc *service.Service, logger *slog.Logger) http.Handler {
	strict := apigen.NewStrictHandlerWithOptions(&Server{service: svc}, nil, apigen.StrictHTTPServerOptions{
		RequestErrorHandlerFunc:  requestError,
		ResponseErrorHandlerFunc: responseError(logger),
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/", notFound)

	routes := apigen.HandlerWithOptions(strict, apigen.StdHTTPServerOptions{
		BaseURL:          basePath,
		BaseRouter:       mux,
		ErrorHandlerFunc: requestError,
	})

	return requestID(recoverPanics(logger)(routes))
}
