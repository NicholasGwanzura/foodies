// Code scaffolded by api/codegen/scaffold. Do not edit: write the function in
// another file of this package and this copy disappears on the next run.

package service

import (
	"context"

	"github.com/google/uuid"

	"foodies/api/internal/db"
)

var _ uuid.UUID

// ListFloatCertificate returns a page of float certificates.
func (s *Service) ListFloatCertificate(ctx context.Context, p db.ListFloatCertificateParams) ([]db.FloatCertificate, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListFloatCertificate(ctx, p))("list float certificates")
}

// GetFloatCertificate returns one row of float certificates.
func (s *Service) GetFloatCertificate(ctx context.Context, id uuid.UUID) (db.FloatCertificate, error) {
	row, err := s.q(ctx).GetFloatCertificate(ctx, id)
	if err != nil {
		return row, notFound(err, "float certificates")
	}
	return row, inScope(ctx, &row.SiteID)
}

// SubmitFloatCertificate is the submit action on float certificates.
func (s *Service) SubmitFloatCertificate(ctx context.Context, p db.ActSubmitFloatCertificateParams) (db.FloatCertificate, error) {
	before, err := s.GetFloatCertificate(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cash.float_certificate.submit", "float_certificates", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitFloatCertificate(ctx, p)
	if err != nil {
		return before, stateError(err, "submit float certificates")
	}
	if err := s.audit(ctx, "cash.float_certificate.submit", "float_certificates", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.startApproval(ctx, "float_certificates", row.ID, &row.SiteID, row)
}

// ListFloatIssue returns a page of float issues.
func (s *Service) ListFloatIssue(ctx context.Context, p db.ListFloatIssueParams) ([]db.FloatIssue, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListFloatIssue(ctx, p))("list float issues")
}

// GetFloatIssue returns one row of float issues.
func (s *Service) GetFloatIssue(ctx context.Context, id uuid.UUID) (db.FloatIssue, error) {
	row, err := s.q(ctx).GetFloatIssue(ctx, id)
	if err != nil {
		return row, notFound(err, "float issues")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreateFloatIssue adds one row to float issues.
func (s *Service) CreateFloatIssue(ctx context.Context, p db.CreateFloatIssueParams) (db.FloatIssue, error) {
	var zero db.FloatIssue
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "FLI", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference = ref
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateFloatIssue(ctx, p)
	if err != nil {
		return zero, wrap(err, "create float issues")
	}
	return row, s.audit(ctx, "cash.float_issue.write", "float_issues", row.ID, &p.SiteID, nil, row)
}

// AcknowledgeFloatIssue is the cashier acknowledges the float action on float issues.
func (s *Service) AcknowledgeFloatIssue(ctx context.Context, p db.ActAcknowledgeFloatIssueParams) (db.FloatIssue, error) {
	before, err := s.GetFloatIssue(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cash.float_issue.acknowledge", "float_issues", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActAcknowledgeFloatIssue(ctx, p)
	if err != nil {
		return before, stateError(err, "acknowledge float issues")
	}
	if err := s.audit(ctx, "cash.float_issue.acknowledge", "float_issues", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, nil
}

// ReturnFloatIssue is the record float returned action on float issues.
func (s *Service) ReturnFloatIssue(ctx context.Context, p db.ActReturnFloatIssueParams) (db.FloatIssue, error) {
	before, err := s.GetFloatIssue(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cash.float_issue.return", "float_issues", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActReturnFloatIssue(ctx, p)
	if err != nil {
		return before, stateError(err, "return float issues")
	}
	if err := s.audit(ctx, "cash.float_issue.return", "float_issues", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, nil
}

// ListFloatException returns a page of float exceptions.
func (s *Service) ListFloatException(ctx context.Context, p db.ListFloatExceptionParams) ([]db.FloatException, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListFloatException(ctx, p))("list float exceptions")
}

// GetFloatException returns one row of float exceptions.
func (s *Service) GetFloatException(ctx context.Context, id uuid.UUID) (db.FloatException, error) {
	row, err := s.q(ctx).GetFloatException(ctx, id)
	if err != nil {
		return row, notFound(err, "float exceptions")
	}
	return row, inScope(ctx, &row.SiteID)
}

// SubmitFloatException is the submit action on float exceptions.
func (s *Service) SubmitFloatException(ctx context.Context, p db.ActSubmitFloatExceptionParams) (db.FloatException, error) {
	before, err := s.GetFloatException(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cash.float_exception.submit", "float_exceptions", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitFloatException(ctx, p)
	if err != nil {
		return before, stateError(err, "submit float exceptions")
	}
	if err := s.audit(ctx, "cash.float_exception.submit", "float_exceptions", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.startApproval(ctx, "float_exceptions", row.ID, &row.SiteID, row)
}

// ReimburseFloatException is the record full reimbursement action on float exceptions.
func (s *Service) ReimburseFloatException(ctx context.Context, p db.ActReimburseFloatExceptionParams) (db.FloatException, error) {
	before, err := s.GetFloatException(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cash.float_exception.reimburse", "float_exceptions", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActReimburseFloatException(ctx, p)
	if err != nil {
		return before, stateError(err, "reimburse float exceptions")
	}
	if err := s.audit(ctx, "cash.float_exception.reimburse", "float_exceptions", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, nil
}

// ListTillFlush returns a page of till flushes.
func (s *Service) ListTillFlush(ctx context.Context, p db.ListTillFlushParams) ([]db.TillFlush, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListTillFlush(ctx, p))("list till flushes")
}

// GetTillFlush returns one row of till flushes.
func (s *Service) GetTillFlush(ctx context.Context, id uuid.UUID) (db.TillFlush, error) {
	row, err := s.q(ctx).GetTillFlush(ctx, id)
	if err != nil {
		return row, notFound(err, "till flushes")
	}
	return row, inScope(ctx, &row.SiteID)
}

// UpdateTillFlush changes one row of till flushes.
func (s *Service) UpdateTillFlush(ctx context.Context, p db.UpdateTillFlushParams) (db.TillFlush, error) {
	before, err := s.GetTillFlush(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateTillFlush(ctx, p)
	if err != nil {
		return before, wrap(err, "update till flushes")
	}
	if err := s.voidSignatures(ctx, "till_flushes", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "cash.till_flush.write", "till_flushes", row.ID, &row.SiteID, before, row)
}

// ListCashUp returns a page of cash ups.
func (s *Service) ListCashUp(ctx context.Context, p db.ListCashUpParams) ([]db.CashUp, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListCashUp(ctx, p))("list cash ups")
}

// GetCashUp returns one row of cash ups.
func (s *Service) GetCashUp(ctx context.Context, id uuid.UUID) (db.CashUp, error) {
	row, err := s.q(ctx).GetCashUp(ctx, id)
	if err != nil {
		return row, notFound(err, "cash ups")
	}
	return row, inScope(ctx, &row.SiteID)
}

// UpdateCashUp changes one row of cash ups.
func (s *Service) UpdateCashUp(ctx context.Context, p db.UpdateCashUpParams) (db.CashUp, error) {
	before, err := s.GetCashUp(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateCashUp(ctx, p)
	if err != nil {
		return before, wrap(err, "update cash ups")
	}
	if err := s.voidSignatures(ctx, "cash_ups", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "cash.cash_up.write", "cash_ups", row.ID, &row.SiteID, before, row)
}

// ReturnCashUp is the return for correction action on cash ups.
func (s *Service) ReturnCashUp(ctx context.Context, p db.ActReturnCashUpParams) (db.CashUp, error) {
	before, err := s.GetCashUp(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cash.cash_up.return", "cash_ups", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActReturnCashUp(ctx, p)
	if err != nil {
		return before, stateError(err, "return cash ups")
	}
	if err := s.audit(ctx, "cash.cash_up.return", "cash_ups", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, nil
}

// CashUpChildren returns the rows of a record's child tables, by API field.
func (s *Service) CashUpChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListCashUpTenders(ctx, id)
		if err != nil {
			return nil, wrap(err, "list cash_up_tenders")
		}
		out["tenders"] = childRows(rows)
	}
	return out, nil
}

// ListCashVariance returns a page of cash variances.
func (s *Service) ListCashVariance(ctx context.Context, p db.ListCashVarianceParams) ([]db.CashVariance, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListCashVariance(ctx, p))("list cash variances")
}

// GetCashVariance returns one row of cash variances.
func (s *Service) GetCashVariance(ctx context.Context, id uuid.UUID) (db.CashVariance, error) {
	row, err := s.q(ctx).GetCashVariance(ctx, id)
	if err != nil {
		return row, notFound(err, "cash variances")
	}
	return row, inScope(ctx, row.SiteID)
}

// UpdateCashVariance changes one row of cash variances.
func (s *Service) UpdateCashVariance(ctx context.Context, p db.UpdateCashVarianceParams) (db.CashVariance, error) {
	before, err := s.GetCashVariance(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateCashVariance(ctx, p)
	if err != nil {
		return before, wrap(err, "update cash variances")
	}
	if err := s.voidSignatures(ctx, "cash_variances", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "cash.cash_variance.write", "cash_variances", row.ID, row.SiteID, before, row)
}

// RecordRecoveryCashVariance is the record immediate payment of the variance action on cash variances.
func (s *Service) RecordRecoveryCashVariance(ctx context.Context, p db.ActRecordRecoveryCashVarianceParams) (db.CashVariance, error) {
	before, err := s.GetCashVariance(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cash.cash_variance.record_recovery", "cash_variances", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActRecordRecoveryCashVariance(ctx, p)
	if err != nil {
		return before, stateError(err, "record recovery cash variances")
	}
	if err := s.audit(ctx, "cash.cash_variance.record_recovery", "cash_variances", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, nil
}

// ReferHrCashVariance is the refer to hr for the applicable process action on cash variances.
func (s *Service) ReferHrCashVariance(ctx context.Context, p db.ActReferHrCashVarianceParams) (db.CashVariance, error) {
	before, err := s.GetCashVariance(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cash.cash_variance.refer_hr", "cash_variances", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActReferHrCashVariance(ctx, p)
	if err != nil {
		return before, stateError(err, "refer hr cash variances")
	}
	if err := s.audit(ctx, "cash.cash_variance.refer_hr", "cash_variances", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, nil
}

// CloseCashVariance is the close action on cash variances.
func (s *Service) CloseCashVariance(ctx context.Context, p db.ActCloseCashVarianceParams) (db.CashVariance, error) {
	before, err := s.GetCashVariance(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cash.cash_variance.close", "cash_variances", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "cash_variances", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCloseCashVariance(ctx, p)
	if err != nil {
		return before, stateError(err, "close cash variances")
	}
	if err := s.audit(ctx, "cash.cash_variance.close", "cash_variances", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, nil
}

// ListCitProvider returns a page of cit providers.
func (s *Service) ListCitProvider(ctx context.Context, p db.ListCitProviderParams) ([]db.CitProvider, error) {
	return wrap2(s.q(ctx).ListCitProvider(ctx, p))("list cit providers")
}

// GetCitProvider returns one row of cit providers.
func (s *Service) GetCitProvider(ctx context.Context, id uuid.UUID) (db.CitProvider, error) {
	row, err := s.q(ctx).GetCitProvider(ctx, id)
	if err != nil {
		return row, notFound(err, "cit providers")
	}
	return row, nil
}

// CreateCitProvider adds one row to cit providers.
func (s *Service) CreateCitProvider(ctx context.Context, p db.CreateCitProviderParams) (db.CitProvider, error) {
	var zero db.CitProvider
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateCitProvider(ctx, p)
	if err != nil {
		return zero, wrap(err, "create cit providers")
	}
	return row, s.audit(ctx, "cash.cit_provider.write", "cit_providers", row.ID, nil, nil, row)
}

// UpdateCitProvider changes one row of cit providers.
func (s *Service) UpdateCitProvider(ctx context.Context, p db.UpdateCitProviderParams) (db.CitProvider, error) {
	before, err := s.GetCitProvider(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateCitProvider(ctx, p)
	if err != nil {
		return before, wrap(err, "update cit providers")
	}
	if err := s.voidSignatures(ctx, "cit_providers", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "cash.cit_provider.write", "cit_providers", row.ID, nil, before, row)
}

// ListCitPerson returns a page of cit personnel.
func (s *Service) ListCitPerson(ctx context.Context, p db.ListCitPersonParams) ([]db.CitPersonnel, error) {
	return wrap2(s.q(ctx).ListCitPerson(ctx, p))("list cit personnel")
}

// GetCitPerson returns one row of cit personnel.
func (s *Service) GetCitPerson(ctx context.Context, id uuid.UUID) (db.CitPersonnel, error) {
	row, err := s.q(ctx).GetCitPerson(ctx, id)
	if err != nil {
		return row, notFound(err, "cit personnel")
	}
	return row, nil
}

// CreateCitPerson adds one row to cit personnel.
func (s *Service) CreateCitPerson(ctx context.Context, p db.CreateCitPersonParams) (db.CitPersonnel, error) {
	var zero db.CitPersonnel
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreateCitPerson(ctx, p)
	if err != nil {
		return zero, wrap(err, "create cit personnel")
	}
	return row, s.audit(ctx, "cash.cit_person.write", "cit_personnel", row.ID, nil, nil, row)
}

// UpdateCitPerson changes one row of cit personnel.
func (s *Service) UpdateCitPerson(ctx context.Context, p db.UpdateCitPersonParams) (db.CitPersonnel, error) {
	before, err := s.GetCitPerson(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateCitPerson(ctx, p)
	if err != nil {
		return before, wrap(err, "update cit personnel")
	}
	if err := s.voidSignatures(ctx, "cit_personnel", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "cash.cit_person.write", "cit_personnel", row.ID, nil, before, row)
}

// ListCitTransit returns a page of cit transits.
func (s *Service) ListCitTransit(ctx context.Context, p db.ListCitTransitParams) ([]db.CitTransit, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListCitTransit(ctx, p))("list cit transits")
}

// GetCitTransit returns one row of cit transits.
func (s *Service) GetCitTransit(ctx context.Context, id uuid.UUID) (db.CitTransit, error) {
	row, err := s.q(ctx).GetCitTransit(ctx, id)
	if err != nil {
		return row, notFound(err, "cit transits")
	}
	return row, inScope(ctx, &row.SiteID)
}

// UpdateCitTransit changes one row of cit transits.
func (s *Service) UpdateCitTransit(ctx context.Context, p db.UpdateCitTransitParams) (db.CitTransit, error) {
	before, err := s.GetCitTransit(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateCitTransit(ctx, p)
	if err != nil {
		return before, wrap(err, "update cit transits")
	}
	if err := s.voidSignatures(ctx, "cit_transits", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "cash.cit_transit.write", "cit_transits", row.ID, &row.SiteID, before, row)
}

// CitTransitChildren returns the rows of a record's child tables, by API field.
func (s *Service) CitTransitChildren(ctx context.Context, id uuid.UUID) (map[string]any, error) {
	out := map[string]any{}
	{
		rows, err := s.q(ctx).ListCitBags(ctx, id)
		if err != nil {
			return nil, wrap(err, "list cit_bags")
		}
		out["bags"] = childRows(rows)
	}
	return out, nil
}

// ReplaceCitBags swaps a record's cit bags for the given rows.
func (s *Service) ReplaceCitBags(ctx context.Context, id uuid.UUID, rows []db.InsertCitBagsParams) error {
	if err := s.q(ctx).DeleteCitBags(ctx, id); err != nil {
		return wrap(err, "clear cit_bags")
	}
	for _, row := range rows {
		row.CreatedBy = actorID(ctx)
		if _, err := s.q(ctx).InsertCitBags(ctx, row); err != nil {
			return wrap(err, "add cit_bags")
		}
	}
	return nil
}

// ListBankDeposit returns a page of bank deposits.
func (s *Service) ListBankDeposit(ctx context.Context, p db.ListBankDepositParams) ([]db.BankDeposit, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListBankDeposit(ctx, p))("list bank deposits")
}

// GetBankDeposit returns one row of bank deposits.
func (s *Service) GetBankDeposit(ctx context.Context, id uuid.UUID) (db.BankDeposit, error) {
	row, err := s.q(ctx).GetBankDeposit(ctx, id)
	if err != nil {
		return row, notFound(err, "bank deposits")
	}
	return row, inScope(ctx, &row.SiteID)
}

// UpdateBankDeposit changes one row of bank deposits.
func (s *Service) UpdateBankDeposit(ctx context.Context, p db.UpdateBankDepositParams) (db.BankDeposit, error) {
	before, err := s.GetBankDeposit(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateBankDeposit(ctx, p)
	if err != nil {
		return before, wrap(err, "update bank deposits")
	}
	if err := s.voidSignatures(ctx, "bank_deposits", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "cash.bank_deposit.write", "bank_deposits", row.ID, &row.SiteID, before, row)
}

// VerifyBankDeposit is the finance verifies action on bank deposits.
func (s *Service) VerifyBankDeposit(ctx context.Context, p db.ActVerifyBankDepositParams) (db.BankDeposit, error) {
	before, err := s.GetBankDeposit(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cash.bank_deposit.verify", "bank_deposits", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "bank_deposits", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActVerifyBankDeposit(ctx, p)
	if err != nil {
		return before, stateError(err, "verify bank deposits")
	}
	if err := s.audit(ctx, "cash.bank_deposit.verify", "bank_deposits", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, nil
}

// ListCitReconciliation returns a page of cit reconciliations.
func (s *Service) ListCitReconciliation(ctx context.Context, p db.ListCitReconciliationParams) ([]db.CitReconciliation, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListCitReconciliation(ctx, p))("list cit reconciliations")
}

// GetCitReconciliation returns one row of cit reconciliations.
func (s *Service) GetCitReconciliation(ctx context.Context, id uuid.UUID) (db.CitReconciliation, error) {
	row, err := s.q(ctx).GetCitReconciliation(ctx, id)
	if err != nil {
		return row, notFound(err, "cit reconciliations")
	}
	return row, inScope(ctx, &row.SiteID)
}

// UpdateCitReconciliation changes one row of cit reconciliations.
func (s *Service) UpdateCitReconciliation(ctx context.Context, p db.UpdateCitReconciliationParams) (db.CitReconciliation, error) {
	before, err := s.GetCitReconciliation(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateCitReconciliation(ctx, p)
	if err != nil {
		return before, wrap(err, "update cit reconciliations")
	}
	if err := s.voidSignatures(ctx, "cit_reconciliations", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "cash.cit_reconciliation.write", "cit_reconciliations", row.ID, &row.SiteID, before, row)
}

// CompleteCitReconciliation is the mark reconciled action on cit reconciliations.
func (s *Service) CompleteCitReconciliation(ctx context.Context, p db.ActCompleteCitReconciliationParams) (db.CitReconciliation, error) {
	before, err := s.GetCitReconciliation(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cash.cit_reconciliation.complete", "cit_reconciliations", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "cit_reconciliations", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCompleteCitReconciliation(ctx, p)
	if err != nil {
		return before, stateError(err, "complete cit reconciliations")
	}
	if err := s.audit(ctx, "cash.cit_reconciliation.complete", "cit_reconciliations", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, nil
}

// ListSurpriseCount returns a page of surprise counts.
func (s *Service) ListSurpriseCount(ctx context.Context, p db.ListSurpriseCountParams) ([]db.SurpriseCount, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListSurpriseCount(ctx, p))("list surprise counts")
}

// GetSurpriseCount returns one row of surprise counts.
func (s *Service) GetSurpriseCount(ctx context.Context, id uuid.UUID) (db.SurpriseCount, error) {
	row, err := s.q(ctx).GetSurpriseCount(ctx, id)
	if err != nil {
		return row, notFound(err, "surprise counts")
	}
	return row, inScope(ctx, &row.SiteID)
}

// UpdateSurpriseCount changes one row of surprise counts.
func (s *Service) UpdateSurpriseCount(ctx context.Context, p db.UpdateSurpriseCountParams) (db.SurpriseCount, error) {
	before, err := s.GetSurpriseCount(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdateSurpriseCount(ctx, p)
	if err != nil {
		return before, wrap(err, "update surprise counts")
	}
	if err := s.voidSignatures(ctx, "surprise_counts", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "cash.surprise_count.write", "surprise_counts", row.ID, &row.SiteID, before, row)
}

// CompleteSurpriseCount is the complete the count action on surprise counts.
func (s *Service) CompleteSurpriseCount(ctx context.Context, p db.ActCompleteSurpriseCountParams) (db.SurpriseCount, error) {
	before, err := s.GetSurpriseCount(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cash.surprise_count.complete", "surprise_counts", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "surprise_counts", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActCompleteSurpriseCount(ctx, p)
	if err != nil {
		return before, stateError(err, "complete surprise counts")
	}
	if err := s.audit(ctx, "cash.surprise_count.complete", "surprise_counts", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, nil
}

// ListPettyCashAccount returns a page of petty cash accounts.
func (s *Service) ListPettyCashAccount(ctx context.Context, p db.ListPettyCashAccountParams) ([]db.PettyCashAccount, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListPettyCashAccount(ctx, p))("list petty cash accounts")
}

// GetPettyCashAccount returns one row of petty cash accounts.
func (s *Service) GetPettyCashAccount(ctx context.Context, id uuid.UUID) (db.PettyCashAccount, error) {
	row, err := s.q(ctx).GetPettyCashAccount(ctx, id)
	if err != nil {
		return row, notFound(err, "petty cash accounts")
	}
	return row, inScope(ctx, &row.SiteID)
}

// CreatePettyCashAccount adds one row to petty cash accounts.
func (s *Service) CreatePettyCashAccount(ctx context.Context, p db.CreatePettyCashAccountParams) (db.PettyCashAccount, error) {
	var zero db.PettyCashAccount
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	p.CreatedBy = actorID(ctx)
	row, err := s.q(ctx).CreatePettyCashAccount(ctx, p)
	if err != nil {
		return zero, wrap(err, "create petty cash accounts")
	}
	return row, s.audit(ctx, "cash.petty_cash_account.write", "petty_cash_accounts", row.ID, &p.SiteID, nil, row)
}

// UpdatePettyCashAccount changes one row of petty cash accounts.
func (s *Service) UpdatePettyCashAccount(ctx context.Context, p db.UpdatePettyCashAccountParams) (db.PettyCashAccount, error) {
	before, err := s.GetPettyCashAccount(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdatePettyCashAccount(ctx, p)
	if err != nil {
		return before, wrap(err, "update petty cash accounts")
	}
	if err := s.voidSignatures(ctx, "petty_cash_accounts", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "cash.petty_cash_account.write", "petty_cash_accounts", row.ID, &row.SiteID, before, row)
}

// ListPettyCashVoucher returns a page of petty cash vouchers.
func (s *Service) ListPettyCashVoucher(ctx context.Context, p db.ListPettyCashVoucherParams) ([]db.PettyCashVoucher, error) {
	p.AllSites, p.SiteIds = scopeOf(ctx)
	return wrap2(s.q(ctx).ListPettyCashVoucher(ctx, p))("list petty cash vouchers")
}

// GetPettyCashVoucher returns one row of petty cash vouchers.
func (s *Service) GetPettyCashVoucher(ctx context.Context, id uuid.UUID) (db.PettyCashVoucher, error) {
	row, err := s.q(ctx).GetPettyCashVoucher(ctx, id)
	if err != nil {
		return row, notFound(err, "petty cash vouchers")
	}
	return row, inScope(ctx, row.SiteID)
}

// UpdatePettyCashVoucher changes one row of petty cash vouchers.
func (s *Service) UpdatePettyCashVoucher(ctx context.Context, p db.UpdatePettyCashVoucherParams) (db.PettyCashVoucher, error) {
	before, err := s.GetPettyCashVoucher(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := checkVersion(before.RowVersion, p.RowVersion); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).UpdatePettyCashVoucher(ctx, p)
	if err != nil {
		return before, wrap(err, "update petty cash vouchers")
	}
	if err := s.voidSignatures(ctx, "petty_cash_vouchers", row.ID); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "cash.petty_cash_voucher.write", "petty_cash_vouchers", row.ID, row.SiteID, before, row)
}

// SubmitPettyCashVoucher is the submit action on petty cash vouchers.
func (s *Service) SubmitPettyCashVoucher(ctx context.Context, p db.ActSubmitPettyCashVoucherParams) (db.PettyCashVoucher, error) {
	before, err := s.GetPettyCashVoucher(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cash.petty_cash_voucher.submit", "petty_cash_vouchers", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitPettyCashVoucher(ctx, p)
	if err != nil {
		return before, stateError(err, "submit petty cash vouchers")
	}
	if err := s.audit(ctx, "cash.petty_cash_voucher.submit", "petty_cash_vouchers", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, s.startApproval(ctx, "petty_cash_vouchers", row.ID, row.SiteID, row)
}

// AttachReceiptPettyCashVoucher is the attach proof of purchase action on petty cash vouchers.
func (s *Service) AttachReceiptPettyCashVoucher(ctx context.Context, p db.ActAttachReceiptPettyCashVoucherParams) (db.PettyCashVoucher, error) {
	before, err := s.GetPettyCashVoucher(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cash.petty_cash_voucher.attach_receipt", "petty_cash_vouchers", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActAttachReceiptPettyCashVoucher(ctx, p)
	if err != nil {
		return before, stateError(err, "attach receipt petty cash vouchers")
	}
	if err := s.audit(ctx, "cash.petty_cash_voucher.attach_receipt", "petty_cash_vouchers", row.ID, row.SiteID, before, row); err != nil {
		return row, err
	}
	return row, nil
}

// ListPettyCashHandover returns a page of petty cash handovers.
func (s *Service) ListPettyCashHandover(ctx context.Context, p db.ListPettyCashHandoverParams) ([]db.PettyCashHandover, error) {
	return wrap2(s.q(ctx).ListPettyCashHandover(ctx, p))("list petty cash handovers")
}

// GetPettyCashHandover returns one row of petty cash handovers.
func (s *Service) GetPettyCashHandover(ctx context.Context, id uuid.UUID) (db.PettyCashHandover, error) {
	row, err := s.q(ctx).GetPettyCashHandover(ctx, id)
	if err != nil {
		return row, notFound(err, "petty cash handovers")
	}
	return row, nil
}

// SubmitPettyCashHandover is the submit action on petty cash handovers.
func (s *Service) SubmitPettyCashHandover(ctx context.Context, p db.ActSubmitPettyCashHandoverParams) (db.PettyCashHandover, error) {
	before, err := s.GetPettyCashHandover(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cash.petty_cash_handover.submit", "petty_cash_handovers", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitPettyCashHandover(ctx, p)
	if err != nil {
		return before, stateError(err, "submit petty cash handovers")
	}
	if err := s.audit(ctx, "cash.petty_cash_handover.submit", "petty_cash_handovers", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, nil
}

// ListPettyCashReconciliation returns a page of petty cash reconciliations.
func (s *Service) ListPettyCashReconciliation(ctx context.Context, p db.ListPettyCashReconciliationParams) ([]db.PettyCashReconciliation, error) {
	return wrap2(s.q(ctx).ListPettyCashReconciliation(ctx, p))("list petty cash reconciliations")
}

// GetPettyCashReconciliation returns one row of petty cash reconciliations.
func (s *Service) GetPettyCashReconciliation(ctx context.Context, id uuid.UUID) (db.PettyCashReconciliation, error) {
	row, err := s.q(ctx).GetPettyCashReconciliation(ctx, id)
	if err != nil {
		return row, notFound(err, "petty cash reconciliations")
	}
	return row, nil
}

// SubmitPettyCashReconciliation is the submit action on petty cash reconciliations.
func (s *Service) SubmitPettyCashReconciliation(ctx context.Context, p db.ActSubmitPettyCashReconciliationParams) (db.PettyCashReconciliation, error) {
	before, err := s.GetPettyCashReconciliation(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cash.petty_cash_reconciliation.submit", "petty_cash_reconciliations", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitPettyCashReconciliation(ctx, p)
	if err != nil {
		return before, stateError(err, "submit petty cash reconciliations")
	}
	if err := s.audit(ctx, "cash.petty_cash_reconciliation.submit", "petty_cash_reconciliations", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, nil
}

// ReviewPettyCashReconciliation is the review action on petty cash reconciliations.
func (s *Service) ReviewPettyCashReconciliation(ctx context.Context, p db.ActReviewPettyCashReconciliationParams) (db.PettyCashReconciliation, error) {
	before, err := s.GetPettyCashReconciliation(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cash.petty_cash_reconciliation.review", "petty_cash_reconciliations", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	if err := s.requireSignatures(ctx, "petty_cash_reconciliations", p.ID); err != nil {
		return before, err
	}
	row, err := s.q(ctx).ActReviewPettyCashReconciliation(ctx, p)
	if err != nil {
		return before, stateError(err, "review petty cash reconciliations")
	}
	if err := s.audit(ctx, "cash.petty_cash_reconciliation.review", "petty_cash_reconciliations", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, nil
}

// ListPettyCashReimbursement returns a page of petty cash reimbursements.
func (s *Service) ListPettyCashReimbursement(ctx context.Context, p db.ListPettyCashReimbursementParams) ([]db.PettyCashReimbursement, error) {
	return wrap2(s.q(ctx).ListPettyCashReimbursement(ctx, p))("list petty cash reimbursements")
}

// GetPettyCashReimbursement returns one row of petty cash reimbursements.
func (s *Service) GetPettyCashReimbursement(ctx context.Context, id uuid.UUID) (db.PettyCashReimbursement, error) {
	row, err := s.q(ctx).GetPettyCashReimbursement(ctx, id)
	if err != nil {
		return row, notFound(err, "petty cash reimbursements")
	}
	return row, nil
}

// SubmitPettyCashReimbursement is the submit action on petty cash reimbursements.
func (s *Service) SubmitPettyCashReimbursement(ctx context.Context, p db.ActSubmitPettyCashReimbursementParams) (db.PettyCashReimbursement, error) {
	before, err := s.GetPettyCashReimbursement(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if err := s.checkSod(ctx, "cash.petty_cash_reimbursement.submit", "petty_cash_reimbursements", p.ID); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitPettyCashReimbursement(ctx, p)
	if err != nil {
		return before, stateError(err, "submit petty cash reimbursements")
	}
	if err := s.audit(ctx, "cash.petty_cash_reimbursement.submit", "petty_cash_reimbursements", row.ID, nil, before, row); err != nil {
		return row, err
	}
	return row, nil
}
