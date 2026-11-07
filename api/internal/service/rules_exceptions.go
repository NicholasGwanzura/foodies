package service

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/google/uuid"

	"foodies/api/internal/auth"
	"foodies/api/internal/db"
)

const posAdminRole = "branch_manager"

func init() {
	onEffect("CreateHold", stopSale)
	onGuard("ReleaseHold", checkHoldRelease)
	onGuard("ExecuteDisposition", checkDispositionHoldOpen)
	onGuard("VerifyCloseNcr", checkNcrVerifier)
	onEffect("CreateNcr", escalateNcr)
	onEffect("CreateRecall", launchRecall)
	onEffect("CreateRecallReconciliation", computeRecallVariance)
	onGuard("ReviewRecallReconciliation", checkRecallReviewer)
	onGuard("CloseRecall", checkRecallClose)
	onGuard("MarkSentRecallCommunication", checkCommunicationApproved)
	onEffect("CreateIncident", escalateIncident)
	onEffect("CreateComplaint", escalateComplaint)
	onGuard("ExecuteServiceRecovery", checkRemedyApproved)
	onGuard("ProposeCorrectionPaymentIncident", checkProviderVerified)
	onGuard("ClosePaymentIncident", checkPaymentReconciled)
	onGuard("CreateManualReceipt", checkManualReceipt)
	onEffect("CreateEquipmentFault", escalateFault)
	onGuard("ClearWorkOrder", checkWorkOrderReport)
	onGuard("CompleteWorkOrder", checkWorkOrderReport)
	onGuard("CreateKeyMovement", checkKeyAvailable)
	onEffect("CreateKeyMovement", markKeyOut)
	onEffect("ReturnKeyMovement", markKeyBack)
	onGuard("SubmitKeyDuplicationRequest", checkKeyMemo)
	onEffect("CreateLostPropertyItem", setLostPropertyDue)
	onGuard("ReturnLostPropertyItem", checkClaimant)
	onGuard("DisposeLostPropertyItem", checkLostPropertyDue)
	onGuard("RecordOpenedStoreOpening", checkOpeningGate)
}

// stopSale: Foodies cannot stop the till, so a HOLD creates a posting and a task
// for the POS administrator, with an escalation until someone confirms (spec 6.3).
func stopSale(ctx context.Context, s *Service, v db.Hold) error {
	if err := s.createPosting(ctx, "gaap_pos", "stop_sale", "holds", v.ID, &v.SiteID, map[string]any{"hold": v.Reference}); err != nil {
		return err
	}
	role, err := s.q(ctx).RoleByKey(ctx, posAdminRole)
	if err != nil {
		return wrap(err, "read role")
	}
	due := s.now().Add(30 * time.Minute)
	if _, err := s.CreateTask(ctx, db.CreateTaskRowParams{
		Title: "Confirm stop-sale for HOLD " + v.Reference, SubjectType: ptr("holds"), SubjectID: &v.ID, SiteID: v.SiteID,
		BusinessDate: s.today(), AssigneeRoleID: &role, DueAt: &due, Priority: ptr("critical"),
	}); err != nil {
		return err
	}
	_, err = s.RaiseEscalation(ctx, "holds", v.ID, "quality", 3, &v.SiteID)
	return err
}

func checkHoldRelease(ctx context.Context, s *Service, v db.Hold) error {
	n, err := s.q(ctx).ApprovedDispositions(ctx, v.ID)
	if err != nil {
		return wrap(err, "check disposition")
	}
	if n == 0 {
		return blocked(http.StatusConflict, "hold_active", "No authorised disposition", "A HOLD needs an authorised disposition before release.",
			Blocker{Kind: "approval_pending", Reference: v.Reference, Message: "Authorised disposition"})
	}
	return nil
}

func checkDispositionHoldOpen(_ context.Context, _ *Service, v db.Disposition) error {
	if v.Status != "approved" {
		return blocked(http.StatusConflict, "approval_required", "Not approved", "A disposition must be approved before it is carried out.")
	}
	return nil
}

func actorIs(ctx context.Context, user *uuid.UUID) bool {
	a := auth.From(ctx)
	return a != nil && user != nil && *user == a.UserID
}

// checkNcrVerifier: effectiveness is verified by someone other than the raiser.
func checkNcrVerifier(ctx context.Context, _ *Service, v db.Ncr) error {
	if actorIs(ctx, v.CreatedBy) {
		return blocked(http.StatusForbidden, "sod_violation", "Someone else must do this", "The person who raised an NCR cannot verify it.")
	}
	return nil
}

func escalateNcr(ctx context.Context, s *Service, v db.Ncr) error {
	if v.SeverityLevel < 3 {
		return nil
	}
	_, err := s.RaiseEscalation(ctx, "ncrs", v.ID, "quality", int(v.SeverityLevel), &v.SiteID)
	return err
}

// launchRecall: any employee can raise one, and it escalates at once. The alert
// goes first and the record follows.
func launchRecall(ctx context.Context, s *Service, v db.Recall) error {
	level := 4
	if v.Kind == "mock" {
		level = 1
	}
	_, err := s.RaiseEscalation(ctx, "recalls", v.ID, "quality", level, v.SiteID)
	return err
}

// computeRecallVariance: produced = stock + distributed + sold + returned + disposed.
func computeRecallVariance(ctx context.Context, s *Service, v db.RecallReconciliation) error {
	accounted := new(big.Rat)
	for _, q := range []string{v.RemainingStock, v.Distributed, v.Sold, v.Returned, v.Disposed} {
		accounted = add(accounted, rat(q))
	}
	return wrap(s.q(ctx).SetRecallVariance(ctx, db.SetRecallVarianceParams{ID: v.ID, Variance: ptr(dec(sub(rat(v.ProducedReceived), accounted)))}), "store variance")
}

func checkRecallReviewer(ctx context.Context, _ *Service, v db.RecallReconciliation) error {
	if actorIs(ctx, v.CreatedBy) {
		return blocked(http.StatusForbidden, "sod_violation", "Independent reviewer needed", "Recall reconciliation has an independent reviewer.",
			Blocker{Kind: "sod_conflict", Reference: "recall_reviewer", Message: "Reviewer must differ from the preparer"})
	}
	return nil
}

// checkRecallClose: closure needs reconciled quantities and started CAPA.
func checkRecallClose(ctx context.Context, s *Service, v db.Recall) error {
	q := s.q(ctx)
	r, err := q.ReviewedRecallReconciliation(ctx, v.ID)
	var blockers []Blocker
	if err != nil || ratp(r.Variance).Sign() != 0 {
		blockers = append(blockers, Blocker{Kind: "approval_pending", Reference: v.Reference, Message: "Reconciled quantities, independently reviewed"})
	}
	capa, err := q.ActionsFrom(ctx, db.ActionsFromParams{SourceType: ptr("recalls"), SourceID: &v.ID})
	if err != nil {
		return wrap(err, "read actions")
	}
	if capa == 0 {
		blockers = append(blockers, Blocker{Kind: "evidence_missing", Reference: v.Reference, Message: "A corrective action started from this recall"})
	}
	if len(blockers) > 0 {
		return blocked(http.StatusConflict, "conflict", "Not ready to close", "This recall cannot close yet.", blockers...)
	}
	return nil
}

func checkCommunicationApproved(_ context.Context, _ *Service, v db.RecallCommunication) error {
	if v.Audience != "internal" && v.Status != "pending_approval" && v.ApprovedBy == nil {
		return blocked(http.StatusConflict, "approval_required", "Approval needed", "External communication must be approved first.")
	}
	return nil
}

// escalateIncident stamps the notification and escalates serious incidents.
func escalateIncident(ctx context.Context, s *Service, v db.Incident) error {
	if err := s.q(ctx).StampIncidentNotified(ctx, v.ID); err != nil {
		return wrap(err, "stamp notification")
	}
	if v.SeverityLevel < 3 {
		return nil
	}
	_, err := s.RaiseEscalation(ctx, "incidents", v.ID, "she", int(v.SeverityLevel), &v.SiteID)
	return err
}

// escalateComplaint: illness, injury or exposure escalates at once, and a
// serious case needs a named owner.
func escalateComplaint(ctx context.Context, s *Service, v db.Complaint) error {
	urgent := v.SeverityLevel >= 3 || slices.Contains([]string{"illness", "injury", "foreign_object", "food_safety"}, v.Category)
	if !urgent {
		return nil
	}
	_, err := s.RaiseEscalation(ctx, "complaints", v.ID, "customer", max(int(v.SeverityLevel), 3), &v.SiteID)
	return err
}

func checkRemedyApproved(_ context.Context, _ *Service, v db.ServiceRecovery) error {
	if v.Status != "approved" {
		return blocked(http.StatusConflict, "approval_required", "Above your authority", "This remedy needs approval before it is given.")
	}
	return nil
}

// checkProviderVerified: a customer's debit alert alone is not proof.
func checkProviderVerified(_ context.Context, _ *Service, v db.PaymentIncident) error {
	if v.ProviderStatus == nil || *v.ProviderStatus == "unknown" {
		return blocked(http.StatusConflict, "conflict", "Check with the provider", "Confirm the status with the provider before proposing a correction.")
	}
	return nil
}

func checkPaymentReconciled(_ context.Context, _ *Service, v db.PaymentIncident) error {
	if !v.SettlementReconciled {
		return blocked(http.StatusConflict, "conflict", "Not reconciled", "This closes after settlement reconciliation.")
	}
	return nil
}

// checkManualReceipt: manual receipting needs the outage authorised, and the
// number must fall inside the issued book.
func checkManualReceipt(ctx context.Context, s *Service, p *db.CreateManualReceiptParams) error {
	q := s.q(ctx)
	outage, err := q.GetPosOutage(ctx, p.PosOutageID)
	if err != nil {
		return notFound(err, "POS outage")
	}
	if outage.Status != "authorised" {
		return blocked(http.StatusConflict, "approval_required", "Not authorised", "Manual receipting starts only after authorisation.")
	}
	book, err := q.GetReceiptBook(ctx, p.ReceiptBookID)
	if err != nil {
		return notFound(err, "receipt book")
	}
	if book.Status != "issued" || p.ReceiptNumber < book.FirstNumber || p.ReceiptNumber > book.LastNumber {
		return errInvalid("/receipt_number", "This number is not in an issued book.")
	}
	return nil
}

func escalateFault(ctx context.Context, s *Service, v db.EquipmentFault) error {
	if v.SeverityLevel < 3 && !v.SafetyImpact {
		return nil
	}
	_, err := s.RaiseEscalation(ctx, "equipment_faults", v.ID, "equipment", max(int(v.SeverityLevel), 3), v.SiteID)
	return err
}

// ConvertEquipmentFault turns a fault into a work order.
func (s *Service) ConvertEquipmentFault(ctx context.Context, p db.ActConvertEquipmentFaultParams) (db.EquipmentFault, error) {
	before, err := s.GetEquipmentFault(ctx, p.ID)
	if err != nil {
		return before, err
	}
	ref, err := s.nextReference(ctx, "WO", before.SiteID)
	if err != nil {
		return before, err
	}
	if _, err := s.q(ctx).CreateWorkOrder(ctx, db.CreateWorkOrderParams{
		Reference: ref, AssetID: before.AssetID, SiteID: before.SiteID, Kind: "corrective", EquipmentFaultID: &before.ID,
		Priority: "high", Description: before.Symptoms, CreatedBy: actorID(ctx),
	}); err != nil {
		return before, wrap(err, "create work order")
	}
	row, err := s.q(ctx).SetEquipmentFaultStatus(ctx, db.SetEquipmentFaultStatusParams{ID: p.ID, Status: "assigned", UpdatedBy: actorID(ctx)})
	if err != nil {
		return before, wrap(err, "update fault")
	}
	return row, s.audit(ctx, "maintenance.equipment_fault.convert", "equipment_faults", row.ID, before.SiteID, before, row)
}

func checkWorkOrderReport(_ context.Context, _ *Service, v db.WorkOrder) error {
	if v.Kind != "preventive" && v.ServiceReportAttachmentID == nil && v.Status == "pending_clearance" {
		return blocked(http.StatusUnprocessableEntity, "evidence_required", "Report needed", "Attach the service report first.")
	}
	return nil
}

func checkKeyAvailable(ctx context.Context, s *Service, p *db.CreateKeyMovementParams) error {
	key, err := s.q(ctx).GetKeyItem(ctx, p.KeyItemID)
	if err != nil {
		return notFound(err, "key")
	}
	if key.Status != "active" {
		return errState("This key is not available.")
	}
	return nil
}

func markKeyOut(ctx context.Context, s *Service, v db.KeyMovement) error {
	_, err := s.q(ctx).SetKeyItemStatus(ctx, db.SetKeyItemStatusParams{ID: v.KeyItemID, Status: "in_use", UpdatedBy: actorID(ctx)})
	return wrap(err, "mark key out")
}

func markKeyBack(ctx context.Context, s *Service, v db.KeyMovement) error {
	_, err := s.q(ctx).SetKeyItemStatus(ctx, db.SetKeyItemStatusParams{ID: v.KeyItemID, Status: "active", UpdatedBy: actorID(ctx)})
	return wrap(err, "mark key back")
}

// checkKeyMemo: duplication is blocked without the signed memo.
func checkKeyMemo(_ context.Context, _ *Service, v db.KeyDuplicationRequest) error {
	if v.MemoAttachmentID == nil {
		return blocked(http.StatusUnprocessableEntity, "evidence_required", "Memo needed", "Attach the signed memo before asking for a duplicate key.")
	}
	return nil
}

// setLostPropertyDue: bank cards 30 days, personal items 90.
func setLostPropertyDue(ctx context.Context, s *Service, v db.LostPropertyItem) error {
	days := 90
	if strings.Contains(v.ItemType, "card") {
		days = 30
	}
	due := dateOf(v.FoundAt.AddDate(0, 0, days))
	return wrap(s.q(ctx).SetLostPropertyDue(ctx, db.SetLostPropertyDueParams{ID: v.ID, DispositionDue: due}), "set due date")
}

func checkClaimant(_ context.Context, _ *Service, v db.LostPropertyItem) error {
	if strings.TrimSpace(deref(v.ClaimantName)) == "" || strings.TrimSpace(deref(v.ClaimantEvidence)) == "" {
		return errInvalid("/claimant_name", "Record who claimed it and how they were verified.")
	}
	return nil
}

func checkLostPropertyDue(_ context.Context, s *Service, v db.LostPropertyItem) error {
	if v.DispositionDue.Valid && s.now().Before(v.DispositionDue.Time) {
		return errState(fmt.Sprintf("This item must be kept until %s.", v.DispositionDue.Time.Format("2 Jan 2006")))
	}
	return nil
}

// checkOpeningGate: a recorded decision is needed before doors open.
func checkOpeningGate(_ context.Context, _ *Service, v db.StoreOpening) error {
	if v.Status != "go" {
		return blocked(http.StatusConflict, "conflict", "Not cleared to open", "This store has no go decision.")
	}
	return nil
}

// DecideStoreOpening records the management decision on the gate. Red never
// opens, and overriding amber needs a rationale.
func (s *Service) DecideStoreOpening(ctx context.Context, id uuid.UUID, decision string, rationale *string, signature uuid.UUID) (db.StoreOpening, error) {
	before, err := s.GetStoreOpening(ctx, id)
	if err != nil {
		return before, err
	}
	if err := s.signedOn(ctx, "store_openings", id, signature); err != nil {
		return before, err
	}
	colour := deref(before.StatusColour)
	if colour == "red" && decision == "go" {
		return before, blocked(http.StatusConflict, "conflict", "Red means no go", "A store on red cannot open until its critical controls pass.")
	}
	if colour == "amber" && decision == "go" && strings.TrimSpace(deref(rationale)) == "" {
		return before, errInvalid("/rationale", "Opening on amber needs a rationale.")
	}
	row, err := s.q(ctx).DecideOpening(ctx, db.DecideOpeningParams{ID: id, Status: decision, Rationale: rationale, DecidedBy: actorID(ctx)})
	if err != nil {
		return before, wrap(err, "record decision")
	}
	return row, s.audit(ctx, "store_ops.store_opening.decide", "store_openings", row.ID, &row.SiteID, before, row)
}

// TraceNode is one step of a trace.
type TraceNode struct {
	Type, ID, Label string
	Site            *uuid.UUID
}

// Trace walks from a batch back to supplier lots and forward to sites. A missing
// link is reported as a gap, and a gap raises an escalation.
func (s *Service) Trace(ctx context.Context, batch uuid.UUID) ([]TraceNode, []string, error) {
	q := s.q(ctx)
	b, err := q.GetBatch(ctx, batch)
	if err != nil {
		return nil, nil, notFound(err, "batch")
	}
	nodes := []TraceNode{{"batch", b.ID.String(), b.BatchNumber, &b.SiteID}}
	var gaps []string
	ingredients, err := q.ListBatchIngredients(ctx, batch)
	if err != nil {
		return nil, nil, wrap(err, "read ingredients")
	}
	if len(ingredients) == 0 {
		gaps = append(gaps, "No ingredients are recorded on this batch.")
	}
	for _, i := range ingredients {
		nodes = append(nodes, TraceNode{Type: "supplier_lot", ID: i.ID.String(), Label: i.SupplierLot})
		if i.GoodsReceiptLineID == nil {
			gaps = append(gaps, "Lot "+i.SupplierLot+" is not linked to a goods receipt.")
			continue
		}
		line, err := q.GetGoodsReceiptOfLine(ctx, *i.GoodsReceiptLineID)
		if err != nil {
			gaps = append(gaps, "Lot "+i.SupplierLot+" has a broken goods receipt link.")
			continue
		}
		nodes = append(nodes, TraceNode{"goods_receipt", line.ID.String(), line.Reference, &line.SiteID})
	}
	sent, err := q.DispatchesOfBatch(ctx, batch)
	if err != nil {
		return nil, nil, wrap(err, "read dispatches")
	}
	for _, d := range sent {
		nodes = append(nodes, TraceNode{"dispatch_note", d.ID.String(), d.Reference, &d.ToSiteID})
	}
	if len(sent) == 0 && b.Status == "closed" {
		gaps = append(gaps, "This released batch has no dispatch record.")
	}
	if len(gaps) > 0 {
		if _, err := s.RaiseEscalation(ctx, "batches", batch, "quality", 3, &b.SiteID); err != nil {
			return nil, nil, err
		}
	}
	return nodes, gaps, nil
}

// ConsolidatedOrders sums on-time orders by product for planning. Late orders
// are logged separately.
func (s *Service) ConsolidatedOrders(ctx context.Context, p db.ConsolidatedOrdersParams) ([]db.ConsolidatedOrdersRow, error) {
	return wrap2(s.q(ctx).ConsolidatedOrders(ctx, p))("consolidate orders")
}
