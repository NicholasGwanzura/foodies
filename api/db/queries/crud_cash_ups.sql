-- name: GetCashUp :one
SELECT * FROM cash_ups WHERE id = $1;

-- name: ListCashUp :many
SELECT * FROM cash_ups
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(till_id)::uuid IS NULL OR till_id = sqlc.narg(till_id)::uuid)
  AND (sqlc.narg(cashier_employee_id)::uuid IS NULL OR cashier_employee_id = sqlc.narg(cashier_employee_id)::uuid)
  AND (sqlc.narg(business_date)::date IS NULL OR business_date = sqlc.narg(business_date)::date)
  AND (sqlc.narg(kind)::text IS NULL OR kind = sqlc.narg(kind)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR business_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR business_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateCashUp :one
INSERT INTO cash_ups (reference, site_id, till_id, cashier_employee_id, supervisor_employee_id, witness_employee_id, business_date, shift, kind, x_read_id, z_report_id, fx_rate_id, safe_id, safe_deposit_seal, closed_at, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(till_id), sqlc.arg(cashier_employee_id), sqlc.arg(supervisor_employee_id), sqlc.narg(witness_employee_id), sqlc.arg(business_date), sqlc.arg(shift), sqlc.arg(kind), sqlc.narg(x_read_id), sqlc.narg(z_report_id), sqlc.narg(fx_rate_id), sqlc.narg(safe_id), sqlc.narg(safe_deposit_seal), sqlc.narg(closed_at), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateCashUp :one
UPDATE cash_ups
SET site_id = COALESCE(sqlc.narg(site_id), site_id), till_id = COALESCE(sqlc.narg(till_id), till_id), cashier_employee_id = COALESCE(sqlc.narg(cashier_employee_id), cashier_employee_id), supervisor_employee_id = COALESCE(sqlc.narg(supervisor_employee_id), supervisor_employee_id), witness_employee_id = COALESCE(sqlc.narg(witness_employee_id), witness_employee_id), business_date = COALESCE(sqlc.narg(business_date), business_date), shift = COALESCE(sqlc.narg(shift), shift), kind = COALESCE(sqlc.narg(kind), kind), x_read_id = COALESCE(sqlc.narg(x_read_id), x_read_id), z_report_id = COALESCE(sqlc.narg(z_report_id), z_report_id), fx_rate_id = COALESCE(sqlc.narg(fx_rate_id), fx_rate_id), safe_id = COALESCE(sqlc.narg(safe_id), safe_id), safe_deposit_seal = COALESCE(sqlc.narg(safe_deposit_seal), safe_deposit_seal), closed_at = COALESCE(sqlc.narg(closed_at), closed_at), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetCashUpStatus :one
UPDATE cash_ups
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListCashUpTenders :many
SELECT * FROM cash_up_tenders WHERE cash_up_id = $1 ORDER BY id;

-- name: InsertCashUpTenders :one
INSERT INTO cash_up_tenders (cash_up_id, currency, tender, expected, actual, variance, created_by)
VALUES (sqlc.arg(cash_up_id), sqlc.arg(currency), sqlc.arg(tender), sqlc.arg(expected), sqlc.arg(actual), sqlc.narg(variance), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteCashUpTenders :exec
DELETE FROM cash_up_tenders WHERE cash_up_id = $1;

-- name: ActSubmitCashUp :one
UPDATE cash_ups
SET status = 'pending_review', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'returned', 'counted')
RETURNING *;

-- name: ActApproveCashUp :one
UPDATE cash_ups
SET status = 'approved', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'counted', 'pending_signatures', 'pending_review')
RETURNING *;

-- name: ActReturnCashUp :one
UPDATE cash_ups
SET status = 'returned', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'counted', 'pending_signatures', 'pending_review', 'approved', 'exception')
RETURNING *;
