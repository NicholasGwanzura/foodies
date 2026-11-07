-- name: GetPettyCashReconciliation :one
SELECT * FROM petty_cash_reconciliations WHERE id = $1;

-- name: ListPettyCashReconciliation :many
SELECT * FROM petty_cash_reconciliations
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(account_id)::uuid IS NULL OR account_id = sqlc.narg(account_id)::uuid)
  AND (sqlc.narg(business_date)::date IS NULL OR business_date = sqlc.narg(business_date)::date)
  AND (sqlc.narg(kind)::text IS NULL OR kind = sqlc.narg(kind)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR business_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR business_date::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreatePettyCashReconciliation :one
INSERT INTO petty_cash_reconciliations (reference, account_id, kind, business_date, opening_balance_amount, opening_balance_currency, receipts_amount, receipts_currency, disbursements_amount, disbursements_currency, physical_closing_amount, physical_closing_currency, explanation, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(account_id), sqlc.arg(kind), sqlc.arg(business_date), sqlc.arg(opening_balance_amount), sqlc.arg(opening_balance_currency), sqlc.arg(receipts_amount), sqlc.arg(receipts_currency), sqlc.arg(disbursements_amount), sqlc.arg(disbursements_currency), sqlc.arg(physical_closing_amount), sqlc.arg(physical_closing_currency), sqlc.narg(explanation), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetPettyCashReconciliationStatus :one
UPDATE petty_cash_reconciliations
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitPettyCashReconciliation :one
UPDATE petty_cash_reconciliations
SET status = 'submitted', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActReviewPettyCashReconciliation :one
UPDATE petty_cash_reconciliations
SET status = 'reviewed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('submitted', 'reviewed')
RETURNING *;
