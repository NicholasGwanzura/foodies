-- name: GetDispositionBatch :one
SELECT * FROM disposition_batches WHERE id = $1;

-- name: ListDispositionBatch :many
SELECT * FROM disposition_batches
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR cutoff_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR cutoff_date::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateDispositionBatch :one
INSERT INTO disposition_batches (retention_class_id, cutoff_date, created_by, updated_by)
VALUES (sqlc.arg(retention_class_id), sqlc.arg(cutoff_date), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetDispositionBatchStatus :one
UPDATE disposition_batches
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitDispositionBatch :one
UPDATE disposition_batches
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('proposed')
RETURNING *;

-- name: ActCompleteDispositionBatch :one
UPDATE disposition_batches
SET completed_at = now(), status = 'completed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('proposed', 'pending_approval', 'approved')
RETURNING *;
