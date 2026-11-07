-- name: GetDocumentAcknowledgement :one
SELECT * FROM document_acknowledgements WHERE id = $1;

-- name: ListDocumentAcknowledgement :many
SELECT * FROM document_acknowledgements
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(document_revision_id)::uuid IS NULL OR document_revision_id = sqlc.narg(document_revision_id)::uuid)
  AND (sqlc.narg(employee_id)::uuid IS NULL OR employee_id = sqlc.narg(employee_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR due_on::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR due_on::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: SetDocumentAcknowledgementStatus :one
UPDATE document_acknowledgements
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActAcknowledgeDocumentAcknowledgement :one
UPDATE document_acknowledgements
SET signature_id = COALESCE(sqlc.narg(signature_id), signature_id), acknowledged_at = now(), status = 'acknowledged', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('pending')
RETURNING *;
