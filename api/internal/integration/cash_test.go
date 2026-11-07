package integration

import (
	"context"
	"net/http"
	"testing"
)

// goLive gives every rule a workflow needs an active value and switches it on,
// so a test can start from a live workflow.
func goLive(t *testing.T, workflow string, site string) {
	t.Helper()
	ctx := context.Background()
	_, err := pool.Exec(ctx, `
		INSERT INTO rule_versions (rule_definition_id, scope_type, value, effective_from, change_reason, status)
		SELECT d.id, 'group', COALESCE((SELECT v.value FROM rule_versions v WHERE v.rule_definition_id = d.id LIMIT 1), '"set"'::jsonb), now() - interval '1 day', 'test', 'active'
		FROM workflow_requirements r JOIN rule_definitions d ON d."key" = r.rule_key
		WHERE r.workflow_key = $1
		  AND NOT EXISTS (SELECT 1 FROM rule_versions x WHERE x.rule_definition_id = d.id AND x.status = 'active')`, workflow)
	if err != nil {
		t.Fatal(err)
	}
	if err := svc.RefreshReadiness(ctx); err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `UPDATE workflow_activations SET status = 'live' WHERE workflow_key = $1 AND site_id = (SELECT id FROM sites WHERE code = $2)`, workflow, site); err != nil {
		t.Fatal(err)
	}
}

func usd(amount string) map[string]any { return map[string]any{"amount": amount, "currency": "USD"} }

func siteID(t *testing.T, code string) string {
	t.Helper()
	return q1(t, `SELECT id::text FROM sites WHERE code = $1`, code)
}

func activateRule(t *testing.T, key string) {
	t.Helper()
	if _, err := pool.Exec(context.Background(), `
		INSERT INTO rule_versions (rule_definition_id, scope_type, value, effective_from, change_reason, status)
		SELECT d.id, 'group', COALESCE((SELECT v.value FROM rule_versions v WHERE v.rule_definition_id = d.id LIMIT 1), '"set"'::jsonb), now() - interval '1 day', 'test', 'active'
		FROM rule_definitions d WHERE d."key" = $1
		  AND NOT EXISTS (SELECT 1 FROM rule_versions x WHERE x.rule_definition_id = d.id AND x.status = 'active')`, key); err != nil {
		t.Fatal(err)
	}
}

func TestPettyCashFromRequestToReimbursement(t *testing.T) {
	ctx := context.Background()
	site := siteID(t, "HRE01")
	supervisor, complexMgr, md := signIn(t, "E1010"), signIn(t, "E1008"), signIn(t, "E1002")
	finance := mfaToken(t, "E1004")

	account := q1(t, `INSERT INTO petty_cash_accounts (site_id, currency, book_balance_amount, book_balance_currency) VALUES ($1, 'USD', 500, 'USD') RETURNING id::text`, site)
	voucher := map[string]any{"account_id": account, "supervisor_employee_id": employeeIDOf(t, "E1010"), "reason": "Light bulbs", "amount": usd("20.00"), "expense_category": "repairs", "recipient": "Hardware shop"}

	// The reimbursement day has no value at all, so nothing can start yet.
	if err := svc.RefreshReadiness(ctx); err != nil {
		t.Fatal(err)
	}
	r := call(t, supervisor, http.MethodPost, "/cash/petty-cash/vouchers", voucher)
	wantCode(t, r, http.StatusConflict, "workflow_not_live")
	blockers, _ := r.body["blockers"].([]any)
	if len(blockers) == 0 {
		t.Fatalf("no blockers named: %s", r.raw)
	}

	goLive(t, "cash.petty_cash", "HRE01")

	bad := map[string]any{}
	for k, v := range voucher {
		bad[k] = v
	}
	bad["expense_category"] = "Salaries"
	wantCode(t, call(t, supervisor, http.MethodPost, "/cash/petty-cash/vouchers", bad), http.StatusUnprocessableEntity, "limit_exceeded")
	bad["expense_category"], bad["amount"] = "repairs", usd("1000.01")
	wantCode(t, call(t, supervisor, http.MethodPost, "/cash/petty-cash/vouchers", bad), http.StatusUnprocessableEntity, "limit_exceeded")

	v := want(t, call(t, supervisor, http.MethodPost, "/cash/petty-cash/vouchers", voucher), http.StatusCreated)
	id := v.str("id")
	path := "/cash/petty-cash/vouchers/" + id
	wantCode(t, call(t, supervisor, http.MethodPost, path+"/issue", map[string]any{"issued_amount": usd("20.00")}), http.StatusForbidden, "forbidden")
	want(t, call(t, supervisor, http.MethodPost, path+"/submit", nil), http.StatusOK)

	// The MD cannot approve until the Complex Manager has verified.
	request := approvalFor(t, id)
	sig := sign(t, md, employeeIDOf(t, "E1002"), "petty_cash_vouchers", id)
	wantCode(t, call(t, md, http.MethodPost, "/approval-requests/"+request+"/approve", map[string]any{"signature_id": sig}), http.StatusConflict, "approval_required")
	wantCode(t, call(t, complexMgr, http.MethodPost, path+"/issue", map[string]any{"issued_amount": usd("20.00")}), http.StatusConflict, "approval_required")
	want(t, call(t, complexMgr, http.MethodPost, path+"/verify", nil), http.StatusOK)
	want(t, call(t, md, http.MethodPost, "/approval-requests/"+request+"/approve", map[string]any{"signature_id": sig}), http.StatusOK)

	wantCode(t, call(t, complexMgr, http.MethodPost, path+"/issue", map[string]any{"issued_amount": usd("25.00")}), http.StatusUnprocessableEntity, "validation_failed")
	issued := want(t, call(t, complexMgr, http.MethodPost, path+"/issue", map[string]any{"issued_amount": usd("20.00")}), http.StatusOK)
	if got := issued.body["balance_after"].(map[string]any)["amount"]; got != "480.0000" {
		t.Fatalf("balance after = %v", got)
	}
	if got := q1(t, `SELECT book_balance_amount::text FROM petty_cash_accounts WHERE id = $1`, account); got != "480.0000" {
		t.Fatalf("register = %s", got)
	}

	receipt := q1(t, `INSERT INTO attachments (filename, content_type, size_bytes, sha256, status) VALUES ('r.jpg', 'image/jpeg', 1, 'x', 'available') RETURNING id::text`)
	want(t, call(t, complexMgr, http.MethodPost, path+"/attach-receipt", map[string]any{"receipt_attachment_id": receipt}), http.StatusOK)

	claim := want(t, call(t, finance, http.MethodPost, "/cash/petty-cash/reimbursements", map[string]any{"account_id": account, "voucher_ids": []string{id}}), http.StatusCreated)
	if got := claim.body["total"].(map[string]any)["amount"]; got != "20.0000" {
		t.Fatalf("total = %v", got)
	}
}

func TestCashUpNeverNetsCurrenciesAndNeedsAnIndependentReviewer(t *testing.T) {
	ctx := context.Background()
	site := siteID(t, "HRE02")
	goLive(t, "cash.cash_up", "HRE02")
	if _, err := pool.Exec(ctx, `INSERT INTO fx_rates (from_currency, to_currency, rate, rate_date, source) VALUES ('ZWG', 'USD', 0.04, '2026-01-01', 'manual')`); err != nil {
		t.Fatal(err)
	}
	cashier := employeeIDOf(t, "E1012")
	reviewer := signIn(t, "E1008")
	if _, err := pool.Exec(ctx, `UPDATE employees SET employment_status = 'active' WHERE id = $1`, cashier); err != nil {
		t.Fatal(err)
	}
	till := q1(t, `SELECT id::text FROM tills WHERE site_id = $1 LIMIT 1`, site)
	finance := mfaToken(t, "E1004")
	up := want(t, call(t, finance, http.MethodPost, "/cash/cash-ups", map[string]any{
		"site_id": site, "till_id": till, "cashier_employee_id": cashier, "supervisor_employee_id": employeeIDOf(t, "E1009"),
		"business_date": "2026-11-05", "shift": "AM", "kind": "shift",
		"tenders": []map[string]any{
			{"currency": "USD", "tender": "cash", "expected": "100.00", "actual": "90.00"},
			{"currency": "ZWG", "tender": "cash", "expected": "500.00", "actual": "750.00"},
		},
	}), http.StatusCreated)
	id, path := up.str("id"), "/cash/cash-ups/"+up.str("id")
	tenders, _ := up.body["tenders"].([]any)
	if len(tenders) != 2 {
		t.Fatalf("tenders = %s", up.raw)
	}

	// -10 USD plus +250 ZiG at 0.04 is 0 in base currency, yet both variances stand.
	got := want(t, call(t, finance, http.MethodGet, path, nil), http.StatusOK)
	if amount := got.body["variance_base"].(map[string]any)["amount"]; amount != "0" {
		t.Fatalf("variance_base = %v", amount)
	}
	if got.str("fx_rate_value") != "0.0400" {
		t.Fatalf("rate not frozen: %s", got.raw)
	}

	// A draft has nothing to decide on, even for the right role.
	draftSig := sign(t, reviewer, employeeIDOf(t, "E1008"), "cash_ups", id)
	wantCode(t, call(t, reviewer, http.MethodPost, path+"/approve", map[string]any{"signature_id": draftSig}), http.StatusConflict, "invalid_state_transition")

	want(t, call(t, finance, http.MethodPost, path+"/submit", nil), http.StatusOK)

	// CIT cannot start before the day's cash is dealt with.
	activateRule(t, "cash.cit_transit_limit")
	goLive(t, "cash.cit", "HRE02")
	provider := q1(t, `INSERT INTO cit_providers (name) VALUES ('Armour Co') RETURNING id::text`)
	transit := map[string]any{"site_id": site, "provider_id": provider, "transit_date": "2026-11-05", "total": usd("100.00")}
	wantCode(t, call(t, finance, http.MethodPost, "/cash/cit/transits", transit), http.StatusConflict, "conflict")

	// Reviewers must be independent of the people on the cash-up.
	if _, err := pool.Exec(ctx, `UPDATE cash_ups SET supervisor_employee_id = $1 WHERE id = $2`, employeeIDOf(t, "E1008"), id); err != nil {
		t.Fatal(err)
	}
	wantCode(t, call(t, reviewer, http.MethodPost, path+"/approve", map[string]any{"signature_id": id}), http.StatusForbidden, "sod_violation")
	if _, err := pool.Exec(ctx, `UPDATE cash_ups SET supervisor_employee_id = $1 WHERE id = $2`, employeeIDOf(t, "E1009"), id); err != nil {
		t.Fatal(err)
	}
	sig := sign(t, reviewer, employeeIDOf(t, "E1008"), "cash_ups", id)
	approved := want(t, call(t, reviewer, http.MethodPost, path+"/approve", map[string]any{"signature_id": sig}), http.StatusOK)
	if approved.str("status") != "approved" {
		t.Fatalf("got %s", approved.raw)
	}
	if n := q1(t, `SELECT count(*)::text FROM cash_variances WHERE cash_up_id = $1`, id); n != "2" {
		t.Fatalf("variances = %s", n)
	}
	if n := q1(t, `SELECT count(*)::text FROM escalations e JOIN cash_variances v ON v.id = e.subject_id WHERE v.cash_up_id = $1 AND v.amount_currency = 'USD'`, id); n != "1" {
		t.Fatalf("a US$10 shortage should escalate over the US$5 rule, got %s", n)
	}
	// The open variances still hold up the transit.
	wantCode(t, call(t, finance, http.MethodPost, "/cash/cit/transits", transit), http.StatusConflict, "conflict")
	if _, err := pool.Exec(ctx, `UPDATE cash_variances SET status = 'closed' WHERE cash_up_id = $1`, id); err != nil {
		t.Fatal(err)
	}

	over := map[string]any{"site_id": site, "provider_id": provider, "transit_date": "2026-11-05", "total": usd("10000.01")}
	wantCode(t, call(t, finance, http.MethodPost, "/cash/cit/transits", over), http.StatusUnprocessableEntity, "limit_exceeded")

	cit := want(t, call(t, finance, http.MethodPost, "/cash/cit/transits", map[string]any{
		"site_id": site, "provider_id": provider, "transit_date": "2026-11-05", "total": usd("100.00"),
		"bags": []map[string]any{{"bag_number": "B1", "seal_number": "S-1", "amount": usd("100.00")}},
	}), http.StatusCreated)
	citPath := "/cash/cit/transits/" + cit.str("id")
	want(t, call(t, finance, http.MethodPost, citPath+"/seal", nil), http.StatusOK)

	other := q1(t, `INSERT INTO cit_providers (name) VALUES ('Other') RETURNING id::text`)
	stranger := q1(t, `INSERT INTO cit_personnel (cit_provider_id, full_name, id_number) VALUES ($1, 'Stranger', 'X') RETURNING id::text`, other)
	wantCode(t, call(t, finance, http.MethodPost, citPath+"/verify-identity", map[string]any{"cit_person_id": stranger}), http.StatusConflict, "conflict")
	crew := q1(t, `INSERT INTO cit_personnel (cit_provider_id, full_name, id_number) VALUES ($1, 'Crew', 'Y') RETURNING id::text`, provider)
	want(t, call(t, finance, http.MethodPost, citPath+"/verify-identity", map[string]any{"cit_person_id": crew}), http.StatusOK)

	wantCode(t, call(t, finance, http.MethodPost, citPath+"/hand-over", map[string]any{"cit_receipt_number": "R1", "armoured_vehicle_confirmed": false}), http.StatusUnprocessableEntity, "validation_failed")
	want(t, call(t, finance, http.MethodPost, citPath+"/hand-over", map[string]any{"cit_receipt_number": "R1", "armoured_vehicle_confirmed": true}), http.StatusOK)

	// A short vault count starts the clock.
	vault := want(t, call(t, finance, http.MethodPost, citPath+"/receive-vault", map[string]any{"amount": usd("99.00"), "seal_results": []map[string]any{}}), http.StatusOK)
	if vault.str("status") != "discrepancy" {
		t.Fatalf("vault status: %s", vault.raw)
	}
	if n := q1(t, `SELECT count(*)::text FROM tasks WHERE subject_id = $1 AND due_at IS NOT NULL`, cit.str("id")); n != "1" {
		t.Fatalf("no discrepancy task: %s", n)
	}
}

func TestBranchCannotSeeSurpriseCounts(t *testing.T) {
	wantCode(t, call(t, signIn(t, "E1009"), http.MethodGet, "/cash/surprise-counts", nil), http.StatusForbidden, "forbidden")
}
