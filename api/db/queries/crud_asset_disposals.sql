-- name: GetAssetDisposal :one
SELECT * FROM asset_disposals WHERE id = $1;

-- name: ListAssetDisposal :many
SELECT * FROM asset_disposals
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateAssetDisposal :one
INSERT INTO asset_disposals (reference, asset_id, "method", "reason", proceeds_amount, proceeds_currency, invoice_reference, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(asset_id), sqlc.arg(method), sqlc.arg(reason), sqlc.narg(proceeds_amount), sqlc.narg(proceeds_currency), sqlc.narg(invoice_reference), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateAssetDisposal :one
UPDATE asset_disposals
SET asset_id = COALESCE(sqlc.narg(asset_id), asset_id), "method" = COALESCE(sqlc.narg(method), "method"), "reason" = COALESCE(sqlc.narg(reason), "reason"), proceeds_amount = COALESCE(sqlc.narg(proceeds_amount), proceeds_amount), proceeds_currency = COALESCE(sqlc.narg(proceeds_currency), proceeds_currency), invoice_reference = COALESCE(sqlc.narg(invoice_reference), invoice_reference), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetAssetDisposalStatus :one
UPDATE asset_disposals
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitAssetDisposal :one
UPDATE asset_disposals
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActConfirmRemovalAssetDisposal :one
UPDATE asset_disposals
SET status = 'removed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('approved')
RETURNING *;
