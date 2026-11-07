-- name: GetStockAdjustment :one
SELECT * FROM stock_adjustments WHERE id = $1;

-- name: ListStockAdjustment :many
SELECT * FROM stock_adjustments
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(item_id)::uuid IS NULL OR item_id = sqlc.narg(item_id)::uuid)
  AND (sqlc.narg(reason_tag)::text IS NULL OR reason_tag = sqlc.narg(reason_tag)::text)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateStockAdjustment :one
INSERT INTO stock_adjustments (reference, site_id, item_id, quantity, value_amount, value_currency, reason_tag, root_cause, stock_variance_id, supporting_attachment_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(item_id), sqlc.arg(quantity), sqlc.narg(value_amount), sqlc.narg(value_currency), sqlc.arg(reason_tag), sqlc.narg(root_cause), sqlc.narg(stock_variance_id), sqlc.narg(supporting_attachment_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateStockAdjustment :one
UPDATE stock_adjustments
SET site_id = COALESCE(sqlc.narg(site_id), site_id), item_id = COALESCE(sqlc.narg(item_id), item_id), quantity = COALESCE(sqlc.narg(quantity), quantity), value_amount = COALESCE(sqlc.narg(value_amount), value_amount), value_currency = COALESCE(sqlc.narg(value_currency), value_currency), reason_tag = COALESCE(sqlc.narg(reason_tag), reason_tag), root_cause = COALESCE(sqlc.narg(root_cause), root_cause), stock_variance_id = COALESCE(sqlc.narg(stock_variance_id), stock_variance_id), supporting_attachment_id = COALESCE(sqlc.narg(supporting_attachment_id), supporting_attachment_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetStockAdjustmentStatus :one
UPDATE stock_adjustments
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitStockAdjustment :one
UPDATE stock_adjustments
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
