-- name: GetDeliveryDiscrepancy :one
SELECT * FROM delivery_discrepancies WHERE id = $1;

-- name: ListDeliveryDiscrepancy :many
SELECT * FROM delivery_discrepancies
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%' OR description ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateDeliveryDiscrepancy :one
INSERT INTO delivery_discrepancies (reference, goods_receipt_id, item_id, discrepancy_type, description, supplier_notified_at, resolution, resolution_reference, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(goods_receipt_id), sqlc.narg(item_id), sqlc.arg(discrepancy_type), sqlc.arg(description), sqlc.narg(supplier_notified_at), sqlc.narg(resolution), sqlc.narg(resolution_reference), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateDeliveryDiscrepancy :one
UPDATE delivery_discrepancies
SET goods_receipt_id = COALESCE(sqlc.narg(goods_receipt_id), goods_receipt_id), item_id = COALESCE(sqlc.narg(item_id), item_id), discrepancy_type = COALESCE(sqlc.narg(discrepancy_type), discrepancy_type), description = COALESCE(sqlc.narg(description), description), supplier_notified_at = COALESCE(sqlc.narg(supplier_notified_at), supplier_notified_at), resolution = COALESCE(sqlc.narg(resolution), resolution), resolution_reference = COALESCE(sqlc.narg(resolution_reference), resolution_reference), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetDeliveryDiscrepancyStatus :one
UPDATE delivery_discrepancies
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActCloseDeliveryDiscrepancy :one
UPDATE delivery_discrepancies
SET closed_at = now(), status = 'closed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('open', 'supplier_notified', 'resolved')
RETURNING *;
