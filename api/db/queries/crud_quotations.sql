-- name: GetQuotation :one
SELECT * FROM quotations WHERE id = $1;

-- name: ListQuotation :many
SELECT * FROM quotations
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(from_date)::date IS NULL OR quote_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR quote_date::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateQuotation :one
INSERT INTO quotations (purchase_requisition_id, supplier_id, quote_number, quote_date, valid_until, total_amount, total_currency, charges_amount, charges_currency, lead_time_days, payment_terms, specification_compliant, selected, selection_reason, attachment_id, created_by, updated_by)
VALUES (sqlc.arg(purchase_requisition_id), sqlc.arg(supplier_id), sqlc.arg(quote_number), sqlc.arg(quote_date), sqlc.narg(valid_until), sqlc.arg(total_amount), sqlc.arg(total_currency), sqlc.narg(charges_amount), sqlc.narg(charges_currency), sqlc.narg(lead_time_days), sqlc.narg(payment_terms), COALESCE(sqlc.narg(specification_compliant)::boolean, false), COALESCE(sqlc.narg(selected)::boolean, false), sqlc.narg(selection_reason), sqlc.narg(attachment_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateQuotation :one
UPDATE quotations
SET supplier_id = COALESCE(sqlc.narg(supplier_id), supplier_id), quote_number = COALESCE(sqlc.narg(quote_number), quote_number), quote_date = COALESCE(sqlc.narg(quote_date), quote_date), valid_until = COALESCE(sqlc.narg(valid_until), valid_until), total_amount = COALESCE(sqlc.narg(total_amount), total_amount), total_currency = COALESCE(sqlc.narg(total_currency), total_currency), charges_amount = COALESCE(sqlc.narg(charges_amount), charges_amount), charges_currency = COALESCE(sqlc.narg(charges_currency), charges_currency), lead_time_days = COALESCE(sqlc.narg(lead_time_days), lead_time_days), payment_terms = COALESCE(sqlc.narg(payment_terms), payment_terms), specification_compliant = COALESCE(sqlc.narg(specification_compliant), specification_compliant), selected = COALESCE(sqlc.narg(selected), selected), selection_reason = COALESCE(sqlc.narg(selection_reason), selection_reason), attachment_id = COALESCE(sqlc.narg(attachment_id), attachment_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: DeleteQuotation :exec
DELETE FROM quotations WHERE id = $1;
