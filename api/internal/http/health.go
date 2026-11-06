package httpapi

import (
	"context"

	"foodies/api/internal/http/apigen"
)

// GetHealth reports whether the API can reach its database.
func (s *Server) GetHealth(ctx context.Context, _ apigen.GetHealthRequestObject) (apigen.GetHealthResponseObject, error) {
	if err := s.service.Ping(ctx); err != nil {
		return nil, err
	}
	return apigen.GetHealth204Response{}, nil
}
