-- name: GetRole :one
SELECT * FROM roles WHERE id = $1;

-- name: ListRole :many
SELECT * FROM roles
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR "name" ILIKE '%' || sqlc.narg(q)::text || '%' OR description ILIKE '%' || sqlc.narg(q)::text || '%')
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateRole :one
INSERT INTO roles ("key", "name", description, permission_keys, requires_mfa, created_by, updated_by)
VALUES (sqlc.arg(key), sqlc.arg(name), sqlc.narg(description), sqlc.arg(permission_keys), COALESCE(sqlc.narg(requires_mfa)::boolean, false), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateRole :one
UPDATE roles
SET "key" = COALESCE(sqlc.narg(key), "key"), "name" = COALESCE(sqlc.narg(name), "name"), description = COALESCE(sqlc.narg(description), description), permission_keys = COALESCE(sqlc.narg(permission_keys), permission_keys), requires_mfa = COALESCE(sqlc.narg(requires_mfa), requires_mfa), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: DeleteRole :exec
DELETE FROM roles WHERE id = $1;
