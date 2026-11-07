-- name: GetEmployeeDocument :one
SELECT * FROM employee_documents WHERE id = $1;

-- name: ListEmployeeDocument :many
SELECT * FROM employee_documents
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(from_date)::date IS NULL OR issue_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR issue_date::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateEmployeeDocument :one
INSERT INTO employee_documents (employee_id, document_type, issue_date, expiry_date, attachment_id, visibility, created_by, updated_by)
VALUES (sqlc.arg(employee_id), sqlc.arg(document_type), sqlc.narg(issue_date), sqlc.narg(expiry_date), sqlc.arg(attachment_id), sqlc.narg(visibility), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteEmployeeDocument :exec
DELETE FROM employee_documents WHERE id = $1;
