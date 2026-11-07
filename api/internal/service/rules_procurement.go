package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"slices"
	"strings"

	"github.com/google/uuid"

	"foodies/api/internal/db"
)

func init() {
	onGuard("SubmitSupplier", checkSupplierDocuments)
	onGuard("SubmitPurchaseRequisition", checkRequisitionLines)
	onGuard("CreatePurchaseOrder", checkPurchaseOrderSource)
	onEffect("CreatePurchaseOrder", stampPurchaseOrder)
	onGuard("SubmitGoodsReceipt", checkReceiptTemperatures)
	onEffect("SubmitGoodsReceipt", applyReceipt)
	onGuard("SubmitPaymentRequisition", checkPaymentDocuments)
	onEffect("SubmitPaymentRequisition", totalPayment)
	onEffect("CreatePaymentBatch", totalBatch)
	onEffect("RecordExecutionPaymentBatch", settleBatch)
	onEffect("PostSupplierInvoice", postInvoiceToSap)
	onEffect("CreateBudgetVariance", assessBudgetVariance)
}

func blocked(status int, code, title, detail string, blockers ...Blocker) error {
	return &Error{Status: status, Code: code, Title: title, Detail: detail, Blockers: blockers}
}

// checkSupplierDocuments: mandatory documents block activation.
func checkSupplierDocuments(_ context.Context, _ *Service, v db.Supplier) error {
	var missing []Blocker
	need := func(ok bool, what string) {
		if !ok {
			missing = append(missing, Blocker{Kind: "evidence_missing", Message: what})
		}
	}
	need(strings.TrimSpace(deref(v.Tin)) != "", "Tax identification number")
	need(v.TaxClearanceExpiry.Valid, "Tax clearance expiry date")
	need(strings.HasPrefix(deref(v.SapCode), "LOSUP") || strings.HasPrefix(deref(v.SapCode), "FOSUP"), "SAP code in the LOSUP or FOSUP series")
	if len(missing) > 0 {
		return blocked(http.StatusUnprocessableEntity, "evidence_required", "Documents missing", "This supplier cannot be activated yet.", missing...)
	}
	return nil
}

func checkRequisitionLines(ctx context.Context, s *Service, v db.PurchaseRequisition) error {
	lines, err := s.q(ctx).ListPurchaseRequisitionLines(ctx, v.ID)
	if err != nil {
		return wrap(err, "read lines")
	}
	if len(lines) == 0 {
		return errInvalid("/lines", "Add at least one line.")
	}
	return nil
}

// checkPurchaseOrderSource: no PR, no PO. The quote count comes from a rule, and
// choosing anything but the lowest quote needs a reason.
func checkPurchaseOrderSource(ctx context.Context, s *Service, p *db.CreatePurchaseOrderParams) error {
	q := s.q(ctx)
	pr, err := q.GetPurchaseRequisition(ctx, p.PurchaseRequisitionID)
	if err != nil {
		return notFound(err, "purchase requisition")
	}
	if pr.Status != "approved" && pr.Status != "sourcing" {
		return blocked(http.StatusConflict, "approval_required", "Requisition not approved", "No approved PR, no purchase order.")
	}
	rule, err := s.RuleValue(ctx, "procurement.quotation_count", &pr.SiteID)
	if err != nil {
		return err
	}
	var count int
	if err := json.Unmarshal(rule.Value, &count); err != nil {
		return fmt.Errorf("decode quotation count: %w", err)
	}
	quotes, err := q.QuotationsFor(ctx, pr.ID)
	if err != nil {
		return wrap(err, "read quotations")
	}
	if len(quotes) < count {
		return blocked(http.StatusUnprocessableEntity, "evidence_required", "Quotations missing",
			fmt.Sprintf("This needs %d quotations, and %d are recorded.", count, len(quotes)))
	}
	var chosen *db.Quotation
	lowest := (*big.Rat)(nil)
	for i := range quotes {
		if quotes[i].Selected {
			chosen = &quotes[i]
		}
		if t := rat(quotes[i].TotalAmount); lowest == nil || t.Cmp(lowest) < 0 {
			lowest = t
		}
	}
	if chosen == nil {
		return errInvalid("/quotations", "Select the winning quotation first.")
	}
	if rat(chosen.TotalAmount).Cmp(lowest) > 0 && strings.TrimSpace(deref(chosen.SelectionReason)) == "" {
		return errInvalid("/selection_reason", "Give a reason for not choosing the lowest quote.")
	}
	return nil
}

// stampPurchaseOrder sets the order date on the server, so a PO cannot be backdated.
func stampPurchaseOrder(ctx context.Context, s *Service, v db.PurchaseOrder) error {
	return wrap(s.q(ctx).SetOrderDate(ctx, db.SetOrderDateParams{ID: v.ID, OrderDate: s.today()}), "stamp order date")
}

// ReplacePurchaseOrderLines stores lines and works out the totals here.
func (s *Service) ReplacePurchaseOrderLines(ctx context.Context, id uuid.UUID, rows []db.InsertPurchaseOrderLinesParams) error {
	q := s.q(ctx)
	if err := q.DeletePurchaseOrderLines(ctx, id); err != nil {
		return wrap(err, "clear lines")
	}
	po, err := q.GetPurchaseOrder(ctx, id)
	if err != nil {
		return wrap(err, "read purchase order")
	}
	total := new(big.Rat)
	for _, row := range rows {
		line := mul(rat(row.Quantity), rat(row.UnitPrice))
		row.LineTotal, row.CreatedBy = ptr(dec(line)), actorID(ctx)
		total = add(total, line)
		if _, err := q.InsertPurchaseOrderLines(ctx, row); err != nil {
			return wrap(err, "add line")
		}
	}
	return wrap(q.SetPurchaseOrderTotal(ctx, db.SetPurchaseOrderTotalParams{ID: id, TotalAmount: ptr(dec(total)), TotalCurrency: &po.Currency}), "store total")
}

// checkReceiptTemperatures: chilled and frozen goods over their limit cannot be
// accepted. Limits come from rules.
func checkReceiptTemperatures(ctx context.Context, s *Service, v db.GoodsReceipt) error {
	lines, err := s.q(ctx).ListGoodsReceiptLines(ctx, v.ID)
	if err != nil {
		return wrap(err, "read lines")
	}
	if len(lines) == 0 {
		return errInvalid("/lines", "Record what arrived.")
	}
	for _, l := range lines {
		if l.Outcome != "accepted" || l.TemperatureC == nil {
			continue
		}
		item, err := s.q(ctx).GetItem(ctx, l.ItemID)
		if err != nil {
			return wrap(err, "read item")
		}
		key := map[string]string{"chilled": "stock.chilled_max_c", "frozen": "stock.frozen_max_c"}[item.StockClass]
		if key == "" {
			continue
		}
		rule, err := s.RuleValue(ctx, key, &v.SiteID)
		if err != nil {
			return err
		}
		var limit float64
		if json.Unmarshal(rule.Value, &limit) == nil {
			if t, _ := ratp(l.TemperatureC).Float64(); t > limit {
				return blocked(http.StatusUnprocessableEntity, "limit_exceeded", "Too warm to accept",
					fmt.Sprintf("%s arrived at %sC, over the limit of %vC. Reject or hold it.", item.Name, *l.TemperatureC, limit),
					Blocker{Kind: "limit_exceeded", Reference: key, Message: "Temperature over the limit"})
			}
		}
	}
	return nil
}

// applyReceipt moves received quantities onto the PO and logs any discrepancy.
func applyReceipt(ctx context.Context, s *Service, v db.GoodsReceipt) error {
	if err := checkBuyerIsNotReceiver(ctx, s, v); err != nil {
		return err
	}
	if v.PurchaseOrderID == nil {
		return nil
	}
	q := s.q(ctx)
	lines, err := q.ListGoodsReceiptLines(ctx, v.ID)
	if err != nil {
		return wrap(err, "read lines")
	}
	for _, l := range lines {
		if l.Outcome == "accepted" || l.Outcome == "partial" {
			if err := q.AddReceivedQuantity(ctx, db.AddReceivedQuantityParams{PurchaseOrderID: *v.PurchaseOrderID, ItemID: &l.ItemID, Quantity: l.ReceivedQuantity}); err != nil {
				return wrap(err, "update PO line")
			}
		}
	}
	return wrap(q.RefreshPurchaseOrderReceipt(ctx, *v.PurchaseOrderID), "update PO status")
}

// checkBuyerIsNotReceiver keeps ordering and receiving apart.
func checkBuyerIsNotReceiver(ctx context.Context, s *Service, v db.GoodsReceipt) error {
	if v.PurchaseOrderID == nil {
		return nil
	}
	po, err := s.q(ctx).GetPurchaseOrder(ctx, *v.PurchaseOrderID)
	if err != nil {
		return wrap(err, "read purchase order")
	}
	if actorIs(ctx, po.CreatedBy) {
		return blocked(http.StatusForbidden, "sod_violation", "Someone else must do this", "The buyer cannot receive their own order.",
			Blocker{Kind: "sod_conflict", Message: "Buyer and receiver are separate"})
	}
	return nil
}

// MatchSupplierInvoice does the three-way match. An unmatched invoice is held.
func (s *Service) MatchSupplierInvoice(ctx context.Context, p db.ActMatchSupplierInvoiceParams) (db.SupplierInvoice, error) {
	before, err := s.GetSupplierInvoice(ctx, p.ID)
	if err != nil {
		return before, err
	}
	if before.Status != "received" && before.Status != "on_hold" {
		return before, errState("This invoice has already been matched.")
	}
	matched, reason, err := s.threeWay(ctx, before)
	if err != nil {
		return before, err
	}
	status, match := "matched", "matched"
	if !matched {
		status, match = "on_hold", "exception"
	}
	row, err := s.q(ctx).SetInvoiceMatch(ctx, db.SetInvoiceMatchParams{ID: before.ID, Status: status, MatchStatus: &match, UpdatedBy: actorID(ctx)})
	if err != nil {
		return before, wrap(err, "store match")
	}
	return row, s.audit(ctx, "procurement.supplier_invoice.match", "supplier_invoices", row.ID, nil, before, map[string]any{"row": row, "reason": reason})
}

func (s *Service) threeWay(ctx context.Context, inv db.SupplierInvoice) (bool, string, error) {
	q := s.q(ctx)
	if inv.PurchaseOrderID == nil {
		return false, "No purchase order", nil
	}
	po, err := q.GetPurchaseOrder(ctx, *inv.PurchaseOrderID)
	if err != nil {
		return false, "", wrap(err, "read PO")
	}
	if po.SupplierID != inv.SupplierID {
		return false, "The supplier differs from the PO", nil
	}
	if ratp(po.TotalAmount).Cmp(rat(inv.TotalAmount)) != 0 {
		return false, "The total differs from the PO", nil
	}
	if len(inv.GoodsReceiptIds) == 0 {
		return false, "No goods receipt", nil
	}
	for _, id := range inv.GoodsReceiptIds {
		grv, err := q.GetGoodsReceipt(ctx, id)
		if err != nil {
			return false, "", wrap(err, "read GRV")
		}
		if grv.PurchaseOrderID == nil || *grv.PurchaseOrderID != po.ID || grv.Status == "draft" || grv.Status == "rejected" {
			return false, "A goods receipt is not valid for this PO", nil
		}
	}
	return true, "", nil
}

// checkPaymentDocuments: payment is blocked without documents, approvals and match.
func checkPaymentDocuments(ctx context.Context, s *Service, v db.PaymentRequisition) error {
	if len(v.InvoiceIds) == 0 {
		return errInvalid("/invoice_ids", "Choose the invoices to pay.")
	}
	var blockers []Blocker
	for _, id := range v.InvoiceIds {
		inv, err := s.q(ctx).GetSupplierInvoice(ctx, id)
		if err != nil {
			return notFound(err, "invoice")
		}
		if inv.SupplierID != v.SupplierID {
			blockers = append(blockers, Blocker{Kind: "evidence_missing", Reference: inv.Reference, Message: "Belongs to another supplier"})
		}
		if !slices.Contains([]string{"matched", "approved_exception", "posted"}, inv.Status) {
			blockers = append(blockers, Blocker{Kind: "evidence_missing", Reference: inv.Reference, Message: "Invoice is not matched"})
		}
	}
	if len(blockers) > 0 {
		return blocked(http.StatusConflict, "conflict", "Not ready to pay", "Every invoice must be matched first.", blockers...)
	}
	return nil
}

// totalPayment sums the invoices and checks tax clearance, which decides WHT.
func totalPayment(ctx context.Context, s *Service, v db.PaymentRequisition) error {
	q := s.q(ctx)
	total, currency := new(big.Rat), ""
	for _, id := range v.InvoiceIds {
		inv, err := q.GetSupplierInvoice(ctx, id)
		if err != nil {
			return wrap(err, "read invoice")
		}
		if currency != "" && currency != inv.TotalCurrency {
			return errInvalid("/invoice_ids", "Pay one currency at a time. Amounts are never netted across currencies.")
		}
		total, currency = add(total, rat(inv.TotalAmount)), inv.TotalCurrency
	}
	supplier, err := q.GetSupplier(ctx, v.SupplierID)
	if err != nil {
		return wrap(err, "read supplier")
	}
	valid := supplier.TaxClearanceExpiry.Valid && supplier.TaxClearanceExpiry.Time.After(s.now())
	return wrap(q.SetPaymentFigures(ctx, db.SetPaymentFiguresParams{ID: v.ID, TotalAmount: ptr(dec(total)), TotalCurrency: &currency, TaxClearanceValid: valid}), "store figures")
}

func totalBatch(ctx context.Context, s *Service, v db.PaymentBatch) error {
	q := s.q(ctx)
	total := new(big.Rat)
	for _, id := range v.RequisitionIds {
		r, err := q.GetPaymentRequisition(ctx, id)
		if err != nil {
			return notFound(err, "payment requisition")
		}
		if r.Status != "approved" {
			return errInvalid("/requisition_ids", r.Reference+" is not approved.")
		}
		if r.TotalCurrency != nil && *r.TotalCurrency != v.Currency {
			return errInvalid("/requisition_ids", r.Reference+" is in another currency.")
		}
		total = add(total, ratp(r.TotalAmount))
	}
	return wrap(q.SetBatchTotal(ctx, db.SetBatchTotalParams{ID: v.ID, TotalAmount: ptr(dec(total)), TotalCurrency: &v.Currency}), "store total")
}

// settleBatch records the bank execution. Payment happens in the portal, and
// this is the proof of it.
func settleBatch(ctx context.Context, s *Service, v db.PaymentBatch) error {
	q := s.q(ctx)
	if v.ProofAttachmentID == nil {
		return errInvalid("/proof_attachment_id", "Attach the bank proof.")
	}
	for _, id := range v.RequisitionIds {
		r, err := q.SetPaymentRequisitionStatus(ctx, db.SetPaymentRequisitionStatusParams{ID: id, Status: "paid", UpdatedBy: actorID(ctx)})
		if err != nil {
			return wrap(err, "mark paid")
		}
		if err := q.MarkInvoicesPaid(ctx, r.InvoiceIds); err != nil {
			return wrap(err, "mark invoices paid")
		}
	}
	return nil
}

func postInvoiceToSap(ctx context.Context, s *Service, v db.SupplierInvoice) error {
	return s.createPosting(ctx, "sap_b1", "supplier_invoice", "supplier_invoices", v.ID, nil, map[string]any{"reference": v.Reference, "total": v.TotalAmount})
}

// assessBudgetVariance works out the percent and escalates what needs it.
func assessBudgetVariance(ctx context.Context, s *Service, v db.BudgetVariance) error {
	budget := rat(v.BudgetAmount)
	if budget.Sign() == 0 {
		return nil
	}
	diff := new(big.Rat).Abs(sub(rat(v.ActualAmount), budget))
	percent := mul(new(big.Rat).Quo(diff, budget), big.NewRat(100, 1))
	q := s.q(ctx)
	if err := q.SetVariancePercent(ctx, db.SetVariancePercentParams{ID: v.ID, VariancePercent: ptr(dec(percent))}); err != nil {
		return wrap(err, "store percent")
	}
	rule, err := s.RuleValue(ctx, "budget.variance_threshold_percent", nil)
	if err != nil {
		return nil //nolint:nilerr // an unset threshold leaves the variance for review
	}
	var limit float64
	if json.Unmarshal(rule.Value, &limit) != nil {
		return nil
	}
	if p, _ := percent.Float64(); p > limit {
		due := s.now().AddDate(0, 0, 7)
		return wrap(q.EscalateBudgetVariance(ctx, db.EscalateBudgetVarianceParams{ID: v.ID, EscalationDue: dateOf(due)}), "escalate variance")
	}
	return nil
}
