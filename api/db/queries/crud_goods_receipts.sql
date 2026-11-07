-- name: GetGoodsReceipt :one
SELECT * FROM goods_receipts WHERE id = $1;

-- name: ListGoodsReceipt :many
SELECT * FROM goods_receipts
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(purchase_order_id)::uuid IS NULL OR purchase_order_id = sqlc.narg(purchase_order_id)::uuid)
  AND (sqlc.narg(supplier_id)::uuid IS NULL OR supplier_id = sqlc.narg(supplier_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateGoodsReceipt :one
INSERT INTO goods_receipts (reference, purchase_order_id, supplier_id, site_id, delivery_note_number, supplier_invoice_number, received_at, receiver_employee_id, delivery_note_attachment_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.narg(purchase_order_id), sqlc.arg(supplier_id), sqlc.arg(site_id), sqlc.narg(delivery_note_number), sqlc.narg(supplier_invoice_number), sqlc.arg(received_at), sqlc.arg(receiver_employee_id), sqlc.narg(delivery_note_attachment_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateGoodsReceipt :one
UPDATE goods_receipts
SET purchase_order_id = COALESCE(sqlc.narg(purchase_order_id), purchase_order_id), supplier_id = COALESCE(sqlc.narg(supplier_id), supplier_id), site_id = COALESCE(sqlc.narg(site_id), site_id), delivery_note_number = COALESCE(sqlc.narg(delivery_note_number), delivery_note_number), supplier_invoice_number = COALESCE(sqlc.narg(supplier_invoice_number), supplier_invoice_number), received_at = COALESCE(sqlc.narg(received_at), received_at), receiver_employee_id = COALESCE(sqlc.narg(receiver_employee_id), receiver_employee_id), delivery_note_attachment_id = COALESCE(sqlc.narg(delivery_note_attachment_id), delivery_note_attachment_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetGoodsReceiptStatus :one
UPDATE goods_receipts
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListGoodsReceiptLines :many
SELECT * FROM goods_receipt_lines WHERE goods_receipt_id = $1 ORDER BY id;

-- name: InsertGoodsReceiptLines :one
INSERT INTO goods_receipt_lines (goods_receipt_id, item_id, ordered_quantity, received_quantity, uom, batch_lot, expiry_date, temperature_c, thermometer_id, pack_condition, outcome, rejection_reason, hold_id, created_by)
VALUES (sqlc.arg(goods_receipt_id), sqlc.arg(item_id), sqlc.narg(ordered_quantity), sqlc.arg(received_quantity), sqlc.arg(uom), sqlc.narg(batch_lot), sqlc.narg(expiry_date), sqlc.narg(temperature_c), sqlc.narg(thermometer_id), sqlc.narg(pack_condition), sqlc.arg(outcome), sqlc.narg(rejection_reason), sqlc.narg(hold_id), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteGoodsReceiptLines :exec
DELETE FROM goods_receipt_lines WHERE goods_receipt_id = $1;

-- name: ActSubmitGoodsReceipt :one
UPDATE goods_receipts
SET status = 'inspected', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
