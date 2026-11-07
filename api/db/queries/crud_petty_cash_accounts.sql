-- name: GetPettyCashAccount :one
SELECT * FROM petty_cash_accounts WHERE id = $1;

-- name: ListPettyCashAccount :many
SELECT * FROM petty_cash_accounts
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreatePettyCashAccount :one
INSERT INTO petty_cash_accounts (site_id, currency, custodian_employee_id, safe_id, created_by, updated_by)
VALUES (sqlc.arg(site_id), sqlc.arg(currency), sqlc.narg(custodian_employee_id), sqlc.narg(safe_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdatePettyCashAccount :one
UPDATE petty_cash_accounts
SET site_id = COALESCE(sqlc.narg(site_id), site_id), currency = COALESCE(sqlc.narg(currency), currency), custodian_employee_id = COALESCE(sqlc.narg(custodian_employee_id), custodian_employee_id), safe_id = COALESCE(sqlc.narg(safe_id), safe_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
