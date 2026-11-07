-- name: GetEmployeeStatusChange :one
SELECT * FROM employee_status_changes WHERE id = $1;

-- name: ListEmployeeStatusChange :many
SELECT * FROM employee_status_changes
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(from_date)::date IS NULL OR effective_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR effective_date::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateEmployeeStatusChange :one
INSERT INTO employee_status_changes (employee_id, "status", effective_date, "reason", supporting_attachment_id, created_by, updated_by)
VALUES (sqlc.arg(employee_id), sqlc.arg(status), sqlc.arg(effective_date), sqlc.arg(reason), sqlc.narg(supporting_attachment_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;
