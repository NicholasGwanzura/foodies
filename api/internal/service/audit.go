package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"foodies/api/internal/auth"
	"foodies/api/internal/db"
)

// audit appends a hash-chained event in the caller's transaction, so the event
// and the change it describes commit together.
func (s *Service) audit(ctx context.Context, action, subjectType string, subjectID uuid.UUID, site *uuid.UUID, before, after any) error {
	return s.InTx(ctx, func(ctx context.Context) error {
		q := s.q(ctx)
		if err := q.LockAuditChain(ctx); err != nil {
			return fmt.Errorf("lock audit chain: %w", err)
		}
		last, err := q.LastAuditHash(ctx)
		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("read audit chain: %w", err)
		}

		beforeJSON, err := auditJSON(before)
		if err != nil {
			return err
		}
		afterJSON, err := auditJSON(after)
		if err != nil {
			return err
		}

		channel, role, ip := "system", "", ""
		var device *uuid.UUID
		p := auth.From(ctx)
		if p != nil {
			channel, ip, device = p.Channel, p.IP, p.DeviceID
			if len(p.Grants) > 0 {
				role = p.Grants[0].RoleKey
			}
		}

		at := s.now().UTC()
		prev := ""
		if last != nil {
			prev = *last
		}
		sum := sha256.New()
		sum.Write(fmt.Appendf(nil, "%s|%s|%s|%s|%s|%d|", prev, action, subjectType, subjectID, channel, at.UnixMicro()))
		sum.Write(beforeJSON)
		sum.Write([]byte{'|'})
		sum.Write(afterJSON)
		hash := hex.EncodeToString(sum.Sum(nil))

		err = q.InsertAuditEvent(ctx, db.InsertAuditEventParams{
			OccurredAt: &at, ActorUserID: actorID(ctx), ActorRole: nonEmpty(role), Action: &action,
			SubjectType: &subjectType, SubjectID: &subjectID, SiteID: site, Channel: &channel, DeviceID: device,
			Ip: nonEmpty(ip), Before: beforeJSON, After: afterJSON, PrevHash: &prev, Hash: &hash,
		})
		if err != nil {
			return fmt.Errorf("write audit event: %w", err)
		}
		return nil
	})
}

func auditJSON(v any) (json.RawMessage, error) {
	if v == nil {
		return nil, nil
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("encode audit value: %w", err)
	}
	return b, nil
}

func nonEmpty(s string) *string {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	return &s
}

// LogRead records a restricted read. The operation has no transaction of its
// own, so the event commits on its own.
func (s *Service) LogRead(ctx context.Context, operation, path string) error {
	id, err := uuid.NewV7()
	if err != nil {
		return fmt.Errorf("new audit id: %w", err)
	}
	return s.audit(ctx, "read."+operation, "restricted_read", id, nil, nil, map[string]string{"path": path})
}
