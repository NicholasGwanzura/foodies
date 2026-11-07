-- name: GetCitReconciliation :one
SELECT * FROM cit_reconciliations WHERE id = $1;

-- name: ListCitReconciliation :many
SELECT * FROM cit_reconciliations
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(reconciliation_date)::date IS NULL OR reconciliation_date = sqlc.narg(reconciliation_date)::date)
  AND (sqlc.narg(from_date)::date IS NULL OR reconciliation_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR reconciliation_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateCitReconciliation :one
INSERT INTO cit_reconciliations (reference, reconciliation_date, site_id, dispatched_amount, dispatched_currency, received_amount, received_currency, deposited_amount, deposited_currency, investigation_owner_employee_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(reconciliation_date), sqlc.arg(site_id), sqlc.arg(dispatched_amount), sqlc.arg(dispatched_currency), sqlc.narg(received_amount), sqlc.narg(received_currency), sqlc.narg(deposited_amount), sqlc.narg(deposited_currency), sqlc.narg(investigation_owner_employee_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateCitReconciliation :one
UPDATE cit_reconciliations
SET reconciliation_date = COALESCE(sqlc.narg(reconciliation_date), reconciliation_date), site_id = COALESCE(sqlc.narg(site_id), site_id), dispatched_amount = COALESCE(sqlc.narg(dispatched_amount), dispatched_amount), dispatched_currency = COALESCE(sqlc.narg(dispatched_currency), dispatched_currency), received_amount = COALESCE(sqlc.narg(received_amount), received_amount), received_currency = COALESCE(sqlc.narg(received_currency), received_currency), deposited_amount = COALESCE(sqlc.narg(deposited_amount), deposited_amount), deposited_currency = COALESCE(sqlc.narg(deposited_currency), deposited_currency), investigation_owner_employee_id = COALESCE(sqlc.narg(investigation_owner_employee_id), investigation_owner_employee_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetCitReconciliationStatus :one
UPDATE cit_reconciliations
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActCompleteCitReconciliation :one
UPDATE cit_reconciliations
SET status = 'reconciled', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('open')
RETURNING *;
