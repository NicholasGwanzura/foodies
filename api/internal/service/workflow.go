package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"foodies/api/internal/auth"
	"foodies/api/internal/db"
)

// stateError explains an action that matched no row: the record exists, since
// the caller just read it, so its status must not allow the action.
func stateError(err error, what string) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return errState("You cannot " + what + " in its current state.")
	}
	return wrap(err, what)
}

// checkSod enforces segregation of duties from the audit trail. It compares
// what the actor did to this record, so delegation cannot get around it.
func (s *Service) checkSod(ctx context.Context, action, subjectType string, subject uuid.UUID) error {
	actor := actorID(ctx)
	if actor == nil {
		return nil
	}
	rows, err := s.q(ctx).SodConflicts(ctx, db.SodConflictsParams{Action: action, SubjectType: &subjectType, SubjectID: &subject, Actor: actor})
	if err != nil {
		return fmt.Errorf("check segregation of duties: %w", err)
	}
	var blockers []Blocker
	for _, r := range rows {
		if r.Enforcement == "block" {
			blockers = append(blockers, Blocker{Kind: "sod_conflict", Reference: r.Key, Message: r.Name})
		}
	}
	if len(blockers) > 0 {
		return &Error{Status: http.StatusForbidden, Code: "sod_violation", Title: "Someone else must do this",
			Detail: "You already acted on this record in a role that must stay separate.", Blockers: blockers}
	}
	return nil
}

// RequirePin confirms the signed-in user with their PIN.
func (s *Service) RequirePin(ctx context.Context, pin string) error {
	return s.CheckPin(ctx, auth.From(ctx).UserID, pin)
}

// RequireSigned confirms a signature the caller made on this record. Any action
// that takes a signature_id comes through here.
func (s *Service) RequireSigned(ctx context.Context, table string, id uuid.UUID, signature *uuid.UUID) error {
	if signature == nil {
		return Fail(http.StatusUnprocessableEntity, "signature_required", "Sign first", "Confirm with your PIN first.")
	}
	return s.signedOn(ctx, table, id, *signature)
}

// childRows folds child rows for an API response.
func childRows[T any](rows []T) []map[string]any {
	out := make([]map[string]any, 0, len(rows))
	for _, r := range rows {
		raw, err := json.Marshal(r)
		if err != nil {
			continue
		}
		m := map[string]any{}
		if json.Unmarshal(raw, &m) == nil {
			out = append(out, m)
		}
	}
	return out
}
