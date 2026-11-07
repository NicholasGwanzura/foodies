package service

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"foodies/api/internal/auth"
	"foodies/api/internal/db"
)

const baseCurrency = "USD"

func errLimit(rule, detail string) error {
	return &Error{Status: http.StatusUnprocessableEntity, Code: "limit_exceeded", Title: "Over the limit", Detail: detail,
		Blockers: []Blocker{{Kind: "limit_exceeded", Reference: rule, Message: detail}}}
}

// overLimit compares an amount to a money rule. Rules are held in USD, and an
// amount in another currency converts at the latest rate, because limits are
// never compared across currencies by face value.
func (s *Service) overLimit(ctx context.Context, rule string, site *uuid.UUID, amount, currency string) (bool, moneyRule, error) {
	v, err := s.RuleValue(ctx, rule, site)
	if err != nil {
		return false, moneyRule{}, err
	}
	limit, err := decodeMoney(v.Value)
	if err != nil {
		return false, limit, err
	}
	value := rat(amount)
	if currency != limit.Currency {
		fx, err := s.q(ctx).LatestRateBetween(ctx, db.LatestRateBetweenParams{FromCurrency: currency, ToCurrency: limit.Currency})
		if err != nil {
			return false, limit, Fail(http.StatusConflict, "rule_not_configured", "No exchange rate",
				fmt.Sprintf("There is no %s to %s rate, so this cannot be checked against its limit.", currency, limit.Currency))
		}
		value = mul(value, rat(fx.Rate))
	}
	return value.Cmp(rat(limit.Amount)) > 0, limit, nil
}

// CreateFloatCertificate records approved float amounts, within the limits.
func (s *Service) CreateFloatCertificate(ctx context.Context, p db.CreateFloatCertificateParams) (db.FloatCertificate, error) {
	var zero db.FloatCertificate
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.requireLive(ctx, "cash.float", p.SiteID); err != nil {
		return zero, err
	}
	for _, c := range []struct{ rule, amount, currency, label string }{
		{"cash.change_float_max", p.ChangeFloatAmount, p.ChangeFloatCurrency, "The change float"},
		{"cash.petty_cash_max", p.PettyCashAmount, p.PettyCashCurrency, "The petty cash amount"},
	} {
		over, limit, err := s.overLimit(ctx, c.rule, &p.SiteID, c.amount, c.currency)
		if err != nil {
			return zero, err
		}
		if over {
			return zero, errLimit(c.rule, fmt.Sprintf("%s cannot exceed %s %s.", c.label, limit.Currency, limit.Amount))
		}
	}
	ref, err := s.nextReference(ctx, "FLC", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference, p.CreatedBy = ref, actorID(ctx)
	row, err := s.q(ctx).CreateFloatCertificate(ctx, p)
	if err != nil {
		return zero, wrap(err, "create certificate")
	}
	return row, s.audit(ctx, "cash.float_certificate.write", "float_certificates", row.ID, &row.SiteID, nil, row)
}

// CreateFloatException records exceptional float use, due back the next day.
func (s *Service) CreateFloatException(ctx context.Context, p db.CreateFloatExceptionParams) (db.FloatException, error) {
	var zero db.FloatException
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "FLX", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference, p.CreatedBy = ref, actorID(ctx)
	row, err := s.q(ctx).CreateFloatException(ctx, p)
	if err != nil {
		return zero, wrap(err, "create float exception")
	}
	due := pgtype.Date{Time: p.TransactionDate.Time.AddDate(0, 0, 1), Valid: true}
	if err := s.q(ctx).SetExceptionDue(ctx, db.SetExceptionDueParams{ID: row.ID, ReimbursementDue: due}); err != nil {
		return zero, wrap(err, "set due date")
	}
	row.ReimbursementDue = due
	return row, s.audit(ctx, "cash.float_exception.write", "float_exceptions", row.ID, &row.SiteID, nil, row)
}

// CreateTillFlush records a flush. The cashier and supervisor count together,
// so they must be different people.
func (s *Service) CreateTillFlush(ctx context.Context, p db.CreateTillFlushParams) (db.TillFlush, error) {
	var zero db.TillFlush
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.requireLive(ctx, "cash.till_flush", p.SiteID); err != nil {
		return zero, err
	}
	if p.CashierEmployeeID == p.SupervisorEmployeeID {
		return zero, errInvalid("/supervisor_employee_id", "The cashier and supervisor must be different people.")
	}
	ref, err := s.nextReference(ctx, "FLS", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference, p.CreatedBy = ref, actorID(ctx)
	row, err := s.q(ctx).CreateTillFlush(ctx, p)
	if err != nil {
		return zero, wrap(err, "create flush")
	}
	return row, s.audit(ctx, "cash.till_flush.write", "till_flushes", row.ID, &row.SiteID, nil, row)
}

func (s *Service) requestSignature(ctx context.Context, table string, id, employee uuid.UUID, meaning string) error {
	_, err := s.CreateSignatureRequest(ctx, db.CreateSignatureRequestRowParams{
		SubjectType: table, SubjectID: id, Meaning: meaning, SignerKind: "employee", SignerEmployeeID: &employee,
	})
	return err
}

// SubmitTillFlush asks each party to sign. A flush cannot close without them.
func (s *Service) SubmitTillFlush(ctx context.Context, p db.ActSubmitTillFlushParams) (db.TillFlush, error) {
	before, err := s.GetTillFlush(ctx, p.ID)
	if err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitTillFlush(ctx, p)
	if err != nil {
		return before, stateError(err, "submit this flush")
	}
	if err := s.audit(ctx, "cash.till_flush.submit", "till_flushes", row.ID, &row.SiteID, before, row); err != nil {
		return row, err
	}
	signers := []struct {
		who     uuid.UUID
		meaning string
	}{{row.CashierEmployeeID, "counted"}, {row.SupervisorEmployeeID, "verified"}}
	if row.SecurityWitnessEmployeeID != nil {
		signers = append(signers, struct {
			who     uuid.UUID
			meaning string
		}{*row.SecurityWitnessEmployeeID, "witnessed"})
	}
	for _, sig := range signers {
		if err := s.requestSignature(ctx, "till_flushes", row.ID, sig.who, sig.meaning); err != nil {
			return row, err
		}
	}
	return row, nil
}

// CreateCashUp opens a cash-up and freezes the exchange rate it will use.
func (s *Service) CreateCashUp(ctx context.Context, p db.CreateCashUpParams) (db.CashUp, error) {
	var zero db.CashUp
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.requireLive(ctx, "cash.cash_up", p.SiteID); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "CUP", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference, p.CreatedBy = ref, actorID(ctx)
	row, err := s.q(ctx).CreateCashUp(ctx, p)
	if err != nil {
		return zero, wrap(err, "create cash-up")
	}
	return row, s.audit(ctx, "cash.cash_up.write", "cash_ups", row.ID, &row.SiteID, nil, row)
}

// ReplaceCashUpTenders stores the counts. Each tender's variance is by its own
// currency. Only the base-currency total converts, at the rate stored on the
// record, so USD and ZiG are never netted.
func (s *Service) ReplaceCashUpTenders(ctx context.Context, id uuid.UUID, rows []db.InsertCashUpTendersParams) error {
	q := s.q(ctx)
	up, err := q.GetCashUp(ctx, id)
	if err != nil {
		return wrap(err, "read cash-up")
	}
	if err := q.DeleteCashUpTenders(ctx, id); err != nil {
		return wrap(err, "clear tenders")
	}
	total := new(big.Rat)
	var rateID *uuid.UUID
	var rate *string
	for _, row := range rows {
		variance := sub(rat(row.Actual), rat(row.Expected))
		text := dec(variance)
		row.Variance, row.CreatedBy = &text, actorID(ctx)
		if _, err := q.InsertCashUpTenders(ctx, row); err != nil {
			return wrap(err, "add tender")
		}
		if row.Currency == baseCurrency {
			total = add(total, variance)
			continue
		}
		fx, err := q.LatestRateBetween(ctx, db.LatestRateBetweenParams{FromCurrency: row.Currency, ToCurrency: baseCurrency})
		if err != nil {
			return Fail(http.StatusConflict, "rule_not_configured", "No exchange rate", "Add a "+row.Currency+" to USD rate before counting in "+row.Currency+".")
		}
		rateID, rate = &fx.ID, &fx.Rate
		total = add(total, mul(variance, rat(fx.Rate)))
	}
	amount := dec(total)
	_, err = q.SetCashUpFigures(ctx, db.SetCashUpFiguresParams{
		ID: up.ID, FxRateID: rateID, FxRateValue: rate, VarianceBaseAmount: &amount, VarianceBaseCurrency: ptr(baseCurrency),
	})
	return wrap(err, "store totals")
}

// SubmitCashUp sends a count for review. A close-of-day needs its Z-report,
// because one that failed to print stops the cash-up until instructed.
func (s *Service) SubmitCashUp(ctx context.Context, p db.ActSubmitCashUpParams) (db.CashUp, error) {
	before, err := s.GetCashUp(ctx, p.ID)
	if err != nil {
		return before, err
	}
	tenders, err := s.q(ctx).ListCashUpTenders(ctx, p.ID)
	if err != nil {
		return before, wrap(err, "read tenders")
	}
	if len(tenders) == 0 {
		return before, errInvalid("/tenders", "Enter the count before submitting.")
	}
	if before.Kind == "close_of_day" && before.ZReportID == nil {
		return before, &Error{Status: http.StatusConflict, Code: "conflict", Title: "Z-report needed",
			Detail:   "A close-of-day cash-up needs its Z-report. Stop here if it failed to print.",
			Blockers: []Blocker{{Kind: "evidence_missing", Message: "Attach the Z-report"}}}
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitCashUp(ctx, p)
	if err != nil {
		return before, stateError(err, "submit this cash-up")
	}
	return row, s.audit(ctx, "cash.cash_up.submit", "cash_ups", row.ID, &row.SiteID, before, row)
}

// ApproveCashUp needs an independent reviewer. Every non-zero tender becomes a
// logged variance, however small.
func (s *Service) ApproveCashUp(ctx context.Context, p db.ActApproveCashUpParams, signature uuid.UUID) (db.CashUp, error) {
	before, err := s.GetCashUp(ctx, p.ID)
	if err != nil {
		return before, err
	}
	a := auth.From(ctx)
	if a.EmployeeID == before.CashierEmployeeID || a.EmployeeID == before.SupervisorEmployeeID {
		return before, &Error{Status: http.StatusForbidden, Code: "sod_violation", Title: "Someone else must do this",
			Detail:   "An independent reviewer must approve a cash-up.",
			Blockers: []Blocker{{Kind: "sod_conflict", Reference: "cash_up_self_approval", Message: "A cashier cannot approve their own cash-up"}}}
	}
	if err := s.checkSod(ctx, "cash.cash_up.approve", "cash_ups", p.ID); err != nil {
		return before, err
	}
	if err := s.signedOn(ctx, "cash_ups", p.ID, signature); err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActApproveCashUp(ctx, p)
	if err != nil {
		return before, stateError(err, "approve this cash-up")
	}
	if err := s.raiseVariances(ctx, row); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "cash.cash_up.approve", "cash_ups", row.ID, &row.SiteID, before, row)
}

// signedOn confirms the actor signed this exact record.
func (s *Service) signedOn(ctx context.Context, table string, subject, signature uuid.UUID) error {
	sig, err := s.q(ctx).GetSignatureRequest(ctx, signature)
	if err != nil {
		return notFound(err, "signature")
	}
	a := auth.From(ctx)
	if sig.Status != "signed" || sig.SubjectType != table || sig.SubjectID != subject ||
		sig.SignerEmployeeID == nil || *sig.SignerEmployeeID != a.EmployeeID {
		return Fail(http.StatusUnprocessableEntity, "signature_required", "Sign first", "Confirm with your PIN before approving.")
	}
	return nil
}

func (s *Service) raiseVariances(ctx context.Context, up db.CashUp) error {
	q := s.q(ctx)
	tenders, err := q.ListCashUpTenders(ctx, up.ID)
	if err != nil {
		return wrap(err, "read tenders")
	}
	for _, t := range tenders {
		variance := ratp(t.Variance)
		if variance.Sign() == 0 {
			continue
		}
		kind, abs := "shortage", new(big.Rat).Abs(variance)
		if variance.Sign() > 0 {
			kind = "overage"
		}
		counts, err := q.CountVariances(ctx, &up.CashierEmployeeID)
		if err != nil {
			return wrap(err, "count variances")
		}
		week, month := counts.Week+1, counts.Month+1
		ref, err := s.nextReference(ctx, "VAR", &up.SiteID)
		if err != nil {
			return err
		}
		amount := dec(abs)
		row, err := q.CreateCashVarianceRow(ctx, db.CreateCashVarianceRowParams{
			Reference: ref, CashUpID: up.ID, SiteID: &up.SiteID, CashierEmployeeID: &up.CashierEmployeeID, VarianceType: &kind,
			AmountAmount: &amount, AmountCurrency: &t.Currency, OccurrenceWeek: &week, OccurrenceMonth: &month,
		})
		if err != nil {
			return wrap(err, "create variance")
		}
		over, _, err := s.overLimit(ctx, "cash.variance_escalation", &up.SiteID, amount, t.Currency)
		if err != nil {
			return err
		}
		if over {
			if _, err := s.RaiseEscalation(ctx, "cash_variances", row.ID, "cash", 2, &up.SiteID); err != nil {
				return err
			}
		}
	}
	return nil
}

// requireCashUpsSettled enforces that banking follows a settled day.
func (s *Service) requireCashUpsSettled(ctx context.Context, site uuid.UUID, date pgtype.Date) error {
	q := s.q(ctx)
	open, err := q.UnsettledCashUps(ctx, db.UnsettledCashUpsParams{SiteID: site, BusinessDate: date})
	if err != nil {
		return wrap(err, "check cash-ups")
	}
	variances, err := q.OpenVariances(ctx, &site)
	if err != nil {
		return wrap(err, "check variances")
	}
	var blockers []Blocker
	if open > 0 {
		blockers = append(blockers, Blocker{Kind: "approval_pending", Message: fmt.Sprintf("%d cash-up(s) not yet approved", open)})
	}
	if variances > 0 {
		blockers = append(blockers, Blocker{Kind: "approval_pending", Message: fmt.Sprintf("%d variance(s) still open", variances)})
	}
	if len(blockers) > 0 {
		return &Error{Status: http.StatusConflict, Code: "conflict", Title: "Cash-ups come first",
			Detail: "CIT preparation cannot start until the cash-up and its variances are dealt with.", Blockers: blockers}
	}
	return nil
}

// CreateCitTransit prepares a transit, once the day's cash is settled.
func (s *Service) CreateCitTransit(ctx context.Context, p db.CreateCitTransitParams) (db.CitTransit, error) {
	var zero db.CitTransit
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.requireLive(ctx, "cash.cit", p.SiteID); err != nil {
		return zero, err
	}
	if err := s.requireCashUpsSettled(ctx, p.SiteID, p.TransitDate); err != nil {
		return zero, err
	}
	over, limit, err := s.overLimit(ctx, "cash.cit_transit_limit", &p.SiteID, p.TotalAmount, p.TotalCurrency)
	if err != nil {
		return zero, err
	}
	if over {
		return zero, errLimit("cash.cit_transit_limit", fmt.Sprintf("A transit cannot exceed %s %s.", limit.Currency, limit.Amount))
	}
	ref, err := s.nextReference(ctx, "CIT", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference, p.CreatedBy = ref, actorID(ctx)
	row, err := s.q(ctx).CreateCitTransit(ctx, p)
	if err != nil {
		return zero, wrap(err, "create transit")
	}
	return row, s.audit(ctx, "cash.cit_transit.write", "cit_transits", row.ID, &row.SiteID, nil, row)
}

// SealCitTransit seals the bags. Their total must equal the transit total.
func (s *Service) SealCitTransit(ctx context.Context, p db.ActSealCitTransitParams) (db.CitTransit, error) {
	before, err := s.GetCitTransit(ctx, p.ID)
	if err != nil {
		return before, err
	}
	bags, err := s.q(ctx).ListCitBags(ctx, p.ID)
	if err != nil {
		return before, wrap(err, "read bags")
	}
	if len(bags) == 0 {
		return before, errInvalid("/bags", "Add the bags and seal numbers first.")
	}
	total := new(big.Rat)
	for _, b := range bags {
		total = add(total, rat(b.AmountAmount))
	}
	if total.Cmp(rat(before.TotalAmount)) != 0 {
		return before, errInvalid("/bags", "The bags do not add up to the transit total.")
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSealCitTransit(ctx, p)
	if err != nil {
		return before, stateError(err, "seal this transit")
	}
	return row, s.audit(ctx, "cash.cit_transit.seal", "cit_transits", row.ID, &row.SiteID, before, row)
}

// VerifyIdentityCitTransit matches the crew to the approved list. A failed
// match blocks the handover.
func (s *Service) VerifyIdentityCitTransit(ctx context.Context, id, person uuid.UUID) (db.CitTransit, error) {
	before, err := s.GetCitTransit(ctx, id)
	if err != nil {
		return before, err
	}
	q := s.q(ctx)
	crew, err := q.GetCitPerson(ctx, person)
	if errors.Is(err, pgx.ErrNoRows) || (err == nil && (crew.CitProviderID != before.ProviderID || !crew.Active)) {
		return before, &Error{Status: http.StatusConflict, Code: "conflict", Title: "Identity not matched",
			Detail: "This person is not on the provider's approved list. Do not hand over."}
	}
	if err != nil {
		return before, wrap(err, "read crew")
	}
	row, err := q.SetCitIdentity(ctx, db.SetCitIdentityParams{ID: id, Person: &person, UpdatedBy: actorID(ctx)})
	if err != nil {
		return before, stateError(err, "verify identity")
	}
	return row, s.audit(ctx, "cash.cit_transit.verify_identity", "cit_transits", row.ID, &row.SiteID, before, row)
}

// HandOverCitTransit records the crew taking the cash, which needs the formal
// receipt and the armoured vehicle.
func (s *Service) HandOverCitTransit(ctx context.Context, id uuid.UUID, receipt string, armoured bool) (db.CitTransit, error) {
	before, err := s.GetCitTransit(ctx, id)
	if err != nil {
		return before, err
	}
	if strings.TrimSpace(receipt) == "" {
		return before, errInvalid("/cit_receipt_number", "Enter the CIT receipt number.")
	}
	if !armoured {
		return before, errInvalid("/armoured_vehicle_confirmed", "Confirm the armoured vehicle before handing over.")
	}
	row, err := s.q(ctx).SetCitHandover(ctx, db.SetCitHandoverParams{ID: id, CitReceiptNumber: &receipt, ArmouredVehicleConfirmed: armoured, UpdatedBy: actorID(ctx)})
	if err != nil {
		return before, stateError(err, "hand this over")
	}
	return row, s.audit(ctx, "cash.cit_transit.hand_over", "cit_transits", row.ID, &row.SiteID, before, row)
}

// ReceiveVaultCitTransit checks the amount at the vault. A mismatch starts the
// discrepancy clock from this moment.
func (s *Service) ReceiveVaultCitTransit(ctx context.Context, id uuid.UUID, amount, currency string) (db.CitTransit, error) {
	before, err := s.GetCitTransit(ctx, id)
	if err != nil {
		return before, err
	}
	status := "received"
	if currency != before.TotalCurrency || rat(amount).Cmp(rat(before.TotalAmount)) != 0 {
		status = "discrepancy"
	}
	row, err := s.q(ctx).SetCitVault(ctx, db.SetCitVaultParams{ID: id, Status: status, VaultAmountAmount: &amount, VaultAmountCurrency: &currency, ReceivedBy: actorID(ctx)})
	if err != nil {
		return before, stateError(err, "receive this at the vault")
	}
	if status == "discrepancy" {
		if err := s.citDiscrepancy(ctx, row); err != nil {
			return row, err
		}
	}
	return row, s.audit(ctx, "cash.cit_transit.receive_vault", "cit_transits", row.ID, &row.SiteID, before, row)
}

// citDiscrepancy raises the alert and a task due within the reporting window.
func (s *Service) citDiscrepancy(ctx context.Context, t db.CitTransit) error {
	if _, err := s.RaiseEscalation(ctx, "cit_transits", t.ID, "cash", 3, &t.SiteID); err != nil {
		return err
	}
	hours := int64(2)
	if v, err := s.RuleValue(ctx, "cash.cit_incident_report_hours", &t.SiteID); err == nil {
		if n, err := strconv.ParseInt(string(v.Value), 10, 64); err == nil {
			hours = n
		}
	}
	role, err := s.q(ctx).RoleByKey(ctx, "finance_manager")
	if err != nil {
		return wrap(err, "read role")
	}
	due := s.now().Add(time.Duration(hours) * time.Hour)
	_, err = s.CreateTask(ctx, db.CreateTaskRowParams{
		Title: "Report the CIT discrepancy on " + t.Reference, SubjectType: ptr("cit_transits"), SubjectID: &t.ID, SiteID: t.SiteID,
		BusinessDate: s.today(), AssigneeRoleID: &role, DueAt: &due, Priority: ptr("critical"),
	})
	return err
}

// RaiseDiscrepancyCitTransit records a mismatch found outside the vault check.
func (s *Service) RaiseDiscrepancyCitTransit(ctx context.Context, p db.ActRaiseDiscrepancyCitTransitParams) (db.CitTransit, error) {
	before, err := s.GetCitTransit(ctx, p.ID)
	if err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActRaiseDiscrepancyCitTransit(ctx, p)
	if err != nil {
		return before, stateError(err, "raise a discrepancy")
	}
	if err := s.citDiscrepancy(ctx, row); err != nil {
		return row, err
	}
	return row, s.audit(ctx, "cash.cit_transit.raise_discrepancy", "cit_transits", row.ID, &row.SiteID, before, row)
}

// CreateBankDeposit records banking. Takings are banked intact, so a partial
// deposit needs a reason, and direct banking needs Head Office's instruction.
func (s *Service) CreateBankDeposit(ctx context.Context, p db.CreateBankDepositParams) (db.BankDeposit, error) {
	var zero db.BankDeposit
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	if err := s.requireLive(ctx, "cash.banking", p.SiteID); err != nil {
		return zero, err
	}
	if p.Route == "direct_bank" && strings.TrimSpace(deref(p.FinanceInstructionRef)) == "" {
		return zero, errInvalid("/finance_instruction_ref", "Direct banking needs a Head Office Finance instruction.")
	}
	partial := p.TakingsCurrency != p.DepositAmountCurrency || rat(p.DepositAmountAmount).Cmp(rat(p.TakingsAmount)) != 0
	if partial && strings.TrimSpace(deref(p.DelayReason)) == "" {
		return zero, errInvalid("/delay_reason", "Takings are banked intact. Explain why this differs.")
	}
	ref, err := s.nextReference(ctx, "DEP", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference, p.CreatedBy = ref, actorID(ctx)
	row, err := s.q(ctx).CreateBankDeposit(ctx, p)
	if err != nil {
		return zero, wrap(err, "create deposit")
	}
	if partial {
		if err := s.q(ctx).SetDepositException(ctx, row.ID); err != nil {
			return row, wrap(err, "flag deposit")
		}
		row.Status = "exception"
	}
	return row, s.audit(ctx, "cash.bank_deposit.write", "bank_deposits", row.ID, &row.SiteID, nil, row)
}

// SubmitProofBankDeposit stores the proof and when it arrived.
func (s *Service) SubmitProofBankDeposit(ctx context.Context, p db.ActSubmitProofBankDepositParams) (db.BankDeposit, error) {
	before, err := s.GetBankDeposit(ctx, p.ID)
	if err != nil {
		return before, err
	}
	p.UpdatedBy = actorID(ctx)
	row, err := s.q(ctx).ActSubmitProofBankDeposit(ctx, p)
	if err != nil {
		return before, stateError(err, "submit proof")
	}
	if err := s.q(ctx).SetDepositProofTime(ctx, row.ID); err != nil {
		return row, wrap(err, "store proof time")
	}
	return row, s.audit(ctx, "cash.bank_deposit.submit_proof", "bank_deposits", row.ID, &row.SiteID, before, row)
}

// prohibitedPettyCash are the spec's categories that go through procurement.
var prohibitedPettyCash = []string{"creditors", "salaries", "wages", "major", "personal"}

// CreatePettyCashVoucher raises a request. Prohibited categories are rejected
// and routed to procurement, and the amount must fit the limit.
func (s *Service) CreatePettyCashVoucher(ctx context.Context, p db.CreatePettyCashVoucherParams) (db.PettyCashVoucher, error) {
	var zero db.PettyCashVoucher
	q := s.q(ctx)
	account, err := q.GetPettyCashAccount(ctx, p.AccountID)
	if err != nil {
		return zero, notFound(err, "petty cash account")
	}
	if err := canWrite(ctx, &account.SiteID); err != nil {
		return zero, err
	}
	if err := s.requireLive(ctx, "cash.petty_cash", account.SiteID); err != nil {
		return zero, err
	}
	if slices.Contains(prohibitedPettyCash, strings.ToLower(strings.TrimSpace(p.ExpenseCategory))) {
		return zero, &Error{Status: http.StatusUnprocessableEntity, Code: "limit_exceeded", Title: "Not a petty cash item",
			Detail: "Petty cash cannot pay for this. Raise a purchase requisition with Procurement instead."}
	}
	over, limit, err := s.overLimit(ctx, "cash.petty_cash_max", &account.SiteID, p.AmountAmount, p.AmountCurrency)
	if err != nil {
		return zero, err
	}
	if over {
		return zero, errLimit("cash.petty_cash_max", fmt.Sprintf("Petty cash cannot exceed %s %s.", limit.Currency, limit.Amount))
	}
	ref, err := s.nextReference(ctx, "PCV", &account.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference, p.CreatedBy, p.SiteID = ref, actorID(ctx), &account.SiteID
	row, err := q.CreatePettyCashVoucher(ctx, p)
	if err != nil {
		return zero, wrap(err, "create voucher")
	}
	return row, s.audit(ctx, "cash.petty_cash_voucher.write", "petty_cash_vouchers", row.ID, &account.SiteID, nil, row)
}

// VerifyPettyCashVoucher is the Complex Manager's check before the MD approves.
func (s *Service) VerifyPettyCashVoucher(ctx context.Context, p db.ActVerifyPettyCashVoucherParams) (db.PettyCashVoucher, error) {
	before, err := s.GetPettyCashVoucher(ctx, p.ID)
	if err != nil {
		return before, err
	}
	row, err := s.q(ctx).SetVoucherVerified(ctx, db.SetVoucherVerifiedParams{ID: p.ID, VerifiedBy: actorID(ctx)})
	if err != nil {
		return before, stateError(err, "verify this voucher")
	}
	return row, s.audit(ctx, "cash.petty_cash_voucher.verify", "petty_cash_vouchers", row.ID, row.SiteID, before, row)
}

// beforeApprove holds table-specific checks that must pass before a decision.
func (s *Service) beforeApprove(ctx context.Context, r db.ApprovalRequest) error {
	if r.SubjectType == nil || r.SubjectID == nil || *r.SubjectType != "petty_cash_vouchers" {
		return nil
	}
	v, err := s.q(ctx).GetPettyCashVoucher(ctx, *r.SubjectID)
	if err != nil {
		return wrap(err, "read voucher")
	}
	if v.VerifiedBy == nil {
		return &Error{Status: http.StatusConflict, Code: "approval_required", Title: "Verification first",
			Detail: "The Complex Manager must verify this voucher before it can be approved."}
	}
	return nil
}

// IssuePettyCashVoucher pays out exactly the approved amount and updates the
// register at once.
func (s *Service) IssuePettyCashVoucher(ctx context.Context, id uuid.UUID, amount, currency string) (db.PettyCashVoucher, error) {
	before, err := s.GetPettyCashVoucher(ctx, id)
	if err != nil {
		return before, err
	}
	if before.Status != "approved" {
		return before, Fail(http.StatusConflict, "approval_required", "Not approved", "No cash can be issued before approval.")
	}
	if currency != before.AmountCurrency || rat(amount).Cmp(rat(before.AmountAmount)) != 0 {
		return before, errInvalid("/issued_amount", "Only the approved amount can be issued.")
	}
	q := s.q(ctx)
	account, err := q.GetPettyCashAccount(ctx, before.AccountID)
	if err != nil {
		return before, wrap(err, "read account")
	}
	balance := ratp(account.BookBalanceAmount)
	after := sub(balance, rat(amount))
	if after.Sign() < 0 {
		return before, errLimit("cash.petty_cash_max", "There is not enough petty cash to issue this.")
	}
	row, err := q.IssueVoucher(ctx, db.IssueVoucherParams{
		ID: id, Issued: &amount, Currency: &currency, BeforeBalance: ptr(dec(balance)), AfterBalance: ptr(dec(after)), UpdatedBy: actorID(ctx),
	})
	if err != nil {
		return before, stateError(err, "issue this voucher")
	}
	err = q.SetAccountBalance(ctx, db.SetAccountBalanceParams{ID: account.ID, BookBalanceAmount: ptr(dec(after)), BookBalanceCurrency: &currency})
	if err != nil {
		return row, wrap(err, "update register")
	}
	return row, s.audit(ctx, "cash.petty_cash_voucher.issue", "petty_cash_vouchers", row.ID, row.SiteID, before, row)
}

// CreatePettyCashReimbursement claims receipted vouchers. The total is summed
// here, never taken from the client.
func (s *Service) CreatePettyCashReimbursement(ctx context.Context, p db.CreatePettyCashReimbursementParams) (db.PettyCashReimbursement, error) {
	var zero db.PettyCashReimbursement
	q := s.q(ctx)
	sum, err := q.VoucherSum(ctx, p.VoucherIds)
	if err != nil {
		return zero, wrap(err, "sum vouchers")
	}
	if int(sum.Matched) != len(p.VoucherIds) || len(p.VoucherIds) == 0 {
		return zero, errInvalid("/voucher_ids", "Every voucher must be issued and have its receipt attached.")
	}
	account, err := q.GetPettyCashAccount(ctx, p.AccountID)
	if err != nil {
		return zero, notFound(err, "petty cash account")
	}
	ref, err := s.nextReference(ctx, "PCF", &account.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference, p.CreatedBy = ref, actorID(ctx)
	row, err := q.CreatePettyCashReimbursement(ctx, p)
	if err != nil {
		return zero, wrap(err, "create reimbursement")
	}
	row, err = q.SetReimbursementTotal(ctx, db.SetReimbursementTotalParams{ID: row.ID, TotalAmount: &sum.Total, TotalCurrency: &account.Currency})
	if err != nil {
		return zero, wrap(err, "store total")
	}
	return row, s.audit(ctx, "cash.petty_cash_reimbursement.write", "petty_cash_reimbursements", row.ID, &account.SiteID, nil, row)
}

// RecordPaymentPettyCashReimbursement tops the account up and closes the vouchers.
func (s *Service) RecordPaymentPettyCashReimbursement(ctx context.Context, p db.ActRecordPaymentPettyCashReimbursementParams) (db.PettyCashReimbursement, error) {
	before, err := s.GetPettyCashReimbursement(ctx, p.ID)
	if err != nil {
		return before, err
	}
	q := s.q(ctx)
	p.UpdatedBy = actorID(ctx)
	row, err := q.ActRecordPaymentPettyCashReimbursement(ctx, p)
	if err != nil {
		return before, stateError(err, "record this payment")
	}
	if err := q.MarkVouchersReimbursed(ctx, row.VoucherIds); err != nil {
		return row, wrap(err, "close vouchers")
	}
	account, err := q.GetPettyCashAccount(ctx, row.AccountID)
	if err != nil {
		return row, wrap(err, "read account")
	}
	balance := add(ratp(account.BookBalanceAmount), ratp(row.TotalAmount))
	err = q.SetAccountBalance(ctx, db.SetAccountBalanceParams{ID: account.ID, BookBalanceAmount: ptr(dec(balance)), BookBalanceCurrency: &account.Currency})
	if err != nil {
		return row, wrap(err, "update register")
	}
	return row, s.audit(ctx, "cash.petty_cash_reimbursement.record_payment", "petty_cash_reimbursements", row.ID, &account.SiteID, before, row)
}

// CreatePettyCashHandover records the joint count at a change of custodian.
func (s *Service) CreatePettyCashHandover(ctx context.Context, p db.CreatePettyCashHandoverParams) (db.PettyCashHandover, error) {
	var zero db.PettyCashHandover
	if p.OutgoingEmployeeID == p.IncomingEmployeeID {
		return zero, errInvalid("/incoming_employee_id", "The outgoing and incoming custodian must differ.")
	}
	q := s.q(ctx)
	account, err := q.GetPettyCashAccount(ctx, p.AccountID)
	if err != nil {
		return zero, notFound(err, "petty cash account")
	}
	ref, err := s.nextReference(ctx, "PCH", &account.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference, p.CreatedBy = ref, actorID(ctx)
	row, err := q.CreatePettyCashHandover(ctx, p)
	if err != nil {
		return zero, wrap(err, "create handover")
	}
	variance := dec(sub(rat(p.PhysicalBalanceAmount), rat(p.BookBalanceAmount)))
	err = q.SetVarianceFigures(ctx, db.SetVarianceFiguresParams{ID: row.ID, VarianceAmount: &variance, VarianceCurrency: &p.BookBalanceCurrency})
	if err != nil {
		return zero, wrap(err, "store variance")
	}
	row.VarianceAmount, row.VarianceCurrency = &variance, &p.BookBalanceCurrency
	return row, s.audit(ctx, "cash.petty_cash_handover.write", "petty_cash_handovers", row.ID, &account.SiteID, nil, row)
}

// CreatePettyCashReconciliation works out the expected balance and the difference.
func (s *Service) CreatePettyCashReconciliation(ctx context.Context, p db.CreatePettyCashReconciliationParams) (db.PettyCashReconciliation, error) {
	var zero db.PettyCashReconciliation
	q := s.q(ctx)
	account, err := q.GetPettyCashAccount(ctx, p.AccountID)
	if err != nil {
		return zero, notFound(err, "petty cash account")
	}
	ref, err := s.nextReference(ctx, "PCR", &account.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference, p.CreatedBy = ref, actorID(ctx)
	row, err := q.CreatePettyCashReconciliation(ctx, p)
	if err != nil {
		return zero, wrap(err, "create reconciliation")
	}
	expected := sub(add(rat(p.OpeningBalanceAmount), rat(p.ReceiptsAmount)), rat(p.DisbursementsAmount))
	variance := sub(rat(p.PhysicalClosingAmount), expected)
	if err := q.SetReconciliationFigures(ctx, db.SetReconciliationFiguresParams{
		ID: row.ID, ExpectedClosingAmount: ptr(dec(expected)), ExpectedClosingCurrency: &account.Currency, VarianceAmount: ptr(dec(variance)),
	}); err != nil {
		return zero, wrap(err, "store figures")
	}
	row.ExpectedClosingAmount, row.ExpectedClosingCurrency = ptr(dec(expected)), &account.Currency
	row.VarianceAmount, row.VarianceCurrency = ptr(dec(variance)), &account.Currency
	return row, s.audit(ctx, "cash.petty_cash_reconciliation.write", "petty_cash_reconciliations", row.ID, &account.SiteID, nil, row)
}

// CreateSurpriseCount plans a count. Branch roles do not hold this permission,
// so a branch cannot see or pre-schedule one.
func (s *Service) CreateSurpriseCount(ctx context.Context, p db.CreateSurpriseCountParams) (db.SurpriseCount, error) {
	var zero db.SurpriseCount
	ref, err := s.nextReference(ctx, "SCC", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference, p.CreatedBy = ref, actorID(ctx)
	row, err := s.q(ctx).CreateSurpriseCount(ctx, p)
	if err != nil {
		return zero, wrap(err, "create count")
	}
	if p.PosBalanceAmount != nil && p.PhysicalCashAmount != nil {
		variance := dec(sub(ratp(p.PhysicalCashAmount), ratp(p.PosBalanceAmount)))
		if err := s.q(ctx).SetSurpriseVariance(ctx, db.SetSurpriseVarianceParams{ID: row.ID, VarianceAmount: &variance, VarianceCurrency: p.PosBalanceCurrency}); err != nil {
			return zero, wrap(err, "store variance")
		}
		row.VarianceAmount, row.VarianceCurrency = &variance, p.PosBalanceCurrency
	}
	return row, s.audit(ctx, "cash.surprise_count.write", "surprise_counts", row.ID, &row.SiteID, nil, row)
}

// CreateCitReconciliation compares what left, arrived and banked.
func (s *Service) CreateCitReconciliation(ctx context.Context, p db.CreateCitReconciliationParams) (db.CitReconciliation, error) {
	var zero db.CitReconciliation
	if err := canWrite(ctx, &p.SiteID); err != nil {
		return zero, err
	}
	ref, err := s.nextReference(ctx, "CRC", &p.SiteID)
	if err != nil {
		return zero, err
	}
	p.Reference, p.CreatedBy = ref, actorID(ctx)
	row, err := s.q(ctx).CreateCitReconciliation(ctx, p)
	if err != nil {
		return zero, wrap(err, "create reconciliation")
	}
	if p.DepositedAmount != nil {
		variance := dec(sub(rat(p.DispatchedAmount), ratp(p.DepositedAmount)))
		if err := s.q(ctx).SetCitReconVariance(ctx, db.SetCitReconVarianceParams{ID: row.ID, VarianceAmount: &variance, VarianceCurrency: &p.DispatchedCurrency}); err != nil {
			return zero, wrap(err, "store variance")
		}
		row.VarianceAmount, row.VarianceCurrency = &variance, &p.DispatchedCurrency
	}
	return row, s.audit(ctx, "cash.cit_reconciliation.write", "cit_reconciliations", row.ID, &row.SiteID, nil, row)
}
