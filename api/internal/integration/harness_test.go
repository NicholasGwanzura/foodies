// Package integration drives the whole API over HTTP against a real Postgres
// started by testcontainers. Tests share one database and create what they
// need, so each uses its own records.
package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go/modules/postgres"

	"foodies/api/internal/auth"
	"foodies/api/internal/db"
	httpapi "foodies/api/internal/http"
	"foodies/api/internal/service"
)

const (
	password = "test-password"
	pin      = "246810"
	secret   = "integration-test-secret-integration-test-secret"
)

var (
	api  *httptest.Server
	pool *pgxpool.Pool
	svc  *service.Service
)

func TestMain(m *testing.M) {
	code, err := setup(m)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	os.Exit(code)
}

func setup(m *testing.M) (int, error) {
	ctx := context.Background()
	container, err := postgres.Run(ctx, "postgres:18-alpine",
		postgres.WithDatabase("foodies"), postgres.WithUsername("foodies"), postgres.WithPassword("foodies"),
		postgres.BasicWaitStrategies())
	if err != nil {
		return 0, fmt.Errorf("start postgres: %w", err)
	}
	defer container.Terminate(ctx) //nolint:errcheck // the process is exiting

	url, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return 0, fmt.Errorf("connection string: %w", err)
	}
	pool, err = pgxpool.New(ctx, url)
	if err != nil {
		return 0, fmt.Errorf("open pool: %w", err)
	}
	defer pool.Close()
	if err := migrate(ctx); err != nil {
		return 0, err
	}
	if err := seed(ctx); err != nil {
		return 0, err
	}

	uploads, err := os.MkdirTemp("", "foodies-uploads")
	if err != nil {
		return 0, fmt.Errorf("upload dir: %w", err)
	}
	defer os.RemoveAll(uploads) //nolint:errcheck // temp dir
	svc = service.New(db.New(pool), pool, auth.NewSigner(secret))
	svc.Configure(service.Config{UploadDir: uploads, SigningSecret: secret})
	api = httptest.NewServer(httpapi.NewHandler(svc, slog.New(slog.DiscardHandler)))
	defer api.Close()
	return m.Run(), nil
}

// migrate applies the Up half of each goose migration.
func migrate(ctx context.Context) error {
	files, err := filepath.Glob("../../db/migrations/*.sql")
	if err != nil {
		return fmt.Errorf("list migrations: %w", err)
	}
	sort.Strings(files)
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("acquire: %w", err)
	}
	defer conn.Release()
	for _, f := range files {
		raw, err := os.ReadFile(f)
		if err != nil {
			return fmt.Errorf("read %s: %w", f, err)
		}
		up, _, _ := strings.Cut(string(raw), "-- +goose Down")
		if _, err := conn.Conn().PgConn().Exec(ctx, up).ReadAll(); err != nil {
			return fmt.Errorf("apply %s: %w", f, err)
		}
	}
	return nil
}

func seed(ctx context.Context) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("acquire: %w", err)
	}
	defer conn.Release()
	for _, f := range []string{"../../cmd/seed/sql/reference.sql", "../../cmd/seed/sql/rules.sql", "../../cmd/seed/sql/demo.sql"} {
		raw, err := os.ReadFile(f)
		if err != nil {
			return fmt.Errorf("read %s: %w", f, err)
		}
		if _, err := conn.Conn().PgConn().Exec(ctx, string(raw)).ReadAll(); err != nil {
			return fmt.Errorf("apply %s: %w", f, err)
		}
	}
	q := db.New(pool)
	ids, err := q.ListUsersWithoutPassword(ctx)
	if err != nil {
		return fmt.Errorf("list users: %w", err)
	}
	pw, err := auth.Hash(password)
	if err != nil {
		return err
	}
	ph, err := auth.Hash(pin)
	if err != nil {
		return err
	}
	for _, id := range ids {
		if err := q.EnsureCredentials(ctx, id); err != nil {
			return fmt.Errorf("credentials: %w", err)
		}
		if err := q.SetPassword(ctx, db.SetPasswordParams{UserID: id, PasswordHash: &pw}); err != nil {
			return fmt.Errorf("password: %w", err)
		}
		if err := q.SetPin(ctx, db.SetPinParams{UserID: id, PinHash: &ph}); err != nil {
			return fmt.Errorf("pin: %w", err)
		}
	}
	return nil
}

// response is a decoded API reply.
type response struct {
	status int
	header http.Header
	body   map[string]any
	raw    []byte
}

func (r response) str(key string) string {
	v, _ := r.body[key].(string)
	return v
}

func (r response) list() []map[string]any {
	items, _ := r.body["data"].([]any)
	out := make([]map[string]any, 0, len(items))
	for _, it := range items {
		if m, ok := it.(map[string]any); ok {
			out = append(out, m)
		}
	}
	return out
}

// call sends a request. A POST gets a fresh Idempotency-Key unless headers set one.
func call(t *testing.T, token, method, path string, body any, headers ...string) response {
	t.Helper()
	var reader bytes.Reader
	if body != nil {
		raw, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		reader = *bytes.NewReader(raw)
	}
	req, err := http.NewRequestWithContext(context.Background(), method, api.URL+"/api/v1"+path, &reader)
	if err != nil {
		t.Fatal(err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	if method == http.MethodPost {
		req.Header.Set("Idempotency-Key", uuid.NewString())
	}
	for i := 0; i+1 < len(headers); i += 2 {
		req.Header.Set(headers[i], headers[i+1])
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close() //nolint:errcheck // nothing to do with a failed close
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(res.Body); err != nil {
		t.Fatal(err)
	}
	out := response{status: res.StatusCode, header: res.Header, raw: buf.Bytes()}
	if buf.Len() > 0 {
		_ = json.Unmarshal(buf.Bytes(), &out.body)
	}
	return out
}

func want(t *testing.T, r response, status int) response {
	t.Helper()
	if r.status != status {
		t.Fatalf("status = %d, want %d: %s", r.status, status, r.raw)
	}
	return r
}

func wantCode(t *testing.T, r response, status int, code string) {
	t.Helper()
	if r.status != status || r.str("code") != code {
		t.Fatalf("got %d %q, want %d %q: %s", r.status, r.str("code"), status, code, r.raw)
	}
}

// signIn returns an access token for a seeded employee whose role has no MFA.
func signIn(t *testing.T, number string) string {
	t.Helper()
	r := want(t, call(t, "", http.MethodPost, "/auth/login", map[string]any{"employee_number": number, "password": password}), http.StatusOK)
	return r.str("access_token")
}

// sysToken signs in the system admin, enrolling two-factor on first use.
func sysToken(t *testing.T) string {
	t.Helper()
	return mfaToken(t, "E1001")
}
