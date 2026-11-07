package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"foodies/api/internal/db"
)

// groupSiteCode stands in for the site in references with no site of their own.
const groupSiteCode = "GRP"

// nextReference returns PREFIX-SITE-YEAR-SEQUENCE. The counter update is in the
// caller's transaction, so a rollback leaves no gap in the sequence.
func (s *Service) nextReference(ctx context.Context, prefix string, site *uuid.UUID) (string, error) {
	code := groupSiteCode
	q := s.q(ctx)
	if site != nil {
		got, err := q.GetSiteCode(ctx, *site)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return "", fmt.Errorf("read site code: %w", err)
		}
		if err == nil {
			code = got
		}
	}
	year := s.now().In(Harare).Year()
	n, err := q.NextReference(ctx, db.NextReferenceParams{Prefix: prefix, SiteCode: code, Year: int32(year)}) //nolint:gosec // a year fits
	if err != nil {
		return "", fmt.Errorf("next reference: %w", err)
	}
	return fmt.Sprintf("%s-%s-%d-%06d", prefix, code, year, n), nil
}
