-- name: GetPettyCashVoucher :one
SELECT * FROM petty_cash_vouchers WHERE id = $1;

-- name: ListPettyCashVoucher :many
SELECT * FROM petty_cash_vouchers
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(account_id)::uuid IS NULL OR account_id = sqlc.narg(account_id)::uuid)
  AND (sqlc.narg(expense_category)::text IS NULL OR expense_category = sqlc.narg(expense_category)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreatePettyCashVoucher :one
INSERT INTO petty_cash_vouchers (reference, account_id, site_id, supervisor_employee_id, "reason", amount_amount, amount_currency, expense_category, recipient, receipt_attachment_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(account_id), sqlc.narg(site_id), sqlc.arg(supervisor_employee_id), sqlc.arg(reason), sqlc.arg(amount_amount), sqlc.arg(amount_currency), sqlc.arg(expense_category), sqlc.arg(recipient), sqlc.narg(receipt_attachment_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdatePettyCashVoucher :one
UPDATE petty_cash_vouchers
SET account_id = COALESCE(sqlc.narg(account_id), account_id), supervisor_employee_id = COALESCE(sqlc.narg(supervisor_employee_id), supervisor_employee_id), "reason" = COALESCE(sqlc.narg(reason), "reason"), amount_amount = COALESCE(sqlc.narg(amount_amount), amount_amount), amount_currency = COALESCE(sqlc.narg(amount_currency), amount_currency), expense_category = COALESCE(sqlc.narg(expense_category), expense_category), recipient = COALESCE(sqlc.narg(recipient), recipient), receipt_attachment_id = COALESCE(sqlc.narg(receipt_attachment_id), receipt_attachment_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetPettyCashVoucherStatus :one
UPDATE petty_cash_vouchers
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitPettyCashVoucher :one
UPDATE petty_cash_vouchers
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActVerifyPettyCashVoucher :one
UPDATE petty_cash_vouchers
SET verified_by = sqlc.narg(updated_by), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('pending_approval')
RETURNING *;

-- name: ActIssuePettyCashVoucher :one
UPDATE petty_cash_vouchers
SET issued_amount_amount = COALESCE(sqlc.narg(issued_amount_amount), issued_amount_amount), issued_amount_currency = COALESCE(sqlc.narg(issued_amount_currency), issued_amount_currency), issued_at = now(), status = 'issued', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'pending_approval', 'approved')
RETURNING *;

-- name: ActAttachReceiptPettyCashVoucher :one
UPDATE petty_cash_vouchers
SET receipt_attachment_id = COALESCE(sqlc.narg(receipt_attachment_id), receipt_attachment_id), status = 'receipted', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('issued')
RETURNING *;
