// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package service

import (
	"context"

	"github.com/google/uuid"

	"foodies/api/internal/db"
)

var _ uuid.UUID

// ListSyncConflict returns a page of sync conflicts.
func (s *Service) ListSyncConflict(ctx context.Context, p db.ListSyncConflictParams) ([]db.SyncConflict, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListSyncConflict(ctx, p))("list sync conflicts")
}

// GetSyncConflict returns one row of sync conflicts.
func (s *Service) GetSyncConflict(ctx context.Context, id uuid.UUID) (db.SyncConflict, error) {
	row, err := s.q(ctx).GetSyncConflict(ctx, id)
	if err != nil {
		return row, notFound(err, "sync conflicts")
	}
	return row, inScope(ctx, row.SiteID)
}

// ResolveSyncConflict is the accept, discard or replace the offline write action on sync conflicts.
func (s *Service) ResolveSyncConflict(ctx context.Context, p db.ActResolveSyncConflictParams) (db.SyncConflict, error) {
	before, err := s.GetSyncConflict(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "sync.sync_conflict.resolve", "sync_conflicts", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ResolveSyncConflict", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActResolveSyncConflict(ctx, p)
	if err != nil {
		return before, stateError(err, "resolve sync conflicts")
	}
	if err := s.audit(ctx, "sync.sync_conflict.resolve", "sync_conflicts", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ResolveSyncConflict", row)
}

// ListWhatsappIdentity returns a page of whatsapp identities.
func (s *Service) ListWhatsappIdentity(ctx context.Context, p db.ListWhatsappIdentityParams) ([]db.WhatsappIdentity, error) {
	return wrap2(s.q(ctx).ListWhatsappIdentity(ctx, p))("list whatsapp identities")
}

// GetWhatsappIdentity returns one row of whatsapp identities.
func (s *Service) GetWhatsappIdentity(ctx context.Context, id uuid.UUID) (db.WhatsappIdentity, error) {
	row, err := s.q(ctx).GetWhatsappIdentity(ctx, id)
	if err != nil {
		return row, notFound(err, "whatsapp identities")
	}
	return row, nil
}

// CreateWhatsappIdentity adds one row to whatsapp identities.
func (s *Service) CreateWhatsappIdentity(ctx context.Context, p db.CreateWhatsappIdentityParams) (db.WhatsappIdentity, error) {
	var zero db.WhatsappIdentity
	if err := s.guard(ctx, "CreateWhatsappIdentity", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateWhatsappIdentity(ctx, p)
	if err != nil {
		return zero, wrap(err, "create whatsapp identities")
	}
	if err := s.audit(ctx, "whatsapp.whatsapp_identity.write", "whatsapp_identities", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateWhatsappIdentity", row)
}

// RevokeWhatsappIdentity is the unbind the number now action on whatsapp identities.
func (s *Service) RevokeWhatsappIdentity(ctx context.Context, p db.ActRevokeWhatsappIdentityParams) (db.WhatsappIdentity, error) {
	before, err := s.GetWhatsappIdentity(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "whatsapp.whatsapp_identity.revoke", "whatsapp_identities", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "RevokeWhatsappIdentity", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActRevokeWhatsappIdentity(ctx, p)
	if err != nil {
		return before, stateError(err, "revoke whatsapp identities")
	}
	if err := s.audit(ctx, "whatsapp.whatsapp_identity.revoke", "whatsapp_identities", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "RevokeWhatsappIdentity", row)
}

// RebindWhatsappIdentity is the move to a new number action on whatsapp identities.
func (s *Service) RebindWhatsappIdentity(ctx context.Context, p db.ActRebindWhatsappIdentityParams) (db.WhatsappIdentity, error) {
	before, err := s.GetWhatsappIdentity(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "whatsapp.whatsapp_identity.rebind", "whatsapp_identities", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "RebindWhatsappIdentity", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActRebindWhatsappIdentity(ctx, p)
	if err != nil {
		return before, stateError(err, "rebind whatsapp identities")
	}
	if err := s.audit(ctx, "whatsapp.whatsapp_identity.rebind", "whatsapp_identities", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "RebindWhatsappIdentity", row)
}

// ListWhatsappTemplate returns a page of whatsapp templates.
func (s *Service) ListWhatsappTemplate(ctx context.Context, p db.ListWhatsappTemplateParams) ([]db.WhatsappTemplate, error) {
	return wrap2(s.q(ctx).ListWhatsappTemplate(ctx, p))("list whatsapp templates")
}

// GetWhatsappTemplate returns one row of whatsapp templates.
func (s *Service) GetWhatsappTemplate(ctx context.Context, id uuid.UUID) (db.WhatsappTemplate, error) {
	row, err := s.q(ctx).GetWhatsappTemplate(ctx, id)
	if err != nil {
		return row, notFound(err, "whatsapp templates")
	}
	return row, nil
}

// CreateWhatsappTemplate adds one row to whatsapp templates.
func (s *Service) CreateWhatsappTemplate(ctx context.Context, p db.CreateWhatsappTemplateParams) (db.WhatsappTemplate, error) {
	var zero db.WhatsappTemplate
	if err := s.guard(ctx, "CreateWhatsappTemplate", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateWhatsappTemplate(ctx, p)
	if err != nil {
		return zero, wrap(err, "create whatsapp templates")
	}
	if err := s.audit(ctx, "whatsapp.whatsapp_template.write", "whatsapp_templates", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateWhatsappTemplate", row)
}

// UpdateWhatsappTemplate changes one row of whatsapp templates.
func (s *Service) UpdateWhatsappTemplate(ctx context.Context, p db.UpdateWhatsappTemplateParams) (db.WhatsappTemplate, error) {
	before, err := s.GetWhatsappTemplate(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateWhatsappTemplate(ctx, p)
	if err != nil {
		return before, wrap(err, "update whatsapp templates")
	}
	if err := s.voidSignatures(ctx, "whatsapp_templates", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "whatsapp.whatsapp_template.write", "whatsapp_templates", row.ID, nil, before, row)
}

// ListWhatsappMessage returns a page of whatsapp messages.
func (s *Service) ListWhatsappMessage(ctx context.Context, p db.ListWhatsappMessageParams) ([]db.WhatsappMessage, error) {
	return wrap2(s.q(ctx).ListWhatsappMessage(ctx, p))("list whatsapp messages")
}

// GetWhatsappMessage returns one row of whatsapp messages.
func (s *Service) GetWhatsappMessage(ctx context.Context, id uuid.UUID) (db.WhatsappMessage, error) {
	row, err := s.q(ctx).GetWhatsappMessage(ctx, id)
	if err != nil {
		return row, notFound(err, "whatsapp messages")
	}
	return row, nil
}

// ListIntegration returns a page of integrations.
func (s *Service) ListIntegration(ctx context.Context, p db.ListIntegrationParams) ([]db.Integration, error) {
	return wrap2(s.q(ctx).ListIntegration(ctx, p))("list integrations")
}

// GetIntegration returns one row of integrations.
func (s *Service) GetIntegration(ctx context.Context, id uuid.UUID) (db.Integration, error) {
	row, err := s.q(ctx).GetIntegration(ctx, id)
	if err != nil {
		return row, notFound(err, "integrations")
	}
	return row, nil
}

// CreateIntegration adds one row to integrations.
func (s *Service) CreateIntegration(ctx context.Context, p db.CreateIntegrationParams) (db.Integration, error) {
	var zero db.Integration
	if err := s.guard(ctx, "CreateIntegration", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateIntegration(ctx, p)
	if err != nil {
		return zero, wrap(err, "create integrations")
	}
	if err := s.audit(ctx, "integrations.integration.write", "integrations", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateIntegration", row)
}

// UpdateIntegration changes one row of integrations.
func (s *Service) UpdateIntegration(ctx context.Context, p db.UpdateIntegrationParams) (db.Integration, error) {
	before, err := s.GetIntegration(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateIntegration(ctx, p)
	if err != nil {
		return before, wrap(err, "update integrations")
	}
	if err := s.voidSignatures(ctx, "integrations", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "integrations.integration.write", "integrations", row.ID, nil, before, row)
}

// ListSyncRun returns a page of sync runs.
func (s *Service) ListSyncRun(ctx context.Context, p db.ListSyncRunParams) ([]db.SyncRun, error) {
	return wrap2(s.q(ctx).ListSyncRun(ctx, p))("list sync runs")
}

// GetSyncRun returns one row of sync runs.
func (s *Service) GetSyncRun(ctx context.Context, id uuid.UUID) (db.SyncRun, error) {
	row, err := s.q(ctx).GetSyncRun(ctx, id)
	if err != nil {
		return row, notFound(err, "sync runs")
	}
	return row, nil
}

// CreateSyncRun adds one row to sync runs.
func (s *Service) CreateSyncRun(ctx context.Context, p db.CreateSyncRunParams) (db.SyncRun, error) {
	var zero db.SyncRun
	if err := s.guard(ctx, "CreateSyncRun", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateSyncRun(ctx, p)
	if err != nil {
		return zero, wrap(err, "create sync runs")
	}
	if err := s.audit(ctx, "integrations.sync_run.write", "sync_runs", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateSyncRun", row)
}

// ListDataImport returns a page of data imports.
func (s *Service) ListDataImport(ctx context.Context, p db.ListDataImportParams) ([]db.DataImport, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListDataImport(ctx, p))("list data imports")
}

// GetDataImport returns one row of data imports.
func (s *Service) GetDataImport(ctx context.Context, id uuid.UUID) (db.DataImport, error) {
	row, err := s.q(ctx).GetDataImport(ctx, id)
	if err != nil {
		return row, notFound(err, "data imports")
	}
	return row, inScope(ctx, row.SiteID)
}

// CreateDataImport adds one row to data imports.
func (s *Service) CreateDataImport(ctx context.Context, p db.CreateDataImportParams) (db.DataImport, error) {
	var zero db.DataImport
	if err := canWrite(ctx, p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateDataImport", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateDataImport(ctx, p)
	if err != nil {
		return zero, wrap(err, "create data imports")
	}
	if err := s.audit(ctx, "integrations.data_import.write", "data_imports", row.ID, p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateDataImport", row)
}

// ListPosting returns a page of postings.
func (s *Service) ListPosting(ctx context.Context, p db.ListPostingParams) ([]db.Posting, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListPosting(ctx, p))("list postings")
}

// GetPosting returns one row of postings.
func (s *Service) GetPosting(ctx context.Context, id uuid.UUID) (db.Posting, error) {
	row, err := s.q(ctx).GetPosting(ctx, id)
	if err != nil {
		return row, notFound(err, "postings")
	}
	return row, inScope(ctx, row.SiteID)
}

// MarkPostedPosting is the record a manual posting action on postings.
func (s *Service) MarkPostedPosting(ctx context.Context, p db.ActMarkPostedPostingParams) (db.Posting, error) {
	before, err := s.GetPosting(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "integrations.posting.mark_posted", "postings", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "MarkPostedPosting", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActMarkPostedPosting(ctx, p)
	if err != nil {
		return before, stateError(err, "mark posted postings")
	}
	if err := s.audit(ctx, "integrations.posting.mark_posted", "postings", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "MarkPostedPosting", row)
}

// ListPosSalesSummary returns a page of pos sales summaries.
func (s *Service) ListPosSalesSummary(ctx context.Context, p db.ListPosSalesSummaryParams) ([]db.PosSalesSummary, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListPosSalesSummary(ctx, p))("list pos sales summaries")
}

// GetPosSalesSummary returns one row of pos sales summaries.
func (s *Service) GetPosSalesSummary(ctx context.Context, id uuid.UUID) (db.PosSalesSummary, error) {
	row, err := s.q(ctx).GetPosSalesSummary(ctx, id)
	if err != nil {
		return row, notFound(err, "pos sales summaries")
	}
	return row, inScope(ctx, row.SiteID)
}

// ListZReport returns a page of z reports.
func (s *Service) ListZReport(ctx context.Context, p db.ListZReportParams) ([]db.ZReport, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListZReport(ctx, p))("list z reports")
}

// GetZReport returns one row of z reports.
func (s *Service) GetZReport(ctx context.Context, id uuid.UUID) (db.ZReport, error) {
	row, err := s.q(ctx).GetZReport(ctx, id)
	if err != nil {
		return row, notFound(err, "z reports")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateZReport adds one row to z reports.
func (s *Service) CreateZReport(ctx context.Context, p db.CreateZReportParams) (db.ZReport, error) {
	var zero db.ZReport
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.guard(ctx, "CreateZReport", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateZReport(ctx, p)
	if err != nil {
		return zero, wrap(err, "create z reports")
	}
	if err := s.audit(ctx, "integrations.z_report.write", "z_reports", row.ID, &p.SiteID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateZReport", row)
}

// ListKpiDefinition returns a page of kpi definitions.
func (s *Service) ListKpiDefinition(ctx context.Context, p db.ListKpiDefinitionParams) ([]db.KpiDefinition, error) {
	return wrap2(s.q(ctx).ListKpiDefinition(ctx, p))("list kpi definitions")
}

// GetKpiDefinition returns one row of kpi definitions.
func (s *Service) GetKpiDefinition(ctx context.Context, id uuid.UUID) (db.KpiDefinition, error) {
	row, err := s.q(ctx).GetKpiDefinition(ctx, id)
	if err != nil {
		return row, notFound(err, "kpi definitions")
	}
	return row, nil
}

// CreateKpiDefinition adds one row to kpi definitions.
func (s *Service) CreateKpiDefinition(ctx context.Context, p db.CreateKpiDefinitionParams) (db.KpiDefinition, error) {
	var zero db.KpiDefinition
	if err := s.guard(ctx, "CreateKpiDefinition", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateKpiDefinition(ctx, p)
	if err != nil {
		return zero, wrap(err, "create kpi definitions")
	}
	if err := s.audit(ctx, "reporting.kpi_definition.write", "kpi_definitions", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateKpiDefinition", row)
}

// UpdateKpiDefinition changes one row of kpi definitions.
func (s *Service) UpdateKpiDefinition(ctx context.Context, p db.UpdateKpiDefinitionParams) (db.KpiDefinition, error) {
	before, err := s.GetKpiDefinition(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateKpiDefinition(ctx, p)
	if err != nil {
		return before, wrap(err, "update kpi definitions")
	}
	if err := s.voidSignatures(ctx, "kpi_definitions", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "reporting.kpi_definition.write", "kpi_definitions", row.ID, nil, before, row)
}

// ListExportJob returns a page of export jobs.
func (s *Service) ListExportJob(ctx context.Context, p db.ListExportJobParams) ([]db.ExportJob, error) {
	return wrap2(s.q(ctx).ListExportJob(ctx, p))("list export jobs")
}

// GetExportJob returns one row of export jobs.
func (s *Service) GetExportJob(ctx context.Context, id uuid.UUID) (db.ExportJob, error) {
	row, err := s.q(ctx).GetExportJob(ctx, id)
	if err != nil {
		return row, notFound(err, "export jobs")
	}
	return row, nil
}

// CreateExportJob adds one row to export jobs.
func (s *Service) CreateExportJob(ctx context.Context, p db.CreateExportJobParams) (db.ExportJob, error) {
	var zero db.ExportJob
	if err := s.guard(ctx, "CreateExportJob", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateExportJob(ctx, p)
	if err != nil {
		return zero, wrap(err, "create export jobs")
	}
	if err := s.audit(ctx, "reporting.export_job.write", "export_jobs", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateExportJob", row)
}

// ListRetentionClass returns a page of retention classes.
func (s *Service) ListRetentionClass(ctx context.Context, p db.ListRetentionClassParams) ([]db.RetentionClass, error) {
	return wrap2(s.q(ctx).ListRetentionClass(ctx, p))("list retention classes")
}

// GetRetentionClass returns one row of retention classes.
func (s *Service) GetRetentionClass(ctx context.Context, id uuid.UUID) (db.RetentionClass, error) {
	row, err := s.q(ctx).GetRetentionClass(ctx, id)
	if err != nil {
		return row, notFound(err, "retention classes")
	}
	return row, nil
}

// CreateRetentionClass adds one row to retention classes.
func (s *Service) CreateRetentionClass(ctx context.Context, p db.CreateRetentionClassParams) (db.RetentionClass, error) {
	var zero db.RetentionClass
	if err := s.guard(ctx, "CreateRetentionClass", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateRetentionClass(ctx, p)
	if err != nil {
		return zero, wrap(err, "create retention classes")
	}
	if err := s.audit(ctx, "retention.retention_class.write", "retention_classes", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateRetentionClass", row)
}

// UpdateRetentionClass changes one row of retention classes.
func (s *Service) UpdateRetentionClass(ctx context.Context, p db.UpdateRetentionClassParams) (db.RetentionClass, error) {
	before, err := s.GetRetentionClass(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateRetentionClass(ctx, p)
	if err != nil {
		return before, wrap(err, "update retention classes")
	}
	if err := s.voidSignatures(ctx, "retention_classes", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "retention.retention_class.write", "retention_classes", row.ID, nil, before, row)
}

// ListLegalHold returns a page of legal holds.
func (s *Service) ListLegalHold(ctx context.Context, p db.ListLegalHoldParams) ([]db.LegalHold, error) {
	return wrap2(s.q(ctx).ListLegalHold(ctx, p))("list legal holds")
}

// GetLegalHold returns one row of legal holds.
func (s *Service) GetLegalHold(ctx context.Context, id uuid.UUID) (db.LegalHold, error) {
	row, err := s.q(ctx).GetLegalHold(ctx, id)
	if err != nil {
		return row, notFound(err, "legal holds")
	}
	return row, nil
}

// CreateLegalHold adds one row to legal holds.
func (s *Service) CreateLegalHold(ctx context.Context, p db.CreateLegalHoldParams) (db.LegalHold, error) {
	var zero db.LegalHold
	if err := s.guard(ctx, "CreateLegalHold", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateLegalHold(ctx, p)
	if err != nil {
		return zero, wrap(err, "create legal holds")
	}
	if err := s.audit(ctx, "retention.legal_hold.write", "legal_holds", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateLegalHold", row)
}

// ReleaseLegalHold is the release the hold action on legal holds.
func (s *Service) ReleaseLegalHold(ctx context.Context, p db.ActReleaseLegalHoldParams) (db.LegalHold, error) {
	before, err := s.GetLegalHold(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "retention.legal_hold.release", "legal_holds", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "ReleaseLegalHold", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActReleaseLegalHold(ctx, p)
	if err != nil {
		return before, stateError(err, "release legal holds")
	}
	if err := s.audit(ctx, "retention.legal_hold.release", "legal_holds", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "ReleaseLegalHold", row)
}

// ListDispositionBatch returns a page of disposition batches.
func (s *Service) ListDispositionBatch(ctx context.Context, p db.ListDispositionBatchParams) ([]db.DispositionBatch, error) {
	return wrap2(s.q(ctx).ListDispositionBatch(ctx, p))("list disposition batches")
}

// GetDispositionBatch returns one row of disposition batches.
func (s *Service) GetDispositionBatch(ctx context.Context, id uuid.UUID) (db.DispositionBatch, error) {
	row, err := s.q(ctx).GetDispositionBatch(ctx, id)
	if err != nil {
		return row, notFound(err, "disposition batches")
	}
	return row, nil
}

// CreateDispositionBatch adds one row to disposition batches.
func (s *Service) CreateDispositionBatch(ctx context.Context, p db.CreateDispositionBatchParams) (db.DispositionBatch, error) {
	var zero db.DispositionBatch
	if err := s.guard(ctx, "CreateDispositionBatch", &p); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateDispositionBatch(ctx, p)
	if err != nil {
		return zero, wrap(err, "create disposition batches")
	}
	if err := s.audit(ctx, "retention.disposition_batch.write", "disposition_batches", row.ID, nil, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CreateDispositionBatch", row)
}

// SubmitDispositionBatch is the submit for approval action on disposition batches.
func (s *Service) SubmitDispositionBatch(ctx context.Context, p db.ActSubmitDispositionBatchParams) (db.DispositionBatch, error) {
	before, err := s.GetDispositionBatch(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "retention.disposition_batch.submit", "disposition_batches", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "SubmitDispositionBatch", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitDispositionBatch(ctx, p)
	if err != nil {
		return before, stateError(err, "submit disposition batches")
	}
	if err := s.audit(ctx, "retention.disposition_batch.submit", "disposition_batches", row.ID, nil, before, row); err != nil {
		return row, err
	}
	if err := s.startApproval(ctx, "disposition_batches", row.ID, nil, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "SubmitDispositionBatch", row)
}

// CompleteDispositionBatch is the record completion action on disposition batches.
func (s *Service) CompleteDispositionBatch(ctx context.Context, p db.ActCompleteDispositionBatchParams) (db.DispositionBatch, error) {
	before, err := s.GetDispositionBatch(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "retention.disposition_batch.complete", "disposition_batches", p.ID); err != nil {
		return before, err
	}
	if err := s.guard(ctx, "CompleteDispositionBatch", before); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "disposition_batches", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCompleteDispositionBatch(ctx, p)
	if err != nil {
		return before, stateError(err, "complete disposition batches")
	}
	if err := s.audit(ctx, "retention.disposition_batch.complete", "disposition_batches", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, s.effect(ctx, "CompleteDispositionBatch", row)
}
