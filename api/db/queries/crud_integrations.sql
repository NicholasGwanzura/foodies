-- name: GetIntegration :one
SELECT * FROM integrations WHERE id = $1;

-- name: ListIntegration :many
SELECT * FROM integrations
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateIntegration :one
INSERT INTO integrations (system, mode, config, created_by, updated_by)
VALUES (sqlc.arg(system), sqlc.arg(mode), sqlc.narg(config), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateIntegration :one
UPDATE integrations
SET system = COALESCE(sqlc.narg(system), system), mode = COALESCE(sqlc.narg(mode), mode), config = COALESCE(sqlc.narg(config), config), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetIntegrationStatus :one
UPDATE integrations
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
