-- name: GetManualReceipt :one
SELECT * FROM manual_receipts WHERE id = $1;

-- name: ListManualReceipt :many
SELECT * FROM manual_receipts
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateManualReceipt :one
INSERT INTO manual_receipts (pos_outage_id, receipt_book_id, receipt_number, issued_at, cashier_employee_id, items, amount_amount, amount_currency, tender, order_docket, created_by, updated_by)
VALUES (sqlc.arg(pos_outage_id), sqlc.arg(receipt_book_id), sqlc.arg(receipt_number), sqlc.arg(issued_at), sqlc.arg(cashier_employee_id), sqlc.arg(items), sqlc.arg(amount_amount), sqlc.arg(amount_currency), sqlc.arg(tender), sqlc.narg(order_docket), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetManualReceiptStatus :one
UPDATE manual_receipts
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActCancelManualReceipt :one
UPDATE manual_receipts
SET status = 'cancelled', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('issued', 'missing', 'back_captured')
RETURNING *;

-- name: ActBackCaptureManualReceipt :one
UPDATE manual_receipts
SET pos_transaction_ref = COALESCE(sqlc.narg(pos_transaction_ref), pos_transaction_ref), status = 'back_captured', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('issued')
RETURNING *;
