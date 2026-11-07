package integration

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"foodies/api/internal/auth"
	"foodies/api/internal/http/apigen"
)

// TestSeedsNameRealPermissions catches a typo that would quietly deny a role
// its access, or leave a segregation rule enforcing nothing.
func TestSeedsNameRealPermissions(t *testing.T) {
	ctx := context.Background()
	known := map[string]bool{}
	for _, op := range apigen.Operations {
		if op.Permission != "" {
			known[op.Permission] = true
		}
	}
	matchesAny := func(pattern string) bool {
		for key := range known {
			if auth.Matches(pattern, key) {
				return true
			}
		}
		return false
	}
	rows, err := pool.Query(ctx, `SELECT r."key", p FROM roles r, unnest(r.permission_keys) p WHERE r.system`)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var role, pattern string
		if err := rows.Scan(&role, &pattern); err != nil {
			t.Fatal(err)
		}
		if !matchesAny(pattern) {
			t.Errorf("role %s: %q matches no operation", role, pattern)
		}
	}
	sod, err := pool.Query(ctx, `SELECT "key", action_a, action_b FROM sod_rules`)
	if err != nil {
		t.Fatal(err)
	}
	defer sod.Close()
	for sod.Next() {
		var key, a, b string
		if err := sod.Scan(&key, &a, &b); err != nil {
			t.Fatal(err)
		}
		for _, action := range []string{a, b} {
			if !known[action] {
				t.Errorf("sod rule %s: %q is not a permission", key, action)
			}
		}
	}
}

func insertItem(t *testing.T, code, class string) string {
	t.Helper()
	return q1(t, `INSERT INTO items (code, name, category, stock_class, uom, rotation) VALUES ($1, $1, 'food', $2, 'kg', 'fefo') RETURNING id::text`, code, class)
}

func TestHoldStopsStockMovingAndNeedsADisposition(t *testing.T) {
	token := superToken(t)
	site, to := siteID(t, "STR01"), siteID(t, "HRE01")
	item := insertItem(t, "HOLD-ITEM", "dry")

	transfer := want(t, call(t, token, http.MethodPost, "/stock-transfers", map[string]any{
		"from_site_id": site, "to_site_id": to,
		"lines": []map[string]any{{"item_id": item, "requested_quantity": "5", "uom": "kg"}},
	}), http.StatusCreated)
	hold := want(t, call(t, token, http.MethodPost, "/holds", map[string]any{
		"site_id": site, "item_id": item, "quantity": "5", "uom": "kg", "reason": "Suspect", "owner_employee_id": employeeIDOf(t, "E1017"),
	}), http.StatusCreated)

	path := "/stock-transfers/" + transfer.str("id")
	wantCode(t, call(t, token, http.MethodPost, path+"/accept", map[string]any{}), http.StatusConflict, "hold_active")

	// The HOLD told the POS administrator, because Foodies cannot stop the till.
	if n := q1(t, `SELECT count(*)::text FROM postings WHERE subject_id = $1 AND posting_type = 'stop_sale'`, hold.str("id")); n != "1" {
		t.Fatalf("stop-sale postings = %s", n)
	}
	if n := q1(t, `SELECT count(*)::text FROM tasks WHERE subject_id = $1`, hold.str("id")); n != "1" {
		t.Fatalf("stop-sale tasks = %s", n)
	}

	// A signature must be the caller's own, on this record.
	release := "/holds/" + hold.str("id") + "/release"
	wantCode(t, call(t, token, http.MethodPost, release, map[string]any{"signature_id": hold.str("id")}), http.StatusNotFound, "not_found")
	sig := sign(t, token, employeeIDOf(t, "E1007"), "holds", hold.str("id"))
	wantCode(t, call(t, token, http.MethodPost, release, map[string]any{"signature_id": sig}), http.StatusConflict, "hold_active")
}

func TestRecallCannotCloseUnreconciled(t *testing.T) {
	preparer := superToken(t)
	recall := want(t, call(t, preparer, http.MethodPost, "/recalls", map[string]any{
		"kind": "mock", "trigger": "complaint", "description": "Mock",
	}), http.StatusCreated)
	id := recall.str("id")
	wantCode(t, call(t, preparer, http.MethodPost, "/recalls/"+id+"/close", map[string]any{"effectiveness_summary": "Too soon"}), http.StatusConflict, "conflict")

	rec := want(t, call(t, preparer, http.MethodPost, "/recalls/"+id+"/reconciliations", map[string]any{
		"produced_received": "100", "remaining_stock": "10", "distributed": "20", "sold": "50", "returned": "15", "disposed": "4",
	}), http.StatusCreated)
	if got := q1(t, `SELECT variance::text FROM recall_reconciliations WHERE id = $1`, rec.str("id")); got != "1.0000" {
		t.Fatalf("variance = %s", got)
	}
	wantCode(t, call(t, preparer, http.MethodPost, "/recalls/"+id+"/reconciliations/"+rec.str("id")+"/review", map[string]any{}), http.StatusForbidden, "sod_violation")
}

func TestTraceReportsGaps(t *testing.T) {
	token := superToken(t)
	site := siteID(t, "CK01")
	prod := q1(t, `INSERT INTO products (code, name, category) VALUES ('P1', 'Pie', 'bakery') RETURNING id::text`)
	batch := q1(t, `INSERT INTO batches (batch_number, product_id, site_id, production_date, operator_employee_id, uom) VALUES ('B-1', $1, $2, '2026-11-07', $3, 'kg') RETURNING id::text`, prod, site, employeeIDOf(t, "E1014"))
	r := want(t, call(t, token, http.MethodGet, "/traceability?batch_id="+batch, nil), http.StatusOK)
	if r.body["complete"] != false {
		t.Fatalf("got %s", r.raw)
	}
	if gaps, _ := r.body["gaps"].([]any); len(gaps) == 0 {
		t.Fatalf("no gaps: %s", r.raw)
	}
	if n := q1(t, `SELECT count(*)::text FROM escalations WHERE subject_id = $1`, batch); n != "1" {
		t.Fatalf("a gap must escalate, got %s", n)
	}
}

func TestThreeWayMatchHoldsWhatDoesNotMatch(t *testing.T) {
	token := superToken(t)
	supplier := q1(t, `INSERT INTO suppliers (reference, legal_name, supplier_type, status) VALUES ('SUP-T-1', 'Acme', 'local', 'active') RETURNING id::text`)
	due := time.Now().Format("2006-01-02")
	r := want(t, call(t, token, http.MethodPost, "/supplier-invoices", map[string]any{
		"supplier_id": supplier, "invoice_number": "INV-1", "invoice_date": due, "total": usd("100.00"), "invoice_received_at": time.Now().Format(time.RFC3339),
	}), http.StatusCreated)
	wantCode(t, call(t, token, http.MethodPost, "/supplier-invoices", map[string]any{
		"supplier_id": supplier, "invoice_number": "INV-1", "invoice_date": due, "total": usd("100.00"), "invoice_received_at": time.Now().Format(time.RFC3339),
	}), http.StatusConflict, "duplicate_record")
	matched := want(t, call(t, token, http.MethodPost, "/supplier-invoices/"+r.str("id")+"/match", nil), http.StatusOK)
	if matched.str("status") != "on_hold" {
		t.Fatalf("an invoice with no PO must be held: %s", matched.raw)
	}
	// Payment refuses an unmatched invoice.
	pay := want(t, call(t, token, http.MethodPost, "/payment-requisitions", map[string]any{
		"supplier_id": supplier, "invoice_ids": []string{r.str("id")}, "payment_type": "transfer",
	}), http.StatusCreated)
	wantCode(t, call(t, token, http.MethodPost, "/payment-requisitions/"+pay.str("id")+"/submit", nil), http.StatusConflict, "conflict")
}

func TestLostPropertyKeptForItsPeriod(t *testing.T) {
	token := superToken(t)
	item := want(t, call(t, token, http.MethodPost, "/lost-property", map[string]any{
		"site_id": siteID(t, "HRE01"), "item_type": "bank_card", "description": "Card", "found_at": time.Now().Format(time.RFC3339),
		"found_location": "Till 1", "finder_employee_id": employeeIDOf(t, "E1011"), "custodian_employee_id": employeeIDOf(t, "E1010"),
	}), http.StatusCreated)
	want30 := time.Now().AddDate(0, 0, 30).Format("2006-01-02")
	if got := q1(t, `SELECT disposition_due::text FROM lost_property_items WHERE id = $1`, item.str("id")); got != want30 {
		t.Fatalf("due = %s, want %s", got, want30)
	}
	wantCode(t, call(t, token, http.MethodPost, "/lost-property/"+item.str("id")+"/dispose", map[string]any{"disposition": "destroyed"}), http.StatusConflict, "invalid_state_transition")
}

func TestNewReadsNeverFailOnAFilledRow(t *testing.T) {
	token := superToken(t)
	r := want(t, call(t, token, http.MethodGet, "/lost-property?limit=5", nil), http.StatusOK)
	if _, err := json.Marshal(r.list()); err != nil {
		t.Fatal(err)
	}
}
