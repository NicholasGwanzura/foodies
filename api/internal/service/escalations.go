package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"foodies/api/internal/auth"
	"foodies/api/internal/db"
)

// RaiseEscalation starts the policy for a domain and severity. With no policy,
// the record stays visible on the board with nobody notified yet, because the
// recipients and timers are unset rules (C-06.09).
func (s *Service) RaiseEscalation(ctx context.Context, subjectType string, subject uuid.UUID, domain string, severity int, site *uuid.UUID) (db.Escalation, error) {
	sev := int32(severity) //nolint:gosec // 1 to 5
	row, err := s.q(ctx).CreateEscalationRow(ctx, db.CreateEscalationRowParams{
		SubjectType: &subjectType, SubjectID: &subject, Domain: &domain, SeverityLevel: &sev, SiteID: site, CreatedBy: actorID(ctx),
	})
	if err != nil {
		return row, wrap(err, "create escalation")
	}
	if err := s.notifyStep(ctx, row, 1); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "exceptions.escalation.create", "escalations", row.ID, site, nil, row)
}

func (s *Service) policySteps(ctx context.Context, e db.Escalation) ([]db.EscalationPolicyStep, error) {
	if e.Domain == nil || e.SeverityLevel == nil {
		return nil, nil
	}
	policy, err := s.q(ctx).ActiveEscalationPolicy(ctx, db.ActiveEscalationPolicyParams{Domain: *e.Domain, SeverityLevel: *e.SeverityLevel})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read escalation policy: %w", err)
	}
	steps, err := s.q(ctx).ListEscalationPolicySteps(ctx, policy.ID)
	if err != nil {
		return nil, fmt.Errorf("read escalation steps: %w", err)
	}
	return steps, nil
}

func (s *Service) notifyStep(ctx context.Context, e db.Escalation, step int32) error {
	steps, err := s.policySteps(ctx, e)
	if err != nil {
		return err
	}
	q := s.q(ctx)
	for _, st := range steps {
		if st.Sequence != step {
			continue
		}
		users, err := q.UsersWithRoleAtSite(ctx, db.UsersWithRoleAtSiteParams{RoleID: st.NotifyRoleID, SiteID: e.SiteID})
		if err != nil {
			return fmt.Errorf("find recipients: %w", err)
		}
		channel := strings.Join(st.Channels, ",")
		for _, u := range users {
			if _, err := q.AddEscalationRecipient(ctx, db.AddEscalationRecipientParams{EscalationID: e.ID, Step: &step, UserID: &u, Channel: &channel}); err != nil {
				return wrap(err, "add recipient")
			}
		}
		title := fmt.Sprintf("Level %d escalation", deref32(e.SeverityLevel))
		return s.notify(ctx, users, "escalations", title, deref(e.SubjectType), "escalations", e.ID)
	}
	return nil
}

func deref32(p *int32) int32 {
	if p == nil {
		return 0
	}
	return *p
}

// AcknowledgeEscalation stops the timer.
func (s *Service) AcknowledgeEscalation(ctx context.Context, p db.ActAcknowledgeEscalationParams) (db.Escalation, error) {
	before, err := s.GetEscalation(ctx, p.ID)
	if err != nil {
		return before, err
	}
	q := s.q(ctx)
	if _, err := q.AckEscalationRecipient(ctx, db.AckEscalationRecipientParams{EscalationID: p.ID, UserID: actorID(ctx)}); err != nil {
		return before, wrap(err, "acknowledge")
	}
	row, err := q.AckEscalationRow(ctx, db.AckEscalationRowParams{ID: p.ID, UpdatedBy: actorID(ctx)})
	if err != nil {
		return before, stateError(err, "acknowledge this")
	}
	return row, s.audit(ctx, "exceptions.escalation.acknowledge", "escalations", row.ID, row.SiteID, before, row)
}

// AdvanceEscalation moves to the next level by hand.
func (s *Service) AdvanceEscalation(ctx context.Context, p db.ActAdvanceEscalationParams) (db.Escalation, error) {
	before, err := s.GetEscalation(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if before.Status == "closed" {
		return before, errState("This escalation is closed.")
	}
	row, err := s.moveUp(ctx, before, "open")
	if err != nil {
		return row, err
	}
	return row, s.audit(ctx, "exceptions.escalation.advance", "escalations", row.ID, row.SiteID, before, row)
}

func (s *Service) moveUp(ctx context.Context, e db.Escalation, status string) (db.Escalation, error) {
	next := deref32(e.CurrentStep) + 1
	steps, err := s.policySteps(ctx, e)
	if err != nil {
		return e, err
	}
	if int(next) > len(steps) {
		next = deref32(e.CurrentStep)
		status = "breached"
	}
	row, err := s.q(ctx).SetEscalationStep(ctx, db.SetEscalationStepParams{ID: e.ID, CurrentStep: &next, Status: status})
	if err != nil {
		return e, wrap(err, "move escalation")
	}
	if status != "breached" {
		return row, s.notifyStep(ctx, row, next)
	}
	return row, nil
}

// checkEscalationTimers moves up anything not acknowledged in time.
func (s *Service) checkEscalationTimers(ctx context.Context) error {
	rows, err := s.q(ctx).EscalationsToCheck(ctx)
	if err != nil {
		return fmt.Errorf("find late escalations: %w", err)
	}
	for _, r := range rows {
		e := db.Escalation{ID: r.ID, Domain: r.Domain, SeverityLevel: r.SeverityLevel, SiteID: r.SiteID, CurrentStep: r.CurrentStep, SubjectType: r.SubjectType}
		status := "open"
		if !r.NextOnTimeout {
			status = "breached"
			if _, err := s.q(ctx).SetEscalationStep(ctx, db.SetEscalationStepParams{ID: r.ID, CurrentStep: r.CurrentStep, Status: status}); err != nil {
				return wrap(err, "breach escalation")
			}
			continue
		}
		if _, err := s.moveUp(ctx, e, status); err != nil {
			return err
		}
	}
	return nil
}

// VerifyAction records the effectiveness check, by someone other than the owner.
func (s *Service) VerifyAction(ctx context.Context, p db.ActVerifyActionParams) (db.Action, error) {
	before, err := s.GetAction(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "exceptions.action.verify", "actions", p.ID); err != nil {
		return before, err
	}
	if a := auth.From(ctx); a != nil && before.OwnerEmployeeID == a.EmployeeID {
		return before, &Error{Status: 403, Code: "sod_violation", Title: "Someone else must do this", Detail: "The action owner cannot verify their own work."}
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActVerifyAction(ctx, p)
	if err != nil {
		return before, stateError(err, "verify this action")
	}
	return row, s.audit(ctx, "exceptions.action.verify", "actions", row.ID, row.SiteID, before, row)
}
