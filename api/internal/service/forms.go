package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"

	"foodies/api/internal/db"
)

// FormField is one question in a template definition.
type FormField struct {
	Key      string   `json:"key"`
	Label    string   `json:"label"`
	Type     string   `json:"type"` // text, number, reading, date, choice, pass_fail, ok_nok, photo
	Required bool     `json:"required"`
	Options  []string `json:"options"`
	Min      *float64 `json:"min"`
	Max      *float64 `json:"max"`
	// RuleMax names a rule holding the limit, so limits are never hard-coded.
	RuleMax string `json:"rule_max"`
	// OnFail says what a failed answer does.
	OnFail *struct {
		Create   string `json:"create"` // action, task, escalation
		Severity int    `json:"severity"`
		Domain   string `json:"domain"`
		Block    bool   `json:"block"`
	} `json:"on_fail"`
}

// FormDefinition is the template's JSON.
type FormDefinition struct {
	Sections []struct {
		Title  string      `json:"title"`
		Fields []FormField `json:"fields"`
	} `json:"sections"`
	RequireReview bool `json:"require_review"`
}

func decodeDefinition(raw json.RawMessage) (FormDefinition, error) {
	var d FormDefinition
	if err := json.Unmarshal(raw, &d); err != nil {
		return d, errInvalid("/definition", "The form definition is not valid.")
	}
	return d, nil
}

func validDefinition(raw json.RawMessage) error {
	d, err := decodeDefinition(raw)
	if err != nil {
		return err
	}
	count, seen := 0, map[string]bool{}
	for _, sec := range d.Sections {
		for _, f := range sec.Fields {
			if f.Key == "" || seen[f.Key] {
				return errInvalid("/definition", "Every question needs its own key.")
			}
			seen[f.Key] = true
			count++
		}
	}
	if count == 0 {
		return errInvalid("/definition", "Add at least one question.")
	}
	return nil
}

// CreateFormTemplateVersion adds the next draft of a template.
func (s *Service) CreateFormTemplateVersion(ctx context.Context, p db.CreateFormTemplateVersionParams) (db.FormTemplateVersion, error) {
	if err := validDefinition(p.Definition); err != nil {
		return db.FormTemplateVersion{}, err
	}
	n, err := s.q(ctx).NextTemplateVersion(ctx, p.FormTemplateID)
	if err != nil {
		return db.FormTemplateVersion{}, wrap(err, "next version")
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateFormTemplateVersionNumbered(ctx, db.CreateFormTemplateVersionNumberedParams{
		FormTemplateID: p.FormTemplateID, Version: &n, Definition: p.Definition, OfflineAllowed: p.OfflineAllowed, Channels: p.Channels,
		RetentionClassID: p.RetentionClassID, RequiredRuleKeys: p.RequiredRuleKeys, EffectiveFrom: p.EffectiveFrom, CreatedBy: p.CreatedBy,
	})
	if err != nil {
		return row, wrap(err, "create version")
	}
	return row, s.audit(ctx, "forms.form_template_version.write", "form_template_versions", row.ID, nil, nil, row)
}

// PublishFormTemplateVersion makes a version the one staff fill in. A template
// that needs rules with no value cannot publish (spec D-10).
func (s *Service) PublishFormTemplateVersion(ctx context.Context, p db.ActPublishFormTemplateVersionParams) (db.FormTemplateVersion, error) {
	before, err := s.GetFormTemplateVersion(ctx, p.ID)
	if err != nil {
		return before, err
	}
	var blockers []Blocker
	for _, key := range before.RequiredRuleKeys {
		if _, err := s.RuleValue(ctx, key, nil); err != nil {
			blockers = append(blockers, Blocker{Kind: "rule_missing", Reference: key, Message: key + " has no approved value"})
		}
	}
	if len(blockers) > 0 {
		return before, &Error{Status: http.StatusConflict, Code: "rule_not_configured", Title: "Rules missing", Detail: "This form depends on rules with no approved value.", Blockers: blockers}
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActPublishFormTemplateVersion(ctx, p)
	if err != nil {
		return before, stateError(err, "publish this version")
	}
	if err := s.publishVersion(ctx, row); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "forms.form_template_version.publish", "form_template_versions", row.ID, nil, before, row)
}

func (s *Service) publishVersion(ctx context.Context, v db.FormTemplateVersion) error {
	q := s.q(ctx)
	if err := q.RetireOtherVersions(ctx, db.RetireOtherVersionsParams{FormTemplateID: v.FormTemplateID, ID: v.ID}); err != nil {
		return wrap(err, "retire old versions")
	}
	return wrap(q.SetCurrentTemplateVersion(ctx, db.SetCurrentTemplateVersionParams{ID: v.FormTemplateID, CurrentVersionID: &v.ID}), "set current version")
}

func failed(ctx context.Context, s *Service, f FormField, value any, site uuid.UUID) (bool, error) {
	switch f.Type {
	case "pass_fail":
		return value == "fail", nil
	case "ok_nok":
		return value == "nok", nil
	case "number", "reading":
		n, ok := toFloat(value)
		if !ok {
			return false, nil
		}
		limit := f.Max
		if f.RuleMax != "" {
			v, err := s.RuleValue(ctx, f.RuleMax, &site)
			if err != nil {
				return false, err
			}
			var parsed float64
			if json.Unmarshal(v.Value, &parsed) == nil {
				limit = &parsed
			}
		}
		return (limit != nil && n > *limit) || (f.Min != nil && n < *f.Min), nil
	}
	return false, nil
}

func toFloat(v any) (float64, bool) {
	switch x := v.(type) {
	case float64:
		return x, true
	case string:
		n, err := strconv.ParseFloat(x, 64)
		return n, err == nil
	}
	return 0, false
}

// SubmitFormSubmission validates the answers, sets the result and runs any
// consequences of failed answers.
func (s *Service) SubmitFormSubmission(ctx context.Context, p db.ActSubmitFormSubmissionParams) (db.FormSubmission, error) {
	before, err := s.GetFormSubmission(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if before.Status != "draft" && before.Status != "returned" {
		return before, errState("This form has already been submitted.")
	}
	q := s.q(ctx)
	version, err := q.GetFormTemplateVersion(ctx, before.TemplateVersionID)
	if err != nil {
		return before, wrap(err, "read version")
	}
	def, err := decodeDefinition(version.Definition)
	if err != nil {
		return before, err
	}
	answers := map[string]any{}
	if err := json.Unmarshal(before.Answers, &answers); err != nil {
		return before, errInvalid("/answers", "The answers are not valid.")
	}

	var issues []FieldIssue
	var fails []FormField
	blocked := false
	for _, sec := range def.Sections {
		for _, f := range sec.Fields {
			value, present := answers[f.Key]
			if f.Required && (!present || value == "" || value == nil) {
				issues = append(issues, FieldIssue{Pointer: "/answers/" + f.Key, Code: "required", Message: f.Label + " is required."})
				continue
			}
			bad, err := failed(ctx, s, f, value, before.SiteID)
			if err != nil {
				return before, err
			}
			if bad {
				fails = append(fails, f)
				blocked = blocked || (f.OnFail != nil && f.OnFail.Block)
			}
		}
	}
	if len(issues) > 0 {
		return before, &Error{Status: http.StatusUnprocessableEntity, Code: "validation_failed", Title: "Validation failed", Detail: "Some answers are missing.", Fields: issues}
	}

	result, status := "pass", "closed"
	if len(fails) > 0 {
		result = "fail"
	}
	if def.RequireReview || blocked {
		status = "pending_review"
	}
	if n, err := q.CountPendingSignatures(ctx, db.CountPendingSignaturesParams{SubjectType: "form_submissions", SubjectID: before.ID}); err == nil && n > 0 {
		status = "pending_signature"
	}
	row, err := q.SetSubmissionOutcome(ctx, db.SetSubmissionOutcomeParams{ID: before.ID, Status: status, Result: &result, UpdatedBy: actorID(ctx)})
	if err != nil {
		return before, wrap(err, "store outcome")
	}
	if err := s.audit(ctx, "forms.form_submission.submit", "form_submissions", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	for _, f := range fails {
		if f.OnFail != nil {
			if err := s.formConsequence(ctx, row, f); err != nil {
				return row, err
			}
		}
	}
	if row.TaskID != nil {
		if _, err := s.CompleteTask(ctx, *row.TaskID); err != nil {
			return row, err
		}
	}
	return row, nil
}

func (s *Service) formConsequence(ctx context.Context, sub db.FormSubmission, f FormField) error {
	what := fmt.Sprintf("%s failed on %s", f.Label, sub.Reference)
	switch f.OnFail.Create {
	case "action":
		ref, err := s.nextReference(ctx, "ACT", &sub.SiteID)
		if err != nil {
			return err
		}
		owner, err := s.q(ctx).GetUser(ctx, *actorID(ctx))
		if err != nil {
			return wrap(err, "read owner")
		}
		_, err = s.q(ctx).CreateAction(ctx, db.CreateActionParams{
			Reference: ref, SourceType: ptr("form_submissions"), SourceID: &sub.ID, Kind: "corrective", Description: what,
			OwnerEmployeeID: owner.EmployeeID, SiteID: &sub.SiteID, DueDate: s.today(), Priority: ptr("high"), CreatedBy: actorID(ctx),
		})
		return wrap(err, "create action")
	case "escalation":
		domain := f.OnFail.Domain
		if domain == "" {
			domain = "general"
		}
		_, err := s.RaiseEscalation(ctx, "form_submissions", sub.ID, domain, f.OnFail.Severity, &sub.SiteID)
		return err
	case "task":
		_, err := s.CreateTask(ctx, db.CreateTaskRowParams{Title: what, SubjectType: ptr("form_submissions"), SubjectID: &sub.ID, SiteID: sub.SiteID, BusinessDate: s.today(), Priority: ptr("high"), AssigneeUserID: actorID(ctx)})
		return err
	}
	return nil
}

// ReviewFormSubmission closes a reviewed submission.
func (s *Service) ReviewFormSubmission(ctx context.Context, p db.ActReviewFormSubmissionParams) (db.FormSubmission, error) {
	before, err := s.GetFormSubmission(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if before.CreatedBy != nil && actorID(ctx) != nil && *before.CreatedBy == *actorID(ctx) {
		return before, &Error{Status: http.StatusForbidden, Code: "sod_violation", Title: "Someone else must do this", Detail: "You cannot review your own form."}
	}
	if err := s.requireSignatures(ctx, "form_submissions", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActReviewFormSubmission(ctx, p)
	if err != nil {
		return before, stateError(err, "review this form")
	}
	return row, s.audit(ctx, "forms.form_submission.review", "form_submissions", row.ID, &row.SiteID, before, row)
}
