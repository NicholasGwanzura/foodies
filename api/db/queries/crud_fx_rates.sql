-- name: GetFxRate :one
SELECT * FROM fx_rates WHERE id = $1;

-- name: ListFxRate :many
SELECT * FROM fx_rates
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(from_currency)::text IS NULL OR from_currency = sqlc.narg(from_currency)::text)
  AND (sqlc.narg(to_currency)::text IS NULL OR to_currency = sqlc.narg(to_currency)::text)
  AND (sqlc.narg(rate_date)::date IS NULL OR rate_date = sqlc.narg(rate_date)::date)
  AND (sqlc.narg(from_date)::date IS NULL OR rate_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR rate_date::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateFxRate :one
INSERT INTO fx_rates (from_currency, to_currency, rate, rate_date, "source", change_reason, created_by, updated_by)
VALUES (sqlc.arg(from_currency), sqlc.arg(to_currency), sqlc.arg(rate), sqlc.arg(rate_date), sqlc.arg(source), sqlc.narg(change_reason), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateFxRate :one
UPDATE fx_rates
SET from_currency = COALESCE(sqlc.narg(from_currency), from_currency), to_currency = COALESCE(sqlc.narg(to_currency), to_currency), rate = COALESCE(sqlc.narg(rate), rate), rate_date = COALESCE(sqlc.narg(rate_date), rate_date), "source" = COALESCE(sqlc.narg(source), "source"), change_reason = COALESCE(sqlc.narg(change_reason), change_reason), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
