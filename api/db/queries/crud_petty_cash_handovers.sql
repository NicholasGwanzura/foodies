-- name: GetPettyCashHandover :one
SELECT * FROM petty_cash_handovers WHERE id = $1;

-- name: ListPettyCashHandover :many
SELECT * FROM petty_cash_handovers
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(account_id)::uuid IS NULL OR account_id = sqlc.narg(account_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreatePettyCashHandover :one
INSERT INTO petty_cash_handovers (reference, account_id, outgoing_employee_id, incoming_employee_id, book_balance_amount, book_balance_currency, physical_balance_amount, physical_balance_currency, handed_over_at, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(account_id), sqlc.arg(outgoing_employee_id), sqlc.arg(incoming_employee_id), sqlc.arg(book_balance_amount), sqlc.arg(book_balance_currency), sqlc.arg(physical_balance_amount), sqlc.arg(physical_balance_currency), sqlc.arg(handed_over_at), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetPettyCashHandoverStatus :one
UPDATE petty_cash_handovers
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitPettyCashHandover :one
UPDATE petty_cash_handovers
SET status = 'pending_signatures', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
