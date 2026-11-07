package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"foodies/api/internal/auth"
)

// Blocker is one thing stopping an action.
type Blocker struct {
	Kind      string
	Reference string
	Message   string
}

// FieldIssue is one invalid field, addressed by JSON pointer.
type FieldIssue struct {
	Pointer string
	Code    string
	Message string
}

// Error is a failure the client can act on. Anything else is an internal error.
type Error struct {
	Status   int
	Code     string
	Title    string
	Detail   string
	Blockers []Blocker
	Fields   []FieldIssue
}

func (e *Error) Error() string { return e.Code + ": " + e.Detail }

// Fail builds an Error.
func Fail(status int, code, title, detail string) *Error {
	return &Error{Status: status, Code: code, Title: title, Detail: detail}
}

// Common failures.
func errNotFound(thing string) *Error {
	return Fail(http.StatusNotFound, "not_found", "Not found", "No such "+thing+".")
}

func errState(detail string) *Error {
	return Fail(http.StatusConflict, "invalid_state_transition", "Not allowed in this state", detail)
}

func errInvalid(pointer, message string) *Error {
	return &Error{
		Status: http.StatusUnprocessableEntity, Code: "validation_failed", Title: "Validation failed", Detail: message,
		Fields: []FieldIssue{{Pointer: pointer, Code: "invalid", Message: message}},
	}
}

// wrap adds context to a database error, turning constraint failures into
// errors the client can act on.
func wrap(err error, what string) error {
	if err == nil {
		return nil
	}
	var svc *Error
	if errors.As(err, &svc) {
		return svc
	}
	var pg *pgconn.PgError
	if errors.As(err, &pg) {
		switch pg.Code {
		case "23505":
			return Fail(http.StatusConflict, "duplicate_record", "Already exists", "A record with these details already exists.")
		case "23503":
			return errInvalid("", "A linked record does not exist.")
		case "23514":
			return errInvalid("", "A value is not one of the allowed options.")
		case "23502":
			return errInvalid("/"+pg.ColumnName, "This field is required.")
		case "22P02", "22007", "22008", "22003":
			return errInvalid("", "A value is not in the expected format.")
		}
	}
	return fmt.Errorf("%s: %w", what, err)
}

// wrap2 curries wrap for calls returning a value and an error.
func wrap2[T any](v T, err error) func(string) (T, error) {
	return func(what string) (T, error) { return v, wrap(err, what) }
}

// notFound maps a missing row to a 404.
func notFound(err error, thing string) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return errNotFound(thing)
	}
	return wrap(err, "get "+thing)
}

func checkVersion(have, want int32) error {
	if have != want {
		return Fail(http.StatusPreconditionFailed, "stale_version", "Changed since you opened it",
			"This record was changed by someone else. Reload it and try again.")
	}
	return nil
}

// actorID returns the signed-in user, or nil for system work.
func actorID(ctx context.Context) *uuid.UUID {
	if p := auth.From(ctx); p != nil {
		return &p.UserID
	}
	return nil
}

// scopeOf returns the sites the current operation may touch.
func scopeOf(ctx context.Context) (bool, []uuid.UUID) {
	p := auth.From(ctx)
	if p == nil {
		return true, nil
	}
	if p.Active.Sites == nil {
		return p.Active.All, []uuid.UUID{}
	}
	return p.Active.All, p.Active.Sites
}

func allowed(ctx context.Context, site *uuid.UUID) bool {
	if site == nil {
		return true
	}
	all, sites := scopeOf(ctx)
	if all {
		return true
	}
	for _, id := range sites {
		if id == *site {
			return true
		}
	}
	return false
}

// inScope hides a record outside the caller's scope, so a 404 reveals nothing.
func inScope(ctx context.Context, site *uuid.UUID) error {
	if !allowed(ctx, site) {
		return errNotFound("record")
	}
	return nil
}

// canWrite refuses a write at a site outside the caller's scope.
func canWrite(ctx context.Context, site *uuid.UUID) error {
	if !allowed(ctx, site) {
		return Fail(http.StatusForbidden, "out_of_scope", "Outside your sites", "You cannot do this at that site.")
	}
	return nil
}

// Etag formats a row version for the ETag header.
func Etag(version int32) string { return strconv.Itoa(int(version)) }
