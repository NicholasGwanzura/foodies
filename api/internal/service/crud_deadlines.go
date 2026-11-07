// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package service

import (
	"context"

	"github.com/google/uuid"

	"foodies/api/internal/db"
)

var _ uuid.UUID

// ListReportDefinition returns a page of report definitions.
func (s *Service) ListReportDefinition(ctx context.Context, p db.ListReportDefinitionParams) ([]db.ReportDefinition, error) {
	return wrap2(s.q(ctx).ListReportDefinition(ctx, p))("list report definitions")
}

// GetReportDefinition returns one row of report definitions.
func (s *Service) GetReportDefinition(ctx context.Context, id uuid.UUID) (db.ReportDefinition, error) {
	row, err := s.q(ctx).GetReportDefinition(ctx, id)
	if err != nil {
		return row, notFound(err, "report definitions")
	}
	return row, nil
}

// CreateReportDefinition adds one row to report definitions.
func (s *Service) CreateReportDefinition(ctx context.Context, p db.CreateReportDefinitionParams) (db.ReportDefinition, error) {
	var zero db.ReportDefinition
	if err := s.guard(ctx, "CreateReportDefinition", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateReportDefinition(ctx, p)
	if err != nil {
		return zero, wrap(err, "create report definitions")
	}
	if err := s.audit(ctx, "mgmt_reports.report_definition.write", "report_definitions", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateReportDefinition", row)
}

// UpdateReportDefinition changes one row of report definitions.
func (s *Service) UpdateReportDefinition(ctx context.Context, p db.UpdateReportDefinitionParams) (db.ReportDefinition, error) {
	before, err := s.GetReportDefinition(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateReportDefinition(ctx, p)
	if err != nil {
		return before, wrap(err, "update report definitions")
	}
	if err := s.voidSignatures(ctx, "report_definitions", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "mgmt_reports.report_definition.write", "report_definitions", row.ID, nil, before, row)
}

// ListReport returns a page of reports.
func (s *Service) ListReport(ctx context.Context, p db.ListReportParams) ([]db.Report, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListReport(ctx, p))("list reports")
}

// GetReport returns one row of reports.
func (s *Service) GetReport(ctx context.Context, id uuid.UUID) (db.Report, error) {
	row, err := s.q(ctx).GetReport(ctx, id)
	if err != nil {
		return row, notFound(err, "reports")
	}
	return row, inScope(ctx, row.SiteID)
}

// UpdateReport changes one row of reports.
func (s *Service) UpdateReport(ctx context.Context, p db.UpdateReportParams) (db.Report, error) {
	before, err := s.GetReport(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateReport(ctx, p)
	if err != nil {
		return before, wrap(err, "update reports")
	}
	if err := s.voidSignatures(ctx, "reports", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "mgmt_reports.report.write", "reports", row.ID, row.SiteID, before, row)
}

// SubmitReport is the submit action on reports.
func (s *Service) SubmitReport(ctx context.Context, p db.ActSubmitReportParams) (db.Report, error) {
	before, err := s.GetReport(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "mgmt_reports.report.submit", "reports", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitReport", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitReport(ctx, p)
	if err != nil {
		return before, stateError(err, "submit reports")
	}
	if err := s.audit(ctx, "mgmt_reports.report.submit", "reports", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitReport", row)
}

// ReviewReport is the review action on reports.
func (s *Service) ReviewReport(ctx context.Context, p db.ActReviewReportParams) (db.Report, error) {
	before, err := s.GetReport(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "mgmt_reports.report.review", "reports", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ReviewReport", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "reports", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActReviewReport(ctx, p)
	if err != nil {
		return before, stateError(err, "review reports")
	}
	if err := s.audit(ctx, "mgmt_reports.report.review", "reports", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ReviewReport", row)
}

// ExplainLateReport is the record the reason for delay action on reports.
func (s *Service) ExplainLateReport(ctx context.Context, p db.ActExplainLateReportParams) (db.Report, error) {
	before, err := s.GetReport(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "mgmt_reports.report.explain_late", "reports", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ExplainLateReport", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActExplainLateReport(ctx, p)
	if err != nil {
		return before, stateError(err, "explain late reports")
	}
	if err := s.audit(ctx, "mgmt_reports.report.explain_late", "reports", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ExplainLateReport", row)
}

// ListReportDeadlineNotice returns a page of report deadline notices.
func (s *Service) ListReportDeadlineNotice(ctx context.Context, p db.ListReportDeadlineNoticeParams) ([]db.ReportDeadlineNotice, error) {
	return wrap2(s.q(ctx).ListReportDeadlineNotice(ctx, p))("list report deadline notices")
}

// GetReportDeadlineNotice returns one row of report deadline notices.
func (s *Service) GetReportDeadlineNotice(ctx context.Context, id uuid.UUID) (db.ReportDeadlineNotice, error) {
	row, err := s.q(ctx).GetReportDeadlineNotice(ctx, id)
	if err != nil {
		return row, notFound(err, "report deadline notices")
	}
	return row, nil
}

// CreateReportDeadlineNotice adds one row to report deadline notices.
func (s *Service) CreateReportDeadlineNotice(ctx context.Context, p db.CreateReportDeadlineNoticeParams) (db.ReportDeadlineNotice, error) {
	var zero db.ReportDeadlineNotice
	if err := s.guard(ctx, "CreateReportDeadlineNotice", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateReportDeadlineNotice(ctx, p)
	if err != nil {
		return zero, wrap(err, "create report deadline notices")
	}
	if err := s.audit(ctx, "mgmt_reports.report_deadline_notice.write", "report_deadline_notices", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateReportDeadlineNotice", row)
}

// ListManagementMeeting returns a page of management meetings.
func (s *Service) ListManagementMeeting(ctx context.Context, p db.ListManagementMeetingParams) ([]db.ManagementMeeting, error) {
	return wrap2(s.q(ctx).ListManagementMeeting(ctx, p))("list management meetings")
}

// GetManagementMeeting returns one row of management meetings.
func (s *Service) GetManagementMeeting(ctx context.Context, id uuid.UUID) (db.ManagementMeeting, error) {
	row, err := s.q(ctx).GetManagementMeeting(ctx, id)
	if err != nil {
		return row, notFound(err, "management meetings")
	}
	return row, nil
}

// CreateManagementMeeting adds one row to management meetings.
func (s *Service) CreateManagementMeeting(ctx context.Context, p db.CreateManagementMeetingParams) (db.ManagementMeeting, error) {
	var zero db.ManagementMeeting
	if err := s.guard(ctx, "CreateManagementMeeting", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "MTG", nil)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateManagementMeeting(ctx, p)
	if err != nil {
		return zero, wrap(err, "create management meetings")
	}
	if err := s.audit(ctx, "mgmt_reports.management_meeting.write", "management_meetings", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateManagementMeeting", row)
}

// UpdateManagementMeeting changes one row of management meetings.
func (s *Service) UpdateManagementMeeting(ctx context.Context, p db.UpdateManagementMeetingParams) (db.ManagementMeeting, error) {
	before, err := s.GetManagementMeeting(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateManagementMeeting(ctx, p)
	if err != nil {
		return before, wrap(err, "update management meetings")
	}
	if err := s.voidSignatures(ctx, "management_meetings", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "mgmt_reports.management_meeting.write", "management_meetings", row.ID, nil, before, row)
}

// ListComplianceObligation returns a page of compliance obligations.
func (s *Service) ListComplianceObligation(ctx context.Context, p db.ListComplianceObligationParams) ([]db.ComplianceObligation, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListComplianceObligation(ctx, p))("list compliance obligations")
}

// GetComplianceObligation returns one row of compliance obligations.
func (s *Service) GetComplianceObligation(ctx context.Context, id uuid.UUID) (db.ComplianceObligation, error) {
	row, err := s.q(ctx).GetComplianceObligation(ctx, id)
	if err != nil {
		return row, notFound(err, "compliance obligations")
	}
	return row, inScope(ctx, row.SiteID)
}

// CreateComplianceObligation adds one row to compliance obligations.
func (s *Service) CreateComplianceObligation(ctx context.Context, p db.CreateComplianceObligationParams) (db.ComplianceObligation, error) {
	var zero db.ComplianceObligation
	if err := canWrite(ctx, p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateComplianceObligation", &p); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "OBL", p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateComplianceObligation(ctx, p)
	if err != nil {
		return zero, wrap(err, "create compliance obligations")
	}
	if err := s.audit(ctx, "compliance.compliance_obligation.write", "compliance_obligations", row.ID, p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateComplianceObligation", row)
}

// UpdateComplianceObligation changes one row of compliance obligations.
func (s *Service) UpdateComplianceObligation(ctx context.Context, p db.UpdateComplianceObligationParams) (db.ComplianceObligation, error) {
	before, err := s.GetComplianceObligation(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateComplianceObligation(ctx, p)
	if err != nil {
		return before, wrap(err, "update compliance obligations")
	}
	if err := s.voidSignatures(ctx, "compliance_obligations", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "compliance.compliance_obligation.write", "compliance_obligations", row.ID, row.SiteID, before, row)
}

// CompleteComplianceObligation is the complete with evidence action on compliance obligations.
func (s *Service) CompleteComplianceObligation(ctx context.Context, p db.ActCompleteComplianceObligationParams) (db.ComplianceObligation, error) {
	before, err := s.GetComplianceObligation(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "compliance.compliance_obligation.complete", "compliance_obligations", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CompleteComplianceObligation", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "compliance_obligations", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCompleteComplianceObligation(ctx, p)
	if err != nil {
		return before, stateError(err, "complete compliance obligations")
	}
	if err := s.audit(ctx, "compliance.compliance_obligation.complete", "compliance_obligations", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CompleteComplianceObligation", row)
}
