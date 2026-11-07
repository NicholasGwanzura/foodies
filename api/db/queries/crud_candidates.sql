-- name: GetCandidate :one
SELECT * FROM candidates WHERE id = $1;

-- name: ListCandidate :many
SELECT * FROM candidates
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateCandidate :one
INSERT INTO candidates (recruitment_requisition_id, full_name, contact, cv_attachment_id, assessment, offer_status, created_by, updated_by)
VALUES (sqlc.arg(recruitment_requisition_id), sqlc.arg(full_name), sqlc.narg(contact), sqlc.narg(cv_attachment_id), sqlc.narg(assessment), sqlc.narg(offer_status), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateCandidate :one
UPDATE candidates
SET recruitment_requisition_id = COALESCE(sqlc.narg(recruitment_requisition_id), recruitment_requisition_id), full_name = COALESCE(sqlc.narg(full_name), full_name), contact = COALESCE(sqlc.narg(contact), contact), cv_attachment_id = COALESCE(sqlc.narg(cv_attachment_id), cv_attachment_id), assessment = COALESCE(sqlc.narg(assessment), assessment), offer_status = COALESCE(sqlc.narg(offer_status), offer_status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetCandidateStatus :one
UPDATE candidates
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
