package integration

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

func q1(t *testing.T, sql string, args ...any) string {
	t.Helper()
	var out string
	if err := pool.QueryRow(context.Background(), sql, args...).Scan(&out); err != nil {
		t.Fatalf("%s: %v", sql, err)
	}
	return out
}

func userIDOf(t *testing.T, number string) string {
	t.Helper()
	return q1(t, `SELECT u.id::text FROM users u JOIN employees e ON e.id = u.employee_id WHERE e.employee_number = $1`, number)
}

func employeeIDOf(t *testing.T, number string) string {
	t.Helper()
	return q1(t, `SELECT id::text FROM employees WHERE employee_number = $1`, number)
}

// signAndDecide signs a subject with the approver's PIN and returns the signature id.
func sign(t *testing.T, token, employee, subjectType, subjectID string) string {
	t.Helper()
	sig := want(t, call(t, token, http.MethodPost, "/signature-requests", map[string]any{
		"subject_type": subjectType, "subject_id": subjectID, "meaning": "approved", "signer_kind": "employee", "signer_employee_id": employee,
	}), http.StatusCreated)
	want(t, call(t, token, http.MethodPost, "/signature-requests/"+sig.str("id")+"/sign", map[string]any{"pin": pin}), http.StatusOK)
	return sig.str("id")
}

// approvalFor finds the pending request for a subject.
func approvalFor(t *testing.T, subjectID string) string {
	t.Helper()
	return q1(t, `SELECT id::text FROM approval_requests WHERE subject_id = $1 AND status = 'pending'`, subjectID)
}

func TestRuleGoesLiveOnlyAfterItsOwnerApproves(t *testing.T) {
	ctx := context.Background()
	admin, gm := sysToken(t), signIn(t, "E1003")
	if err := svc.RefreshReadiness(ctx); err != nil {
		t.Fatal(err)
	}

	site := q1(t, `SELECT id::text FROM sites WHERE code = 'HRE01'`)
	activation := q1(t, `SELECT id::text FROM workflow_activations WHERE workflow_key = 'cash.till_flush' AND site_id = $1`, site)
	wantCode(t, call(t, admin, http.MethodPost, "/workflow-activations/"+activation+"/activate", nil), http.StatusConflict, "workflow_not_live")

	// The seeded proposal does nothing until submitted and approved.
	version := q1(t, `SELECT v.id::text FROM rule_versions v JOIN rule_definitions d ON d.id = v.rule_definition_id WHERE d."key" = 'cash.till_flush_threshold'`)
	definition := q1(t, `SELECT id::text FROM rule_definitions WHERE "key" = 'cash.till_flush_threshold'`)
	want(t, call(t, admin, http.MethodPost, "/rules/"+definition+"/versions/"+version+"/submit", nil), http.StatusOK)

	request := approvalFor(t, version)
	gmEmployee := employeeIDOf(t, "E1003")
	sig := sign(t, gm, gmEmployee, "rule_versions", version)
	want(t, call(t, gm, http.MethodPost, "/approval-requests/"+request+"/approve", map[string]any{"signature_id": sig}), http.StatusOK)

	if got := q1(t, `SELECT status FROM rule_versions WHERE id = $1`, version); got != "active" {
		t.Fatalf("rule version status = %s", got)
	}
	r := want(t, call(t, admin, http.MethodPost, "/workflow-activations/"+activation+"/activate", nil), http.StatusOK)
	if r.str("status") != "live" {
		t.Fatalf("activation = %s", r.raw)
	}
}

func TestApprovalRoutingAndSeparation(t *testing.T) {
	admin, gm, cashier := sysToken(t), signIn(t, "E1003"), signIn(t, "E1011")

	// A rule version gives a real approval to work with.
	definition := q1(t, `SELECT id::text FROM rule_definitions WHERE "key" = 'budget.escalation_working_days'`)
	v := want(t, call(t, admin, http.MethodPost, "/rules/"+definition+"/versions", map[string]any{
		"scope_type": "group", "value": 5, "effective_from": "2026-01-01T00:00:00Z", "change_reason": "test",
	}), http.StatusCreated)
	path := "/rules/" + definition + "/versions/" + v.str("id")
	want(t, call(t, admin, http.MethodPost, path+"/submit", nil), http.StatusOK)
	request := approvalFor(t, v.str("id"))

	// Someone without the step's role cannot decide.
	wantCode(t, call(t, cashier, http.MethodPost, "/approval-requests/"+request+"/approve", map[string]any{"signature_id": v.str("id")}), http.StatusForbidden, "forbidden")
	// The approver must have signed first.
	wantCode(t, call(t, gm, http.MethodPost, "/approval-requests/"+request+"/approve", map[string]any{"signature_id": v.str("id")}), http.StatusNotFound, "not_found")

	// It shows in the approver's inbox, not the requester's.
	inbox := want(t, call(t, gm, http.MethodGet, "/approval-requests?assignee=me", nil), http.StatusOK).list()
	if !strings.Contains(string(mustJSON(t, inbox)), request) {
		t.Fatalf("not in the inbox: %v", inbox)
	}

	wantCode(t, call(t, gm, http.MethodPost, "/approval-requests/"+request+"/return", map[string]any{"comment": ""}), http.StatusUnprocessableEntity, "validation_failed")
	want(t, call(t, gm, http.MethodPost, "/approval-requests/"+request+"/return", map[string]any{"comment": "Check the value"}), http.StatusOK)
	if got := q1(t, `SELECT status FROM rule_versions WHERE id = $1`, v.str("id")); got == "active" {
		t.Fatalf("a returned rule went active: %s", got)
	}
}

func mustJSON(t *testing.T, v any) []byte {
	t.Helper()
	raw, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	return raw
}

func TestSignatureNeedsTheSignersOwnPin(t *testing.T) {
	supervisor, cashier := signIn(t, "E1010"), signIn(t, "E1011")
	form := q1(t, `SELECT id::text FROM form_template_versions LIMIT 1`)
	site := q1(t, `SELECT id::text FROM sites WHERE code = 'HRE01'`)
	sub := want(t, call(t, cashier, http.MethodPost, "/form-submissions", map[string]any{
		"template_version_id": form, "site_id": site, "business_date": "2026-11-07", "captured_at": "2026-11-07T06:00:00Z", "answers": map[string]any{},
	}), http.StatusCreated)

	sig := want(t, call(t, cashier, http.MethodPost, "/signature-requests", map[string]any{
		"subject_type": "form_submissions", "subject_id": sub.str("id"), "meaning": "witnessed", "signer_kind": "employee", "signer_employee_id": employeeIDOf(t, "E1010"),
	}), http.StatusCreated)
	path := "/signature-requests/" + sig.str("id") + "/sign"

	wantCode(t, call(t, cashier, http.MethodPost, path, map[string]any{"pin": pin}), http.StatusForbidden, "forbidden")
	wantCode(t, call(t, supervisor, http.MethodPost, path, map[string]any{"pin": "000000"}), http.StatusUnprocessableEntity, "pin_invalid")
	signed := want(t, call(t, supervisor, http.MethodPost, path, map[string]any{"pin": pin}), http.StatusOK)
	if signed.str("signed_at") == "" || signed.str("content_hash") == "" {
		t.Fatalf("signature not recorded: %s", signed.raw)
	}

	// Editing the record voids what was signed.
	want(t, call(t, cashier, http.MethodPatch, "/form-submissions/"+sub.str("id"), map[string]any{"answers": map[string]any{"notes": "changed"}}, "If-Match", sub.header.Get("ETag")), http.StatusOK)
	if got := q1(t, `SELECT status FROM signature_requests WHERE id = $1`, sig.str("id")); got != "expired" {
		t.Fatalf("signature status after edit = %s", got)
	}
}

func TestPinLocksAfterFiveWrongTries(t *testing.T) {
	token := signIn(t, "E1016")
	form := q1(t, `SELECT id::text FROM form_template_versions LIMIT 1`)
	site := q1(t, `SELECT id::text FROM sites WHERE code = 'HRE01'`)
	sub := want(t, call(t, token, http.MethodPost, "/form-submissions", map[string]any{
		"template_version_id": form, "site_id": site, "business_date": "2026-11-07", "captured_at": "2026-11-07T06:00:00Z", "answers": map[string]any{},
	}), http.StatusCreated)
	sig := want(t, call(t, token, http.MethodPost, "/signature-requests", map[string]any{
		"subject_type": "form_submissions", "subject_id": sub.str("id"), "meaning": "counted", "signer_kind": "employee", "signer_employee_id": employeeIDOf(t, "E1016"),
	}), http.StatusCreated)
	path := "/signature-requests/" + sig.str("id") + "/sign"
	for range 4 {
		wantCode(t, call(t, token, http.MethodPost, path, map[string]any{"pin": "111111"}), http.StatusUnprocessableEntity, "pin_invalid")
	}
	wantCode(t, call(t, token, http.MethodPost, path, map[string]any{"pin": "111111"}), http.StatusUnprocessableEntity, "pin_locked")
	wantCode(t, call(t, token, http.MethodPost, path, map[string]any{"pin": pin}), http.StatusUnprocessableEntity, "pin_locked")
}

func TestScheduleToChecklistWithConsequences(t *testing.T) {
	ctx := context.Background()
	if err := svc.Tick(ctx); err != nil {
		t.Fatal(err)
	}
	if err := svc.Tick(ctx); err != nil { // Running twice must not duplicate.
		t.Fatal(err)
	}
	supervisor := signIn(t, "E1010")
	tasks := want(t, call(t, supervisor, http.MethodGet, "/tasks?assignee=me&status=open", nil), http.StatusOK).list()
	if len(tasks) != 1 {
		t.Fatalf("got %d tasks, want 1: %v", len(tasks), tasks)
	}
	task, _ := tasks[0]["id"].(string)
	site, _ := tasks[0]["site_id"].(string)
	want(t, call(t, supervisor, http.MethodPost, "/tasks/"+task+"/claim", nil), http.StatusOK)

	version := q1(t, `SELECT id::text FROM form_template_versions LIMIT 1`)
	create := map[string]any{
		"template_version_id": version, "site_id": site, "business_date": "2026-11-07", "task_id": task, "captured_at": "2026-11-07T06:00:00Z",
		"answers": map[string]any{"floor_clean": "pass"},
	}
	sub := want(t, call(t, supervisor, http.MethodPost, "/form-submissions", create), http.StatusCreated)
	path := "/form-submissions/" + sub.str("id")
	wantCode(t, call(t, supervisor, http.MethodPost, path+"/submit", nil), http.StatusUnprocessableEntity, "validation_failed")

	answers := map[string]any{"floor_clean": "fail", "tills_on": "ok", "float_counted": "pass", "fridge_temp": 9, "oil_quality": "Good"}
	want(t, call(t, supervisor, http.MethodPatch, path, map[string]any{"answers": answers}, "If-Match", sub.header.Get("ETag")), http.StatusOK)
	done := want(t, call(t, supervisor, http.MethodPost, path+"/submit", nil), http.StatusOK)
	if done.str("result") != "fail" {
		t.Fatalf("result = %s", done.raw)
	}
	if got := q1(t, `SELECT status FROM tasks WHERE id = $1`, task); got != "done" {
		t.Fatalf("task status = %s", got)
	}
	if got := q1(t, `SELECT count(*)::text FROM actions WHERE source_id = $1`, sub.str("id")); got != "1" {
		t.Fatalf("actions created = %s", got)
	}
	esc := q1(t, `SELECT id::text FROM escalations WHERE subject_id = $1`, sub.str("id"))
	if got := q1(t, `SELECT count(*)::text FROM escalation_recipients WHERE escalation_id = $1`, esc); got == "0" {
		t.Fatal("nobody was told about the escalation")
	}
	want(t, call(t, supervisor, http.MethodPost, "/escalations/"+esc+"/acknowledge", map[string]any{}), http.StatusOK)
}

func TestEscalationMovesUpWhenNotAcknowledged(t *testing.T) {
	ctx := context.Background()
	site := q1(t, `SELECT id::text FROM sites WHERE code = 'HRE01'`)
	row := q1(t, `SELECT id::text FROM form_submissions LIMIT 1`)
	_ = site
	var esc string
	err := pool.QueryRow(ctx, `INSERT INTO escalations (subject_type, subject_id, domain, severity_level, site_id, current_step, triggered_at)
		VALUES ('form_submissions', $1, 'cash', 3, $2, 1, now()) RETURNING id::text`, row, site).Scan(&esc)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := pool.Exec(ctx, `INSERT INTO escalation_recipients (escalation_id, step, user_id, channel, notified_at) VALUES ($1, 1, $2, 'in_app', now() - interval '1 hour')`, esc, userIDOf(t, "E1010")); err != nil {
		t.Fatal(err)
	}
	if err := svc.Tick(ctx); err != nil {
		t.Fatal(err)
	}
	if got := q1(t, `SELECT current_step::text FROM escalations WHERE id = $1`, esc); got != "2" {
		t.Fatalf("step = %s, want 2", got)
	}
	if got := q1(t, `SELECT count(*)::text FROM notifications WHERE subject_id = $1 AND recipient_user_id = $2`, esc, userIDOf(t, "E1008")); got == "0" {
		t.Fatal("the next level was not notified")
	}
}

func TestActionOwnerCannotVerifyOwnWork(t *testing.T) {
	branch, complexMgr := signIn(t, "E1009"), signIn(t, "E1008")
	a := want(t, call(t, branch, http.MethodPost, "/actions", map[string]any{
		"kind": "corrective", "description": "Fix the door", "owner_employee_id": employeeIDOf(t, "E1009"), "due_date": "2026-12-01",
	}), http.StatusCreated)
	path := "/actions/" + a.str("id")
	want(t, call(t, branch, http.MethodPost, path+"/complete", map[string]any{}), http.StatusOK)
	verify := map[string]any{"effectiveness_result": "Door closes"}
	wantCode(t, call(t, branch, http.MethodPost, path+"/verify", verify), http.StatusForbidden, "sod_violation")
	done := want(t, call(t, complexMgr, http.MethodPost, path+"/verify", verify), http.StatusOK)
	if done.str("status") != "closed" || done.str("effectiveness_result") != "Door closes" {
		t.Fatalf("got %s", done.raw)
	}
}

func TestDashboardAndSearch(t *testing.T) {
	branch := signIn(t, "E1009")
	d := want(t, call(t, branch, http.MethodGet, "/dashboards/my-day", nil), http.StatusOK)
	if tiles, _ := d.body["tiles"].([]any); len(tiles) == 0 {
		t.Fatalf("no tiles: %s", d.raw)
	}
	r := want(t, call(t, branch, http.MethodGet, "/search?q=FRM", nil), http.StatusOK)
	if len(r.list()) == 0 {
		t.Fatalf("no hits: %s", r.raw)
	}
}

func TestOfflineReplayIsIdempotentAndConflictsGoToSupervisors(t *testing.T) {
	admin := sysToken(t)
	device := want(t, call(t, admin, http.MethodPost, "/devices", map[string]any{"label": "Front tablet", "shared": true}), http.StatusCreated)
	token := signIn(t, "E1010")
	version := q1(t, `SELECT id::text FROM form_template_versions LIMIT 1`)
	site := q1(t, `SELECT id::text FROM sites WHERE code = 'HRE01'`)
	const client = "01890000-0000-7000-8000-000000000001"
	batch := map[string]any{"device_id": device.str("id"), "items": []map[string]any{
		{"client_record_id": client, "operation_id": "createFormSubmission", "captured_at": "2026-11-07T06:00:00Z", "payload": map[string]any{
			"template_version_id": version, "site_id": site, "business_date": "2026-11-06", "captured_at": "2026-11-07T06:00:00Z", "answers": map[string]any{"notes": "offline"},
		}},
		{"client_record_id": "01890000-0000-7000-8000-000000000002", "operation_id": "createFormSubmission", "captured_at": "2026-11-07T06:00:00Z", "payload": map[string]any{"site_id": site}},
	}}
	before := q1(t, `SELECT count(*)::text FROM form_submissions`)
	first := want(t, call(t, token, http.MethodPost, "/sync/batches", batch), http.StatusAccepted)
	items, _ := first.body["items"].([]any)
	if len(items) != 2 || items[0].(map[string]any)["status"] != "applied" || items[1].(map[string]any)["status"] != "conflict" {
		t.Fatalf("items = %s", first.raw)
	}
	if id, _ := items[0].(map[string]any)["record_id"].(string); id == "" {
		t.Fatalf("no record_id: %s", first.raw)
	}
	again := want(t, call(t, token, http.MethodPost, "/sync/batches", batch), http.StatusAccepted)
	repeat, _ := again.body["items"].([]any)
	if repeat[0].(map[string]any)["status"] != "duplicate" {
		t.Fatalf("replay status: %s", again.raw)
	}
	after := q1(t, `SELECT count(*)::text FROM form_submissions`)
	if before == after || q1(t, `SELECT (count(*) = $1::int + 1)::text FROM form_submissions`, before) != "true" {
		t.Fatalf("submissions went from %s to %s", before, after)
	}
	if got := q1(t, `SELECT count(*)::text FROM sync_conflicts WHERE client_record_id = '01890000-0000-7000-8000-000000000002'`); got == "0" {
		t.Fatal("the rejected write was dropped")
	}
}

func TestEvidenceUploadChecksTheChecksum(t *testing.T) {
	token := signIn(t, "E1010")
	a := want(t, call(t, token, http.MethodPost, "/attachments", map[string]any{
		"filename": "a.txt", "content_type": "text/plain", "size_bytes": 5,
		"sha256": "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
	}), http.StatusCreated)
	path := "/attachments/" + a.str("id") + "/complete"
	wantCode(t, call(t, token, http.MethodPost, path, nil), http.StatusConflict, "invalid_state_transition")
}

func TestOffboardingEndsAccessAtOnce(t *testing.T) {
	hr := mfaToken(t, "E1019")
	target := signIn(t, "E1012")
	want(t, call(t, target, http.MethodGet, "/me", nil), http.StatusOK)
	want(t, call(t, hr, http.MethodPost, "/employees/"+employeeIDOf(t, "E1012")+"/status-changes", map[string]any{
		"status": "resigned", "effective_date": "2026-11-07", "reason": "Left the company",
	}), http.StatusCreated)
	wantCode(t, call(t, target, http.MethodGet, "/me", nil), http.StatusUnauthorized, "unauthenticated")
}
