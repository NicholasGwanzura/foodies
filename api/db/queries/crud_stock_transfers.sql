-- name: GetStockTransfer :one
SELECT * FROM stock_transfers WHERE id = $1;

-- name: ListStockTransfer :many
SELECT * FROM stock_transfers
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_site_id)::uuid IS NULL OR from_site_id = sqlc.narg(from_site_id)::uuid)
  AND (sqlc.narg(to_site_id)::uuid IS NULL OR to_site_id = sqlc.narg(to_site_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateStockTransfer :one
INSERT INTO stock_transfers (reference, from_site_id, to_site_id, vehicle, dispatch_security_employee_id, receiving_security_employee_id, invoice_attachment_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(from_site_id), sqlc.arg(to_site_id), sqlc.narg(vehicle), sqlc.narg(dispatch_security_employee_id), sqlc.narg(receiving_security_employee_id), sqlc.narg(invoice_attachment_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateStockTransfer :one
UPDATE stock_transfers
SET from_site_id = COALESCE(sqlc.narg(from_site_id), from_site_id), to_site_id = COALESCE(sqlc.narg(to_site_id), to_site_id), vehicle = COALESCE(sqlc.narg(vehicle), vehicle), dispatch_security_employee_id = COALESCE(sqlc.narg(dispatch_security_employee_id), dispatch_security_employee_id), receiving_security_employee_id = COALESCE(sqlc.narg(receiving_security_employee_id), receiving_security_employee_id), invoice_attachment_id = COALESCE(sqlc.narg(invoice_attachment_id), invoice_attachment_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetStockTransferStatus :one
UPDATE stock_transfers
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListStockTransferLines :many
SELECT * FROM stock_transfer_lines WHERE stock_transfer_id = $1 ORDER BY id;

-- name: InsertStockTransferLines :one
INSERT INTO stock_transfer_lines (stock_transfer_id, item_id, requested_quantity, approved_quantity, dispatched_quantity, received_quantity, uom, batch_lot, condition, created_by)
VALUES (sqlc.arg(stock_transfer_id), sqlc.arg(item_id), sqlc.arg(requested_quantity), sqlc.narg(approved_quantity), sqlc.narg(dispatched_quantity), sqlc.narg(received_quantity), sqlc.arg(uom), sqlc.narg(batch_lot), sqlc.narg(condition), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteStockTransferLines :exec
DELETE FROM stock_transfer_lines WHERE stock_transfer_id = $1;

-- name: ActAcceptStockTransfer :one
UPDATE stock_transfers
SET status = 'accepted', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('requested')
RETURNING *;

-- name: ActRejectStockTransfer :one
UPDATE stock_transfers
SET status = 'rejected', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('requested', 'accepted')
RETURNING *;

-- name: ActDispatchStockTransfer :one
UPDATE stock_transfers
SET dispatched_at = now(), status = 'dispatched', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('requested', 'accepted')
RETURNING *;

-- name: ActReceiveStockTransfer :one
UPDATE stock_transfers
SET received_by = sqlc.narg(updated_by), status = 'received', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('requested', 'accepted', 'dispatched')
RETURNING *;

-- name: ActCloseStockTransfer :one
UPDATE stock_transfers
SET status = 'closed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('requested', 'accepted', 'dispatched', 'received', 'discrepancy')
RETURNING *;
