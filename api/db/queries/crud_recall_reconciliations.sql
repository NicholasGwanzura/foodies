-- name: GetRecallReconciliation :one
SELECT * FROM recall_reconciliations WHERE id = $1;

-- name: ListRecallReconciliation :many
SELECT * FROM recall_reconciliations
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateRecallReconciliation :one
INSERT INTO recall_reconciliations (recall_id, produced_received, remaining_stock, distributed, sold, returned, disposed, explanation, independent_reviewer_id, created_by, updated_by)
VALUES (sqlc.arg(recall_id), sqlc.arg(produced_received), sqlc.arg(remaining_stock), sqlc.arg(distributed), sqlc.arg(sold), sqlc.arg(returned), sqlc.arg(disposed), sqlc.narg(explanation), sqlc.narg(independent_reviewer_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateRecallReconciliation :one
UPDATE recall_reconciliations
SET produced_received = COALESCE(sqlc.narg(produced_received), produced_received), remaining_stock = COALESCE(sqlc.narg(remaining_stock), remaining_stock), distributed = COALESCE(sqlc.narg(distributed), distributed), sold = COALESCE(sqlc.narg(sold), sold), returned = COALESCE(sqlc.narg(returned), returned), disposed = COALESCE(sqlc.narg(disposed), disposed), explanation = COALESCE(sqlc.narg(explanation), explanation), independent_reviewer_id = COALESCE(sqlc.narg(independent_reviewer_id), independent_reviewer_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetRecallReconciliationStatus :one
UPDATE recall_reconciliations
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActReviewRecallReconciliation :one
UPDATE recall_reconciliations
SET status = 'reviewed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
