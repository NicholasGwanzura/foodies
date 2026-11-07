package service

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"foodies/api/internal/auth"
	"foodies/api/internal/db"
)

// Tile is one dashboard figure.
type Tile struct {
	Key, Label string
	Value      int64
	Status     string
	Drilldown  string
}

// Dashboard returns the figures for a dashboard key.
func (s *Service) Dashboard(ctx context.Context) ([]Tile, string, error) {
	q := s.q(ctx)
	p := auth.From(ctx)
	all, sites := scopeOf(ctx)
	roles, err := q.UserRoleIds(ctx, p.UserID)
	if err != nil {
		return nil, "", fmt.Errorf("read roles: %w", err)
	}
	open, err := q.CountOpenTasksFor(ctx, db.CountOpenTasksForParams{UserID: &p.UserID, RoleIds: roles, AllSites: all, SiteIds: sites})
	if err != nil {
		return nil, "", fmt.Errorf("count tasks: %w", err)
	}
	overdue, err := q.CountOverdueTasksFor(ctx, db.CountOverdueTasksForParams{UserID: &p.UserID, RoleIds: roles, AllSites: all, SiteIds: sites})
	if err != nil {
		return nil, "", fmt.Errorf("count overdue: %w", err)
	}
	inbox, err := s.ApprovalInbox(ctx, nil, 101)
	if err != nil {
		return nil, "", err
	}
	escalations, err := q.CountOpenEscalations(ctx, db.CountOpenEscalationsParams{AllSites: all, SiteIds: sites})
	if err != nil {
		return nil, "", fmt.Errorf("count escalations: %w", err)
	}
	unread, err := q.CountUnreadNotifications(ctx, &p.UserID)
	if err != nil {
		return nil, "", fmt.Errorf("count notifications: %w", err)
	}
	status := func(n int64) string {
		if n > 0 {
			return "warn"
		}
		return "ok"
	}
	return []Tile{
		{"tasks_due", "Tasks due", open, status(overdue), "/tasks"},
		{"overdue", "Overdue", overdue, status(overdue), "/tasks?overdue=true"},
		{"approvals_waiting", "Approvals waiting", int64(len(inbox)), status(int64(len(inbox))), "/approvals"},
		{"open_escalations", "Open escalations", escalations, status(escalations), "/escalations"},
		{"unread", "Unread", unread, "ok", "/notifications"},
	}, "My Day", nil
}

// ApprovalInbox lists what waits on the signed-in user.
func (s *Service) ApprovalInbox(ctx context.Context, cursor *uuid.UUID, limit int32) ([]db.ApprovalRequest, error) {
	return wrap2(s.q(ctx).ApprovalInbox(ctx, db.ApprovalInboxParams{UserID: ptr(auth.From(ctx).UserID), CursorID: cursor, RowLimit: limit}))("read inbox")
}

// Search finds records by reference or title, within scope.
func (s *Service) Search(ctx context.Context, term string, limit int32) ([]db.SearchRecordsRow, error) {
	term = strings.TrimSpace(term)
	if term == "" {
		return nil, errInvalid("/q", "Enter something to find.")
	}
	all, sites := scopeOf(ctx)
	return wrap2(s.q(ctx).SearchRecords(ctx, db.SearchRecordsParams{Term: term, AllSites: all, SiteIds: sites, RowLimit: limit}))("search")
}

// EffectiveFxRate returns the rate on a date: the latest at or before it.
func (s *Service) EffectiveFxRate(ctx context.Context, from, to string, on pgtype.Date) (db.FxRate, error) {
	row, err := s.q(ctx).EffectiveFxRate(ctx, db.EffectiveFxRateParams{FromCurrency: from, ToCurrency: to, RateDate: on})
	if err != nil {
		return row, notFound(err, "exchange rate")
	}
	return row, nil
}

// UpdateFxRate changes a rate, which must say why (spec D-15).
func (s *Service) UpdateFxRate(ctx context.Context, p db.UpdateFxRateParams) (db.FxRate, error) {
	before, err := s.GetFxRate(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	if p.Rate != nil && *p.Rate != before.Rate && strings.TrimSpace(deref(p.ChangeReason)) == "" {
		return before, errInvalid("/change_reason", "Say why this rate is changing.")
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateFxRate(ctx, p)
	if err != nil {
		return before, wrap(err, "update rate")
	}
	return row, s.audit(ctx, "org.fx_rate.write", "fx_rates", row.ID, nil, before, row)
}

// SyncManifest gives a device what it needs to work offline.
func (s *Service) SyncManifest(ctx context.Context) (map[string]any, error) {
	q := s.q(ctx)
	rules, err := q.ActiveRuleSnapshot(ctx)
	if err != nil {
		return nil, wrap(err, "read rules")
	}
	values := map[string]json.RawMessage{}
	for _, r := range rules {
		values[r.Key] = r.Value
	}
	forms, err := q.OfflineTemplates(ctx)
	if err != nil {
		return nil, wrap(err, "read templates")
	}
	raw, err := json.Marshal(values)
	if err != nil {
		return nil, fmt.Errorf("encode rules: %w", err)
	}
	sum := sha256.Sum256(raw)
	return map[string]any{
		"generated_at": s.now().UTC(), "rule_snapshot_id": hex.EncodeToString(sum[:8]),
		"rule_values": values, "form_template_versions": forms,
	}, nil
}

// RecordConflict sends a write the server could not apply to a supervisor.
func (s *Service) RecordConflict(ctx context.Context, batch string, device, clientRecord uuid.UUID, reason string, payload json.RawMessage) error {
	_, err := s.q(ctx).CreateSyncConflict(ctx, db.CreateSyncConflictParams{
		BatchID: &batch, DeviceID: &device, ClientRecordID: &clientRecord, Reason: &reason, ClientPayload: payload, SiteID: &auth.From(ctx).SiteID,
	})
	return wrap(err, "record conflict")
}

// VerifyWhatsAppToken answers Meta's webhook verification.
func (s *Service) VerifyWhatsAppToken(token string) bool {
	return s.config.WhatsAppToken != "" && hmac.Equal([]byte(token), []byte(s.config.WhatsAppToken))
}

// VerifyWhatsAppSignature checks X-Hub-Signature-256.
func (s *Service) VerifyWhatsAppSignature(body []byte, header string) bool {
	if s.config.WhatsAppKey == "" {
		return false
	}
	mac := hmac.New(sha256.New, []byte(s.config.WhatsAppKey))
	mac.Write(body)
	return hmac.Equal([]byte(strings.TrimPrefix(header, "sha256=")), []byte(hex.EncodeToString(mac.Sum(nil))))
}

type whatsappWebhook struct {
	Entry []struct {
		Changes []struct {
			Value struct {
				Messages []struct {
					From string `json:"from"`
					ID   string `json:"id"`
					Type string `json:"type"`
				} `json:"messages"`
			} `json:"value"`
		} `json:"changes"`
	} `json:"entry"`
}

// ReceiveWhatsApp logs inbound messages, ignoring repeats by provider id. The
// body holds no personal detail beyond the message type, per spec 9.3.
func (s *Service) ReceiveWhatsApp(ctx context.Context, body []byte) error {
	var hook whatsappWebhook
	if err := json.Unmarshal(body, &hook); err != nil {
		return errInvalid("", "The webhook body is not valid.")
	}
	q := s.q(ctx)
	inbound := "inbound"
	for _, e := range hook.Entry {
		for _, c := range e.Changes {
			for _, m := range c.Value.Messages {
				at := s.now()
				identity, err := q.IdentityByPhone(ctx, "+"+strings.TrimPrefix(m.From, "+"))
				var id *uuid.UUID
				if err == nil {
					id = &identity.ID
				}
				err = q.InsertWhatsappMessage(ctx, db.InsertWhatsappMessageParams{
					Direction: &inbound, IdentityID: id, ProviderMessageID: &m.ID, MessageType: &m.Type, SentAt: &at, DeliveryStatus: ptr("received"),
				})
				if err != nil {
					return wrap(err, "log message")
				}
			}
		}
	}
	return nil
}

// RecordIntegrationEvents authenticates a system by key and logs a sync run.
func (s *Service) RecordIntegrationEvents(ctx context.Context, system, key string) error {
	denied := Fail(http.StatusUnauthorized, "unauthenticated", "Unknown system", "The integration key is not valid.")
	row, err := s.q(ctx).IntegrationBySystem(ctx, system)
	if err != nil || key == "" {
		return denied
	}
	var config struct {
		KeyHash string `json:"api_key_hash"`
	}
	if json.Unmarshal(row.Config, &config) != nil || config.KeyHash == "" ||
		!hmac.Equal([]byte(config.KeyHash), []byte(auth.HashOpaque(key))) {
		return denied
	}
	return wrap(s.q(ctx).TouchIntegration(ctx, row.ID), "record event")
}

// RetryPosting queues a failed posting again.
func (s *Service) RetryPosting(ctx context.Context, p db.ActRetryPostingParams) (db.Posting, error) {
	before, err := s.GetPosting(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if before.Status != "failed" && before.Status != "manual_required" {
		return before, errState("Only a failed posting can be retried.")
	}
	row, err := s.q(ctx).RequeuePosting(ctx, p.ID)
	if err != nil {
		return before, wrap(err, "requeue posting")
	}
	return row, s.audit(ctx, "integrations.posting.retry", "postings", row.ID, before.SiteID, before, row)
}
