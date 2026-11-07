-- name: GetRoleAssignment :one
SELECT * FROM role_assignments WHERE id = $1;

-- name: ListRoleAssignment :many
SELECT * FROM role_assignments
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(user_id)::uuid IS NULL OR user_id = sqlc.narg(user_id)::uuid)
  AND (sqlc.narg(role_id)::uuid IS NULL OR role_id = sqlc.narg(role_id)::uuid)
  AND (sqlc.narg(scope_id)::uuid IS NULL OR scope_id = sqlc.narg(scope_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR effective_from::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR effective_from::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateRoleAssignment :one
INSERT INTO role_assignments (user_id, role_id, scope_type, scope_id, effective_from, effective_to, created_by, updated_by)
VALUES (sqlc.arg(user_id), sqlc.arg(role_id), sqlc.arg(scope_type), sqlc.narg(scope_id), sqlc.arg(effective_from), sqlc.narg(effective_to), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateRoleAssignment :one
UPDATE role_assignments
SET user_id = COALESCE(sqlc.narg(user_id), user_id), role_id = COALESCE(sqlc.narg(role_id), role_id), scope_type = COALESCE(sqlc.narg(scope_type), scope_type), scope_id = COALESCE(sqlc.narg(scope_id), scope_id), effective_from = COALESCE(sqlc.narg(effective_from), effective_from), effective_to = COALESCE(sqlc.narg(effective_to), effective_to), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: DeleteRoleAssignment :exec
DELETE FROM role_assignments WHERE id = $1;
