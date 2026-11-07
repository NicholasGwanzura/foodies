-- name: GetDelegation :one
SELECT * FROM delegations WHERE id = $1;

-- name: ListDelegation :many
SELECT * FROM delegations
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateDelegation :one
INSERT INTO delegations (from_user_id, to_user_id, role_id, workflow_keys, valid_from, valid_to, "reason", created_by, updated_by)
VALUES (sqlc.arg(from_user_id), sqlc.arg(to_user_id), sqlc.narg(role_id), sqlc.narg(workflow_keys), sqlc.arg(valid_from), sqlc.arg(valid_to), sqlc.arg(reason), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateDelegation :one
UPDATE delegations
SET from_user_id = COALESCE(sqlc.narg(from_user_id), from_user_id), to_user_id = COALESCE(sqlc.narg(to_user_id), to_user_id), role_id = COALESCE(sqlc.narg(role_id), role_id), workflow_keys = COALESCE(sqlc.narg(workflow_keys), workflow_keys), valid_from = COALESCE(sqlc.narg(valid_from), valid_from), valid_to = COALESCE(sqlc.narg(valid_to), valid_to), "reason" = COALESCE(sqlc.narg(reason), "reason"), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetDelegationStatus :one
UPDATE delegations
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
