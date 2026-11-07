-- name: GetAssetTransfer :one
SELECT * FROM asset_transfers WHERE id = $1;

-- name: ListAssetTransfer :many
SELECT * FROM asset_transfers
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateAssetTransfer :one
INSERT INTO asset_transfers (reference, asset_id, to_site_id, to_custodian_employee_id, to_cost_centre_id, "reason", created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(asset_id), sqlc.arg(to_site_id), sqlc.narg(to_custodian_employee_id), sqlc.narg(to_cost_centre_id), sqlc.arg(reason), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetAssetTransferStatus :one
UPDATE asset_transfers
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitAssetTransfer :one
UPDATE asset_transfers
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActReceiveAssetTransfer :one
UPDATE asset_transfers
SET received_at = now(), status = 'received', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'pending_approval', 'approved')
RETURNING *;
