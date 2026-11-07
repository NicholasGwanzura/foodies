package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"foodies/api/internal/db"
)

// Replay is a stored response to send again.
type Replay struct {
	Status  int
	Headers map[string]string
	Body    []byte
}

// inFlightWindow is how long a claimed key blocks a repeat before it counts as
// abandoned, for example after a crash between commit and storing the response.
const inFlightWindow = time.Minute

// ClaimIdempotency reserves key for actor. It returns a Replay when the key
// already holds a finished response, and nil when the caller should proceed.
func (s *Service) ClaimIdempotency(ctx context.Context, actor uuid.UUID, key, hash string) (*Replay, error) {
	q := s.q(ctx)
	claimed, err := q.ClaimIdempotencyKey(ctx, db.ClaimIdempotencyKeyParams{ActorID: actor, Key: key, RequestHash: hash})
	if err != nil {
		return nil, fmt.Errorf("claim idempotency key: %w", err)
	}
	if claimed {
		return nil, nil
	}
	row, err := q.GetIdempotencyKey(ctx, db.GetIdempotencyKeyParams{ActorID: actor, Key: key})
	if err != nil {
		return nil, fmt.Errorf("read idempotency key: %w", err)
	}
	if row.RequestHash != hash {
		return nil, Fail(http.StatusConflict, "idempotency_key_reused", "Key already used",
			"This Idempotency-Key was used with a different request.")
	}
	if row.Status == nil {
		if s.now().Sub(row.CreatedAt) > inFlightWindow {
			return nil, nil
		}
		return nil, Fail(http.StatusConflict, "conflict", "Already in progress", "This request is still being processed.")
	}
	replay := &Replay{Status: int(*row.Status), Body: row.Body, Headers: map[string]string{}}
	if len(row.Headers) > 0 {
		if err := json.Unmarshal(row.Headers, &replay.Headers); err != nil {
			return nil, fmt.Errorf("decode stored headers: %w", err)
		}
	}
	return replay, nil
}

// FinishIdempotency stores a successful response for replay.
func (s *Service) FinishIdempotency(ctx context.Context, actor uuid.UUID, key string, status int, headers map[string]string, body []byte) error {
	encoded, err := json.Marshal(headers)
	if err != nil {
		return fmt.Errorf("encode headers: %w", err)
	}
	code := int32(status) //nolint:gosec // an HTTP status fits
	err = s.q(ctx).FinishIdempotencyKey(ctx, db.FinishIdempotencyKeyParams{ActorID: actor, Key: key, Status: &code, Headers: encoded, Body: body})
	if err != nil {
		return fmt.Errorf("store idempotent response: %w", err)
	}
	return nil
}

// ReleaseIdempotency frees a key whose request failed.
func (s *Service) ReleaseIdempotency(ctx context.Context, actor uuid.UUID, key string) error {
	if err := s.q(ctx).ReleaseIdempotencyKey(ctx, db.ReleaseIdempotencyKeyParams{ActorID: actor, Key: key}); err != nil {
		return fmt.Errorf("release idempotency key: %w", err)
	}
	return nil
}
