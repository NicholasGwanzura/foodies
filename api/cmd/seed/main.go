// Command seed loads reference data, and with -demo a made-up organisation
// with credentials for every demo user. It is safe to run repeatedly.
package main

import (
	"context"
	"embed"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"foodies/api/internal/auth"
	"foodies/api/internal/db"
)

//go:embed sql/*.sql
var seeds embed.FS

func main() {
	demo := flag.Bool("demo", false, "also load the demo organisation")
	flag.Parse()
	if err := run(context.Background(), *demo); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, demo bool) error {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return fmt.Errorf("open database pool: %w", err)
	}
	defer pool.Close()

	names := []string{"reference.sql", "rules.sql"}
	if demo {
		names = append(names, "demo.sql", "pilot.sql")
	}
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("acquire connection: %w", err)
	}
	defer conn.Release()
	for _, name := range names {
		sql, err := seeds.ReadFile("sql/" + name)
		if err != nil {
			return fmt.Errorf("read %s: %w", name, err)
		}
		if _, err := conn.Conn().PgConn().Exec(ctx, string(sql)).ReadAll(); err != nil {
			return fmt.Errorf("run %s: %w", name, err)
		}
		fmt.Println("loaded", name)
	}
	if demo {
		return demoCredentials(ctx, pool)
	}
	return nil
}

// demoCredentials gives every user without a password the demo password and
// PIN from .env.
func demoCredentials(ctx context.Context, pool *pgxpool.Pool) error {
	password, pin := os.Getenv("SEED_PASSWORD"), os.Getenv("SEED_PIN")
	if password == "" || pin == "" {
		return fmt.Errorf("SEED_PASSWORD and SEED_PIN must be set to load the demo")
	}
	q := db.New(pool)
	ids, err := q.ListUsersWithoutPassword(ctx)
	if err != nil {
		return fmt.Errorf("list users: %w", err)
	}
	passwordHash, err := auth.Hash(password)
	if err != nil {
		return err
	}
	pinHash, err := auth.Hash(pin)
	if err != nil {
		return err
	}
	for _, id := range ids {
		if err := q.EnsureCredentials(ctx, id); err != nil {
			return fmt.Errorf("create credentials: %w", err)
		}
		if err := q.SetPassword(ctx, db.SetPasswordParams{UserID: id, PasswordHash: &passwordHash}); err != nil {
			return fmt.Errorf("set password: %w", err)
		}
		if err := q.SetPin(ctx, db.SetPinParams{UserID: id, PinHash: &pinHash}); err != nil {
			return fmt.Errorf("set pin: %w", err)
		}
	}
	fmt.Println("credentials set for", len(ids), "users")
	return nil
}
