-- name: GetPettyCashReimbursement :one
SELECT * FROM petty_cash_reimbursements WHERE id = $1;

-- name: ListPettyCashReimbursement :many
SELECT * FROM petty_cash_reimbursements
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(account_id)::uuid IS NULL OR account_id = sqlc.narg(account_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreatePettyCashReimbursement :one
INSERT INTO petty_cash_reimbursements (reference, account_id, voucher_ids, payment_reference, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(account_id), sqlc.arg(voucher_ids), sqlc.narg(payment_reference), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetPettyCashReimbursementStatus :one
UPDATE petty_cash_reimbursements
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitPettyCashReimbursement :one
UPDATE petty_cash_reimbursements
SET submitted_at = now(), status = 'submitted', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActRecordPaymentPettyCashReimbursement :one
UPDATE petty_cash_reimbursements
SET payment_reference = COALESCE(sqlc.narg(payment_reference), payment_reference), status = 'paid', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('approved')
RETURNING *;
