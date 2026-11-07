package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	"foodies/api/internal/auth"
	"foodies/api/internal/db"
)

func init() {
	onGuard("IssueStockRequisition", checkIssue)
	onGuard("AcceptStockTransfer", checkTransferHolds)
	onGuard("DispatchStockTransfer", checkTransferHolds)
	onEffect("ReceiveStockTransfer", checkTransferReceipt)
	onEffect("CreateStockReturn", clockReturn)
	onGuard("RequestSkipStocktake", checkSkip)
	onEffect("ApproveStocktake", raiseStockVariances)
	onGuard("CreateDisposal", checkDisposal)
	onGuard("RecordDisposalDisposal", checkDisposalHold)
	onGuard("CompleteBatch", checkBatchComplete)
	onEffect("CompleteBatch", computeYield)
	onGuard("ReleaseBatch", checkBatchRelease)
	onGuard("VerifyDispatchNote", checkDispatchHolds)
	onGuard("SecurityClearDispatchNote", checkDispatchSigned)
}

func dateOf(t time.Time) pgtype.Date { return pgtype.Date{Time: t, Valid: true} }

// requireNoHold stops anything held from moving (spec 6.3).
func (s *Service) requireNoHold(ctx context.Context, site *uuid.UUID, item, batch *uuid.UUID, lot *string) error {
	n, err := s.q(ctx).ActiveHolds(ctx, db.ActiveHoldsParams{SiteID: site, ItemID: item, BatchID: batch, BatchLot: lot})
	if err != nil {
		return wrap(err, "check holds")
	}
	if n > 0 {
		return blocked(http.StatusConflict, "hold_active", "On HOLD", "This stock is on HOLD and cannot be used or moved.",
			Blocker{Kind: "hold_active", Message: "Stock under HOLD"})
	}
	return nil
}

// checkIssue: no issue without an authorised requisition, and held stock stays put.
func checkIssue(ctx context.Context, s *Service, v db.StockRequisition) error {
	if v.Status != "approved" {
		return blocked(http.StatusConflict, "approval_required", "Not authorised", "No issue without an authorised requisition.")
	}
	lines, err := s.q(ctx).ListStockRequisitionLines(ctx, v.ID)
	if err != nil {
		return wrap(err, "read lines")
	}
	for _, l := range lines {
		if err := s.requireNoHold(ctx, &v.SiteID, &l.ItemID, nil, l.BatchLot); err != nil {
			return err
		}
	}
	return nil
}

func checkTransferHolds(ctx context.Context, s *Service, v db.StockTransfer) error {
	lines, err := s.q(ctx).ListStockTransferLines(ctx, v.ID)
	if err != nil {
		return wrap(err, "read lines")
	}
	for _, l := range lines {
		if err := s.requireNoHold(ctx, &v.FromSiteID, &l.ItemID, nil, l.BatchLot); err != nil {
			return err
		}
	}
	return nil
}

// checkTransferReceipt: a difference moves to discrepancy and starts the 24-hour
// TBA follow-up. It must be investigated before confirmation.
func checkTransferReceipt(ctx context.Context, s *Service, v db.StockTransfer) error {
	q := s.q(ctx)
	lines, err := q.ListStockTransferLines(ctx, v.ID)
	if err != nil {
		return wrap(err, "read lines")
	}
	for _, l := range lines {
		if ratp(l.DispatchedQuantity).Cmp(ratp(l.ReceivedQuantity)) != 0 {
			due := s.now().Add(24 * time.Hour)
			return wrap(q.FlagTransferDiscrepancy(ctx, db.FlagTransferDiscrepancyParams{ID: v.ID, TbaFollowupDue: &due}), "flag discrepancy")
		}
	}
	return nil
}

// clockReturn sets the 24-hour follow-up and 48-hour closure clocks (C-04.06).
func clockReturn(ctx context.Context, s *Service, v db.StockReturn) error {
	follow, closure := s.now().Add(24*time.Hour), s.now().Add(48*time.Hour)
	return wrap(s.q(ctx).SetReturnClocks(ctx, db.SetReturnClocksParams{ID: v.ID, FollowupDue: &follow, ClosureDue: &closure}), "set clocks")
}

func checkSkip(_ context.Context, _ *Service, v db.Stocktake) error {
	if strings.TrimSpace(deref(v.AlternativeControl)) == "" {
		return errInvalid("/alternative_control", "A skipped count needs an agreed alternative control.")
	}
	return nil
}

// ReplaceStocktakeLines stores counts. The count that stands is the recount if
// there is one, compared against GAAP Unity, the authoritative balance.
func (s *Service) ReplaceStocktakeLines(ctx context.Context, id uuid.UUID, rows []db.InsertStocktakeLinesParams) error {
	q := s.q(ctx)
	if err := q.DeleteStocktakeLines(ctx, id); err != nil {
		return wrap(err, "clear lines")
	}
	for _, row := range rows {
		counted := row.Count1
		if row.Recount != nil {
			counted = row.Recount
		}
		if counted != nil && row.GaapQuantity != nil {
			row.VarianceQuantity = ptr(dec(sub(rat(*counted), rat(*row.GaapQuantity))))
		}
		row.CreatedBy = actorID(ctx)
		if _, err := q.InsertStocktakeLines(ctx, row); err != nil {
			return wrap(err, "add line")
		}
	}
	return nil
}

// raiseStockVariances opens a variance for every line that differs.
func raiseStockVariances(ctx context.Context, s *Service, v db.Stocktake) error {
	q := s.q(ctx)
	lines, err := q.ListStocktakeLines(ctx, v.ID)
	if err != nil {
		return wrap(err, "read lines")
	}
	for _, l := range lines {
		if ratp(l.VarianceQuantity).Sign() == 0 {
			continue
		}
		ref, err := s.nextReference(ctx, "SVR", &v.SiteID)
		if err != nil {
			return err
		}
		_, err = q.OpenStockVariance(ctx, db.OpenStockVarianceParams{
			Reference: ref, StocktakeID: v.ID, ItemID: l.ItemID, VarianceQuantity: *l.VarianceQuantity, CreatedBy: actorID(ctx),
		})
		if err != nil {
			return wrap(err, "open variance")
		}
	}
	return nil
}

// checkDisposal: never one person alone.
func checkDisposal(ctx context.Context, _ *Service, p *db.CreateDisposalParams) error {
	if a := auth.From(ctx); a != nil && a.EmployeeID == p.WitnessEmployeeID {
		return blocked(http.StatusForbidden, "sod_violation", "Two people needed", "A disposal needs a witness other than you.",
			Blocker{Kind: "sod_conflict", Reference: "disposal_two_people", Message: "A disposal needs at least two people"})
	}
	return nil
}

func checkDisposalHold(ctx context.Context, s *Service, v db.Disposal) error {
	if v.HoldID == nil {
		return nil
	}
	n, err := s.q(ctx).ApprovedDispositions(ctx, *v.HoldID)
	if err != nil {
		return wrap(err, "check disposition")
	}
	if n == 0 {
		return blocked(http.StatusConflict, "hold_active", "No authorised disposition", "Held stock cannot be disposed of without an authorised disposition.")
	}
	return nil
}

// checkBatchComplete: every ingredient needs its supplier lot. Missing labels mean no use.
func checkBatchComplete(ctx context.Context, s *Service, v db.Batch) error {
	lines, err := s.q(ctx).ListBatchIngredients(ctx, v.ID)
	if err != nil {
		return wrap(err, "read ingredients")
	}
	if len(lines) == 0 {
		return errInvalid("/ingredients", "Record the ingredients and their supplier lots.")
	}
	for _, l := range lines {
		if strings.TrimSpace(l.SupplierLot) == "" {
			return errInvalid("/ingredients", "Every ingredient needs its supplier lot. Missing labels mean no use.")
		}
		if err := s.requireNoHold(ctx, nil, &l.ItemID, nil, &l.SupplierLot); err != nil {
			return err
		}
	}
	if v.ActualOutput == nil {
		return errInvalid("/actual_output", "Enter the actual output.")
	}
	return nil
}

// computeYield: yield = usable output / standard or input quantity x 100.
func computeYield(ctx context.Context, s *Service, v db.Batch) error {
	q := s.q(ctx)
	base := ratp(v.PlannedOutput)
	if base.Sign() == 0 {
		base = ratp(v.InputQuantity)
	}
	if base.Sign() == 0 {
		return nil
	}
	yield := mul(new(big.Rat).Quo(ratp(v.ActualOutput), base), big.NewRat(100, 1))
	standard := ratp(v.StandardYieldPercent)
	var standardText *string
	if v.RecipeVersionID != nil {
		rv, err := q.GetRecipeVersion(ctx, *v.RecipeVersionID)
		if err != nil {
			return wrap(err, "read recipe")
		}
		standard, standardText = rat(rv.StandardYieldPercent), &rv.StandardYieldPercent
	}
	variance := sub(yield, standard)
	err := q.SetBatchYield(ctx, db.SetBatchYieldParams{ID: v.ID, ActualYieldPercent: ptr(dec(yield)), StandardYieldPercent: standardText, YieldVariancePercent: ptr(dec(variance))})
	if err != nil {
		return wrap(err, "store yield")
	}
	if standardText == nil {
		return nil
	}
	rule, err := s.RuleValue(ctx, "production.yield_variance_percent", &v.SiteID)
	if err != nil {
		return nil //nolint:nilerr // An unset rule leaves the variance on the record.
	}
	var limit float64
	if json.Unmarshal(rule.Value, &limit) == nil {
		if p, _ := new(big.Rat).Abs(variance).Float64(); p > limit {
			_, err := s.RaiseEscalation(ctx, "batches", v.ID, "quality", 2, &v.SiteID)
			return err
		}
	}
	return nil
}

// checkBatchRelease: no clear status, no release.
func checkBatchRelease(ctx context.Context, s *Service, v db.Batch) error {
	if deref(v.QcStatus) != "released" {
		return blocked(http.StatusConflict, "approval_required", "QC status needed", "This batch has no clear QC status, so it cannot be released.")
	}
	return s.requireNoHold(ctx, nil, nil, &v.ID, nil)
}

func checkDispatchHolds(ctx context.Context, s *Service, v db.DispatchNote) error {
	lines, err := s.q(ctx).ListDispatchNoteLines(ctx, v.ID)
	if err != nil {
		return wrap(err, "read lines")
	}
	if len(lines) == 0 {
		return errInvalid("/lines", "Pick the items first.")
	}
	for _, l := range lines {
		if err := s.requireNoHold(ctx, nil, nil, &l.BatchID, nil); err != nil {
			return err
		}
	}
	return nil
}

// checkDispatchSigned: nothing leaves without a signed D-Note.
func checkDispatchSigned(ctx context.Context, s *Service, v db.DispatchNote) error {
	return s.requireSignatures(ctx, "dispatch_notes", v.ID)
}

// orderDeadline returns the cut-off of a kitchen order cycle: 10:00 and 19:00.
func orderDeadline(order pgtype.Date, cycle string) time.Time {
	at := "10:00"
	if strings.HasPrefix(cycle, "evening") {
		at = "19:00"
	}
	return atTime(order.Time, at)
}

// SubmitKitchenOrder files an order. After the cut-off it needs a reason and
// waits for acknowledgement, and is not guaranteed same-cycle dispatch.
func (s *Service) SubmitKitchenOrder(ctx context.Context, p db.ActSubmitKitchenOrderParams) (db.KitchenOrder, error) {
	before, err := s.GetKitchenOrder(ctx, p.ID)
	if err != nil {
		return before, err
	}
	deadline := orderDeadline(before.OrderDate, before.Cycle)
	late := s.now().After(deadline)
	if late && strings.TrimSpace(deref(before.LateReason)) == "" {
		return before, errInvalid("/late_reason", fmt.Sprintf("This cycle closed at %s. Give a reason.", deadline.In(Harare).Format("15:04")))
	}
	status := "submitted"
	if late {
		status = "late_pending"
	}
	row, err := s.q(ctx).FileKitchenOrder(ctx, db.FileKitchenOrderParams{ID: p.ID, Status: status, DeadlineAt: &deadline, Late: late, UpdatedBy: actorID(ctx)})
	if err != nil {
		return before, stateError(err, "submit this order")
	}
	return row, s.audit(ctx, "production.kitchen_order.submit", "kitchen_orders", row.ID, &row.SiteID, before, row)
}

// ReplaceKitchenOrderLines flags quantities above the par level for a second look.
func (s *Service) ReplaceKitchenOrderLines(ctx context.Context, id uuid.UUID, rows []db.InsertKitchenOrderLinesParams) error {
	q := s.q(ctx)
	if err := q.DeleteKitchenOrderLines(ctx, id); err != nil {
		return wrap(err, "clear lines")
	}
	for _, row := range rows {
		unusual := row.ParLevel != nil && rat(row.Quantity).Cmp(rat(*row.ParLevel)) > 0
		row.UnusualFlag, row.CreatedBy = &unusual, actorID(ctx)
		if _, err := q.InsertKitchenOrderLines(ctx, row); err != nil {
			return wrap(err, "add line")
		}
	}
	return nil
}
