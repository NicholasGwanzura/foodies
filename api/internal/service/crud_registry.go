// Code scaffolded by api/codegen/scaffold. Do not edit.

package service

import (
	"context"

	"github.com/google/uuid"

	"foodies/api/internal/db"
)

// subjects lets shared services reach any record by table name.
var subjects = map[string]subject{
	"complexes": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetComplex(ctx, id) },
	},
	"sites": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetSite(ctx, id) },
	},
	"departments": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetDepartment(ctx, id) },
	},
	"cost_centres": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetCostCentre(ctx, id) },
	},
	"production_lines": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetProductionLine(ctx, id)
		},
	},
	"tills": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetTill(ctx, id) },
	},
	"safes": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetSafe(ctx, id) },
	},
	"storage_locations": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetStorageLocation(ctx, id)
		},
	},
	"currencies": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetCurrency(ctx, id) },
	},
	"fx_rates": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetFxRate(ctx, id) },
	},
	"public_holidays": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetPublicHoliday(ctx, id) },
	},
	"employees": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetEmployee(ctx, id) },
	},
	"employee_status_changes": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetEmployeeStatusChange(ctx, id)
		},
	},
	"employee_transfers": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetEmployeeTransfer(ctx, id)
		},
		statuses: []string{"draft", "pending_approval", "approved", "rejected", "effective"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetEmployeeTransferStatus(ctx, db.SetEmployeeTransferStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"employee_documents": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetEmployeeDocument(ctx, id)
		},
	},
	"emergency_contacts": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetEmergencyContact(ctx, id)
		},
	},
	"users": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetUser(ctx, id) },
		statuses: []string{"invited", "active", "locked", "disabled"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetUserStatus(ctx, db.SetUserStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"roles": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetRole(ctx, id) },
	},
	"role_assignments": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetRoleAssignment(ctx, id)
		},
	},
	"delegations": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetDelegation(ctx, id) },
		statuses: []string{"draft", "pending_approval", "active", "expired", "revoked"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetDelegationStatus(ctx, db.SetDelegationStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"sod_rules": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetSodRule(ctx, id) },
	},
	"devices": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetDevice(ctx, id) },
		statuses: []string{"active", "revoked"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetDeviceStatus(ctx, db.SetDeviceStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"audit_events": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetAuditEvent(ctx, id) },
	},
	"rule_definitions": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetRuleDefinition(ctx, id)
		},
	},
	"rule_versions": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetRuleVersion(ctx, id) },
		statuses: []string{"draft", "pending_approval", "active", "superseded", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetRuleVersionStatus(ctx, db.SetRuleVersionStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"workflow_activations": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetWorkflowActivation(ctx, id)
		},
		statuses: []string{"not_ready", "ready", "live", "suspended"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetWorkflowActivationStatus(ctx, db.SetWorkflowActivationStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"approval_policies": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetApprovalPolicy(ctx, id)
		},
		statuses: []string{"draft", "pending_approval", "active", "superseded"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetApprovalPolicyStatus(ctx, db.SetApprovalPolicyStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"approval_requests": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetApprovalRequest(ctx, id)
		},
		statuses: []string{"pending", "approved", "rejected", "returned", "cancelled"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetApprovalRequestStatus(ctx, db.SetApprovalRequestStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"signature_requests": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetSignatureRequest(ctx, id)
		},
		statuses: []string{"pending", "signed", "declined", "expired"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetSignatureRequestStatus(ctx, db.SetSignatureRequestStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"attachments": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetAttachment(ctx, id) },
		statuses: []string{"pending", "available", "quarantined"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetAttachmentStatus(ctx, db.SetAttachmentStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"comments": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetComment(ctx, id) },
	},
	"record_links": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetRecordLink(ctx, id) },
	},
	"schedules": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetSchedule(ctx, id) },
		statuses: []string{"active", "paused"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetScheduleStatus(ctx, db.SetScheduleStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"tasks": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetTask(ctx, id) },
		statuses: []string{"open", "in_progress", "done", "skipped", "cancelled"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetTaskStatus(ctx, db.SetTaskStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"form_templates": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetFormTemplate(ctx, id) },
		statuses: []string{"draft", "active", "retired"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetFormTemplateStatus(ctx, db.SetFormTemplateStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"form_template_versions": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetFormTemplateVersion(ctx, id)
		},
		statuses: []string{"draft", "pending_approval", "published", "retired"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetFormTemplateVersionStatus(ctx, db.SetFormTemplateVersionStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"form_submissions": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetFormSubmission(ctx, id)
		},
		statuses: []string{"draft", "submitted", "pending_signature", "pending_review", "closed", "returned"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetFormSubmissionStatus(ctx, db.SetFormSubmissionStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"escalation_policies": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetEscalationPolicy(ctx, id)
		},
		statuses: []string{"draft", "active", "superseded"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetEscalationPolicyStatus(ctx, db.SetEscalationPolicyStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"escalations": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetEscalation(ctx, id) },
		statuses: []string{"open", "acknowledged", "breached", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetEscalationStatus(ctx, db.SetEscalationStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"actions": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetAction(ctx, id) },
		statuses: []string{"open", "in_progress", "pending_verification", "closed", "cancelled"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetActionStatus(ctx, db.SetActionStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"notifications": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetNotification(ctx, id) },
		statuses: []string{"queued", "sent", "delivered", "read", "failed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetNotificationStatus(ctx, db.SetNotificationStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"sync_conflicts": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetSyncConflict(ctx, id) },
		statuses: []string{"open", "resolved"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetSyncConflictStatus(ctx, db.SetSyncConflictStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"whatsapp_identities": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetWhatsappIdentity(ctx, id)
		},
		statuses: []string{"pending", "verified", "revoked"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetWhatsappIdentityStatus(ctx, db.SetWhatsappIdentityStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"whatsapp_templates": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetWhatsappTemplate(ctx, id)
		},
	},
	"whatsapp_messages": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetWhatsappMessage(ctx, id)
		},
	},
	"integrations": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetIntegration(ctx, id) },
		statuses: []string{"disabled", "active", "degraded"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetIntegrationStatus(ctx, db.SetIntegrationStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"sync_runs": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetSyncRun(ctx, id) },
		statuses: []string{"running", "succeeded", "partial", "failed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetSyncRunStatus(ctx, db.SetSyncRunStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"data_imports": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetDataImport(ctx, id) },
		statuses: []string{"uploaded", "validating", "applied", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetDataImportStatus(ctx, db.SetDataImportStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"postings": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetPosting(ctx, id) },
		statuses: []string{"pending", "sent", "posted", "failed", "manual_required", "cancelled"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetPostingStatus(ctx, db.SetPostingStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"pos_sales_summaries": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetPosSalesSummary(ctx, id)
		},
	},
	"z_reports": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetZReport(ctx, id) },
	},
	"kpi_definitions": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetKpiDefinition(ctx, id) },
	},
	"export_jobs": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetExportJob(ctx, id) },
		statuses: []string{"queued", "running", "ready", "failed", "expired"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetExportJobStatus(ctx, db.SetExportJobStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"retention_classes": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetRetentionClass(ctx, id)
		},
		statuses: []string{"unset", "draft", "active", "superseded"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetRetentionClassStatus(ctx, db.SetRetentionClassStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"legal_holds": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetLegalHold(ctx, id) },
		statuses: []string{"active", "released"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetLegalHoldStatus(ctx, db.SetLegalHoldStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"disposition_batches": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetDispositionBatch(ctx, id)
		},
		statuses: []string{"proposed", "pending_approval", "approved", "completed", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetDispositionBatchStatus(ctx, db.SetDispositionBatchStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
}
