// Package service holds the business logic. HTTP handlers and the WhatsApp
// channel both call it, so rules, approvals and audit behave the same.
package service

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"foodies/api/internal/auth"
	"foodies/api/internal/db"
)

// Harare is the business time zone. Africa/Harare has no daylight saving.
var Harare = time.FixedZone("Africa/Harare", 2*60*60)

// Service is the entry point to the business logic.
type Service struct {
	pool    *pgxpool.Pool
	queries db.Querier
	signer  *auth.Signer
	now     func() time.Time
	config  Config
}

// Config holds settings that only some features need.
type Config struct {
	UploadDir     string
	PublicBaseURL string
	SigningSecret string
	WhatsAppToken string
	WhatsAppKey   string
}

// Configure applies settings after New.
func (s *Service) Configure(c Config) { s.config = c }

// New returns a Service backed by queries. The pool may be nil in unit tests,
// which makes InTx run its function without a transaction.
func New(queries db.Querier, pool *pgxpool.Pool, signer *auth.Signer) *Service {
	return &Service{pool: pool, queries: queries, signer: signer, now: time.Now}
}

// Ping confirms the database answers.
func (s *Service) Ping(ctx context.Context) error {
	if _, err := s.queries.Ping(ctx); err != nil {
		return fmt.Errorf("ping database: %w", err)
	}
	return nil
}

type txKey struct{}

// q returns the queries for ctx, inside its transaction if it has one.
func (s *Service) q(ctx context.Context) db.Querier {
	if tx, ok := ctx.Value(txKey{}).(pgx.Tx); ok {
		return db.New(tx)
	}
	return s.queries
}

// InTx runs fn in one transaction, joining any already open in ctx.
func (s *Service) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	if s.pool == nil || ctx.Value(txKey{}) != nil {
		return fn(ctx)
	}
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin: %w", err)
	}
	// Rollback after a commit is a no-op, and a failed rollback has no remedy.
	defer tx.Rollback(context.WithoutCancel(ctx)) //nolint:errcheck // see above

	if err := fn(context.WithValue(ctx, txKey{}, tx)); err != nil {
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit: %w", err)
	}
	return nil
}
