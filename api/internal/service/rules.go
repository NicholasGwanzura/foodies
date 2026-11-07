package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"foodies/api/internal/db"
)

// RuleValue returns the active value of a rule at a site, or an error naming it.
func (s *Service) RuleValue(ctx context.Context, key string, site *uuid.UUID) (db.RuleVersion, error) {
	row, err := s.q(ctx).ActiveRuleValue(ctx, db.ActiveRuleValueParams{Key: key, SiteID: site})
	if errors.Is(err, pgx.ErrNoRows) {
		return row, &Error{Status: http.StatusConflict, Code: "rule_not_configured", Title: "Rule not set",
			Detail:   "The rule " + key + " has no approved value yet.",
			Blockers: []Blocker{{Kind: "rule_missing", Reference: key, Message: key + " has no approved value"}}}
	}
	return row, wrap(err, "read rule")
}

// requireLive stops a workflow that is not live at a site (spec 8.1).
func (s *Service) requireLive(ctx context.Context, workflow string, site uuid.UUID) error {
	row, err := s.q(ctx).GetActivation(ctx, db.GetActivationParams{WorkflowKey: workflow, SiteID: site})
	if errors.Is(err, pgx.ErrNoRows) {
		return nil // No rules are required, so nothing gates this workflow.
	}
	if err != nil {
		return fmt.Errorf("read activation: %w", err)
	}
	if row.Status == "live" {
		return nil
	}
	blockers := []Blocker{}
	for _, key := range row.MissingRuleKeys {
		blockers = append(blockers, Blocker{Kind: "rule_missing", Reference: key, Message: key + " has no approved value"})
	}
	if len(blockers) == 0 {
		blockers = append(blockers, Blocker{Kind: "workflow_not_live", Reference: workflow, Message: "This workflow has not been switched on here"})
	}
	return &Error{Status: http.StatusConflict, Code: "workflow_not_live", Title: "Workflow not live at this site",
		Detail: fmt.Sprintf("This cannot start at this site until %d item(s) are dealt with.", len(blockers)), Blockers: blockers}
}

// RefreshReadiness recomputes every workflow at every active site.
func (s *Service) RefreshReadiness(ctx context.Context) error {
	q := s.q(ctx)
	workflows, err := q.ListWorkflowKeys(ctx)
	if err != nil {
		return fmt.Errorf("list workflows: %w", err)
	}
	sites, err := q.ListActiveSiteIds(ctx)
	if err != nil {
		return fmt.Errorf("list sites: %w", err)
	}
	for _, w := range workflows {
		for _, site := range sites {
			if err := q.EnsureActivation(ctx, db.EnsureActivationParams{WorkflowKey: w, SiteID: site}); err != nil {
				return fmt.Errorf("ensure activation: %w", err)
			}
			missing, err := q.MissingRules(ctx, db.MissingRulesParams{WorkflowKey: w, SiteID: &site})
			if err != nil {
				return fmt.Errorf("find missing rules: %w", err)
			}
			keys := make([]string, 0, len(missing))
			for _, m := range missing {
				keys = append(keys, m.RuleKey)
			}
			err = q.SetActivationReadiness(ctx, db.SetActivationReadinessParams{WorkflowKey: w, SiteID: site, MissingRuleKeys: keys})
			if err != nil {
				return fmt.Errorf("store readiness: %w", err)
			}
		}
	}
	return nil
}

// ActivateWorkflowActivation switches a ready workflow on at a site.
func (s *Service) ActivateWorkflowActivation(ctx context.Context, p db.ActActivateWorkflowActivationParams) (db.WorkflowActivation, error) {
	if err := s.RefreshReadiness(ctx); err != nil {
		return db.WorkflowActivation{}, err
	}
	before, err := s.GetWorkflowActivation(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if len(before.MissingRuleKeys) > 0 {
		blockers := make([]Blocker, 0, len(before.MissingRuleKeys))
		for _, k := range before.MissingRuleKeys {
			blockers = append(blockers, Blocker{Kind: "rule_missing", Reference: k, Message: k + " has no approved value"})
		}
		return before, &Error{Status: http.StatusConflict, Code: "workflow_not_live", Title: "Not ready",
			Detail: "This workflow still has rules without an approved value.", Blockers: blockers}
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActActivateWorkflowActivation(ctx, p)
	if err != nil {
		return before, stateError(err, "activate this workflow")
	}
	return row, s.audit(ctx, "rules.workflow_activation.activate", "workflow_activations", row.ID, &row.SiteID, before, row)
}

// afterDecision does the work a table needs beyond a status change.
func (s *Service) afterDecision(ctx context.Context, table string, id uuid.UUID, outcome string) error {
	if outcome != "approved" {
		return nil
	}
	q := s.q(ctx)
	switch table {
	case "rule_versions":
		v, err := q.GetRuleVersion(ctx, id)
		if err != nil {
			return wrap(err, "read rule version")
		}
		err = q.SupersedeRuleVersions(ctx, db.SupersedeRuleVersionsParams{RuleDefinitionID: v.RuleDefinitionID, ScopeType: v.ScopeType, ScopeID: v.ScopeID, ID: v.ID})
		if err != nil {
			return wrap(err, "supersede rule versions")
		}
		return s.RefreshReadiness(ctx)
	case "approval_policies":
		pol, err := q.GetApprovalPolicy(ctx, id)
		if err != nil {
			return wrap(err, "read policy")
		}
		return wrap(q.SupersedePolicies(ctx, db.SupersedePoliciesParams{WorkflowKey: pol.WorkflowKey, ID: pol.ID}), "supersede policies")
	case "form_template_versions":
		v, err := q.GetFormTemplateVersion(ctx, id)
		if err != nil {
			return wrap(err, "read version")
		}
		return s.publishVersion(ctx, v)
	case "employee_transfers":
		return s.applyTransfer(ctx, id)
	}
	return nil
}

// Readiness returns the go-live state of workflows in the caller's scope.
func (s *Service) Readiness(ctx context.Context, site *uuid.UUID) ([]db.WorkflowActivation, error) {
	if err := s.RefreshReadiness(ctx); err != nil {
		return nil, err
	}
	rows, err := s.q(ctx).ListAllActivations(ctx)
	if err != nil {
		return nil, wrap(err, "list activations")
	}
	out := make([]db.WorkflowActivation, 0, len(rows))
	for _, r := range rows {
		if (site == nil || r.SiteID == *site) && allowed(ctx, &r.SiteID) {
			out = append(out, r)
		}
	}
	return out, nil
}

// RuleOwner names the role that owns a rule.
func (s *Service) RuleOwner(ctx context.Context, key string) (name, owner string) {
	d, err := s.q(ctx).GetRuleDefinitionByKey(ctx, key)
	if err != nil {
		return key, ""
	}
	if d.OwnerRoleID != nil {
		if n, err := s.q(ctx).OwnerRoleOfRule(ctx, *d.OwnerRoleID); err == nil {
			owner = n
		}
	}
	return d.Name, owner
}

// SubmitRuleVersion sends a value to its owner. It keeps the actor so the same
// person cannot approve it.
func (s *Service) SubmitRuleVersion(ctx context.Context, p db.ActSubmitRuleVersionParams) (db.RuleVersion, error) {
	before, err := s.GetRuleVersion(ctx, p.ID)
	if err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitRuleVersion(ctx, p)
	if err != nil {
		return before, stateError(err, "submit this rule value")
	}
	if err := s.audit(ctx, "rules.rule_version.submit", "rule_versions", row.ID, nil, before, row); err != nil {
		return row, err
	}
	// The approver needs to see which rule and what value, not just a row.
	def, err := s.q(ctx).GetRuleDefinition(ctx, row.RuleDefinitionID)
	if err != nil {
		return row, wrap(err, "read rule")
	}
	return row, s.startApproval(ctx, "rule_versions", row.ID, nil, map[string]string{"name": def.Key + " = " + string(row.Value)})
}
