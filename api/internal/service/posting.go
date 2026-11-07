package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"

	"foodies/api/internal/db"
)

// createPosting records that a decision has to reach a system of record.
//
//nolint:unused // First called by the money trail slice.
func (s *Service) createPosting(ctx context.Context, system, kind, table string, id uuid.UUID, site *uuid.UUID, payload any) error {
	raw, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("encode posting: %w", err)
	}
	return wrap(s.q(ctx).CreatePostingRow(ctx, db.CreatePostingRowParams{
		TargetSystem: &system, PostingType: &kind, SubjectType: &table, SubjectID: &id, SiteID: site, Payload: raw, CreatedBy: actorID(ctx),
	}), "create posting")
}
