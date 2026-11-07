-- name: GetCurrency :one
SELECT * FROM currencies WHERE id = $1;

-- name: ListCurrency :many
SELECT * FROM currencies
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR code ILIKE '%' || sqlc.narg(q)::text || '%' OR "name" ILIKE '%' || sqlc.narg(q)::text || '%')
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateCurrency :one
INSERT INTO currencies (code, "name", symbol, decimal_places, is_base, active, created_by, updated_by)
VALUES (sqlc.arg(code), sqlc.arg(name), sqlc.narg(symbol), sqlc.narg(decimal_places), COALESCE(sqlc.narg(is_base)::boolean, false), COALESCE(sqlc.narg(active)::boolean, true), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateCurrency :one
UPDATE currencies
SET code = COALESCE(sqlc.narg(code), code), "name" = COALESCE(sqlc.narg(name), "name"), symbol = COALESCE(sqlc.narg(symbol), symbol), decimal_places = COALESCE(sqlc.narg(decimal_places), decimal_places), is_base = COALESCE(sqlc.narg(is_base), is_base), active = COALESCE(sqlc.narg(active), active), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
