// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package service

import (
	"context"

	"github.com/google/uuid"

	"foodies/api/internal/db"
)

var _ uuid.UUID

// ListRuleDefinition returns a page of rule definitions.
func (s *Service) ListRuleDefinition(ctx context.Context, p db.ListRuleDefinitionParams) ([]db.RuleDefinition, error) {
	return wrap2(s.q(ctx).ListRuleDefinition(ctx, p))("list rule definitions")
}

// GetRuleDefinition returns one row of rule definitions.
func (s *Service) GetRuleDefinition(ctx context.Context, id uuid.UUID) (db.RuleDefinition, error) {
	row, err := s.q(ctx).GetRuleDefinition(ctx, id)
	if err != nil {
		return row, notFound(err, "rule definitions")
	}
	return row, nil
}

// UpdateRuleDefinition changes one row of rule definitions.
func (s *Service) UpdateRuleDefinition(ctx context.Context, p db.UpdateRuleDefinitionParams) (db.RuleDefinition, error) {
	before, err := s.GetRuleDefinition(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateRuleDefinition(ctx, p)
	if err != nil {
		return before, wrap(err, "update rule definitions")
	}
	if err := s.voidSignatures(ctx, "rule_definitions", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "rules.rule_definition.write", "rule_definitions", row.ID, nil, before, row)
}

// ListRuleVersion returns a page of rule versions.
func (s *Service) ListRuleVersion(ctx context.Context, p db.ListRuleVersionParams) ([]db.RuleVersion, error) {
	return wrap2(s.q(ctx).ListRuleVersion(ctx, p))("list rule versions")
}

// GetRuleVersion returns one row of rule versions.
func (s *Service) GetRuleVersion(ctx context.Context, id uuid.UUID) (db.RuleVersion, error) {
	row, err := s.q(ctx).GetRuleVersion(ctx, id)
	if err != nil {
		return row, notFound(err, "rule versions")
	}
	return row, nil
}

// CreateRuleVersion adds one row to rule versions.
func (s *Service) CreateRuleVersion(ctx context.Context, p db.CreateRuleVersionParams) (db.RuleVersion, error) {
	var zero db.RuleVersion
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateRuleVersion(ctx, p)
	if err != nil {
		return zero, wrap(err, "create rule versions")
	}
	return row, s.audit(ctx, "rules.rule_version.write", "rule_versions", row.ID, nil, nil, row)
}

// ListWorkflowActivation returns a page of workflow activations.
func (s *Service) ListWorkflowActivation(ctx context.Context, p db.ListWorkflowActivationParams) ([]db.WorkflowActivation, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListWorkflowActivation(ctx, p))("list workflow activations")
}

// GetWorkflowActivation returns one row of workflow activations.
func (s *Service) GetWorkflowActivation(ctx context.Context, id uuid.UUID) (db.WorkflowActivation, error) {
	row, err := s.q(ctx).GetWorkflowActivation(ctx, id)
	if err != nil {
		return row, notFound(err, "workflow activations")
	}
	return row, inScope(ctx, &row.SiteID)
}

// SuspendWorkflowActivation is the suspend the workflow at the site action on workflow activations.
func (s *Service) SuspendWorkflowActivation(ctx context.Context, p db.ActSuspendWorkflowActivationParams) (db.WorkflowActivation, error) {
	before, err := s.GetWorkflowActivation(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "rules.workflow_activation.suspend", "workflow_activations", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSuspendWorkflowActivation(ctx, p)
	if err != nil {
		return before, stateError(err, "suspend workflow activations")
	}
	if err := s.audit(ctx, "rules.workflow_activation.suspend", "workflow_activations", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, nil
}

// ListApprovalPolicy returns a page of approval policies.
func (s *Service) ListApprovalPolicy(ctx context.Context, p db.ListApprovalPolicyParams) ([]db.ApprovalPolicy, error) {
	return wrap2(s.q(ctx).ListApprovalPolicy(ctx, p))("list approval policies")
}

// GetApprovalPolicy returns one row of approval policies.
func (s *Service) GetApprovalPolicy(ctx context.Context, id uuid.UUID) (db.ApprovalPolicy, error) {
	row, err := s.q(ctx).GetApprovalPolicy(ctx, id)
	if err != nil {
		return row, notFound(err, "approval policies")
	}
	return row, nil
}

// CreateApprovalPolicy adds one row to approval policies.
func (s *Service) CreateApprovalPolicy(ctx context.Context, p db.CreateApprovalPolicyParams) (db.ApprovalPolicy, error) {
	var zero db.ApprovalPolicy
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateApprovalPolicy(ctx, p)
	if err != nil {
		return zero, wrap(err, "create approval policies")
	}
	return row, s.audit(ctx, "approvals.approval_policy.write", "approval_policies", row.ID, nil, nil, row)
}

// UpdateApprovalPolicy changes one row of approval policies.
func (s *Service) UpdateApprovalPolicy(ctx context.Context, p db.UpdateApprovalPolicyParams) (db.ApprovalPolicy, error) {
	before, err := s.GetApprovalPolicy(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateApprovalPolicy(ctx, p)
	if err != nil {
		return before, wrap(err, "update approval policies")
	}
	if err := s.voidSignatures(ctx, "approval_policies", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "approvals.approval_policy.write", "approval_policies", row.ID, nil, before, row)
}

// SubmitApprovalPolicy is the submit for approval action on approval policies.
func (s *Service) SubmitApprovalPolicy(ctx context.Context, p db.ActSubmitApprovalPolicyParams) (db.ApprovalPolicy, error) {
	before, err := s.GetApprovalPolicy(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "approvals.approval_policy.submit", "approval_policies", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitApprovalPolicy(ctx, p)
	if err != nil {
		return before, stateError(err, "submit approval policies")
	}
	if err := s.audit(ctx, "approvals.approval_policy.submit", "approval_policies", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.startApproval(ctx, "approval_policies", row.ID, nil, row)
}

// ApprovalPolicyChildren returns the rows of a record's child tables, by API field.
func (s *Service) ApprovalPolicyChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListApprovalPolicySteps(ctx, id)
		if err != nil {
			return nil, wrap(err, "list approval_policy_steps")
		}
		out["steps"] = childRows(rows)
	}
	return out, nil
}

// ReplaceApprovalPolicySteps swaps a record's approval policy steps for the given rows.
func (s *Service) ReplaceApprovalPolicySteps(ctx context.Context, id uuid.UUID, rows []db.InsertApprovalPolicyStepsParams) error {
	if err := s.q(ctx).DeleteApprovalPolicySteps(ctx, id); err != nil {
		return wrap(err, "clear approval_policy_steps")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertApprovalPolicySteps(ctx, row); err != nil {
			return wrap(err, "add approval_policy_steps")
		}
	}
	return nil
}

// ListApprovalRequest returns a page of approval requests.
func (s *Service) ListApprovalRequest(ctx context.Context, p db.ListApprovalRequestParams) ([]db.ApprovalRequest, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListApprovalRequest(ctx, p))("list approval requests")
}

// GetApprovalRequest returns one row of approval requests.
func (s *Service) GetApprovalRequest(ctx context.Context, id uuid.UUID) (db.ApprovalRequest, error) {
	row, err := s.q(ctx).GetApprovalRequest(ctx, id)
	if err != nil {
		return row, notFound(err, "approval requests")
	}
	return row, inScope(ctx, row.SiteID)
}

// ApprovalRequestChildren returns the rows of a record's child tables, by API field.
func (s *Service) ApprovalRequestChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListApprovalRequestSteps(ctx, id)
		if err != nil {
			return nil, wrap(err, "list approval_request_steps")
		}
		out["steps"] = childRows(rows)
	}
	return out, nil
}

// ReplaceApprovalRequestSteps swaps a record's approval request steps for the given rows.
func (s *Service) ReplaceApprovalRequestSteps(ctx context.Context, id uuid.UUID, rows []db.InsertApprovalRequestStepsParams) error {
	if err := s.q(ctx).DeleteApprovalRequestSteps(ctx, id); err != nil {
		return wrap(err, "clear approval_request_steps")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertApprovalRequestSteps(ctx, row); err != nil {
			return wrap(err, "add approval_request_steps")
		}
	}
	return nil
}

// ListSignatureRequest returns a page of signature requests.
func (s *Service) ListSignatureRequest(ctx context.Context, p db.ListSignatureRequestParams) ([]db.SignatureRequest, error) {
	return wrap2(s.q(ctx).ListSignatureRequest(ctx, p))("list signature requests")
}

// GetSignatureRequest returns one row of signature requests.
func (s *Service) GetSignatureRequest(ctx context.Context, id uuid.UUID) (db.SignatureRequest, error) {
	row, err := s.q(ctx).GetSignatureRequest(ctx, id)
	if err != nil {
		return row, notFound(err, "signature requests")
	}
	return row, nil
}

// ListAttachment returns a page of attachments.
func (s *Service) ListAttachment(ctx context.Context, p db.ListAttachmentParams) ([]db.Attachment, error) {
	return wrap2(s.q(ctx).ListAttachment(ctx, p))("list attachments")
}

// GetAttachment returns one row of attachments.
func (s *Service) GetAttachment(ctx context.Context, id uuid.UUID) (db.Attachment, error) {
	row, err := s.q(ctx).GetAttachment(ctx, id)
	if err != nil {
		return row, notFound(err, "attachments")
	}
	return row, nil
}

// ListComment returns a page of comments.
func (s *Service) ListComment(ctx context.Context, p db.ListCommentParams) ([]db.Comment, error) {
	return wrap2(s.q(ctx).ListComment(ctx, p))("list comments")
}

// GetComment returns one row of comments.
func (s *Service) GetComment(ctx context.Context, id uuid.UUID) (db.Comment, error) {
	row, err := s.q(ctx).GetComment(ctx, id)
	if err != nil {
		return row, notFound(err, "comments")
	}
	return row, nil
}

// CreateComment adds one row to comments.
func (s *Service) CreateComment(ctx context.Context, p db.CreateCommentParams) (db.Comment, error) {
	var zero db.Comment
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateComment(ctx, p)
	if err != nil {
		return zero, wrap(err, "create comments")
	}
	return row, s.audit(ctx, "evidence.comment.write", "comments", row.ID, nil, nil, row)
}

// ListRecordLink returns a page of record links.
func (s *Service) ListRecordLink(ctx context.Context, p db.ListRecordLinkParams) ([]db.RecordLink, error) {
	return wrap2(s.q(ctx).ListRecordLink(ctx, p))("list record links")
}

// GetRecordLink returns one row of record links.
func (s *Service) GetRecordLink(ctx context.Context, id uuid.UUID) (db.RecordLink, error) {
	row, err := s.q(ctx).GetRecordLink(ctx, id)
	if err != nil {
		return row, notFound(err, "record links")
	}
	return row, nil
}

// CreateRecordLink adds one row to record links.
func (s *Service) CreateRecordLink(ctx context.Context, p db.CreateRecordLinkParams) (db.RecordLink, error) {
	var zero db.RecordLink
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateRecordLink(ctx, p)
	if err != nil {
		return zero, wrap(err, "create record links")
	}
	return row, s.audit(ctx, "evidence.record_link.write", "record_links", row.ID, nil, nil, row)
}

// DeleteRecordLink removes one row of record links.
func (s *Service) DeleteRecordLink(ctx context.Context, id uuid.UUID) error {
	before, err := s.GetRecordLink(ctx, id)
	if err != nil {
		return err
	}
	if err := s.q(ctx).DeleteRecordLink(ctx, id); err != nil {
		return wrap(err, "delete record links")
	}
	return s.audit(ctx, "record_links.delete", "record_links", id, nil, before, nil)
}
