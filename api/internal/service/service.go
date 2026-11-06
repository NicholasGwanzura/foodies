// Package service holds the business logic. HTTP handlers and the WhatsApp
// channel both call it, so rules, approvals and audit behave the same.
package service

import (
	"context"
	"fmt"

	"foodies/api/internal/db"
)

// Service is the entry point to the business logic.
type Service struct {
	queries db.Querier
}

// New returns a Service backed by queries.
func New(queries db.Querier) *Service {
	return &Service{queries: queries}
}

// Ping confirms the database answers.
func (s *Service) Ping(ctx context.Context) error {
	if _, err := s.queries.Ping(ctx); err != nil {
		return fmt.Errorf("ping database: %w", err)
	}
	return nil
}
