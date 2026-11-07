-- name: GetWasteRecord :one
SELECT * FROM waste_records WHERE id = $1;

-- name: ListWasteRecord :many
SELECT * FROM waste_records
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(category)::text IS NULL OR category = sqlc.narg(category)::text)
  AND (sqlc.narg(batch_id)::uuid IS NULL OR batch_id = sqlc.narg(batch_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateWasteRecord :one
INSERT INTO waste_records (reference, site_id, occurred_at, item_id, product_id, batch_id, quantity, uom, category, cause, value_amount, value_currency, photo_attachment_id, disposition, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(occurred_at), sqlc.narg(item_id), sqlc.narg(product_id), sqlc.narg(batch_id), sqlc.arg(quantity), sqlc.arg(uom), sqlc.arg(category), sqlc.narg(cause), sqlc.narg(value_amount), sqlc.narg(value_currency), sqlc.narg(photo_attachment_id), sqlc.narg(disposition), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateWasteRecord :one
UPDATE waste_records
SET site_id = COALESCE(sqlc.narg(site_id), site_id), occurred_at = COALESCE(sqlc.narg(occurred_at), occurred_at), item_id = COALESCE(sqlc.narg(item_id), item_id), product_id = COALESCE(sqlc.narg(product_id), product_id), batch_id = COALESCE(sqlc.narg(batch_id), batch_id), quantity = COALESCE(sqlc.narg(quantity), quantity), uom = COALESCE(sqlc.narg(uom), uom), category = COALESCE(sqlc.narg(category), category), cause = COALESCE(sqlc.narg(cause), cause), value_amount = COALESCE(sqlc.narg(value_amount), value_amount), value_currency = COALESCE(sqlc.narg(value_currency), value_currency), photo_attachment_id = COALESCE(sqlc.narg(photo_attachment_id), photo_attachment_id), disposition = COALESCE(sqlc.narg(disposition), disposition), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetWasteRecordStatus :one
UPDATE waste_records
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActReviewWasteRecord :one
UPDATE waste_records
SET status = 'reviewed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('reviewed')
RETURNING *;
