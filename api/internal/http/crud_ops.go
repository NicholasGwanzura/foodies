// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package httpapi

import (
	"context"

	"foodies/api/internal/db"
	"foodies/api/internal/http/apigen"
)

var _ db.Querier

// ListSyncConflicts implements the listSyncConflicts operation.
func (s *Server) ListSyncConflicts(ctx context.Context, req apigen.ListSyncConflictsRequestObject) (apigen.ListSyncConflictsResponseObject, error) {
	return listPage[apigen.ListSyncConflicts200JSONResponse](ctx, req.Params, nil, s.service.ListSyncConflict)
}

// GetSyncConflict implements the getSyncConflict operation.
func (s *Server) GetSyncConflict(ctx context.Context, req apigen.GetSyncConflictRequestObject) (apigen.GetSyncConflictResponseObject, error) {
	row, err := s.service.GetSyncConflict(ctx, req.SyncConflictId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SyncConflict](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetSyncConflict200JSONResponse{Body: body, Headers: apigen.GetSyncConflict200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ResolveSyncConflict implements the resolveSyncConflict operation.
func (s *Server) ResolveSyncConflict(ctx context.Context, req apigen.ResolveSyncConflictRequestObject) (apigen.ResolveSyncConflictResponseObject, error) {
	in, err := input[db.ActResolveSyncConflictParams](req.Body, map[string]any{"id": req.SyncConflictId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ResolveSyncConflict(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SyncConflict](row)
	if err != nil {
		return nil, err
	}
	return apigen.ResolveSyncConflict200JSONResponse{Body: body, Headers: apigen.ResolveSyncConflict200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListWhatsappIdentities implements the listWhatsappIdentities operation.
func (s *Server) ListWhatsappIdentities(ctx context.Context, req apigen.ListWhatsappIdentitiesRequestObject) (apigen.ListWhatsappIdentitiesResponseObject, error) {
	return listPage[apigen.ListWhatsappIdentities200JSONResponse](ctx, req.Params, nil, s.service.ListWhatsappIdentity)
}

// GetWhatsappIdentity implements the getWhatsappIdentity operation.
func (s *Server) GetWhatsappIdentity(ctx context.Context, req apigen.GetWhatsappIdentityRequestObject) (apigen.GetWhatsappIdentityResponseObject, error) {
	row, err := s.service.GetWhatsappIdentity(ctx, req.WhatsappIdentityId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WhatsappIdentity](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetWhatsappIdentity200JSONResponse{Body: body, Headers: apigen.GetWhatsappIdentity200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateWhatsappIdentity implements the createWhatsappIdentity operation.
func (s *Server) CreateWhatsappIdentity(ctx context.Context, req apigen.CreateWhatsappIdentityRequestObject) (apigen.CreateWhatsappIdentityResponseObject, error) {
	in, err := input[db.CreateWhatsappIdentityParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateWhatsappIdentity(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WhatsappIdentity](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateWhatsappIdentity201JSONResponse{Body: body, Headers: apigen.CreateWhatsappIdentity201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// RevokeWhatsappIdentity implements the revokeWhatsappIdentity operation.
func (s *Server) RevokeWhatsappIdentity(ctx context.Context, req apigen.RevokeWhatsappIdentityRequestObject) (apigen.RevokeWhatsappIdentityResponseObject, error) {
	in, err := input[db.ActRevokeWhatsappIdentityParams](req.Body, map[string]any{"id": req.WhatsappIdentityId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.RevokeWhatsappIdentity(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WhatsappIdentity](row)
	if err != nil {
		return nil, err
	}
	return apigen.RevokeWhatsappIdentity200JSONResponse{Body: body, Headers: apigen.RevokeWhatsappIdentity200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// RebindWhatsappIdentity implements the rebindWhatsappIdentity operation.
func (s *Server) RebindWhatsappIdentity(ctx context.Context, req apigen.RebindWhatsappIdentityRequestObject) (apigen.RebindWhatsappIdentityResponseObject, error) {
	in, err := input[db.ActRebindWhatsappIdentityParams](req.Body, map[string]any{"id": req.WhatsappIdentityId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.RebindWhatsappIdentity(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WhatsappIdentity](row)
	if err != nil {
		return nil, err
	}
	return apigen.RebindWhatsappIdentity200JSONResponse{Body: body, Headers: apigen.RebindWhatsappIdentity200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListWhatsappTemplates implements the listWhatsappTemplates operation.
func (s *Server) ListWhatsappTemplates(ctx context.Context, req apigen.ListWhatsappTemplatesRequestObject) (apigen.ListWhatsappTemplatesResponseObject, error) {
	return listPage[apigen.ListWhatsappTemplates200JSONResponse](ctx, req.Params, nil, s.service.ListWhatsappTemplate)
}

// GetWhatsappTemplate implements the getWhatsappTemplate operation.
func (s *Server) GetWhatsappTemplate(ctx context.Context, req apigen.GetWhatsappTemplateRequestObject) (apigen.GetWhatsappTemplateResponseObject, error) {
	row, err := s.service.GetWhatsappTemplate(ctx, req.WhatsappTemplateId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WhatsappTemplate](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetWhatsappTemplate200JSONResponse{Body: body, Headers: apigen.GetWhatsappTemplate200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateWhatsappTemplate implements the createWhatsappTemplate operation.
func (s *Server) CreateWhatsappTemplate(ctx context.Context, req apigen.CreateWhatsappTemplateRequestObject) (apigen.CreateWhatsappTemplateResponseObject, error) {
	in, err := input[db.CreateWhatsappTemplateParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateWhatsappTemplate(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WhatsappTemplate](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateWhatsappTemplate201JSONResponse{Body: body, Headers: apigen.CreateWhatsappTemplate201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateWhatsappTemplate implements the updateWhatsappTemplate operation.
func (s *Server) UpdateWhatsappTemplate(ctx context.Context, req apigen.UpdateWhatsappTemplateRequestObject) (apigen.UpdateWhatsappTemplateResponseObject, error) {
	in, err := input[db.UpdateWhatsappTemplateParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.WhatsappTemplateId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateWhatsappTemplate(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WhatsappTemplate](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateWhatsappTemplate200JSONResponse{Body: body, Headers: apigen.UpdateWhatsappTemplate200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListWhatsappMessages implements the listWhatsappMessages operation.
func (s *Server) ListWhatsappMessages(ctx context.Context, req apigen.ListWhatsappMessagesRequestObject) (apigen.ListWhatsappMessagesResponseObject, error) {
	return listPage[apigen.ListWhatsappMessages200JSONResponse](ctx, req.Params, nil, s.service.ListWhatsappMessage)
}

// GetWhatsappMessage implements the getWhatsappMessage operation.
func (s *Server) GetWhatsappMessage(ctx context.Context, req apigen.GetWhatsappMessageRequestObject) (apigen.GetWhatsappMessageResponseObject, error) {
	row, err := s.service.GetWhatsappMessage(ctx, req.WhatsappMessageId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.WhatsappMessage](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetWhatsappMessage200JSONResponse{Body: body, Headers: apigen.GetWhatsappMessage200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListIntegrations implements the listIntegrations operation.
func (s *Server) ListIntegrations(ctx context.Context, req apigen.ListIntegrationsRequestObject) (apigen.ListIntegrationsResponseObject, error) {
	return listPage[apigen.ListIntegrations200JSONResponse](ctx, req.Params, nil, s.service.ListIntegration)
}

// GetIntegration implements the getIntegration operation.
func (s *Server) GetIntegration(ctx context.Context, req apigen.GetIntegrationRequestObject) (apigen.GetIntegrationResponseObject, error) {
	row, err := s.service.GetIntegration(ctx, req.IntegrationId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Integration](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetIntegration200JSONResponse{Body: body, Headers: apigen.GetIntegration200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateIntegration implements the createIntegration operation.
func (s *Server) CreateIntegration(ctx context.Context, req apigen.CreateIntegrationRequestObject) (apigen.CreateIntegrationResponseObject, error) {
	in, err := input[db.CreateIntegrationParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateIntegration(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Integration](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateIntegration201JSONResponse{Body: body, Headers: apigen.CreateIntegration201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateIntegration implements the updateIntegration operation.
func (s *Server) UpdateIntegration(ctx context.Context, req apigen.UpdateIntegrationRequestObject) (apigen.UpdateIntegrationResponseObject, error) {
	in, err := input[db.UpdateIntegrationParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.IntegrationId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateIntegration(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Integration](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateIntegration200JSONResponse{Body: body, Headers: apigen.UpdateIntegration200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListSyncRuns implements the listSyncRuns operation.
func (s *Server) ListSyncRuns(ctx context.Context, req apigen.ListSyncRunsRequestObject) (apigen.ListSyncRunsResponseObject, error) {
	return listPage[apigen.ListSyncRuns200JSONResponse](ctx, req.Params, map[string]any{"integration_id": req.IntegrationId}, s.service.ListSyncRun)
}

// GetSyncRun implements the getSyncRun operation.
func (s *Server) GetSyncRun(ctx context.Context, req apigen.GetSyncRunRequestObject) (apigen.GetSyncRunResponseObject, error) {
	row, err := s.service.GetSyncRun(ctx, req.SyncRunId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SyncRun](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetSyncRun200JSONResponse{Body: body, Headers: apigen.GetSyncRun200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateSyncRun implements the createSyncRun operation.
func (s *Server) CreateSyncRun(ctx context.Context, req apigen.CreateSyncRunRequestObject) (apigen.CreateSyncRunResponseObject, error) {
	in, err := input[db.CreateSyncRunParams](req.Body, map[string]any{"integration_id": req.IntegrationId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateSyncRun(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.SyncRun](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateSyncRun201JSONResponse{Body: body, Headers: apigen.CreateSyncRun201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListDataImports implements the listDataImports operation.
func (s *Server) ListDataImports(ctx context.Context, req apigen.ListDataImportsRequestObject) (apigen.ListDataImportsResponseObject, error) {
	return listPage[apigen.ListDataImports200JSONResponse](ctx, req.Params, nil, s.service.ListDataImport)
}

// GetDataImport implements the getDataImport operation.
func (s *Server) GetDataImport(ctx context.Context, req apigen.GetDataImportRequestObject) (apigen.GetDataImportResponseObject, error) {
	row, err := s.service.GetDataImport(ctx, req.DataImportId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DataImport](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetDataImport200JSONResponse{Body: body, Headers: apigen.GetDataImport200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateDataImport implements the createDataImport operation.
func (s *Server) CreateDataImport(ctx context.Context, req apigen.CreateDataImportRequestObject) (apigen.CreateDataImportResponseObject, error) {
	in, err := input[db.CreateDataImportParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateDataImport(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DataImport](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateDataImport201JSONResponse{Body: body, Headers: apigen.CreateDataImport201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListPostings implements the listPostings operation.
func (s *Server) ListPostings(ctx context.Context, req apigen.ListPostingsRequestObject) (apigen.ListPostingsResponseObject, error) {
	return listPage[apigen.ListPostings200JSONResponse](ctx, req.Params, nil, s.service.ListPosting)
}

// GetPosting implements the getPosting operation.
func (s *Server) GetPosting(ctx context.Context, req apigen.GetPostingRequestObject) (apigen.GetPostingResponseObject, error) {
	row, err := s.service.GetPosting(ctx, req.PostingId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Posting](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetPosting200JSONResponse{Body: body, Headers: apigen.GetPosting200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// RetryPosting implements the retryPosting operation.
func (s *Server) RetryPosting(ctx context.Context, req apigen.RetryPostingRequestObject) (apigen.RetryPostingResponseObject, error) {
	in, err := input[db.ActRetryPostingParams](struct{}{}, map[string]any{"id": req.PostingId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.RetryPosting(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Posting](row)
	if err != nil {
		return nil, err
	}
	return apigen.RetryPosting200JSONResponse{Body: body, Headers: apigen.RetryPosting200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// MarkPostedPosting implements the markPostedPosting operation.
func (s *Server) MarkPostedPosting(ctx context.Context, req apigen.MarkPostedPostingRequestObject) (apigen.MarkPostedPostingResponseObject, error) {
	in, err := input[db.ActMarkPostedPostingParams](req.Body, map[string]any{"id": req.PostingId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.MarkPostedPosting(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.Posting](row)
	if err != nil {
		return nil, err
	}
	return apigen.MarkPostedPosting200JSONResponse{Body: body, Headers: apigen.MarkPostedPosting200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListPosSalesSummaries implements the listPosSalesSummaries operation.
func (s *Server) ListPosSalesSummaries(ctx context.Context, req apigen.ListPosSalesSummariesRequestObject) (apigen.ListPosSalesSummariesResponseObject, error) {
	return listPage[apigen.ListPosSalesSummaries200JSONResponse](ctx, req.Params, nil, s.service.ListPosSalesSummary)
}

// GetPosSalesSummary implements the getPosSalesSummary operation.
func (s *Server) GetPosSalesSummary(ctx context.Context, req apigen.GetPosSalesSummaryRequestObject) (apigen.GetPosSalesSummaryResponseObject, error) {
	row, err := s.service.GetPosSalesSummary(ctx, req.PosSalesSummaryId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.PosSalesSummary](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetPosSalesSummary200JSONResponse{Body: body, Headers: apigen.GetPosSalesSummary200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListZReports implements the listZReports operation.
func (s *Server) ListZReports(ctx context.Context, req apigen.ListZReportsRequestObject) (apigen.ListZReportsResponseObject, error) {
	return listPage[apigen.ListZReports200JSONResponse](ctx, req.Params, nil, s.service.ListZReport)
}

// GetZReport implements the getZReport operation.
func (s *Server) GetZReport(ctx context.Context, req apigen.GetZReportRequestObject) (apigen.GetZReportResponseObject, error) {
	row, err := s.service.GetZReport(ctx, req.ZReportId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ZReport](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetZReport200JSONResponse{Body: body, Headers: apigen.GetZReport200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateZReport implements the createZReport operation.
func (s *Server) CreateZReport(ctx context.Context, req apigen.CreateZReportRequestObject) (apigen.CreateZReportResponseObject, error) {
	in, err := input[db.CreateZReportParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateZReport(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ZReport](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateZReport201JSONResponse{Body: body, Headers: apigen.CreateZReport201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListKpiDefinitions implements the listKpiDefinitions operation.
func (s *Server) ListKpiDefinitions(ctx context.Context, req apigen.ListKpiDefinitionsRequestObject) (apigen.ListKpiDefinitionsResponseObject, error) {
	return listPage[apigen.ListKpiDefinitions200JSONResponse](ctx, req.Params, nil, s.service.ListKpiDefinition)
}

// GetKpiDefinition implements the getKpiDefinition operation.
func (s *Server) GetKpiDefinition(ctx context.Context, req apigen.GetKpiDefinitionRequestObject) (apigen.GetKpiDefinitionResponseObject, error) {
	row, err := s.service.GetKpiDefinition(ctx, req.KpiDefinitionId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KpiDefinition](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetKpiDefinition200JSONResponse{Body: body, Headers: apigen.GetKpiDefinition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateKpiDefinition implements the createKpiDefinition operation.
func (s *Server) CreateKpiDefinition(ctx context.Context, req apigen.CreateKpiDefinitionRequestObject) (apigen.CreateKpiDefinitionResponseObject, error) {
	in, err := input[db.CreateKpiDefinitionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateKpiDefinition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KpiDefinition](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateKpiDefinition201JSONResponse{Body: body, Headers: apigen.CreateKpiDefinition201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateKpiDefinition implements the updateKpiDefinition operation.
func (s *Server) UpdateKpiDefinition(ctx context.Context, req apigen.UpdateKpiDefinitionRequestObject) (apigen.UpdateKpiDefinitionResponseObject, error) {
	in, err := input[db.UpdateKpiDefinitionParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.KpiDefinitionId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateKpiDefinition(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.KpiDefinition](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateKpiDefinition200JSONResponse{Body: body, Headers: apigen.UpdateKpiDefinition200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListExportJobs implements the listExportJobs operation.
func (s *Server) ListExportJobs(ctx context.Context, req apigen.ListExportJobsRequestObject) (apigen.ListExportJobsResponseObject, error) {
	return listPage[apigen.ListExportJobs200JSONResponse](ctx, req.Params, nil, s.service.ListExportJob)
}

// GetExportJob implements the getExportJob operation.
func (s *Server) GetExportJob(ctx context.Context, req apigen.GetExportJobRequestObject) (apigen.GetExportJobResponseObject, error) {
	row, err := s.service.GetExportJob(ctx, req.ExportJobId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ExportJob](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetExportJob200JSONResponse{Body: body, Headers: apigen.GetExportJob200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateExportJob implements the createExportJob operation.
func (s *Server) CreateExportJob(ctx context.Context, req apigen.CreateExportJobRequestObject) (apigen.CreateExportJobResponseObject, error) {
	in, err := input[db.CreateExportJobParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateExportJob(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.ExportJob](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateExportJob201JSONResponse{Body: body, Headers: apigen.CreateExportJob201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListRetentionClasses implements the listRetentionClasses operation.
func (s *Server) ListRetentionClasses(ctx context.Context, req apigen.ListRetentionClassesRequestObject) (apigen.ListRetentionClassesResponseObject, error) {
	return listPage[apigen.ListRetentionClasses200JSONResponse](ctx, req.Params, nil, s.service.ListRetentionClass)
}

// GetRetentionClass implements the getRetentionClass operation.
func (s *Server) GetRetentionClass(ctx context.Context, req apigen.GetRetentionClassRequestObject) (apigen.GetRetentionClassResponseObject, error) {
	row, err := s.service.GetRetentionClass(ctx, req.RetentionClassId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RetentionClass](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetRetentionClass200JSONResponse{Body: body, Headers: apigen.GetRetentionClass200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateRetentionClass implements the createRetentionClass operation.
func (s *Server) CreateRetentionClass(ctx context.Context, req apigen.CreateRetentionClassRequestObject) (apigen.CreateRetentionClassResponseObject, error) {
	in, err := input[db.CreateRetentionClassParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateRetentionClass(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RetentionClass](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateRetentionClass201JSONResponse{Body: body, Headers: apigen.CreateRetentionClass201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// UpdateRetentionClass implements the updateRetentionClass operation.
func (s *Server) UpdateRetentionClass(ctx context.Context, req apigen.UpdateRetentionClassRequestObject) (apigen.UpdateRetentionClassResponseObject, error) {
	in, err := input[db.UpdateRetentionClassParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	in.ID = req.RetentionClassId
	if in.RowVersion, err = ifMatch(req.Params.IfMatch); err != nil {
		return nil, err
	}
	row, err := s.service.UpdateRetentionClass(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.RetentionClass](row)
	if err != nil {
		return nil, err
	}
	return apigen.UpdateRetentionClass200JSONResponse{Body: body, Headers: apigen.UpdateRetentionClass200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListLegalHolds implements the listLegalHolds operation.
func (s *Server) ListLegalHolds(ctx context.Context, req apigen.ListLegalHoldsRequestObject) (apigen.ListLegalHoldsResponseObject, error) {
	return listPage[apigen.ListLegalHolds200JSONResponse](ctx, req.Params, nil, s.service.ListLegalHold)
}

// GetLegalHold implements the getLegalHold operation.
func (s *Server) GetLegalHold(ctx context.Context, req apigen.GetLegalHoldRequestObject) (apigen.GetLegalHoldResponseObject, error) {
	row, err := s.service.GetLegalHold(ctx, req.LegalHoldId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.LegalHold](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetLegalHold200JSONResponse{Body: body, Headers: apigen.GetLegalHold200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateLegalHold implements the createLegalHold operation.
func (s *Server) CreateLegalHold(ctx context.Context, req apigen.CreateLegalHoldRequestObject) (apigen.CreateLegalHoldResponseObject, error) {
	in, err := input[db.CreateLegalHoldParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateLegalHold(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.LegalHold](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateLegalHold201JSONResponse{Body: body, Headers: apigen.CreateLegalHold201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ReleaseLegalHold implements the releaseLegalHold operation.
func (s *Server) ReleaseLegalHold(ctx context.Context, req apigen.ReleaseLegalHoldRequestObject) (apigen.ReleaseLegalHoldResponseObject, error) {
	in, err := input[db.ActReleaseLegalHoldParams](req.Body, map[string]any{"id": req.LegalHoldId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.ReleaseLegalHold(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.LegalHold](row)
	if err != nil {
		return nil, err
	}
	return apigen.ReleaseLegalHold200JSONResponse{Body: body, Headers: apigen.ReleaseLegalHold200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// ListDispositionBatches implements the listDispositionBatches operation.
func (s *Server) ListDispositionBatches(ctx context.Context, req apigen.ListDispositionBatchesRequestObject) (apigen.ListDispositionBatchesResponseObject, error) {
	return listPage[apigen.ListDispositionBatches200JSONResponse](ctx, req.Params, nil, s.service.ListDispositionBatch)
}

// GetDispositionBatch implements the getDispositionBatch operation.
func (s *Server) GetDispositionBatch(ctx context.Context, req apigen.GetDispositionBatchRequestObject) (apigen.GetDispositionBatchResponseObject, error) {
	row, err := s.service.GetDispositionBatch(ctx, req.DispositionBatchId)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DispositionBatch](row)
	if err != nil {
		return nil, err
	}
	return apigen.GetDispositionBatch200JSONResponse{Body: body, Headers: apigen.GetDispositionBatch200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CreateDispositionBatch implements the createDispositionBatch operation.
func (s *Server) CreateDispositionBatch(ctx context.Context, req apigen.CreateDispositionBatchRequestObject) (apigen.CreateDispositionBatchResponseObject, error) {
	in, err := input[db.CreateDispositionBatchParams](req.Body, nil)
	if err != nil {
		return nil, err
	}
	row, err := s.service.CreateDispositionBatch(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DispositionBatch](row)
	if err != nil {
		return nil, err
	}
	return apigen.CreateDispositionBatch201JSONResponse{Body: body, Headers: apigen.CreateDispositionBatch201ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// SubmitDispositionBatch implements the submitDispositionBatch operation.
func (s *Server) SubmitDispositionBatch(ctx context.Context, req apigen.SubmitDispositionBatchRequestObject) (apigen.SubmitDispositionBatchResponseObject, error) {
	in, err := input[db.ActSubmitDispositionBatchParams](struct{}{}, map[string]any{"id": req.DispositionBatchId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.SubmitDispositionBatch(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DispositionBatch](row)
	if err != nil {
		return nil, err
	}
	return apigen.SubmitDispositionBatch200JSONResponse{Body: body, Headers: apigen.SubmitDispositionBatch200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}

// CompleteDispositionBatch implements the completeDispositionBatch operation.
func (s *Server) CompleteDispositionBatch(ctx context.Context, req apigen.CompleteDispositionBatchRequestObject) (apigen.CompleteDispositionBatchResponseObject, error) {
	in, err := input[db.ActCompleteDispositionBatchParams](struct{}{}, map[string]any{"id": req.DispositionBatchId})
	if err != nil {
		return nil, err
	}
	row, err := s.service.CompleteDispositionBatch(ctx, in)
	if err != nil {
		return nil, err
	}
	body, err := output[apigen.DispositionBatch](row)
	if err != nil {
		return nil, err
	}
	return apigen.CompleteDispositionBatch200JSONResponse{Body: body, Headers: apigen.CompleteDispositionBatch200ResponseHeaders{ETag: etag(row.RowVersion)}}, nil
}
