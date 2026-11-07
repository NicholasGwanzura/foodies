-- name: GetDisposition :one
SELECT * FROM dispositions WHERE id = $1;

-- name: ListDisposition :many
SELECT * FROM dispositions
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateDisposition :one
INSERT INTO dispositions (reference, hold_id, disposition_type, quantity, "method", safety_assessment, rework_batch_id, disposal_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(hold_id), sqlc.arg(disposition_type), sqlc.arg(quantity), sqlc.arg(method), sqlc.arg(safety_assessment), sqlc.narg(rework_batch_id), sqlc.narg(disposal_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetDispositionStatus :one
UPDATE dispositions
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitDisposition :one
UPDATE dispositions
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActExecuteDisposition :one
UPDATE dispositions
SET status = 'executed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'pending_approval', 'approved')
RETURNING *;
