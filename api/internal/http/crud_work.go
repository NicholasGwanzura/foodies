// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package httpapi

import (
	"context"

	"foodies/api/internal/db"
	"foodies/api/internal/http/apigen"
)

var _ db.Querier

// ListSchedules implements the listSchedules operation.
func (s *Server) ListSchedules(ctx context.Context, req apigen.ListSchedulesRequestObject) (apigen.ListSchedulesResponseObject, error) {
	return listPage[apigen.ListSchedules200JSONResponse](ctx, req.Params, nil, s.service.ListSchedule)
}

// GetSchedule implements the getSchedule operation.
func (s *Server) GetSchedule(ctx context.Context, req apigen.GetScheduleRequestObject) (apigen.GetScheduleResponseObject, error) {
	row, err := s.service.GetSchedule(ctx, req.ScheduleId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Schedule](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetSchedule200JSONResponse{Body: body, Headers: apigen.GetSchedule200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateSchedule implements the createSchedule operation.
func (s *Server) CreateSchedule(ctx context.Context, req apigen.CreateScheduleRequestObject) (apigen.CreateScheduleResponseObject, error) {
	in, err := input[db.CreateScheduleParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateSchedule(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Schedule](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateSchedule201JSONResponse{Body: body, Headers: apigen.CreateSchedule201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateSchedule implements the updateSchedule operation.
func (s *Server) UpdateSchedule(ctx context.Context, req apigen.UpdateScheduleRequestObject) (apigen.UpdateScheduleResponseObject, error) {
	in, err := input[db.UpdateScheduleParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.ScheduleId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateSchedule(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Schedule](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateSchedule200JSONResponse{Body: body, Headers: apigen.UpdateSchedule200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// PauseSchedule implements the pauseSchedule operation.
func (s *Server) PauseSchedule(ctx context.Context, req apigen.PauseScheduleRequestObject) (apigen.PauseScheduleResponseObject, error) {
	in, err := input[db.ActPauseScheduleParams](struct{}{}, map[string]any{"id": req.ScheduleId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.PauseSchedule(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Schedule](row)
	if err != nil {
		return nil, err
	}
	return apigen.PauseSchedule200JSONResponse{Body: body, Headers: apigen.PauseSchedule200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ResumeSchedule implements the resumeSchedule operation.
func (s *Server) ResumeSchedule(ctx context.Context, req apigen.ResumeScheduleRequestObject) (apigen.ResumeScheduleResponseObject, error) {
	in, err := input[db.ActResumeScheduleParams](struct{}{}, map[string]any{"id": req.ScheduleId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ResumeSchedule(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Schedule](row)
	if err != nil {
		return nil, err
	}
	return apigen.ResumeSchedule200JSONResponse{Body: body, Headers: apigen.ResumeSchedule200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// GetTask implements the getTask operation.
func (s *Server) GetTask(ctx context.Context, req apigen.GetTaskRequestObject) (apigen.GetTaskResponseObject, error) {
	row, err := s.service.GetTask(ctx, req.TaskId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Task](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetTask200JSONResponse{Body: body, Headers: apigen.GetTask200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateTask implements the updateTask operation.
func (s *Server) UpdateTask(ctx context.Context, req apigen.UpdateTaskRequestObject) (apigen.UpdateTaskResponseObject, error) {
	in, err := input[db.UpdateTaskParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.TaskId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateTask(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Task](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateTask200JSONResponse{Body: body, Headers: apigen.UpdateTask200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListFormTemplates implements the listFormTemplates operation.
func (s *Server) ListFormTemplates(ctx context.Context, req apigen.ListFormTemplatesRequestObject) (apigen.ListFormTemplatesResponseObject, error) {
	return listPage[apigen.ListFormTemplates200JSONResponse](ctx, req.Params, nil, s.service.ListFormTemplate)
}

// GetFormTemplate implements the getFormTemplate operation.
func (s *Server) GetFormTemplate(ctx context.Context, req apigen.GetFormTemplateRequestObject) (apigen.GetFormTemplateResponseObject, error) {
	row, err := s.service.GetFormTemplate(ctx, req.FormTemplateId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FormTemplate](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetFormTemplate200JSONResponse{Body: body, Headers: apigen.GetFormTemplate200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateFormTemplate implements the createFormTemplate operation.
func (s *Server) CreateFormTemplate(ctx context.Context, req apigen.CreateFormTemplateRequestObject) (apigen.CreateFormTemplateResponseObject, error) {
	in, err := input[db.CreateFormTemplateParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateFormTemplate(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FormTemplate](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateFormTemplate201JSONResponse{Body: body, Headers: apigen.CreateFormTemplate201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateFormTemplate implements the updateFormTemplate operation.
func (s *Server) UpdateFormTemplate(ctx context.Context, req apigen.UpdateFormTemplateRequestObject) (apigen.UpdateFormTemplateResponseObject, error) {
	in, err := input[db.UpdateFormTemplateParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.FormTemplateId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateFormTemplate(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FormTemplate](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateFormTemplate200JSONResponse{Body: body, Headers: apigen.UpdateFormTemplate200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListFormTemplateVersions implements the listFormTemplateVersions operation.
func (s *Server) ListFormTemplateVersions(ctx context.Context, req apigen.ListFormTemplateVersionsRequestObject) (apigen.ListFormTemplateVersionsResponseObject, error) {
	return listPage[apigen.ListFormTemplateVersions200JSONResponse](ctx, req.Params, map[string]any{"form_template_id": req.FormTemplateId}, s.service.ListFormTemplateVersion)
}

// GetFormTemplateVersion implements the getFormTemplateVersion operation.
func (s *Server) GetFormTemplateVersion(ctx context.Context, req apigen.GetFormTemplateVersionRequestObject) (apigen.GetFormTemplateVersionResponseObject, error) {
	row, err := s.service.GetFormTemplateVersion(ctx, req.FormTemplateVersionId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FormTemplateVersion](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetFormTemplateVersion200JSONResponse{Body: body, Headers: apigen.GetFormTemplateVersion200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateFormTemplateVersion implements the createFormTemplateVersion operation.
func (s *Server) CreateFormTemplateVersion(ctx context.Context, req apigen.CreateFormTemplateVersionRequestObject) (apigen.CreateFormTemplateVersionResponseObject, error) {
	in, err := input[db.CreateFormTemplateVersionParams](req.Body, map[string]any{"form_template_id": req.FormTemplateId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateFormTemplateVersion(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FormTemplateVersion](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateFormTemplateVersion201JSONResponse{Body: body, Headers: apigen.CreateFormTemplateVersion201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateFormTemplateVersion implements the updateFormTemplateVersion operation.
func (s *Server) UpdateFormTemplateVersion(ctx context.Context, req apigen.UpdateFormTemplateVersionRequestObject) (apigen.UpdateFormTemplateVersionResponseObject, error) {
	in, err := input[db.UpdateFormTemplateVersionParams](req.Body, map[string]any{"form_template_id": req.FormTemplateId})
	if err != nil {
		return nil, err
	}
	in.ID = req.FormTemplateVersionId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateFormTemplateVersion(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FormTemplateVersion](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateFormTemplateVersion200JSONResponse{Body: body, Headers: apigen.UpdateFormTemplateVersion200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitFormTemplateVersion implements the submitFormTemplateVersion operation.
func (s *Server) SubmitFormTemplateVersion(ctx context.Context, req apigen.SubmitFormTemplateVersionRequestObject) (apigen.SubmitFormTemplateVersionResponseObject, error) {
	in, err := input[db.ActSubmitFormTemplateVersionParams](struct{}{}, map[string]any{"id": req.FormTemplateVersionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitFormTemplateVersion(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FormTemplateVersion](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitFormTemplateVersion200JSONResponse{Body: body, Headers: apigen.SubmitFormTemplateVersion200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// PublishFormTemplateVersion implements the publishFormTemplateVersion operation.
func (s *Server) PublishFormTemplateVersion(ctx context.Context, req apigen.PublishFormTemplateVersionRequestObject) (apigen.PublishFormTemplateVersionResponseObject, error) {
	in, err := input[db.ActPublishFormTemplateVersionParams](struct{}{}, map[string]any{"id": req.FormTemplateVersionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.PublishFormTemplateVersion(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FormTemplateVersion](row)
	if err != nil {
		return nil, err
	}
	return apigen.PublishFormTemplateVersion200JSONResponse{Body: body, Headers: apigen.PublishFormTemplateVersion200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// RetireFormTemplateVersion implements the retireFormTemplateVersion operation.
func (s *Server) RetireFormTemplateVersion(ctx context.Context, req apigen.RetireFormTemplateVersionRequestObject) (apigen.RetireFormTemplateVersionResponseObject, error) {
	in, err := input[db.ActRetireFormTemplateVersionParams](struct{}{}, map[string]any{"id": req.FormTemplateVersionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.RetireFormTemplateVersion(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FormTemplateVersion](row)
	if err != nil {
		return nil, err
	}
	return apigen.RetireFormTemplateVersion200JSONResponse{Body: body, Headers: apigen.RetireFormTemplateVersion200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListFormSubmissions implements the listFormSubmissions operation.
func (s *Server) ListFormSubmissions(ctx context.Context, req apigen.ListFormSubmissionsRequestObject) (apigen.ListFormSubmissionsResponseObject, error) {
	return listPage[apigen.ListFormSubmissions200JSONResponse](ctx, req.Params, nil, s.service.ListFormSubmission)
}

// GetFormSubmission implements the getFormSubmission operation.
func (s *Server) GetFormSubmission(ctx context.Context, req apigen.GetFormSubmissionRequestObject) (apigen.GetFormSubmissionResponseObject, error) {
	row, err := s.service.GetFormSubmission(ctx, req.FormSubmissionId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FormSubmission](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetFormSubmission200JSONResponse{Body: body, Headers: apigen.GetFormSubmission200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateFormSubmission implements the createFormSubmission operation.
func (s *Server) CreateFormSubmission(ctx context.Context, req apigen.CreateFormSubmissionRequestObject) (apigen.CreateFormSubmissionResponseObject, error) {
	in, err := input[db.CreateFormSubmissionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateFormSubmission(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FormSubmission](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateFormSubmission201JSONResponse{Body: body, Headers: apigen.CreateFormSubmission201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateFormSubmission implements the updateFormSubmission operation.
func (s *Server) UpdateFormSubmission(ctx context.Context, req apigen.UpdateFormSubmissionRequestObject) (apigen.UpdateFormSubmissionResponseObject, error) {
	in, err := input[db.UpdateFormSubmissionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.FormSubmissionId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateFormSubmission(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FormSubmission](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateFormSubmission200JSONResponse{Body: body, Headers: apigen.UpdateFormSubmission200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitFormSubmission implements the submitFormSubmission operation.
func (s *Server) SubmitFormSubmission(ctx context.Context, req apigen.SubmitFormSubmissionRequestObject) (apigen.SubmitFormSubmissionResponseObject, error) {
	in, err := input[db.ActSubmitFormSubmissionParams](struct{}{}, map[string]any{"id": req.FormSubmissionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitFormSubmission(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FormSubmission](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitFormSubmission200JSONResponse{Body: body, Headers: apigen.SubmitFormSubmission200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReviewFormSubmission implements the reviewFormSubmission operation.
func (s *Server) ReviewFormSubmission(ctx context.Context, req apigen.ReviewFormSubmissionRequestObject) (apigen.ReviewFormSubmissionResponseObject, error) {
	in, err := input[db.ActReviewFormSubmissionParams](req.Body, map[string]any{"id": req.FormSubmissionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReviewFormSubmission(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FormSubmission](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReviewFormSubmission200JSONResponse{Body: body, Headers: apigen.ReviewFormSubmission200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReturnFormSubmission implements the returnFormSubmission operation.
func (s *Server) ReturnFormSubmission(ctx context.Context, req apigen.ReturnFormSubmissionRequestObject) (apigen.ReturnFormSubmissionResponseObject, error) {
	in, err := input[db.ActReturnFormSubmissionParams](req.Body, map[string]any{"id": req.FormSubmissionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReturnFormSubmission(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.FormSubmission](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReturnFormSubmission200JSONResponse{Body: body, Headers: apigen.ReturnFormSubmission200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListEscalationPolicies implements the listEscalationPolicies operation.
func (s *Server) ListEscalationPolicies(ctx context.Context, req apigen.ListEscalationPoliciesRequestObject) (apigen.ListEscalationPoliciesResponseObject, error) {
	return listPage[apigen.ListEscalationPolicies200JSONResponse](ctx, req.Params, nil, s.service.ListEscalationPolicy)
}

// GetEscalationPolicy implements the getEscalationPolicy operation.
func (s *Server) GetEscalationPolicy(ctx context.Context, req apigen.GetEscalationPolicyRequestObject) (apigen.GetEscalationPolicyResponseObject, error) {
	row, err := s.service.GetEscalationPolicy(ctx, req.EscalationPolicyId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.EscalationPolicyChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.EscalationPolicy](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetEscalationPolicy200JSONResponse{Body: body, Headers: apigen.GetEscalationPolicy200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateEscalationPolicy implements the createEscalationPolicy operation.
func (s *Server) CreateEscalationPolicy(ctx context.Context, req apigen.CreateEscalationPolicyRequestObject) (apigen.CreateEscalationPolicyResponseObject, error) {
	in, err := input[db.CreateEscalationPolicyParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateEscalationPolicy(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["steps"]; ok {
		rows, err := childInputs[db.InsertEscalationPolicyStepsParams](items, "escalation_policy_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceEscalationPolicySteps(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.EscalationPolicyChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.EscalationPolicy](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.CreateEscalationPolicy201JSONResponse{Body: body, Headers: apigen.CreateEscalationPolicy201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateEscalationPolicy implements the updateEscalationPolicy operation.
func (s *Server) UpdateEscalationPolicy(ctx context.Context, req apigen.UpdateEscalationPolicyRequestObject) (apigen.UpdateEscalationPolicyResponseObject, error) {
	in, err := input[db.UpdateEscalationPolicyParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.EscalationPolicyId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateEscalationPolicy(ctx, in)
	if err != nil {
		return nil, err
	}
	raw, err := toMap(req.Body)
	if err != nil {
		return nil, err
	}
	if items, ok := raw["steps"]; ok {
		rows, err := childInputs[db.InsertEscalationPolicyStepsParams](items, "escalation_policy_id", row.ID)
		if err != nil {
			return nil, err
		}
		if err := s.service.ReplaceEscalationPolicySteps(ctx, row.ID, rows); err != nil {
			return nil, err
		}
	}
	kids, err := s.service.EscalationPolicyChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.EscalationPolicy](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateEscalationPolicy200JSONResponse{Body: body, Headers: apigen.UpdateEscalationPolicy200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListEscalations implements the listEscalations operation.
func (s *Server) ListEscalations(ctx context.Context, req apigen.ListEscalationsRequestObject) (apigen.ListEscalationsResponseObject, error) {
	return listPage[apigen.ListEscalations200JSONResponse](ctx, req.Params, nil, s.service.ListEscalation)
}

// GetEscalation implements the getEscalation operation.
func (s *Server) GetEscalation(ctx context.Context, req apigen.GetEscalationRequestObject) (apigen.GetEscalationResponseObject, error) {
	row, err := s.service.GetEscalation(ctx, req.EscalationId)
	if err != nil {
		return nil, err
	}
	kids, err := s.service.EscalationChildren(ctx, row.ID)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Escalation](row, kids)
	if err != nil {
		return nil, err
	}
	return apigen.GetEscalation200JSONResponse{Body: body, Headers: apigen.GetEscalation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// AcknowledgeEscalation implements the acknowledgeEscalation operation.
func (s *Server) AcknowledgeEscalation(ctx context.Context, req apigen.AcknowledgeEscalationRequestObject) (apigen.AcknowledgeEscalationResponseObject, error) {
	in, err := input[db.ActAcknowledgeEscalationParams](req.Body, map[string]any{"id": req.EscalationId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.AcknowledgeEscalation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Escalation](row)
	if err != nil {
		return nil, err
	}
	return apigen.AcknowledgeEscalation200JSONResponse{Body: body, Headers: apigen.AcknowledgeEscalation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// AdvanceEscalation implements the advanceEscalation operation.
func (s *Server) AdvanceEscalation(ctx context.Context, req apigen.AdvanceEscalationRequestObject) (apigen.AdvanceEscalationResponseObject, error) {
	in, err := input[db.ActAdvanceEscalationParams](req.Body, map[string]any{"id": req.EscalationId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.AdvanceEscalation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Escalation](row)
	if err != nil {
		return nil, err
	}
	return apigen.AdvanceEscalation200JSONResponse{Body: body, Headers: apigen.AdvanceEscalation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListActions implements the listActions operation.
func (s *Server) ListActions(ctx context.Context, req apigen.ListActionsRequestObject) (apigen.ListActionsResponseObject, error) {
	return listPage[apigen.ListActions200JSONResponse](ctx, req.Params, nil, s.service.ListAction)
}

// GetAction implements the getAction operation.
func (s *Server) GetAction(ctx context.Context, req apigen.GetActionRequestObject) (apigen.GetActionResponseObject, error) {
	row, err := s.service.GetAction(ctx, req.ActionId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Action](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetAction200JSONResponse{Body: body, Headers: apigen.GetAction200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateAction implements the createAction operation.
func (s *Server) CreateAction(ctx context.Context, req apigen.CreateActionRequestObject) (apigen.CreateActionResponseObject, error) {
	in, err := input[db.CreateActionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateAction(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Action](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateAction201JSONResponse{Body: body, Headers: apigen.CreateAction201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateAction implements the updateAction operation.
func (s *Server) UpdateAction(ctx context.Context, req apigen.UpdateActionRequestObject) (apigen.UpdateActionResponseObject, error) {
	in, err := input[db.UpdateActionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.ActionId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateAction(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Action](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateAction200JSONResponse{Body: body, Headers: apigen.UpdateAction200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CompleteAction implements the completeAction operation.
func (s *Server) CompleteAction(ctx context.Context, req apigen.CompleteActionRequestObject) (apigen.CompleteActionResponseObject, error) {
	in, err := input[db.ActCompleteActionParams](req.Body, map[string]any{"id": req.ActionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CompleteAction(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Action](row)
	if err != nil {
		return nil, err
	}
	return apigen.CompleteAction200JSONResponse{Body: body, Headers: apigen.CompleteAction200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// VerifyAction implements the verifyAction operation.
func (s *Server) VerifyAction(ctx context.Context, req apigen.VerifyActionRequestObject) (apigen.VerifyActionResponseObject, error) {
	in, err := input[db.ActVerifyActionParams](req.Body, map[string]any{"id": req.ActionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.VerifyAction(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Action](row)
	if err != nil {
		return nil, err
	}
	return apigen.VerifyAction200JSONResponse{Body: body, Headers: apigen.VerifyAction200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReopenAction implements the reopenAction operation.
func (s *Server) ReopenAction(ctx context.Context, req apigen.ReopenActionRequestObject) (apigen.ReopenActionResponseObject, error) {
	in, err := input[db.ActReopenActionParams](req.Body, map[string]any{"id": req.ActionId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReopenAction(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Action](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReopenAction200JSONResponse{Body: body, Headers: apigen.ReopenAction200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}
