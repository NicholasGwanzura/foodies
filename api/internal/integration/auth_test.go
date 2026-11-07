package integration

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"

	"foodies/api/internal/auth"
)

// mfaToken signs in an MFA role, enrolling the authenticator if needed.
func mfaToken(t *testing.T, number string) string {
	t.Helper()
	login := map[string]any{"employee_number": number, "password": password}
	r := want(t, call(t, "", http.MethodPost, "/auth/login", login), http.StatusOK)
	if r.str("status") == "authenticated" {
		limited := r.str("access_token")
		e := want(t, call(t, limited, http.MethodPost, "/auth/mfa/enrol", nil), http.StatusOK)
		code, err := auth.TOTPCode(e.str("secret"), time.Now())
		if err != nil {
			t.Fatal(err)
		}
		want(t, call(t, limited, http.MethodPost, "/auth/mfa/confirm", map[string]any{"code": code}), http.StatusNoContent)
		r = want(t, call(t, "", http.MethodPost, "/auth/login", login), http.StatusOK)
	}
	var secretValue string
	if err := pool.QueryRow(context.Background(),
		`SELECT c.totp_secret FROM user_credentials c JOIN users u ON u.id = c.user_id
		 JOIN employees e ON e.id = u.employee_id WHERE e.employee_number = $1`, number).Scan(&secretValue); err != nil {
		t.Fatal(err)
	}
	code, err := auth.TOTPCode(secretValue, time.Now())
	if err != nil {
		t.Fatal(err)
	}
	v := want(t, call(t, "", http.MethodPost, "/auth/mfa/verify", map[string]any{
		"mfa_challenge_id": r.str("mfa_challenge_id"), "code": code,
	}), http.StatusOK)
	return v.str("access_token")
}

func TestSignInLocksAfterFiveFailures(t *testing.T) {
	bad := map[string]any{"employee_number": "E1022", "password": "nope"}
	for range 4 {
		wantCode(t, call(t, "", http.MethodPost, "/auth/login", bad), http.StatusUnauthorized, "unauthenticated")
	}
	wantCode(t, call(t, "", http.MethodPost, "/auth/login", bad), http.StatusUnauthorized, "account_locked")
	good := map[string]any{"employee_number": "E1022", "password": password}
	wantCode(t, call(t, "", http.MethodPost, "/auth/login", good), http.StatusUnauthorized, "account_locked")
}

func TestUnknownEmployeeLooksLikeWrongPassword(t *testing.T) {
	r := call(t, "", http.MethodPost, "/auth/login", map[string]any{"employee_number": "NOPE", "password": "x"})
	wantCode(t, r, http.StatusUnauthorized, "unauthenticated")
}

func TestRefreshRotatesAndReuseEndsTheSession(t *testing.T) {
	first := want(t, call(t, "", http.MethodPost, "/auth/login", map[string]any{"employee_number": "E1021", "password": password}), http.StatusOK)
	spent := first.str("refresh_token")
	second := want(t, call(t, "", http.MethodPost, "/auth/refresh", map[string]any{"refresh_token": spent}), http.StatusOK)
	if second.str("refresh_token") == spent {
		t.Fatal("refresh token did not rotate")
	}

	wantCode(t, call(t, "", http.MethodPost, "/auth/refresh", map[string]any{"refresh_token": spent}), http.StatusUnauthorized, "token_expired")
	// The thief's replay kills the family, including the newest token.
	wantCode(t, call(t, "", http.MethodPost, "/auth/refresh", map[string]any{"refresh_token": second.str("refresh_token")}), http.StatusUnauthorized, "token_expired")
	wantCode(t, call(t, second.str("access_token"), http.MethodGet, "/me", nil), http.StatusUnauthorized, "unauthenticated")
}

func TestMfaRoleCannotWorkBeforeEnrolling(t *testing.T) {
	r := want(t, call(t, "", http.MethodPost, "/auth/login", map[string]any{"employee_number": "E1005", "password": password}), http.StatusOK)
	limited := r.str("access_token")
	want(t, call(t, limited, http.MethodGet, "/me", nil), http.StatusOK)
	wantCode(t, call(t, limited, http.MethodGet, "/me/permissions", nil), http.StatusUnauthorized, "mfa_required")
	wantCode(t, call(t, limited, http.MethodGet, "/sites", nil), http.StatusUnauthorized, "mfa_required")

	mfaToken(t, "E1005")
	second := want(t, call(t, "", http.MethodPost, "/auth/login", map[string]any{"employee_number": "E1005", "password": password}), http.StatusOK)
	if second.str("status") != "mfa_required" || second.str("access_token") != "" {
		t.Fatalf("enrolled user got %s", second.raw)
	}
	wantCode(t, call(t, "", http.MethodPost, "/auth/mfa/verify", map[string]any{
		"mfa_challenge_id": second.str("mfa_challenge_id"), "code": "000000",
	}), http.StatusUnauthorized, "unauthenticated")
}

func TestLogoutEndsTheToken(t *testing.T) {
	token := signIn(t, "E1009")
	want(t, call(t, token, http.MethodPost, "/auth/logout", nil), http.StatusNoContent)
	wantCode(t, call(t, token, http.MethodGet, "/me", nil), http.StatusUnauthorized, "unauthenticated")
}

func TestPermissionAndScope(t *testing.T) {
	cashier, branchByo := signIn(t, "E1011"), signIn(t, "E1021")

	sites := want(t, call(t, branchByo, http.MethodGet, "/sites", nil), http.StatusOK).list()
	if len(sites) != 1 || sites[0]["code"] != "BYO01" {
		t.Fatalf("a branch manager sees %v", sites)
	}
	wantCode(t, call(t, cashier, http.MethodGet, "/users", nil), http.StatusForbidden, "forbidden")

	// A record outside your scope is a 404, so its existence stays hidden.
	var other string
	own := want(t, call(t, cashier, http.MethodGet, "/tills", nil), http.StatusOK).list()
	if len(own) == 0 {
		t.Fatal("no tills in scope")
	}
	if err := pool.QueryRow(context.Background(), `SELECT t.id FROM tills t JOIN sites s ON s.id = t.site_id WHERE s.code = 'BYO01' LIMIT 1`).Scan(&other); err != nil {
		t.Fatal(err)
	}
	wantCode(t, call(t, cashier, http.MethodGet, "/tills/"+other, nil), http.StatusNotFound, "not_found")
}

func TestSafeWrites(t *testing.T) {
	admin := sysToken(t)
	body := map[string]any{"code": "TST", "name": "Test complex"}
	k1, k2 := uuid.NewString(), uuid.NewString()

	first := want(t, call(t, admin, http.MethodPost, "/complexes", body, "Idempotency-Key", k1), http.StatusCreated)
	again := want(t, call(t, admin, http.MethodPost, "/complexes", body, "Idempotency-Key", k1), http.StatusCreated)
	if first.str("id") != again.str("id") || again.header.Get("Idempotent-Replayed") != "true" {
		t.Fatalf("replay differs: %s vs %s", first.raw, again.raw)
	}
	wantCode(t, call(t, admin, http.MethodPost, "/complexes", map[string]any{"code": "OTHER", "name": "x"}, "Idempotency-Key", k1), http.StatusConflict, "idempotency_key_reused")
	wantCode(t, call(t, admin, http.MethodPost, "/complexes", body), http.StatusConflict, "duplicate_record")

	// A rejected write frees its key, so the corrected request can use it.
	wantCode(t, call(t, admin, http.MethodPost, "/complexes", map[string]any{"name": "No code"}, "Idempotency-Key", k2), http.StatusUnprocessableEntity, "validation_failed")
	want(t, call(t, admin, http.MethodPost, "/complexes", map[string]any{"code": "TS2", "name": "Fixed"}, "Idempotency-Key", k2), http.StatusCreated)

	path := "/complexes/" + first.str("id")
	etag := first.header.Get("ETag")
	updated := want(t, call(t, admin, http.MethodPatch, path, map[string]any{"name": "Renamed"}, "If-Match", etag), http.StatusOK)
	if updated.str("name") != "Renamed" || updated.str("code") != "TST" {
		t.Fatalf("patch changed the wrong fields: %s", updated.raw)
	}
	wantCode(t, call(t, admin, http.MethodPatch, path, map[string]any{"name": "Stale"}, "If-Match", etag), http.StatusPreconditionFailed, "stale_version")
}

func TestPagination(t *testing.T) {
	admin := sysToken(t)
	seen := map[string]bool{}
	path := "/sites?limit=2"
	for pages := 0; ; pages++ {
		if pages > 10 {
			t.Fatal("cursor never ended")
		}
		r := want(t, call(t, admin, http.MethodGet, path, nil), http.StatusOK)
		for _, s := range r.list() {
			id, _ := s["id"].(string)
			if seen[id] {
				t.Fatalf("site %s repeated", id)
			}
			seen[id] = true
		}
		page, _ := r.body["page"].(map[string]any)
		next, _ := page["next_cursor"].(string)
		if next == "" {
			break
		}
		path = "/sites?limit=2&cursor=" + next
	}
	if len(seen) != 6 {
		t.Fatalf("saw %d sites, want 6", len(seen))
	}
}

func TestAuditChainHoldsAndIsAppendOnly(t *testing.T) {
	admin := sysToken(t)
	want(t, call(t, admin, http.MethodPost, "/complexes", map[string]any{"code": "AUD", "name": "Audited"}), http.StatusCreated)

	ctx := context.Background()
	rows, err := pool.Query(ctx, `SELECT prev_hash, hash FROM audit_events ORDER BY seq`)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()
	last, n := "", 0
	for rows.Next() {
		var prev, hash string
		if err := rows.Scan(&prev, &hash); err != nil {
			t.Fatal(err)
		}
		if prev != last {
			t.Fatalf("chain broken at row %d", n)
		}
		last, n = hash, n+1
	}
	if n == 0 {
		t.Fatal("no audit events")
	}
	if _, err := pool.Exec(ctx, `UPDATE audit_events SET action = 'tampered'`); err == nil {
		t.Fatal("audit_events accepted an update")
	}
	if _, err := pool.Exec(ctx, `DELETE FROM audit_events`); err == nil {
		t.Fatal("audit_events accepted a delete")
	}
}
