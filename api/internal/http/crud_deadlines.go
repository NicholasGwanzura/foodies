// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package httpapi

import (
	"context"

	"foodies/api/internal/db"
	"foodies/api/internal/http/apigen"
)

var _ db.Querier

// ListReportDefinitions implements the listReportDefinitions operation.
func (s *Server) ListReportDefinitions(ctx context.Context, req apigen.ListReportDefinitionsRequestObject) (apigen.ListReportDefinitionsResponseObject, error) {
	return listPage[apigen.ListReportDefinitions200JSONResponse](ctx, req.Params, nil, s.service.ListReportDefinition)
}

// GetReportDefinition implements the getReportDefinition operation.
func (s *Server) GetReportDefinition(ctx context.Context, req apigen.GetReportDefinitionRequestObject) (apigen.GetReportDefinitionResponseObject, error) {
	row, err := s.service.GetReportDefinition(ctx, req.ReportDefinitionId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ReportDefinition](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetReportDefinition200JSONResponse{Body: body, Headers: apigen.GetReportDefinition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateReportDefinition implements the createReportDefinition operation.
func (s *Server) CreateReportDefinition(ctx context.Context, req apigen.CreateReportDefinitionRequestObject) (apigen.CreateReportDefinitionResponseObject, error) {
	in, err := input[db.CreateReportDefinitionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateReportDefinition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ReportDefinition](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateReportDefinition201JSONResponse{Body: body, Headers: apigen.CreateReportDefinition201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateReportDefinition implements the updateReportDefinition operation.
func (s *Server) UpdateReportDefinition(ctx context.Context, req apigen.UpdateReportDefinitionRequestObject) (apigen.UpdateReportDefinitionResponseObject, error) {
	in, err := input[db.UpdateReportDefinitionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.ReportDefinitionId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateReportDefinition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ReportDefinition](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateReportDefinition200JSONResponse{Body: body, Headers: apigen.UpdateReportDefinition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListReports implements the listReports operation.
func (s *Server) ListReports(ctx context.Context, req apigen.ListReportsRequestObject) (apigen.ListReportsResponseObject, error) {
	return listPage[apigen.ListReports200JSONResponse](ctx, req.Params, nil, s.service.ListReport)
}

// GetReport implements the getReport operation.
func (s *Server) GetReport(ctx context.Context, req apigen.GetReportRequestObject) (apigen.GetReportResponseObject, error) {
	row, err := s.service.GetReport(ctx, req.ReportId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Report](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetReport200JSONResponse{Body: body, Headers: apigen.GetReport200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateReport implements the updateReport operation.
func (s *Server) UpdateReport(ctx context.Context, req apigen.UpdateReportRequestObject) (apigen.UpdateReportResponseObject, error) {
	in, err := input[db.UpdateReportParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.ReportId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateReport(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Report](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateReport200JSONResponse{Body: body, Headers: apigen.UpdateReport200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitReport implements the submitReport operation.
func (s *Server) SubmitReport(ctx context.Context, req apigen.SubmitReportRequestObject) (apigen.SubmitReportResponseObject, error) {
	in, err := input[db.ActSubmitReportParams](struct{}{}, map[string]any{"id": req.ReportId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitReport(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Report](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitReport200JSONResponse{Body: body, Headers: apigen.SubmitReport200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReviewReport implements the reviewReport operation.
func (s *Server) ReviewReport(ctx context.Context, req apigen.ReviewReportRequestObject) (apigen.ReviewReportResponseObject, error) {
	in, err := input[db.ActReviewReportParams](req.Body, map[string]any{"id": req.ReportId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReviewReport(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Report](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReviewReport200JSONResponse{Body: body, Headers: apigen.ReviewReport200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ExplainLateReport implements the explainLateReport operation.
func (s *Server) ExplainLateReport(ctx context.Context, req apigen.ExplainLateReportRequestObject) (apigen.ExplainLateReportResponseObject, error) {
	in, err := input[db.ActExplainLateReportParams](req.Body, map[string]any{"id": req.ReportId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ExplainLateReport(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Report](row)
	if err != nil {
		return nil, err
	}
	return apigen.ExplainLateReport200JSONResponse{Body: body, Headers: apigen.ExplainLateReport200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListReportDeadlineNotices implements the listReportDeadlineNotices operation.
func (s *Server) ListReportDeadlineNotices(ctx context.Context, req apigen.ListReportDeadlineNoticesRequestObject) (apigen.ListReportDeadlineNoticesResponseObject, error) {
	return listPage[apigen.ListReportDeadlineNotices200JSONResponse](ctx, req.Params, nil, s.service.ListReportDeadlineNotice)
}

// GetReportDeadlineNotice implements the getReportDeadlineNotice operation.
func (s *Server) GetReportDeadlineNotice(ctx context.Context, req apigen.GetReportDeadlineNoticeRequestObject) (apigen.GetReportDeadlineNoticeResponseObject, error) {
	row, err := s.service.GetReportDeadlineNotice(ctx, req.ReportDeadlineNoticeId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ReportDeadlineNotice](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetReportDeadlineNotice200JSONResponse{Body: body, Headers: apigen.GetReportDeadlineNotice200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateReportDeadlineNotice implements the createReportDeadlineNotice operation.
func (s *Server) CreateReportDeadlineNotice(ctx context.Context, req apigen.CreateReportDeadlineNoticeRequestObject) (apigen.CreateReportDeadlineNoticeResponseObject, error) {
	in, err := input[db.CreateReportDeadlineNoticeParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateReportDeadlineNotice(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ReportDeadlineNotice](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateReportDeadlineNotice201JSONResponse{Body: body, Headers: apigen.CreateReportDeadlineNotice201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListManagementMeetings implements the listManagementMeetings operation.
func (s *Server) ListManagementMeetings(ctx context.Context, req apigen.ListManagementMeetingsRequestObject) (apigen.ListManagementMeetingsResponseObject, error) {
	return listPage[apigen.ListManagementMeetings200JSONResponse](ctx, req.Params, nil, s.service.ListManagementMeeting)
}

// GetManagementMeeting implements the getManagementMeeting operation.
func (s *Server) GetManagementMeeting(ctx context.Context, req apigen.GetManagementMeetingRequestObject) (apigen.GetManagementMeetingResponseObject, error) {
	row, err := s.service.GetManagementMeeting(ctx, req.ManagementMeetingId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ManagementMeeting](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetManagementMeeting200JSONResponse{Body: body, Headers: apigen.GetManagementMeeting200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateManagementMeeting implements the createManagementMeeting operation.
func (s *Server) CreateManagementMeeting(ctx context.Context, req apigen.CreateManagementMeetingRequestObject) (apigen.CreateManagementMeetingResponseObject, error) {
	in, err := input[db.CreateManagementMeetingParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateManagementMeeting(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ManagementMeeting](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateManagementMeeting201JSONResponse{Body: body, Headers: apigen.CreateManagementMeeting201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateManagementMeeting implements the updateManagementMeeting operation.
func (s *Server) UpdateManagementMeeting(ctx context.Context, req apigen.UpdateManagementMeetingRequestObject) (apigen.UpdateManagementMeetingResponseObject, error) {
	in, err := input[db.UpdateManagementMeetingParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.ManagementMeetingId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateManagementMeeting(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ManagementMeeting](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateManagementMeeting200JSONResponse{Body: body, Headers: apigen.UpdateManagementMeeting200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListComplianceObligations implements the listComplianceObligations operation.
func (s *Server) ListComplianceObligations(ctx context.Context, req apigen.ListComplianceObligationsRequestObject) (apigen.ListComplianceObligationsResponseObject, error) {
	return listPage[apigen.ListComplianceObligations200JSONResponse](ctx, req.Params, nil, s.service.ListComplianceObligation)
}

// GetComplianceObligation implements the getComplianceObligation operation.
func (s *Server) GetComplianceObligation(ctx context.Context, req apigen.GetComplianceObligationRequestObject) (apigen.GetComplianceObligationResponseObject, error) {
	row, err := s.service.GetComplianceObligation(ctx, req.ComplianceObligationId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ComplianceObligation](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetComplianceObligation200JSONResponse{Body: body, Headers: apigen.GetComplianceObligation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateComplianceObligation implements the createComplianceObligation operation.
func (s *Server) CreateComplianceObligation(ctx context.Context, req apigen.CreateComplianceObligationRequestObject) (apigen.CreateComplianceObligationResponseObject, error) {
	in, err := input[db.CreateComplianceObligationParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateComplianceObligation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ComplianceObligation](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateComplianceObligation201JSONResponse{Body: body, Headers: apigen.CreateComplianceObligation201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateComplianceObligation implements the updateComplianceObligation operation.
func (s *Server) UpdateComplianceObligation(ctx context.Context, req apigen.UpdateComplianceObligationRequestObject) (apigen.UpdateComplianceObligationResponseObject, error) {
	in, err := input[db.UpdateComplianceObligationParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.ComplianceObligationId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateComplianceObligation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ComplianceObligation](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateComplianceObligation200JSONResponse{Body: body, Headers: apigen.UpdateComplianceObligation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CompleteComplianceObligation implements the completeComplianceObligation operation.
func (s *Server) CompleteComplianceObligation(ctx context.Context, req apigen.CompleteComplianceObligationRequestObject) (apigen.CompleteComplianceObligationResponseObject, error) {
	in, err := input[db.ActCompleteComplianceObligationParams](req.Body, map[string]any{"id": req.ComplianceObligationId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CompleteComplianceObligation(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ComplianceObligation](row)
	if err != nil {
		return nil, err
	}
	return apigen.CompleteComplianceObligation200JSONResponse{Body: body, Headers: apigen.CompleteComplianceObligation200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}
