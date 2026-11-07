-- name: GetTillFlush :one
SELECT * FROM till_flushes WHERE id = $1;

-- name: ListTillFlush :many
SELECT * FROM till_flushes
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(till_id)::uuid IS NULL OR till_id = sqlc.narg(till_id)::uuid)
  AND (sqlc.narg(business_date)::date IS NULL OR business_date = sqlc.narg(business_date)::date)
  AND (sqlc.narg(from_date)::date IS NULL OR business_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR business_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateTillFlush :one
INSERT INTO till_flushes (reference, site_id, till_id, cashier_employee_id, supervisor_employee_id, security_witness_employee_id, business_date, shift, threshold_reached_at, flushed_at, amount_amount, amount_currency, denominations, safe_id, safe_reference, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(till_id), sqlc.arg(cashier_employee_id), sqlc.arg(supervisor_employee_id), sqlc.narg(security_witness_employee_id), sqlc.arg(business_date), sqlc.arg(shift), sqlc.narg(threshold_reached_at), sqlc.arg(flushed_at), sqlc.arg(amount_amount), sqlc.arg(amount_currency), sqlc.narg(denominations), sqlc.arg(safe_id), sqlc.narg(safe_reference), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateTillFlush :one
UPDATE till_flushes
SET site_id = COALESCE(sqlc.narg(site_id), site_id), till_id = COALESCE(sqlc.narg(till_id), till_id), cashier_employee_id = COALESCE(sqlc.narg(cashier_employee_id), cashier_employee_id), supervisor_employee_id = COALESCE(sqlc.narg(supervisor_employee_id), supervisor_employee_id), security_witness_employee_id = COALESCE(sqlc.narg(security_witness_employee_id), security_witness_employee_id), business_date = COALESCE(sqlc.narg(business_date), business_date), shift = COALESCE(sqlc.narg(shift), shift), threshold_reached_at = COALESCE(sqlc.narg(threshold_reached_at), threshold_reached_at), flushed_at = COALESCE(sqlc.narg(flushed_at), flushed_at), amount_amount = COALESCE(sqlc.narg(amount_amount), amount_amount), amount_currency = COALESCE(sqlc.narg(amount_currency), amount_currency), denominations = COALESCE(sqlc.narg(denominations), denominations), safe_id = COALESCE(sqlc.narg(safe_id), safe_id), safe_reference = COALESCE(sqlc.narg(safe_reference), safe_reference), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetTillFlushStatus :one
UPDATE till_flushes
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitTillFlush :one
UPDATE till_flushes
SET status = 'pending_signatures', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
