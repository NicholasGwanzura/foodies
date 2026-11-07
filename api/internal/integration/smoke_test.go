package integration

import (
	"context"
	"net/http"
	"sort"
	"strings"
	"testing"

	"foodies/api/internal/http/apigen"
)

// superToken gives a user a role holding every permission, for sweeps.
func superToken(t *testing.T) string {
	t.Helper()
	ctx := context.Background()
	_, err := pool.Exec(ctx, `INSERT INTO roles ("key", name, permission_keys) VALUES ('test_all', 'Test', '{*}') ON CONFLICT ("key") DO NOTHING`)
	if err != nil {
		t.Fatal(err)
	}
	_, err = pool.Exec(ctx, `INSERT INTO role_assignments (user_id, role_id, scope_type, effective_from)
		SELECT $1::uuid, r.id, 'all', '2020-01-01' FROM roles r WHERE r."key" = 'test_all'
		AND NOT EXISTS (SELECT 1 FROM role_assignments a WHERE a.user_id = $1::uuid AND a.role_id = r.id)`, userIDOf(t, "E1007"))
	if err != nil {
		t.Fatal(err)
	}
	return signIn(t, "E1007")
}

// TestEveryListAnswers sweeps each collection read, so a column the mapping
// cannot carry into its API schema fails here and not in front of a user.
func TestEveryListAnswers(t *testing.T) {
	token := superToken(t)
	names := make([]string, 0, len(apigen.Operations))
	for name := range apigen.Operations {
		names = append(names, name)
	}
	sort.Strings(names)
	checked := 0
	for _, name := range names {
		op := apigen.Operations[name]
		if op.Method != http.MethodGet || strings.Contains(op.Path, "{") || op.Public {
			continue
		}
		r := call(t, token, http.MethodGet, op.Path+"?limit=5", nil)
		// A 400 means the operation has required filters, which this sweep lacks.
		if r.status != http.StatusOK && r.status != http.StatusBadRequest {
			t.Errorf("%s %s: %d %s", name, op.Path, r.status, r.raw)
		}
		checked++
	}
	t.Logf("checked %d reads", checked)
}
