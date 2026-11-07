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
	logger  *slog.Logger
	// routes is the API behind the sync replay, which re-enters it per item.
	routes http.Handler
}

// NewHandler returns the API with its middleware applied.
func NewHandler(svc *service.Service, logger *slog.Logger) http.Handler {
	server := &Server{service: svc, logger: logger}
	// Strict middleware applies in list order, so the last entry runs first.
	middlewares := []apigen.StrictMiddlewareFunc{server.unitOfWork, server.authorize}
	strict := apigen.NewStrictHandlerWithOptions(server, middlewares, apigen.StrictHTTPServerOptions{
		RequestErrorHandlerFunc:  requestError,
		ResponseErrorHandlerFunc: responseError(logger),
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/", notFound)
	mux.HandleFunc("PUT /files/{id}", server.serveFile)
	mux.HandleFunc("GET /files/{id}", server.serveFile)

	routes := apigen.HandlerWithOptions(strict, apigen.StdHTTPServerOptions{
		BaseURL:          basePath,
		BaseRouter:       mux,
		ErrorHandlerFunc: requestError,
	})

	server.routes = requestID(recoverPanics(logger)(server.idempotent(webhookBody(routes))))
	return server.routes
}
