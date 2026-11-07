-- name: GetDocumentRevision :one
SELECT * FROM document_revisions WHERE id = $1;

-- name: ListDocumentRevision :many
SELECT * FROM document_revisions
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR revision_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR revision_date::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateDocumentRevision :one
INSERT INTO document_revisions (document_id, revision, revision_date, summary_of_change, file_attachment_id, prepared_by, temporary, expires_on, requires_acknowledgement, training_required, created_by, updated_by)
VALUES (sqlc.arg(document_id), sqlc.arg(revision), sqlc.narg(revision_date), sqlc.arg(summary_of_change), sqlc.arg(file_attachment_id), sqlc.narg(prepared_by), COALESCE(sqlc.narg(temporary)::boolean, false), sqlc.narg(expires_on), COALESCE(sqlc.narg(requires_acknowledgement)::boolean, false), COALESCE(sqlc.narg(training_required)::boolean, false), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateDocumentRevision :one
UPDATE document_revisions
SET revision = COALESCE(sqlc.narg(revision), revision), revision_date = COALESCE(sqlc.narg(revision_date), revision_date), summary_of_change = COALESCE(sqlc.narg(summary_of_change), summary_of_change), file_attachment_id = COALESCE(sqlc.narg(file_attachment_id), file_attachment_id), prepared_by = COALESCE(sqlc.narg(prepared_by), prepared_by), temporary = COALESCE(sqlc.narg(temporary), temporary), expires_on = COALESCE(sqlc.narg(expires_on), expires_on), requires_acknowledgement = COALESCE(sqlc.narg(requires_acknowledgement), requires_acknowledgement), training_required = COALESCE(sqlc.narg(training_required), training_required), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetDocumentRevisionStatus :one
UPDATE document_revisions
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitDocumentRevision :one
UPDATE document_revisions
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActIssueDocumentRevision :one
UPDATE document_revisions
SET status = 'issued', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'in_review', 'qms_review', 'pending_approval', 'approved')
RETURNING *;

-- name: ActWithdrawDocumentRevision :one
UPDATE document_revisions
SET status = 'withdrawn', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('issued', 'approved')
RETURNING *;
