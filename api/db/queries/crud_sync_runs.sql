-- name: GetSyncRun :one
SELECT * FROM sync_runs WHERE id = $1;

-- name: ListSyncRun :many
SELECT * FROM sync_runs
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateSyncRun :one
INSERT INTO sync_runs (integration_id, created_by, updated_by)
VALUES (sqlc.arg(integration_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetSyncRunStatus :one
UPDATE sync_runs
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
