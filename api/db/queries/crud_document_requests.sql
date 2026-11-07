-- name: GetDocumentRequest :one
SELECT * FROM document_requests WHERE id = $1;

-- name: ListDocumentRequest :many
SELECT * FROM document_requests
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR target_on::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR target_on::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateDocumentRequest :one
INSERT INTO document_requests (reference, raised_by_employee_id, department_id, process, gap, trigger, proposed_type, extension_reason, document_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(raised_by_employee_id), sqlc.arg(department_id), sqlc.arg(process), sqlc.arg(gap), sqlc.arg(trigger), sqlc.narg(proposed_type), sqlc.narg(extension_reason), sqlc.narg(document_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateDocumentRequest :one
UPDATE document_requests
SET raised_by_employee_id = COALESCE(sqlc.narg(raised_by_employee_id), raised_by_employee_id), department_id = COALESCE(sqlc.narg(department_id), department_id), process = COALESCE(sqlc.narg(process), process), gap = COALESCE(sqlc.narg(gap), gap), trigger = COALESCE(sqlc.narg(trigger), trigger), proposed_type = COALESCE(sqlc.narg(proposed_type), proposed_type), extension_reason = COALESCE(sqlc.narg(extension_reason), extension_reason), document_id = COALESCE(sqlc.narg(document_id), document_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetDocumentRequestStatus :one
UPDATE document_requests
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActCloseDocumentRequest :one
UPDATE document_requests
SET status = 'closed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('raised', 'classified', 'drafting', 'in_review', 'approved')
RETURNING *;
