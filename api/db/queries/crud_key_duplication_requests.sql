-- name: GetKeyDuplicationRequest :one
SELECT * FROM key_duplication_requests WHERE id = $1;

-- name: ListKeyDuplicationRequest :many
SELECT * FROM key_duplication_requests
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR cut_on::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR cut_on::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateKeyDuplicationRequest :one
INSERT INTO key_duplication_requests (reference, site_id, key_type, "reason", quantity, risk_assessment, memo_attachment_id, locksmith, cut_on, issued_key_ids, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(key_type), sqlc.arg(reason), sqlc.arg(quantity), sqlc.narg(risk_assessment), sqlc.narg(memo_attachment_id), sqlc.narg(locksmith), sqlc.narg(cut_on), sqlc.narg(issued_key_ids), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateKeyDuplicationRequest :one
UPDATE key_duplication_requests
SET site_id = COALESCE(sqlc.narg(site_id), site_id), key_type = COALESCE(sqlc.narg(key_type), key_type), "reason" = COALESCE(sqlc.narg(reason), "reason"), quantity = COALESCE(sqlc.narg(quantity), quantity), risk_assessment = COALESCE(sqlc.narg(risk_assessment), risk_assessment), memo_attachment_id = COALESCE(sqlc.narg(memo_attachment_id), memo_attachment_id), locksmith = COALESCE(sqlc.narg(locksmith), locksmith), cut_on = COALESCE(sqlc.narg(cut_on), cut_on), issued_key_ids = COALESCE(sqlc.narg(issued_key_ids), issued_key_ids), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetKeyDuplicationRequestStatus :one
UPDATE key_duplication_requests
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitKeyDuplicationRequest :one
UPDATE key_duplication_requests
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
