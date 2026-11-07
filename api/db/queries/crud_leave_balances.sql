-- name: GetLeaveBalance :one
SELECT * FROM leave_balances WHERE id = $1;

-- name: ListLeaveBalance :many
SELECT * FROM leave_balances
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(employee_id)::uuid IS NULL OR employee_id = sqlc.narg(employee_id)::uuid)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);
