-- name: GetDocument :one
SELECT * FROM documents WHERE id = $1;

-- name: ListDocument :many
SELECT * FROM documents
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR title ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(document_type)::text IS NULL OR document_type = sqlc.narg(document_type)::text)
  AND (sqlc.narg(department_id)::uuid IS NULL OR department_id = sqlc.narg(department_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR issue_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR issue_date::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateDocument :one
INSERT INTO documents (document_number, title, document_type, department_id, owner_employee_id, next_review_on, storage_location, applicable_site_ids, created_by, updated_by)
VALUES (sqlc.arg(document_number), sqlc.arg(title), sqlc.arg(document_type), sqlc.arg(department_id), sqlc.narg(owner_employee_id), sqlc.narg(next_review_on), sqlc.narg(storage_location), sqlc.narg(applicable_site_ids), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateDocument :one
UPDATE documents
SET document_number = COALESCE(sqlc.narg(document_number), document_number), title = COALESCE(sqlc.narg(title), title), document_type = COALESCE(sqlc.narg(document_type), document_type), department_id = COALESCE(sqlc.narg(department_id), department_id), owner_employee_id = COALESCE(sqlc.narg(owner_employee_id), owner_employee_id), next_review_on = COALESCE(sqlc.narg(next_review_on), next_review_on), storage_location = COALESCE(sqlc.narg(storage_location), storage_location), applicable_site_ids = COALESCE(sqlc.narg(applicable_site_ids), applicable_site_ids), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetDocumentStatus :one
UPDATE documents
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
