-- name: GetAttachment :one
SELECT * FROM attachments WHERE id = $1;

-- name: ListAttachment :many
SELECT * FROM attachments
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(subject_type)::text IS NULL OR subject_type = sqlc.narg(subject_type)::text)
  AND (sqlc.narg(subject_id)::uuid IS NULL OR subject_id = sqlc.narg(subject_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateAttachment :one
INSERT INTO attachments (filename, content_type, size_bytes, sha256, subject_type, subject_id, classification, kind, paper_original_location, captured_at, created_by, updated_by)
VALUES (sqlc.arg(filename), sqlc.arg(content_type), sqlc.arg(size_bytes), sqlc.arg(sha256), sqlc.narg(subject_type), sqlc.narg(subject_id), sqlc.narg(classification), sqlc.narg(kind), sqlc.narg(paper_original_location), sqlc.narg(captured_at), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetAttachmentStatus :one
UPDATE attachments
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActCompleteAttachment :one
UPDATE attachments
SET updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
