-- name: GetPaymentRequisition :one
SELECT * FROM payment_requisitions WHERE id = $1;

-- name: ListPaymentRequisition :many
SELECT * FROM payment_requisitions
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(supplier_id)::uuid IS NULL OR supplier_id = sqlc.narg(supplier_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR due_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR due_date::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreatePaymentRequisition :one
INSERT INTO payment_requisitions (reference, supplier_id, invoice_ids, due_date, priority, payment_type, withholding_tax_amount, withholding_tax_currency, "reason", created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(supplier_id), sqlc.arg(invoice_ids), sqlc.narg(due_date), sqlc.narg(priority), sqlc.arg(payment_type), sqlc.narg(withholding_tax_amount), sqlc.narg(withholding_tax_currency), sqlc.narg(reason), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdatePaymentRequisition :one
UPDATE payment_requisitions
SET supplier_id = COALESCE(sqlc.narg(supplier_id), supplier_id), invoice_ids = COALESCE(sqlc.narg(invoice_ids), invoice_ids), due_date = COALESCE(sqlc.narg(due_date), due_date), priority = COALESCE(sqlc.narg(priority), priority), payment_type = COALESCE(sqlc.narg(payment_type), payment_type), withholding_tax_amount = COALESCE(sqlc.narg(withholding_tax_amount), withholding_tax_amount), withholding_tax_currency = COALESCE(sqlc.narg(withholding_tax_currency), withholding_tax_currency), "reason" = COALESCE(sqlc.narg(reason), "reason"), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetPaymentRequisitionStatus :one
UPDATE payment_requisitions
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitPaymentRequisition :one
UPDATE payment_requisitions
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
