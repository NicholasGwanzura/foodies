-- name: GetSupplierInvoice :one
SELECT * FROM supplier_invoices WHERE id = $1;

-- name: ListSupplierInvoice :many
SELECT * FROM supplier_invoices
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(supplier_id)::uuid IS NULL OR supplier_id = sqlc.narg(supplier_id)::uuid)
  AND (sqlc.narg(match_status)::text IS NULL OR match_status = sqlc.narg(match_status)::text)
  AND (sqlc.narg(purchase_order_id)::uuid IS NULL OR purchase_order_id = sqlc.narg(purchase_order_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR invoice_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR invoice_date::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateSupplierInvoice :one
INSERT INTO supplier_invoices (reference, supplier_id, invoice_number, invoice_date, total_amount, total_currency, vat_amount, vat_currency, due_date, purchase_order_id, goods_receipt_ids, invoice_received_at, forwarded_to_finance_at, qr_verified, attachment_id, paper_original_location, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(supplier_id), sqlc.arg(invoice_number), sqlc.arg(invoice_date), sqlc.arg(total_amount), sqlc.arg(total_currency), sqlc.narg(vat_amount), sqlc.narg(vat_currency), sqlc.narg(due_date), sqlc.narg(purchase_order_id), sqlc.narg(goods_receipt_ids), sqlc.arg(invoice_received_at), sqlc.narg(forwarded_to_finance_at), COALESCE(sqlc.narg(qr_verified)::boolean, false), sqlc.narg(attachment_id), sqlc.narg(paper_original_location), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateSupplierInvoice :one
UPDATE supplier_invoices
SET supplier_id = COALESCE(sqlc.narg(supplier_id), supplier_id), invoice_number = COALESCE(sqlc.narg(invoice_number), invoice_number), invoice_date = COALESCE(sqlc.narg(invoice_date), invoice_date), total_amount = COALESCE(sqlc.narg(total_amount), total_amount), total_currency = COALESCE(sqlc.narg(total_currency), total_currency), vat_amount = COALESCE(sqlc.narg(vat_amount), vat_amount), vat_currency = COALESCE(sqlc.narg(vat_currency), vat_currency), due_date = COALESCE(sqlc.narg(due_date), due_date), purchase_order_id = COALESCE(sqlc.narg(purchase_order_id), purchase_order_id), goods_receipt_ids = COALESCE(sqlc.narg(goods_receipt_ids), goods_receipt_ids), invoice_received_at = COALESCE(sqlc.narg(invoice_received_at), invoice_received_at), forwarded_to_finance_at = COALESCE(sqlc.narg(forwarded_to_finance_at), forwarded_to_finance_at), qr_verified = COALESCE(sqlc.narg(qr_verified), qr_verified), attachment_id = COALESCE(sqlc.narg(attachment_id), attachment_id), paper_original_location = COALESCE(sqlc.narg(paper_original_location), paper_original_location), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetSupplierInvoiceStatus :one
UPDATE supplier_invoices
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListSupplierInvoiceLines :many
SELECT * FROM supplier_invoice_lines WHERE supplier_invoice_id = $1 ORDER BY id;

-- name: InsertSupplierInvoiceLines :one
INSERT INTO supplier_invoice_lines (supplier_invoice_id, description, quantity, unit_price, line_total, po_line_id, match_result, created_by)
VALUES (sqlc.arg(supplier_invoice_id), sqlc.arg(description), sqlc.arg(quantity), sqlc.arg(unit_price), sqlc.narg(line_total), sqlc.narg(po_line_id), sqlc.narg(match_result), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteSupplierInvoiceLines :exec
DELETE FROM supplier_invoice_lines WHERE supplier_invoice_id = $1;

-- name: ActMatchSupplierInvoice :one
UPDATE supplier_invoices
SET status = 'matched', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('received')
RETURNING *;

-- name: ActHoldSupplierInvoice :one
UPDATE supplier_invoices
SET status = 'on_hold', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('received', 'matched', 'approved_exception', 'posted', 'paid', 'void')
RETURNING *;

-- name: ActApproveExceptionSupplierInvoice :one
UPDATE supplier_invoices
SET status = 'approved_exception', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('on_hold')
RETURNING *;

-- name: ActPostSupplierInvoice :one
UPDATE supplier_invoices
SET status = 'posted', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('received', 'matched', 'on_hold', 'approved_exception')
RETURNING *;
