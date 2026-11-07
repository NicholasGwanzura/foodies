// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package httpapi

import (
	"context"

	"foodies/api/internal/db"
	"foodies/api/internal/http/apigen"
)

var _ db.Querier

// ListRuleDefinitions implements the listRuleDefinitions operation.
func (s *Server) ListRuleDefinitions(ctx context.Context, req apigen.ListRuleDefinitionsRequestObject) (apigen.ListRuleDefinitionsResponseObject, error) {
	return listPage[apigen.ListRuleDefinitions200JSONResponse](ctx, req.Params, nil, s.service.ListRuleDefinition)
}

// GetRuleDefinition implements the getRuleDefinition operation.
func (s *Server) GetRuleDefinition(ctx context.Context, req apigen.GetRuleDefinitionRequestObject) (apigen.GetRuleDefinitionResponseObject, error) {
	row, err := s.service.GetRuleDefinition(ctx, req.RuleDefinitionId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RuleDefinition](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetRuleDefinition200JSONResponse{Body: body, Headers: apigen.GetRuleDefinition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateRuleDefinition implements the updateRuleDefinition operation.
func (s *Server) UpdateRuleDefinition(ctx context.Context, req apigen.UpdateRuleDefinitionRequestObject) (apigen.UpdateRuleDefinitionResponseObject, error) {
	in, err := input[db.UpdateRuleDefinitionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.RuleDefinitionId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateRuleDefinition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RuleDefinition](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateRuleDefinition200JSONResponse{Body: body, Headers: apigen.UpdateRuleDefinition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListRuleVersions implements the listRuleVersions operation.
func (s *Server) ListRuleVersions(ctx context.Context, req apigen.ListRuleVersionsRequestObject) (apigen.ListRuleVersionsResponseObject, error) {
	return listPage[apigen.ListRuleVersions200JSONResponse](ctx, req.Params, map[string]any{"rule_definition_id": req.RuleDefinitionId}, s.service.ListRuleVersion)
}

// GetRuleVersion implements the getRuleVersion operation.
func (s *Server) GetRuleVersion(ctx context.Context, req apigen.GetRuleVersionRequestObject) (apigen.GetRuleVersionResponseObject, error) {
	row, err := s.service.GetRuleVersion(ctx, req.RuleVersionId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RuleVersion](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetRuleVersion200JSONResponse{Body: body, Headers: apigen.GetRuleVersion200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateRuleVersion implements the createRuleVersion operation.
func (s *Server) CreateRuleVersion(ctx context.Context, req apigen.CreateRuleVersionRequestObject) (apigen.CreateRuleVersionResponseObject, error) {
	in, err := input[db.CreateRuleVersionParams](req.Body, map[string]any{"rule_definition_id": req.RuleDefinitionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateRuleVersion(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RuleVersion](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateRuleVersion201JSONResponse{Body: body, Headers: apigen.CreateRuleVersion201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitRuleVersion implements the submitRuleVersion operation.
func (s *Server) SubmitRuleVersion(ctx context.Context, req apigen.SubmitRuleVersionRequestObject) (apigen.SubmitRuleVersionResponseObject, error) {
	in, err := input[db.ActSubmitRuleVersionParams](struct{}{}, map[string]any{"id": req.RuleVersionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitRuleVersion(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RuleVersion](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitRuleVersion200JSONResponse{Body: body, Headers: apigen.SubmitRuleVersion200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListWorkflowActivations implements the listWorkflowActivations operation.
func (s *Server) ListWorkflowActivations(ctx context.Context, req apigen.ListWorkflowActivationsRequestObject) (apigen.ListWorkflowActivationsResponseObject, error) {
	return listPage[apigen.ListWorkflowActivations200JSONResponse](ctx, req.Params, nil, s.service.ListWorkflowActivation)
}

// GetWorkflowActivation implements the getWorkflowActivation operation.
func (s *Server) GetWorkflowActivation(ctx context.Context, req apigen.GetWorkflowActivationRequestObject) (apigen.GetWorkflowActivationResponseObject, error) {
	row, err := s.service.GetWorkflowActivation(ctx, req.WorkflowActivationId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WorkflowActivation](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetWorkflowActivation200JSONResponse{Body: body, Headers: apigen.GetWorkflowActivation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ActivateWorkflowActivation implements the activateWorkflowActivation operation.
func (s *Server) ActivateWorkflowActivation(ctx context.Context, req apigen.ActivateWorkflowActivationRequestObject) (apigen.ActivateWorkflowActivationResponseObject, error) {
	in, err := input[db.ActActivateWorkflowActivationParams](struct{}{}, map[string]any{"id": req.WorkflowActivationId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ActivateWorkflowActivation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WorkflowActivation](row)
	if err != nil {
		return nil, err
	}
	return apigen.ActivateWorkflowActivation200JSONResponse{Body: body, Headers: apigen.ActivateWorkflowActivation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SuspendWorkflowActivation implements the suspendWorkflowActivation operation.
func (s *Server) SuspendWorkflowActivation(ctx context.Context, req apigen.SuspendWorkflowActivationRequestObject) (apigen.SuspendWorkflowActivationResponseObject, error) {
	in, err := input[db.ActSuspendWorkflowActivationParams](req.Body, map[string]any{"id": req.WorkflowActivationId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SuspendWorkflowActivation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WorkflowActivation](row)
	if err != nil {
		return nil, err
	}
	return apigen.SuspendWorkflowActivation200JSONResponse{Body: body, Headers: apigen.SuspendWorkflowActivation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListApprovalPolicies implements the listApprovalPolicies operation.
func (s *Server) ListApprovalPolicies(ctx context.Context, req apigen.ListApprovalPoliciesRequestObject) (apigen.ListApprovalPoliciesResponseObject, error) {
	return listPage[apigen.ListApprovalPolicies200JSONResponse](ctx, req.Params, nil, s.service.ListApprovalPolicy)
}

// GetApprovalPolicy implements the getApprovalPolicy operation.
func (s *Server) GetApprovalPolicy(ctx context.Context, req apigen.GetApprovalPolicyRequestObject) (apigen.GetApprovalPolicyResponseObject, error) {
	row, err := s.service.GetApprovalPolicy(ctx, req.ApprovalPolicyId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.ApprovalPolicyChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ApprovalPolicy](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetApprovalPolicy200JSONResponse{Body: body, Headers: apigen.GetApprovalPolicy200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateApprovalPolicy implements the createApprovalPolicy operation.
func (s *Server) CreateApprovalPolicy(ctx context.Context, req apigen.CreateApprovalPolicyRequestObject) (apigen.CreateApprovalPolicyResponseObject, error) {
	in, err := input[db.CreateApprovalPolicyParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateApprovalPolicy(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["steps"]; ok {
		rows, err := childInputs[db.InsertApprovalPolicyStepsParams](items, "approval_policy_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceApprovalPolicySteps(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.ApprovalPolicyChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ApprovalPolicy](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateApprovalPolicy201JSONResponse{Body: body, Headers: apigen.CreateApprovalPolicy201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateApprovalPolicy implements the updateApprovalPolicy operation.
func (s *Server) UpdateApprovalPolicy(ctx context.Context, req apigen.UpdateApprovalPolicyRequestObject) (apigen.UpdateApprovalPolicyResponseObject, error) {
	in, err := input[db.UpdateApprovalPolicyParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.ApprovalPolicyId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateApprovalPolicy(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["steps"]; ok {
		rows, err := childInputs[db.InsertApprovalPolicyStepsParams](items, "approval_policy_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceApprovalPolicySteps(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.ApprovalPolicyChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ApprovalPolicy](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateApprovalPolicy200JSONResponse{Body: body, Headers: apigen.UpdateApprovalPolicy200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitApprovalPolicy implements the submitApprovalPolicy operation.
func (s *Server) SubmitApprovalPolicy(ctx context.Context, req apigen.SubmitApprovalPolicyRequestObject) (apigen.SubmitApprovalPolicyResponseObject, error) {
	in, err := input[db.ActSubmitApprovalPolicyParams](struct{}{}, map[string]any{"id": req.ApprovalPolicyId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitApprovalPolicy(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ApprovalPolicy](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitApprovalPolicy200JSONResponse{Body: body, Headers: apigen.SubmitApprovalPolicy200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// GetApprovalRequest implements the getApprovalRequest operation.
func (s *Server) GetApprovalRequest(ctx context.Context, req apigen.GetApprovalRequestRequestObject) (apigen.GetApprovalRequestResponseObject, error) {
	row, err := s.service.GetApprovalRequest(ctx, req.ApprovalRequestId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.ApprovalRequestChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ApprovalRequest](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetApprovalRequest200JSONResponse{Body: body, Headers: apigen.GetApprovalRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListSignatureRequests implements the listSignatureRequests operation.
func (s *Server) ListSignatureRequests(ctx context.Context, req apigen.ListSignatureRequestsRequestObject) (apigen.ListSignatureRequestsResponseObject, error) {
	return listPage[apigen.ListSignatureRequests200JSONResponse](ctx, req.Params, nil, s.service.ListSignatureRequest)
}

// GetSignatureRequest implements the getSignatureRequest operation.
func (s *Server) GetSignatureRequest(ctx context.Context, req apigen.GetSignatureRequestRequestObject) (apigen.GetSignatureRequestResponseObject, error) {
	row, err := s.service.GetSignatureRequest(ctx, req.SignatureRequestId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SignatureRequest](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetSignatureRequest200JSONResponse{Body: body, Headers: apigen.GetSignatureRequest200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListAttachments implements the listAttachments operation.
func (s *Server) ListAttachments(ctx context.Context, req apigen.ListAttachmentsRequestObject) (apigen.ListAttachmentsResponseObject, error) {
	return listPage[apigen.ListAttachments200JSONResponse](ctx, req.Params, nil, s.service.ListAttachment)
}

// CreateAttachment implements the createAttachment operation.
func (s *Server) CreateAttachment(ctx context.Context, req apigen.CreateAttachmentRequestObject) (apigen.CreateAttachmentResponseObject, error) {
	in, err := input[db.CreateAttachmentParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateAttachment(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Attachment](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateAttachment201JSONResponse{Body: body, Headers: apigen.CreateAttachment201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CompleteAttachment implements the completeAttachment operation.
func (s *Server) CompleteAttachment(ctx context.Context, req apigen.CompleteAttachmentRequestObject) (apigen.CompleteAttachmentResponseObject, error) {
	in, err := input[db.ActCompleteAttachmentParams](struct{}{}, map[string]any{"id": req.AttachmentId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CompleteAttachment(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Attachment](row)
	if err != nil {
		return nil, err
	}
	return apigen.CompleteAttachment200JSONResponse{Body: body, Headers: apigen.CompleteAttachment200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListComments implements the listComments operation.
func (s *Server) ListComments(ctx context.Context, req apigen.ListCommentsRequestObject) (apigen.ListCommentsResponseObject, error) {
	return listPage[apigen.ListComments200JSONResponse](ctx, req.Params, nil, s.service.ListComment)
}

// CreateComment implements the createComment operation.
func (s *Server) CreateComment(ctx context.Context, req apigen.CreateCommentRequestObject) (apigen.CreateCommentResponseObject, error) {
	in, err := input[db.CreateCommentParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateComment(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Comment](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateComment201JSONResponse{Body: body, Headers: apigen.CreateComment201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListRecordLinks implements the listRecordLinks operation.
func (s *Server) ListRecordLinks(ctx context.Context, req apigen.ListRecordLinksRequestObject) (apigen.ListRecordLinksResponseObject, error) {
	return listPage[apigen.ListRecordLinks200JSONResponse](ctx, req.Params, nil, s.service.ListRecordLink)
}

// CreateRecordLink implements the createRecordLink operation.
func (s *Server) CreateRecordLink(ctx context.Context, req apigen.CreateRecordLinkRequestObject) (apigen.CreateRecordLinkResponseObject, error) {
	in, err := input[db.CreateRecordLinkParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateRecordLink(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RecordLink](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateRecordLink201JSONResponse{Body: body, Headers: apigen.CreateRecordLink201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// DeleteRecordLink implements the deleteRecordLink operation.
func (s *Server) DeleteRecordLink(ctx context.Context, req apigen.DeleteRecordLinkRequestObject) (apigen.DeleteRecordLinkResponseObject, error) {
	if err := s.service.DeleteRecordLink(ctx, req.RecordLinkId); err != nil {
		return nil, err
	}
	return apigen.DeleteRecordLink204Response{}, nil
}
