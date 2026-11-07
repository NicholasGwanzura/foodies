-- name: GetKeyMovement :one
SELECT * FROM key_movements WHERE id = $1;

-- name: ListKeyMovement :many
SELECT * FROM key_movements
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(key_item_id)::uuid IS NULL OR key_item_id = sqlc.narg(key_item_id)::uuid)
  AND (sqlc.narg(holder_employee_id)::uuid IS NULL OR holder_employee_id = sqlc.narg(holder_employee_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateKeyMovement :one
INSERT INTO key_movements (key_item_id, holder_employee_id, purpose, issued_at, condition, movement_kind, created_by, updated_by)
VALUES (sqlc.arg(key_item_id), sqlc.arg(holder_employee_id), sqlc.arg(purpose), sqlc.arg(issued_at), sqlc.narg(condition), sqlc.narg(movement_kind), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetKeyMovementStatus :one
UPDATE key_movements
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActReturnKeyMovement :one
UPDATE key_movements
SET condition = COALESCE(sqlc.narg(condition), condition), returned_at = now(), status = 'returned', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('out', 'overdue')
RETURNING *;
