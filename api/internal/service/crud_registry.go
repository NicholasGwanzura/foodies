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
	"float_certificates": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetFloatCertificate(ctx, id)
		},
		statuses: []string{"draft", "pending_approval", "active", "superseded", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetFloatCertificateStatus(ctx, db.SetFloatCertificateStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"float_issues": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetFloatIssue(ctx, id) },
		statuses: []string{"issued", "acknowledged", "returned", "exception"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetFloatIssueStatus(ctx, db.SetFloatIssueStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"float_exceptions": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetFloatException(ctx, id)
		},
		statuses: []string{"draft", "pending_approval", "approved", "reimbursed", "overdue", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetFloatExceptionStatus(ctx, db.SetFloatExceptionStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"till_flushes": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetTillFlush(ctx, id) },
		statuses: []string{"draft", "pending_signatures", "closed", "exception"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetTillFlushStatus(ctx, db.SetTillFlushStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"cash_ups": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetCashUp(ctx, id) },
		statuses: []string{"draft", "counted", "pending_signatures", "pending_review", "approved", "returned", "exception"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetCashUpStatus(ctx, db.SetCashUpStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"cash_variances": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetCashVariance(ctx, id) },
		statuses: []string{"open", "recovered", "under_investigation", "referred", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetCashVarianceStatus(ctx, db.SetCashVarianceStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"cit_providers": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetCitProvider(ctx, id) },
	},
	"cit_personnel": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetCitPerson(ctx, id) },
	},
	"cit_transits": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetCitTransit(ctx, id) },
		statuses: []string{"draft", "sealed", "identity_verified", "handed_over", "in_transit", "received", "reconciled", "discrepancy", "cancelled"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetCitTransitStatus(ctx, db.SetCitTransitStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"bank_deposits": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetBankDeposit(ctx, id) },
		statuses: []string{"draft", "banked", "proof_submitted", "verified", "exception"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetBankDepositStatus(ctx, db.SetBankDepositStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"cit_reconciliations": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetCitReconciliation(ctx, id)
		},
		statuses: []string{"open", "reconciled", "unresolved"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetCitReconciliationStatus(ctx, db.SetCitReconciliationStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"surprise_counts": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetSurpriseCount(ctx, id) },
		statuses: []string{"planned", "in_progress", "completed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetSurpriseCountStatus(ctx, db.SetSurpriseCountStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"petty_cash_accounts": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetPettyCashAccount(ctx, id)
		},
	},
	"petty_cash_vouchers": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetPettyCashVoucher(ctx, id)
		},
		statuses: []string{"draft", "pending_approval", "approved", "issued", "receipted", "reimbursed", "rejected", "cancelled"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetPettyCashVoucherStatus(ctx, db.SetPettyCashVoucherStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"petty_cash_handovers": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetPettyCashHandover(ctx, id)
		},
		statuses: []string{"draft", "pending_signatures", "closed", "exception"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetPettyCashHandoverStatus(ctx, db.SetPettyCashHandoverStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"petty_cash_reconciliations": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetPettyCashReconciliation(ctx, id)
		},
		statuses: []string{"draft", "submitted", "reviewed", "exception"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetPettyCashReconciliationStatus(ctx, db.SetPettyCashReconciliationStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"petty_cash_reimbursements": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetPettyCashReimbursement(ctx, id)
		},
		statuses: []string{"draft", "submitted", "approved", "paid", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetPettyCashReimbursementStatus(ctx, db.SetPettyCashReimbursementStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"suppliers": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetSupplier(ctx, id) },
		statuses: []string{"draft", "pending_verification", "active", "suspended", "blocked"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetSupplierStatus(ctx, db.SetSupplierStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"supplier_amendments": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetSupplierAmendment(ctx, id)
		},
		statuses: []string{"draft", "pending_approval", "approved", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetSupplierAmendmentStatus(ctx, db.SetSupplierAmendmentStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"purchase_requisitions": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetPurchaseRequisition(ctx, id)
		},
		statuses: []string{"draft", "submitted", "returned", "validated", "pending_approval", "approved", "rejected", "sourcing", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetPurchaseRequisitionStatus(ctx, db.SetPurchaseRequisitionStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"quotations": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetQuotation(ctx, id) },
	},
	"purchase_orders": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetPurchaseOrder(ctx, id) },
		statuses: []string{"draft", "pending_approval", "approved", "issued", "part_received", "received", "closed", "cancelled"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetPurchaseOrderStatus(ctx, db.SetPurchaseOrderStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"goods_receipts": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetGoodsReceipt(ctx, id) },
		statuses: []string{"draft", "inspected", "accepted", "part_rejected", "rejected", "posted"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetGoodsReceiptStatus(ctx, db.SetGoodsReceiptStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"delivery_discrepancies": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetDeliveryDiscrepancy(ctx, id)
		},
		statuses: []string{"open", "supplier_notified", "resolved", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetDeliveryDiscrepancyStatus(ctx, db.SetDeliveryDiscrepancyStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"supplier_invoices": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetSupplierInvoice(ctx, id)
		},
		statuses: []string{"received", "matched", "on_hold", "approved_exception", "posted", "paid", "void"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetSupplierInvoiceStatus(ctx, db.SetSupplierInvoiceStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"payment_requisitions": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetPaymentRequisition(ctx, id)
		},
		statuses: []string{"draft", "pending_approval", "approved", "scheduled", "paid", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetPaymentRequisitionStatus(ctx, db.SetPaymentRequisitionStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"payment_batches": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetPaymentBatch(ctx, id) },
		statuses: []string{"draft", "pending_approval", "approved", "executed", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetPaymentBatchStatus(ctx, db.SetPaymentBatchStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"emergency_procurements": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetEmergencyProcurement(ctx, id)
		},
		statuses: []string{"draft", "pending_approval", "approved", "rejected", "reviewed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetEmergencyProcurementStatus(ctx, db.SetEmergencyProcurementStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"unauthorized_purchases": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetUnauthorizedPurchase(ctx, id)
		},
		statuses: []string{"reported", "under_review", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetUnauthorizedPurchaseStatus(ctx, db.SetUnauthorizedPurchaseStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"budgets": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetBudget(ctx, id) },
		statuses: []string{"draft", "submitted", "consolidated", "pending_approval", "approved", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetBudgetStatus(ctx, db.SetBudgetStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"budget_variances": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetBudgetVariance(ctx, id)
		},
		statuses: []string{"open", "explained", "escalated", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetBudgetVarianceStatus(ctx, db.SetBudgetVarianceStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"budget_reallocations": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetBudgetReallocation(ctx, id)
		},
		statuses: []string{"draft", "pending_approval", "approved", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetBudgetReallocationStatus(ctx, db.SetBudgetReallocationStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"items": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetItem(ctx, id) },
	},
	"stock_balances": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetStockBalance(ctx, id) },
	},
	"stock_requisitions": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetStockRequisition(ctx, id)
		},
		statuses: []string{"draft", "submitted", "approved", "issued", "acknowledged", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetStockRequisitionStatus(ctx, db.SetStockRequisitionStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"stock_transfers": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetStockTransfer(ctx, id) },
		statuses: []string{"requested", "accepted", "rejected", "dispatched", "received", "discrepancy", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetStockTransferStatus(ctx, db.SetStockTransferStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"stock_returns": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetStockReturn(ctx, id) },
		statuses: []string{"open", "followed_up", "credited", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetStockReturnStatus(ctx, db.SetStockReturnStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"stocktakes": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetStocktake(ctx, id) },
		statuses: []string{"planned", "counting", "submitted", "reviewed", "approved", "posted", "skipped"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetStocktakeStatus(ctx, db.SetStocktakeStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"stock_variances": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetStockVariance(ctx, id) },
		statuses: []string{"open", "investigating", "resolved", "escalated", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetStockVarianceStatus(ctx, db.SetStockVarianceStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"stock_adjustments": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetStockAdjustment(ctx, id)
		},
		statuses: []string{"draft", "pending_approval", "approved", "posted", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetStockAdjustmentStatus(ctx, db.SetStockAdjustmentStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"waste_records": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetWasteRecord(ctx, id) },
		statuses: []string{"recorded", "reviewed", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetWasteRecordStatus(ctx, db.SetWasteRecordStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"disposals": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetDisposal(ctx, id) },
		statuses: []string{"draft", "pending_approval", "approved", "disposed", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetDisposalStatus(ctx, db.SetDisposalStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"recipes": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetRecipe(ctx, id) },
	},
	"recipe_versions": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetRecipeVersion(ctx, id) },
		statuses: []string{"draft", "pending_approval", "active", "superseded"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetRecipeVersionStatus(ctx, db.SetRecipeVersionStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"kitchen_orders": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetKitchenOrder(ctx, id) },
		statuses: []string{"draft", "submitted", "late_pending", "confirmed", "planned", "dispatched", "received", "cancelled"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetKitchenOrderStatus(ctx, db.SetKitchenOrderStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"production_plans": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetProductionPlan(ctx, id)
		},
		statuses: []string{"draft", "confirmed", "in_progress", "completed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetProductionPlanStatus(ctx, db.SetProductionPlanStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"batches": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetBatch(ctx, id) },
		statuses: []string{"in_progress", "completed", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetBatchStatus(ctx, db.SetBatchStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"dispatch_notes": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetDispatchNote(ctx, id) },
		statuses: []string{"draft", "picked", "verified", "security_cleared", "in_transit", "delivered", "discrepancy", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetDispatchNoteStatus(ctx, db.SetDispatchNoteStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"products": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetProduct(ctx, id) },
	},
	"product_changes": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetProductChange(ctx, id) },
		statuses: []string{"draft", "pending_approval", "approved", "effective", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetProductChangeStatus(ctx, db.SetProductChangeStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"cost_sheets": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetCostSheet(ctx, id) },
		statuses: []string{"draft", "pending_approval", "approved", "superseded"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetCostSheetStatus(ctx, db.SetCostSheetStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"npd_projects": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetNpdProject(ctx, id) },
		statuses: []string{"active", "on_hold", "launched", "cancelled"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetNpdProjectStatus(ctx, db.SetNpdProjectStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"ncrs": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetNcr(ctx, id) },
		statuses: []string{"open", "contained", "investigating", "action_pending", "pending_verification", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetNcrStatus(ctx, db.SetNcrStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"holds": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetHold(ctx, id) },
		statuses: []string{"on_hold", "pending_disposition", "released", "rework", "rejected", "disposed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetHoldStatus(ctx, db.SetHoldStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"dispositions": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetDisposition(ctx, id) },
		statuses: []string{"draft", "pending_approval", "approved", "executed", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetDispositionStatus(ctx, db.SetDispositionStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"recalls": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetRecall(ctx, id) },
		statuses: []string{"initiated", "held", "assessed", "decided", "retrieving", "reconciling", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetRecallStatus(ctx, db.SetRecallStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"recall_communications": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetRecallCommunication(ctx, id)
		},
		statuses: []string{"draft", "pending_approval", "sent", "acknowledged"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetRecallCommunicationStatus(ctx, db.SetRecallCommunicationStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"recall_retrievals": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetRecallRetrieval(ctx, id)
		},
		statuses: []string{"open", "returned", "no_stock", "unaccounted"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetRecallRetrievalStatus(ctx, db.SetRecallRetrievalStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"recall_reconciliations": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetRecallReconciliation(ctx, id)
		},
		statuses: []string{"draft", "reviewed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetRecallReconciliationStatus(ctx, db.SetRecallReconciliationStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"incidents": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetIncident(ctx, id) },
		statuses: []string{"reported", "acknowledged", "investigating", "action_pending", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetIncidentStatus(ctx, db.SetIncidentStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"complaints": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetComplaint(ctx, id) },
		statuses: []string{"new", "triaged", "investigating", "awaiting_customer", "resolved", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetComplaintStatus(ctx, db.SetComplaintStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"service_recoveries": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetServiceRecovery(ctx, id)
		},
		statuses: []string{"draft", "pending_approval", "approved", "executed", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetServiceRecoveryStatus(ctx, db.SetServiceRecoveryStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"payment_incidents": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetPaymentIncident(ctx, id)
		},
		statuses: []string{"open", "verified", "correction_pending", "correction_approved", "reconciling", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetPaymentIncidentStatus(ctx, db.SetPaymentIncidentStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"pos_outages": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetPosOutage(ctx, id) },
		statuses: []string{"declared", "pending_authorisation", "authorised", "restored", "reconciled", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetPosOutageStatus(ctx, db.SetPosOutageStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"receipt_books": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetReceiptBook(ctx, id) },
		statuses: []string{"in_storage", "issued", "returned", "missing"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetReceiptBookStatus(ctx, db.SetReceiptBookStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"manual_receipts": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetManualReceipt(ctx, id) },
		statuses: []string{"issued", "cancelled", "missing", "back_captured"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetManualReceiptStatus(ctx, db.SetManualReceiptStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"equipment_faults": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetEquipmentFault(ctx, id)
		},
		statuses: []string{"reported", "assigned", "resolved", "cancelled"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetEquipmentFaultStatus(ctx, db.SetEquipmentFaultStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"work_orders": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetWorkOrder(ctx, id) },
		statuses: []string{"open", "assigned", "in_progress", "pending_clearance", "closed", "cancelled"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetWorkOrderStatus(ctx, db.SetWorkOrderStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"maintenance_plans": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetMaintenancePlan(ctx, id)
		},
		statuses: []string{"active", "paused"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetMaintenancePlanStatus(ctx, db.SetMaintenancePlanStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"report_definitions": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetReportDefinition(ctx, id)
		},
	},
	"reports": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetReport(ctx, id) },
		statuses: []string{"pending", "draft", "submitted", "reviewed", "late", "missing"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetReportStatus(ctx, db.SetReportStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"report_deadline_notices": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetReportDeadlineNotice(ctx, id)
		},
	},
	"management_meetings": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetManagementMeeting(ctx, id)
		},
		statuses: []string{"planned", "held"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetManagementMeetingStatus(ctx, db.SetManagementMeetingStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"compliance_obligations": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetComplianceObligation(ctx, id)
		},
		statuses: []string{"upcoming", "due", "overdue", "completed", "cancelled"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetComplianceObligationStatus(ctx, db.SetComplianceObligationStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"asset_classes": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetAssetClass(ctx, id) },
	},
	"assets": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetAsset(ctx, id) },
		statuses: []string{"wip", "active", "out_of_service", "disposed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetAssetStatus(ctx, db.SetAssetStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"capex_requests": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetCapexRequest(ctx, id) },
		statuses: []string{"draft", "pending_approval", "approved", "rejected", "ordered", "capitalised"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetCapexRequestStatus(ctx, db.SetCapexRequestStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"asset_transfers": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetAssetTransfer(ctx, id) },
		statuses: []string{"draft", "pending_approval", "approved", "received", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetAssetTransferStatus(ctx, db.SetAssetTransferStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"asset_verifications": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetAssetVerification(ctx, id)
		},
		statuses: []string{"planned", "in_progress", "submitted", "reviewed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetAssetVerificationStatus(ctx, db.SetAssetVerificationStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"asset_disposals": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetAssetDisposal(ctx, id) },
		statuses: []string{"draft", "pending_approval", "approved", "removed", "completed", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetAssetDisposalStatus(ctx, db.SetAssetDisposalStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"key_items": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetKeyItem(ctx, id) },
		statuses: []string{"active", "in_use", "lost", "faulty", "retired"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetKeyItemStatus(ctx, db.SetKeyItemStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"key_movements": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetKeyMovement(ctx, id) },
		statuses: []string{"out", "returned", "overdue"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetKeyMovementStatus(ctx, db.SetKeyMovementStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"key_duplication_requests": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetKeyDuplicationRequest(ctx, id)
		},
		statuses: []string{"draft", "pending_approval", "approved", "cut", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetKeyDuplicationRequestStatus(ctx, db.SetKeyDuplicationRequestStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"key_audits": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetKeyAudit(ctx, id) },
		statuses: []string{"planned", "in_progress", "completed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetKeyAuditStatus(ctx, db.SetKeyAuditStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"visitor_entries": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetVisitorEntry(ctx, id) },
		statuses: []string{"on_site", "departed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetVisitorEntryStatus(ctx, db.SetVisitorEntryStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"lost_property_items": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetLostPropertyItem(ctx, id)
		},
		statuses: []string{"in_custody", "returned", "disposed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetLostPropertyItemStatus(ctx, db.SetLostPropertyItemStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"documents": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetDocument(ctx, id) },
		statuses: []string{"draft", "in_review", "approved", "issued", "obsolete"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetDocumentStatus(ctx, db.SetDocumentStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"document_revisions": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetDocumentRevision(ctx, id)
		},
		statuses: []string{"draft", "in_review", "qms_review", "pending_approval", "approved", "issued", "withdrawn"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetDocumentRevisionStatus(ctx, db.SetDocumentRevisionStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"document_requests": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetDocumentRequest(ctx, id)
		},
		statuses: []string{"raised", "classified", "drafting", "in_review", "approved", "closed", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetDocumentRequestStatus(ctx, db.SetDocumentRequestStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"document_acknowledgements": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetDocumentAcknowledgement(ctx, id)
		},
		statuses: []string{"pending", "acknowledged", "overdue"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetDocumentAcknowledgementStatus(ctx, db.SetDocumentAcknowledgementStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"courses": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetCourse(ctx, id) },
	},
	"training_requirements": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetTrainingRequirement(ctx, id)
		},
	},
	"training_sessions": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetTrainingSession(ctx, id)
		},
		statuses: []string{"planned", "held", "cancelled"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetTrainingSessionStatus(ctx, db.SetTrainingSessionStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"training_assessments": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetTrainingAssessment(ctx, id)
		},
	},
	"competencies": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetCompetency(ctx, id) },
	},
	"certificates": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetCertificate(ctx, id) },
	},
	"training_requests": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetTrainingRequest(ctx, id)
		},
		statuses: []string{"draft", "pending_approval", "approved", "scheduled", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetTrainingRequestStatus(ctx, db.SetTrainingRequestStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"leave_types": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetLeaveType(ctx, id) },
	},
	"leave_balances": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetLeaveBalance(ctx, id) },
	},
	"leave_requests": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetLeaveRequest(ctx, id) },
		statuses: []string{"draft", "pending_approval", "approved", "rejected", "cancelled", "taken"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetLeaveRequestStatus(ctx, db.SetLeaveRequestStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"exit_clearances": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetExitClearance(ctx, id) },
		statuses: []string{"initiated", "in_clearance", "cleared", "completed", "cancelled"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetExitClearanceStatus(ctx, db.SetExitClearanceStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"rosters": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetRoster(ctx, id) },
		statuses: []string{"draft", "pending_approval", "published"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetRosterStatus(ctx, db.SetRosterStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"attendance_records": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetAttendanceRecord(ctx, id)
		},
	},
	"uniform_issues": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetUniformIssue(ctx, id) },
		statuses: []string{"issued", "returned", "cancelled"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetUniformIssueStatus(ctx, db.SetUniformIssueStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"recruitment_requisitions": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetRecruitmentRequisition(ctx, id)
		},
		statuses: []string{"draft", "pending_approval", "approved", "open", "filled", "cancelled"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetRecruitmentRequisitionStatus(ctx, db.SetRecruitmentRequisitionStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"candidates": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetCandidate(ctx, id) },
		statuses: []string{"applied", "shortlisted", "interviewed", "offered", "hired", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetCandidateStatus(ctx, db.SetCandidateStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"performance_reviews": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetPerformanceReview(ctx, id)
		},
		statuses: []string{"draft", "in_review", "completed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetPerformanceReviewStatus(ctx, db.SetPerformanceReviewStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"disciplinary_cases": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetDisciplinaryCase(ctx, id)
		},
		statuses: []string{"registered", "investigating", "hearing", "decided", "appeal", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetDisciplinaryCaseStatus(ctx, db.SetDisciplinaryCaseStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"grievances": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetGrievance(ctx, id) },
		statuses: []string{"received", "investigating", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetGrievanceStatus(ctx, db.SetGrievanceStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"store_openings": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetStoreOpening(ctx, id) },
		statuses: []string{"in_progress", "no_go", "management_decision", "go", "opened"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetStoreOpeningStatus(ctx, db.SetStoreOpeningStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"occurrence_entries": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetOccurrenceEntry(ctx, id)
		},
	},
	"patrol_routes": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetPatrolRoute(ctx, id) },
	},
	"patrol_rounds": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetPatrolRound(ctx, id) },
		statuses: []string{"in_progress", "completed", "incomplete"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetPatrolRoundStatus(ctx, db.SetPatrolRoundStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"cameras": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetCamera(ctx, id) },
	},
	"cctv_outages": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetCctvOutage(ctx, id) },
		statuses: []string{"open", "restored", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetCctvOutageStatus(ctx, db.SetCctvOutageStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"footage_requests": {
		get: func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) {
			return q.GetFootageRequest(ctx, id)
		},
		statuses: []string{"draft", "pending_approval", "approved", "preserved", "exported", "closed", "rejected"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetFootageRequestStatus(ctx, db.SetFootageRequestStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
	"external_visits": {
		get:      func(ctx context.Context, q db.Querier, id uuid.UUID) (any, error) { return q.GetExternalVisit(ctx, id) },
		statuses: []string{"requested", "screening", "pending_approval", "approved", "held", "cancelled", "closed"},
		setStatus: func(ctx context.Context, q db.Querier, id uuid.UUID, status string, actor *uuid.UUID) error {
			_, err := q.SetExternalVisitStatus(ctx, db.SetExternalVisitStatusParams{ID: id, Status: status, UpdatedBy: actor})
			return err
		},
	},
}
