package service

import (
	"context"

	"github.com/google/uuid"

	"foodies/api/internal/db"
)

// subject is what shared services need from any table.
type subject struct {
	get       func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error)
	statuses  []string
	setStatus func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error
}
