// Command server runs the Foodies API.
package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"foodies/api/internal/auth"
	"foodies/api/internal/db"
	httpapi "foodies/api/internal/http"
	"foodies/api/internal/service"
)

const (
	defaultAddr     = ":8080"
	shutdownTimeout = 10 * time.Second
	minSecretLen    = 32
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, logger); err != nil {
		logger.Error("server stopped", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, logger *slog.Logger) error {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return errors.New("DATABASE_URL is not set")
	}
	jwtSecret := os.Getenv("AUTH_JWT_SECRET")
	if len(jwtSecret) < minSecretLen {
		return fmt.Errorf("AUTH_JWT_SECRET must be at least %d characters", minSecretLen)
	}
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = defaultAddr
	}

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return fmt.Errorf("open database pool: %w", err)
	}
	defer pool.Close()

	svc := service.New(db.New(pool), pool, auth.NewSigner(jwtSecret))
	svc.Configure(service.Config{
		UploadDir:     os.Getenv("UPLOAD_DIR"),
		PublicBaseURL: os.Getenv("PUBLIC_API_URL"),
		SigningSecret: jwtSecret,
		WhatsAppToken: os.Getenv("WHATSAPP_VERIFY_TOKEN"),
		WhatsAppKey:   os.Getenv("WHATSAPP_APP_SECRET"),
	})
	go svc.RunBackground(ctx, logger)
	if err := svc.Ping(ctx); err != nil {
		return err
	}

	server := &http.Server{
		Addr:              addr,
		Handler:           httpapi.NewHandler(svc, logger),
		ReadHeaderTimeout: 5 * time.Second,
	}

	serveErr := make(chan error, 1)
	go func() {
		logger.Info("listening", "addr", addr)
		serveErr <- server.ListenAndServe()
	}()

	select {
	case err := <-serveErr:
		return fmt.Errorf("serve: %w", err)
	case <-ctx.Done():
	}

	// The signal context is already cancelled, so shutdown gets its own.
	shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), shutdownTimeout)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("shutdown: %w", err)
	}
	return nil
}
