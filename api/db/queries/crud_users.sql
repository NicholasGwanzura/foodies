-- name: GetUser :one
SELECT * FROM users WHERE id = $1;

-- name: ListUser :many
SELECT * FROM users
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(employee_id)::uuid IS NULL OR employee_id = sqlc.narg(employee_id)::uuid)
  AND (sqlc.narg(channels)::text IS NULL OR channels = sqlc.narg(channels)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateUser :one
INSERT INTO users (employee_id, channels, created_by, updated_by)
VALUES (sqlc.arg(employee_id), sqlc.arg(channels), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET employee_id = COALESCE(sqlc.narg(employee_id), employee_id), channels = COALESCE(sqlc.narg(channels), channels), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetUserStatus :one
UPDATE users
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
