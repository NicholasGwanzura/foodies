// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package service

import (
	"context"

	"github.com/google/uuid"

	"foodies/api/internal/db"
)

var _ uuid.UUID

// ListSchedule returns a page of schedules.
func (s *Service) ListSchedule(ctx context.Context, p db.ListScheduleParams) ([]db.Schedule, error) {
	return wrap2(s.q(ctx).ListSchedule(ctx, p))("list schedules")
}

// GetSchedule returns one row of schedules.
func (s *Service) GetSchedule(ctx context.Context, id uuid.UUID) (db.Schedule, error) {
	row, err := s.q(ctx).GetSchedule(ctx, id)
	if err != nil {
		return row, notFound(err, "schedules")
	}
	return row, nil
}

// CreateSchedule adds one row to schedules.
func (s *Service) CreateSchedule(ctx context.Context, p db.CreateScheduleParams) (db.Schedule, error) {
	var zero db.Schedule
	if err := s.guard(ctx, "CreateSchedule", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateSchedule(ctx, p)
	if err != nil {
		return zero, wrap(err, "create schedules")
	}
	if err := s.audit(ctx, "tasks.schedule.write", "schedules", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateSchedule", row)
}

// UpdateSchedule changes one row of schedules.
func (s *Service) UpdateSchedule(ctx context.Context, p db.UpdateScheduleParams) (db.Schedule, error) {
	before, err := s.GetSchedule(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateSchedule(ctx, p)
	if err != nil {
		return before, wrap(err, "update schedules")
	}
	if err := s.voidSignatures(ctx, "schedules", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "tasks.schedule.write", "schedules", row.ID, nil, before, row)
}

// PauseSchedule is the pause generation action on schedules.
func (s *Service) PauseSchedule(ctx context.Context, p db.ActPauseScheduleParams) (db.Schedule, error) {
	before, err := s.GetSchedule(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "tasks.schedule.pause", "schedules", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "PauseSchedule", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActPauseSchedule(ctx, p)
	if err != nil {
		return before, stateError(err, "pause schedules")
	}
	if err := s.audit(ctx, "tasks.schedule.pause", "schedules", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "PauseSchedule", row)
}

// ResumeSchedule is the resume generation action on schedules.
func (s *Service) ResumeSchedule(ctx context.Context, p db.ActResumeScheduleParams) (db.Schedule, error) {
	before, err := s.GetSchedule(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "tasks.schedule.resume", "schedules", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ResumeSchedule", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActResumeSchedule(ctx, p)
	if err != nil {
		return before, stateError(err, "resume schedules")
	}
	if err := s.audit(ctx, "tasks.schedule.resume", "schedules", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ResumeSchedule", row)
}

// ListTask returns a page of tasks.
func (s *Service) ListTask(ctx context.Context, p db.ListTaskParams) ([]db.Task, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListTask(ctx, p))("list tasks")
}

// GetTask returns one row of tasks.
func (s *Service) GetTask(ctx context.Context, id uuid.UUID) (db.Task, error) {
	row, err := s.q(ctx).GetTask(ctx, id)
	if err != nil {
		return row, notFound(err, "tasks")
	}
	return row, inScope(ctx, &row.SiteID)
}

// UpdateTask changes one row of tasks.
func (s *Service) UpdateTask(ctx context.Context, p db.UpdateTaskParams) (db.Task, error) {
	before, err := s.GetTask(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateTask(ctx, p)
	if err != nil {
		return before, wrap(err, "update tasks")
	}
	if err := s.voidSignatures(ctx, "tasks", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "tasks.task.write", "tasks", row.ID, &row.SiteID, before, row)
}

// ListFormTemplate returns a page of form templates.
func (s *Service) ListFormTemplate(ctx context.Context, p db.ListFormTemplateParams) ([]db.FormTemplate, error) {
	return wrap2(s.q(ctx).ListFormTemplate(ctx, p))("list form templates")
}

// GetFormTemplate returns one row of form templates.
func (s *Service) GetFormTemplate(ctx context.Context, id uuid.UUID) (db.FormTemplate, error) {
	row, err := s.q(ctx).GetFormTemplate(ctx, id)
	if err != nil {
		return row, notFound(err, "form templates")
	}
	return row, nil
}

// CreateFormTemplate adds one row to form templates.
func (s *Service) CreateFormTemplate(ctx context.Context, p db.CreateFormTemplateParams) (db.FormTemplate, error) {
	var zero db.FormTemplate
	if err := s.guard(ctx, "CreateFormTemplate", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateFormTemplate(ctx, p)
	if err != nil {
		return zero, wrap(err, "create form templates")
	}
	if err := s.audit(ctx, "forms.form_template.write", "form_templates", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateFormTemplate", row)
}

// UpdateFormTemplate changes one row of form templates.
func (s *Service) UpdateFormTemplate(ctx context.Context, p db.UpdateFormTemplateParams) (db.FormTemplate, error) {
	before, err := s.GetFormTemplate(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateFormTemplate(ctx, p)
	if err != nil {
		return before, wrap(err, "update form templates")
	}
	if err := s.voidSignatures(ctx, "form_templates", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "forms.form_template.write", "form_templates", row.ID, nil, before, row)
}

// ListFormTemplateVersion returns a page of form template versions.
func (s *Service) ListFormTemplateVersion(ctx context.Context, p db.ListFormTemplateVersionParams) ([]db.FormTemplateVersion, error) {
	return wrap2(s.q(ctx).ListFormTemplateVersion(ctx, p))("list form template versions")
}

// GetFormTemplateVersion returns one row of form template versions.
func (s *Service) GetFormTemplateVersion(ctx context.Context, id uuid.UUID) (db.FormTemplateVersion, error) {
	row, err := s.q(ctx).GetFormTemplateVersion(ctx, id)
	if err != nil {
		return row, notFound(err, "form template versions")
	}
	return row, nil
}

// UpdateFormTemplateVersion changes one row of form template versions.
func (s *Service) UpdateFormTemplateVersion(ctx context.Context, p db.UpdateFormTemplateVersionParams) (db.FormTemplateVersion, error) {
	before, err := s.GetFormTemplateVersion(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateFormTemplateVersion(ctx, p)
	if err != nil {
		return before, wrap(err, "update form template versions")
	}
	if err := s.voidSignatures(ctx, "form_template_versions", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "forms.form_template_version.write", "form_template_versions", row.ID, nil, before, row)
}

// SubmitFormTemplateVersion is the submit for approval action on form template versions.
func (s *Service) SubmitFormTemplateVersion(ctx context.Context, p db.ActSubmitFormTemplateVersionParams) (db.FormTemplateVersion, error) {
	before, err := s.GetFormTemplateVersion(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "forms.form_template_version.submit", "form_template_versions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitFormTemplateVersion", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitFormTemplateVersion(ctx, p)
	if err != nil {
		return before, stateError(err, "submit form template versions")
	}
	if err := s.audit(ctx, "forms.form_template_version.submit", "form_template_versions", row.ID, nil, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "form_template_versions", row.ID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitFormTemplateVersion", row)
}

// RetireFormTemplateVersion is the retire action on form template versions.
func (s *Service) RetireFormTemplateVersion(ctx context.Context, p db.ActRetireFormTemplateVersionParams) (db.FormTemplateVersion, error) {
	before, err := s.GetFormTemplateVersion(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "forms.form_template_version.retire", "form_template_versions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "RetireFormTemplateVersion", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActRetireFormTemplateVersion(ctx, p)
	if err != nil {
		return before, stateError(err, "retire form template versions")
	}
	if err := s.audit(ctx, "forms.form_template_version.retire", "form_template_versions", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "RetireFormTemplateVersion", row)
}

// ListFormSubmission returns a page of form submissions.
func (s *Service) ListFormSubmission(ctx context.Context, p db.ListFormSubmissionParams) ([]db.FormSubmission, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListFormSubmission(ctx, p))("list form submissions")
}

// GetFormSubmission returns one row of form submissions.
func (s *Service) GetFormSubmission(ctx context.Context, id uuid.UUID) (db.FormSubmission, error) {
	row, err := s.q(ctx).GetFormSubmission(ctx, id)
	if err != nil {
		return row, notFound(err, "form submissions")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateFormSubmission adds one row to form submissions.
func (s *Service) CreateFormSubmission(ctx context.Context, p db.CreateFormSubmissionParams) (db.FormSubmission, error) {
	var zero db.FormSubmission
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateFormSubmission", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "FRM", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateFormSubmission(ctx, p)
	if err != nil {
		return zero, wrap(err, "create form submissions")
	}
	if err := s.audit(ctx, "forms.form_submission.write", "form_submissions", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateFormSubmission", row)
}

// UpdateFormSubmission changes one row of form submissions.
func (s *Service) UpdateFormSubmission(ctx context.Context, p db.UpdateFormSubmissionParams) (db.FormSubmission, error) {
	before, err := s.GetFormSubmission(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateFormSubmission(ctx, p)
	if err != nil {
		return before, wrap(err, "update form submissions")
	}
	if err := s.voidSignatures(ctx, "form_submissions", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "forms.form_submission.write", "form_submissions", row.ID, &row.SiteID, before, row)
}

// ReturnFormSubmission is the return for correction action on form submissions.
func (s *Service) ReturnFormSubmission(ctx context.Context, p db.ActReturnFormSubmissionParams) (db.FormSubmission, error) {
	before, err := s.GetFormSubmission(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "forms.form_submission.return", "form_submissions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ReturnFormSubmission", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActReturnFormSubmission(ctx, p)
	if err != nil {
		return before, stateError(err, "return form submissions")
	}
	if err := s.audit(ctx, "forms.form_submission.return", "form_submissions", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ReturnFormSubmission", row)
}

// ListEscalationPolicy returns a page of escalation policies.
func (s *Service) ListEscalationPolicy(ctx context.Context, p db.ListEscalationPolicyParams) ([]db.EscalationPolicy, error) {
	return wrap2(s.q(ctx).ListEscalationPolicy(ctx, p))("list escalation policies")
}

// GetEscalationPolicy returns one row of escalation policies.
func (s *Service) GetEscalationPolicy(ctx context.Context, id uuid.UUID) (db.EscalationPolicy, error) {
	row, err := s.q(ctx).GetEscalationPolicy(ctx, id)
	if err != nil {
		return row, notFound(err, "escalation policies")
	}
	return row, nil
}

// CreateEscalationPolicy adds one row to escalation policies.
func (s *Service) CreateEscalationPolicy(ctx context.Context, p db.CreateEscalationPolicyParams) (db.EscalationPolicy, error) {
	var zero db.EscalationPolicy
	if err := s.guard(ctx, "CreateEscalationPolicy", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateEscalationPolicy(ctx, p)
	if err != nil {
		return zero, wrap(err, "create escalation policies")
	}
	if err := s.audit(ctx, "exceptions.escalation_policy.write", "escalation_policies", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateEscalationPolicy", row)
}

// UpdateEscalationPolicy changes one row of escalation policies.
func (s *Service) UpdateEscalationPolicy(ctx context.Context, p db.UpdateEscalationPolicyParams) (db.EscalationPolicy, error) {
	before, err := s.GetEscalationPolicy(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateEscalationPolicy(ctx, p)
	if err != nil {
		return before, wrap(err, "update escalation policies")
	}
	if err := s.voidSignatures(ctx, "escalation_policies", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "exceptions.escalation_policy.write", "escalation_policies", row.ID, nil, before, row)
}

// EscalationPolicyChildren returns the rows of a record's child tables, by API field.
func (s *Service) EscalationPolicyChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListEscalationPolicySteps(ctx, id)
		if err != nil {
			return nil, wrap(err, "list escalation_policy_steps")
		}
		out["steps"] = childRows(rows)
	}
	return out, nil
}

// ReplaceEscalationPolicySteps swaps a record's escalation policy steps for the given rows.
func (s *Service) ReplaceEscalationPolicySteps(ctx context.Context, id uuid.UUID, rows []db.InsertEscalationPolicyStepsParams) error {
	if err := s.q(ctx).DeleteEscalationPolicySteps(ctx, id); err != nil {
		return wrap(err, "clear escalation_policy_steps")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertEscalationPolicySteps(ctx, row); err != nil {
			return wrap(err, "add escalation_policy_steps")
		}
	}
	return nil
}

// ListEscalation returns a page of escalations.
func (s *Service) ListEscalation(ctx context.Context, p db.ListEscalationParams) ([]db.Escalation, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListEscalation(ctx, p))("list escalations")
}

// GetEscalation returns one row of escalations.
func (s *Service) GetEscalation(ctx context.Context, id uuid.UUID) (db.Escalation, error) {
	row, err := s.q(ctx).GetEscalation(ctx, id)
	if err != nil {
		return row, notFound(err, "escalations")
	}
	return row, inScope(ctx, row.SiteID)
}

// EscalationChildren returns the rows of a record's child tables, by API field.
func (s *Service) EscalationChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListEscalationRecipients(ctx, id)
		if err != nil {
			return nil, wrap(err, "list escalation_recipients")
		}
		out["recipients"] = childRows(rows)
	}
	return out, nil
}

// ReplaceEscalationRecipients swaps a record's escalation recipients for the given rows.
func (s *Service) ReplaceEscalationRecipients(ctx context.Context, id uuid.UUID, rows []db.InsertEscalationRecipientsParams) error {
	if err := s.q(ctx).DeleteEscalationRecipients(ctx, id); err != nil {
		return wrap(err, "clear escalation_recipients")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertEscalationRecipients(ctx, row); err != nil {
			return wrap(err, "add escalation_recipients")
		}
	}
	return nil
}

// ListAction returns a page of actions.
func (s *Service) ListAction(ctx context.Context, p db.ListActionParams) ([]db.Action, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListAction(ctx, p))("list actions")
}

// GetAction returns one row of actions.
func (s *Service) GetAction(ctx context.Context, id uuid.UUID) (db.Action, error) {
	row, err := s.q(ctx).GetAction(ctx, id)
	if err != nil {
		return row, notFound(err, "actions")
	}
	return row, inScope(ctx, row.SiteID)
}

// CreateAction adds one row to actions.
func (s *Service) CreateAction(ctx context.Context, p db.CreateActionParams) (db.Action, error) {
	var zero db.Action
	if err := canWrite(ctx, p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateAction", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "ACT", p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateAction(ctx, p)
	if err != nil {
		return zero, wrap(err, "create actions")
	}
	if err := s.audit(ctx, "exceptions.action.write", "actions", row.ID, p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateAction", row)
}

// UpdateAction changes one row of actions.
func (s *Service) UpdateAction(ctx context.Context, p db.UpdateActionParams) (db.Action, error) {
	before, err := s.GetAction(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateAction(ctx, p)
	if err != nil {
		return before, wrap(err, "update actions")
	}
	if err := s.voidSignatures(ctx, "actions", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "exceptions.action.write", "actions", row.ID, row.SiteID, before, row)
}

// CompleteAction is the submit evidence of completion action on actions.
func (s *Service) CompleteAction(ctx context.Context, p db.ActCompleteActionParams) (db.Action, error) {
	before, err := s.GetAction(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "exceptions.action.complete", "actions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CompleteAction", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "actions", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCompleteAction(ctx, p)
	if err != nil {
		return before, stateError(err, "complete actions")
	}
	if err := s.audit(ctx, "exceptions.action.complete", "actions", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CompleteAction", row)
}

// ReopenAction is the reopen action on actions.
func (s *Service) ReopenAction(ctx context.Context, p db.ActReopenActionParams) (db.Action, error) {
	before, err := s.GetAction(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "exceptions.action.reopen", "actions", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ReopenAction", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActReopenAction(ctx, p)
	if err != nil {
		return before, stateError(err, "reopen actions")
	}
	if err := s.audit(ctx, "exceptions.action.reopen", "actions", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ReopenAction", row)
}

// ListNotification returns a page of notifications.
func (s *Service) ListNotification(ctx context.Context, p db.ListNotificationParams) ([]db.Notification, error) {
	return wrap2(s.q(ctx).ListNotification(ctx, p))("list notifications")
}
