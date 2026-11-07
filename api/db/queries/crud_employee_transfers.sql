-- name: GetEmployeeTransfer :one
SELECT * FROM employee_transfers WHERE id = $1;

-- name: ListEmployeeTransfer :many
SELECT * FROM employee_transfers
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR effective_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR effective_date::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateEmployeeTransfer :one
INSERT INTO employee_transfers (employee_id, from_site_id, to_site_id, from_department_id, to_department_id, new_job_title, new_manager_employee_id, effective_date, "reason", created_by, updated_by)
VALUES (sqlc.arg(employee_id), sqlc.narg(from_site_id), sqlc.narg(to_site_id), sqlc.narg(from_department_id), sqlc.narg(to_department_id), sqlc.narg(new_job_title), sqlc.narg(new_manager_employee_id), sqlc.arg(effective_date), sqlc.arg(reason), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetEmployeeTransferStatus :one
UPDATE employee_transfers
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitEmployeeTransfer :one
UPDATE employee_transfers
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
