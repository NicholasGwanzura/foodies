package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"foodies/api/internal/auth"
	"foodies/api/internal/db"
)

// stepCondition is the JSON a policy step may carry, to apply within a band.
type stepCondition struct {
	MinAmount string `json:"min_amount"`
	MaxAmount string `json:"max_amount"`
	Currency  string `json:"currency"`
}

func (c stepCondition) applies(amount, currency *string) bool {
	if c.MinAmount == "" && c.MaxAmount == "" {
		return true
	}
	if amount == nil {
		return false
	}
	if c.Currency != "" && currency != nil && c.Currency != *currency {
		return false
	}
	value, ok := new(big.Rat).SetString(*amount)
	if !ok {
		return false
	}
	if lo, ok := new(big.Rat).SetString(c.MinAmount); ok && value.Cmp(lo) < 0 {
		return false
	}
	if hi, ok := new(big.Rat).SetString(c.MaxAmount); ok && value.Cmp(hi) > 0 {
		return false
	}
	return true
}

// recordFacts pulls the reference and amount out of a row, whatever its table.
func recordFacts(row any) (label string, amount, currency *string) {
	raw, err := json.Marshal(row)
	if err != nil {
		return "", nil, nil
	}
	m := map[string]any{}
	if json.Unmarshal(raw, &m) != nil {
		return "", nil, nil
	}
	for _, key := range []string{"reference", "name", "title", "code"} {
		if v, ok := m[key].(string); ok && v != "" {
			label = v
			break
		}
	}
	for key, v := range m {
		name, ok := strings.CutSuffix(key, "_amount")
		if !ok {
			continue
		}
		if text, ok := v.(string); ok {
			if c, ok := m[name+"_currency"].(string); ok {
				if amount == nil || name == "amount" || name == "total" {
					amount, currency = &text, &c
				}
			}
		}
	}
	return label, amount, currency
}

// startApproval routes a record through the active policy for its workflow.
// With no active policy the workflow cannot proceed (spec D-10).
func (s *Service) startApproval(ctx context.Context, table string, id uuid.UUID, site *uuid.UUID, row any) error {
	q := s.q(ctx)
	policy, err := q.ActivePolicyFor(ctx, table)
	if errors.Is(err, pgx.ErrNoRows) {
		return &Error{Status: http.StatusConflict, Code: "rule_not_configured", Title: "No approval route",
			Detail:   "No approval policy is active for this workflow, so it cannot be submitted yet.",
			Blockers: []Blocker{{Kind: "rule_missing", Reference: table, Message: "Approval policy for " + strings.ReplaceAll(table, "_", " ")}}}
	}
	if err != nil {
		return fmt.Errorf("read approval policy: %w", err)
	}
	steps, err := q.ListApprovalPolicySteps(ctx, policy.ID)
	if err != nil {
		return fmt.Errorf("read policy steps: %w", err)
	}

	label, amount, currency := recordFacts(row)
	var chosen []db.ApprovalPolicyStep
	for _, st := range steps {
		cond := stepCondition{}
		if len(st.Condition) > 0 {
			if err := json.Unmarshal(st.Condition, &cond); err != nil {
				return fmt.Errorf("decode step condition: %w", err)
			}
		}
		if cond.applies(amount, currency) {
			chosen = append(chosen, st)
		}
	}
	if len(chosen) == 0 {
		return &Error{Status: http.StatusConflict, Code: "rule_not_configured", Title: "No approval route",
			Detail: "No step of the approval policy covers this amount."}
	}

	summary := strings.TrimSpace(humanTable(table) + " " + label)
	request, err := q.CreateApprovalRequest(ctx, db.CreateApprovalRequestParams{
		SubjectType: &table, SubjectID: &id, WorkflowKey: &table, PolicyID: &policy.ID, SiteID: site,
		AmountAmount: amount, AmountCurrency: currency, Summary: &summary, CreatedBy: actorID(ctx),
	})
	if err != nil {
		return wrap(err, "create approval request")
	}
	pending := "pending"
	for i, st := range chosen {
		seq := int32(i + 1) //nolint:gosec // a handful of steps
		var due *time.Time
		if st.SlaMinutes != nil {
			due = ptr(s.now().Add(time.Duration(*st.SlaMinutes) * time.Minute))
		}
		_, err := q.InsertApprovalRequestSteps(ctx, db.InsertApprovalRequestStepsParams{
			ApprovalRequestID: request.ID, Sequence: &seq, RoleID: &st.RoleID, Decision: &pending, DueAt: due, CreatedBy: actorID(ctx),
		})
		if err != nil {
			return wrap(err, "create approval step")
		}
	}
	if err := s.notifyRole(ctx, chosen[0].RoleID, site, "approvals", "Approval needed", summary, "approval_requests", request.ID); err != nil {
		return err
	}
	return s.audit(ctx, "approvals.approval_request.create", "approval_requests", request.ID, site, nil, request)
}

func humanTable(table string) string {
	words := strings.ReplaceAll(strings.TrimSuffix(table, "s"), "_", " ")
	if words == "" {
		return table
	}
	return strings.ToUpper(words[:1]) + words[1:]
}

// currentStep returns the request's open step.
func (s *Service) currentStep(ctx context.Context, r db.ApprovalRequest) (db.ApprovalRequestStep, error) {
	steps, err := s.q(ctx).ListApprovalRequestSteps(ctx, r.ID)
	if err != nil {
		return db.ApprovalRequestStep{}, wrap(err, "read steps")
	}
	for _, st := range steps {
		if st.Sequence != nil && r.CurrentStep != nil && *st.Sequence == *r.CurrentStep {
			return st, nil
		}
	}
	return db.ApprovalRequestStep{}, errState("This approval has no open step.")
}

// mayDecide checks the actor holds the step's role, directly, by reassignment
// or by a live delegation that covers it.
func (s *Service) mayDecide(ctx context.Context, r db.ApprovalRequest, st db.ApprovalRequestStep) error {
	p := auth.From(ctx)
	q := s.q(ctx)
	if st.AssignedUserID != nil {
		if *st.AssignedUserID == p.UserID {
			return nil
		}
		return errForbiddenAt("This approval is assigned to someone else.")
	}
	if st.RoleID == nil {
		return errForbiddenAt("This approval step has no role.")
	}
	holds, err := q.UserHoldsRole(ctx, db.UserHoldsRoleParams{UserID: p.UserID, RoleID: *st.RoleID, SiteID: r.SiteID})
	if err != nil {
		return fmt.Errorf("check role: %w", err)
	}
	if holds {
		return nil
	}
	delegated, err := q.DelegatorsTo(ctx, p.UserID)
	if err != nil {
		return fmt.Errorf("read delegations: %w", err)
	}
	for _, d := range delegated {
		if d.RoleID != nil && *d.RoleID == *st.RoleID {
			return nil
		}
	}
	return errForbiddenAt("Your role does not approve this step.")
}

func errForbiddenAt(detail string) *Error {
	return Fail(http.StatusForbidden, "forbidden", "Not allowed", detail)
}

// loadOpenApproval reads a pending request and checks the caller may act on it.
func (s *Service) loadOpenApproval(ctx context.Context, id uuid.UUID) (db.ApprovalRequest, db.ApprovalRequestStep, error) {
	var st db.ApprovalRequestStep
	r, err := s.q(ctx).GetApprovalRequest(ctx, id)
	if err != nil {
		return r, st, notFound(err, "approval")
	}
	if err := inScope(ctx, r.SiteID); err != nil {
		return r, st, err
	}
	if r.Status != "pending" {
		return r, st, errState("This approval has already been decided.")
	}
	if st, err = s.currentStep(ctx, r); err != nil {
		return r, st, err
	}
	return r, st, s.mayDecide(ctx, r, st)
}

// signedBy confirms a signature exists on this approval's subject, by the actor.
func (s *Service) signedBy(ctx context.Context, r db.ApprovalRequest, signature uuid.UUID) error {
	sig, err := s.q(ctx).GetSignatureRequest(ctx, signature)
	if err != nil {
		return notFound(err, "signature")
	}
	p := auth.From(ctx)
	if sig.Status != "signed" || sig.SignerEmployeeID == nil || *sig.SignerEmployeeID != p.EmployeeID ||
		r.SubjectID == nil || sig.SubjectID != *r.SubjectID {
		return Fail(http.StatusUnprocessableEntity, "signature_required", "Sign first", "Confirm with your PIN before deciding.")
	}
	return nil
}

func (s *Service) decide(ctx context.Context, st db.ApprovalRequestStep, decision string, comment *string, signature *uuid.UUID) error {
	p := auth.From(ctx)
	err := s.q(ctx).DecideApprovalStep(ctx, db.DecideApprovalStepParams{
		ID: st.ID, DecidedBy: &p.UserID, Decision: &decision, Channel: &p.Channel, Comment: comment, SignatureID: signature,
	})
	if err != nil {
		return wrap(err, "record decision")
	}
	return nil
}

// ApproveApprovalRequest records one step's approval, moving to the next step
// or finishing the request. A retrospective approval is impossible here,
// because the subject stays blocked until this runs.
func (s *Service) ApproveApprovalRequest(ctx context.Context, id uuid.UUID, comment *string, signature uuid.UUID) (db.ApprovalRequest, error) {
	r, st, err := s.loadOpenApproval(ctx, id)
	if err != nil {
		return r, err
	}
	if r.CreatedBy != nil && *r.CreatedBy == auth.From(ctx).UserID {
		return r, &Error{Status: http.StatusForbidden, Code: "sod_violation", Title: "Someone else must do this",
			Detail: "You cannot approve your own request.", Blockers: []Blocker{{Kind: "sod_conflict", Message: "Requester and approver must differ"}}}
	}
	if err := s.signedBy(ctx, r, signature); err != nil {
		return r, err
	}
	if err := s.beforeApprove(ctx, r); err != nil {
		return r, err
	}
	if err := s.decide(ctx, st, "approved", comment, &signature); err != nil {
		return r, err
	}
	q := s.q(ctx)
	steps, err := q.ListApprovalRequestSteps(ctx, r.ID)
	if err != nil {
		return r, wrap(err, "read steps")
	}
	var next *db.ApprovalRequestStep
	for i := range steps {
		if steps[i].Sequence != nil && *steps[i].Sequence > *r.CurrentStep {
			next = &steps[i]
			break
		}
	}
	if next != nil {
		row, err := q.AdvanceApproval(ctx, db.AdvanceApprovalParams{ID: r.ID, CurrentStep: next.Sequence})
		if err != nil {
			return r, wrap(err, "advance approval")
		}
		if next.RoleID != nil {
			if err := s.notifyRole(ctx, *next.RoleID, r.SiteID, "approvals", "Approval needed", deref(r.Summary), "approval_requests", r.ID); err != nil {
				return row, err
			}
		}
		return row, s.audit(ctx, "approvals.approval_request.approve", "approval_requests", r.ID, r.SiteID, r, row)
	}
	return s.finishApproval(ctx, r, "approved")
}

func deref(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

// RejectApprovalRequest ends the request.
func (s *Service) RejectApprovalRequest(ctx context.Context, id uuid.UUID, comment string, signature uuid.UUID) (db.ApprovalRequest, error) {
	r, st, err := s.loadOpenApproval(ctx, id)
	if err != nil {
		return r, err
	}
	if strings.TrimSpace(comment) == "" {
		return r, errInvalid("/comment", "Say why you are rejecting this.")
	}
	if err := s.signedBy(ctx, r, signature); err != nil {
		return r, err
	}
	if err := s.decide(ctx, st, "rejected", &comment, &signature); err != nil {
		return r, err
	}
	return s.finishApproval(ctx, r, "rejected")
}

// ReturnApprovalRequest sends the record back to its author.
func (s *Service) ReturnApprovalRequest(ctx context.Context, id uuid.UUID, comment string) (db.ApprovalRequest, error) {
	r, st, err := s.loadOpenApproval(ctx, id)
	if err != nil {
		return r, err
	}
	if strings.TrimSpace(comment) == "" {
		return r, errInvalid("/comment", "Say what needs changing.")
	}
	if err := s.decide(ctx, st, "returned", &comment, nil); err != nil {
		return r, err
	}
	return s.finishApproval(ctx, r, "returned")
}

// ReassignApprovalRequest hands the open step to one named person.
func (s *Service) ReassignApprovalRequest(ctx context.Context, id, user uuid.UUID, reason string) (db.ApprovalRequest, error) {
	r, st, err := s.loadOpenApproval(ctx, id)
	if err != nil {
		return r, err
	}
	if err := s.q(ctx).ReassignApprovalStep(ctx, db.ReassignApprovalStepParams{ID: st.ID, AssignedUserID: &user}); err != nil {
		return r, wrap(err, "reassign step")
	}
	if err := s.notify(ctx, []uuid.UUID{user}, "approvals", "Approval reassigned to you", deref(r.Summary), "approval_requests", r.ID); err != nil {
		return r, err
	}
	return r, s.audit(ctx, "approvals.approval_request.reassign", "approval_requests", r.ID, r.SiteID, nil, map[string]any{"user_id": user, "reason": reason})
}

// finishApproval closes the request and applies the outcome to its subject.
func (s *Service) finishApproval(ctx context.Context, r db.ApprovalRequest, outcome string) (db.ApprovalRequest, error) {
	q := s.q(ctx)
	row, err := q.FinishApproval(ctx, db.FinishApprovalParams{ID: r.ID, Status: outcome})
	if err != nil {
		return r, wrap(err, "finish approval")
	}
	if r.SubjectType != nil && r.SubjectID != nil {
		if err := s.applyOutcome(ctx, *r.SubjectType, *r.SubjectID, outcome); err != nil {
			return row, err
		}
		if r.CreatedBy != nil {
			title := "Approved"
			switch outcome {
			case "rejected":
				title = "Rejected"
			case "returned":
				title = "Returned for changes"
			}
			if err := s.notify(ctx, []uuid.UUID{*r.CreatedBy}, "approvals", title, deref(r.Summary), *r.SubjectType, *r.SubjectID); err != nil {
				return row, err
			}
		}
	}
	return row, s.audit(ctx, "approvals.approval_request."+strings.TrimSuffix(outcome, "ed"), "approval_requests", r.ID, r.SiteID, r, row)
}

// outcomeStatus picks the subject status for a decision, from what the table has.
func outcomeStatus(statuses []string, outcome string) string {
	var wants []string
	switch outcome {
	case "approved":
		wants = []string{"approved", "active", "published", "effective", "issued"}
	case "rejected":
		wants = []string{"rejected", "cancelled", "draft"}
	default:
		wants = []string{"returned", "draft", "open"}
	}
	for _, w := range wants {
		if slices.Contains(statuses, w) {
			return w
		}
	}
	return ""
}

// approvalHooks run when a subject's approval finishes, for tables that need
// more than a status change.
func (s *Service) applyOutcome(ctx context.Context, table string, id uuid.UUID, outcome string) error {
	sub, ok := subjects[table]
	if !ok || sub.setStatus == nil {
		return nil
	}
	status := outcomeStatus(sub.statuses, outcome)
	if status == "" {
		return nil
	}
	if err := sub.setStatus(ctx, s.q(ctx), id, status, actorID(ctx)); err != nil {
		return wrap(err, "apply decision")
	}
	return s.afterDecision(ctx, table, id, outcome)
}
