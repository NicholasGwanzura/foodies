-- name: GetSignatureRequest :one
SELECT * FROM signature_requests WHERE id = $1;

-- name: ListSignatureRequest :many
SELECT * FROM signature_requests
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(subject_type)::text IS NULL OR subject_type = sqlc.narg(subject_type)::text)
  AND (sqlc.narg(subject_id)::uuid IS NULL OR subject_id = sqlc.narg(subject_id)::uuid)
  AND (sqlc.narg(signer_employee_id)::uuid IS NULL OR signer_employee_id = sqlc.narg(signer_employee_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateSignatureRequest :one
INSERT INTO signature_requests (subject_type, subject_id, meaning, signer_kind, signer_employee_id, signer_role_id, external_name, external_organisation, external_id_ref, proof_attachment_id, created_by, updated_by)
VALUES (sqlc.arg(subject_type), sqlc.arg(subject_id), sqlc.arg(meaning), sqlc.arg(signer_kind), sqlc.narg(signer_employee_id), sqlc.narg(signer_role_id), sqlc.narg(external_name), sqlc.narg(external_organisation), sqlc.narg(external_id_ref), sqlc.narg(proof_attachment_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetSignatureRequestStatus :one
UPDATE signature_requests
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSignSignatureRequest :one
UPDATE signature_requests
SET updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActDeclineSignatureRequest :one
UPDATE signature_requests
SET updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
