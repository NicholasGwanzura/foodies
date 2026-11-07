-- name: GetCashVariance :one
SELECT * FROM cash_variances WHERE id = $1;

-- name: ListCashVariance :many
SELECT * FROM cash_variances
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(cashier_employee_id)::uuid IS NULL OR cashier_employee_id = sqlc.narg(cashier_employee_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: UpdateCashVariance :one
UPDATE cash_variances
SET cash_up_id = COALESCE(sqlc.narg(cash_up_id), cash_up_id), explanation = COALESCE(sqlc.narg(explanation), explanation), recovery_status = COALESCE(sqlc.narg(recovery_status), recovery_status), recovered_at = COALESCE(sqlc.narg(recovered_at), recovered_at), recovery_reference = COALESCE(sqlc.narg(recovery_reference), recovery_reference), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetCashVarianceStatus :one
UPDATE cash_variances
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActRecordRecoveryCashVariance :one
UPDATE cash_variances
SET recovery_reference = COALESCE(sqlc.narg(recovery_reference), recovery_reference), recovered_at = COALESCE(sqlc.narg(recovered_at), recovered_at), status = 'recovered', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('open', 'under_investigation')
RETURNING *;

-- name: ActReferHrCashVariance :one
UPDATE cash_variances
SET status = 'referred', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('open', 'under_investigation')
RETURNING *;

-- name: ActCloseCashVariance :one
UPDATE cash_variances
SET status = 'closed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('open', 'recovered', 'under_investigation', 'referred')
RETURNING *;
