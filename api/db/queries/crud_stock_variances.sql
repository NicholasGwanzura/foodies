-- name: GetStockVariance :one
SELECT * FROM stock_variances WHERE id = $1;

-- name: ListStockVariance :many
SELECT * FROM stock_variances
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(stocktake_id)::uuid IS NULL OR stocktake_id = sqlc.narg(stocktake_id)::uuid)
  AND (sqlc.narg(item_id)::uuid IS NULL OR item_id = sqlc.narg(item_id)::uuid)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: UpdateStockVariance :one
UPDATE stock_variances
SET stocktake_id = COALESCE(sqlc.narg(stocktake_id), stocktake_id), item_id = COALESCE(sqlc.narg(item_id), item_id), variance_quantity = COALESCE(sqlc.narg(variance_quantity), variance_quantity), variance_value_amount = COALESCE(sqlc.narg(variance_value_amount), variance_value_amount), variance_value_currency = COALESCE(sqlc.narg(variance_value_currency), variance_value_currency), documents_checked = COALESCE(sqlc.narg(documents_checked), documents_checked), root_cause = COALESCE(sqlc.narg(root_cause), root_cause), responsible_employee_id = COALESCE(sqlc.narg(responsible_employee_id), responsible_employee_id), recovery = COALESCE(sqlc.narg(recovery), recovery), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetStockVarianceStatus :one
UPDATE stock_variances
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActCloseStockVariance :one
UPDATE stock_variances
SET status = 'closed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('open', 'investigating', 'resolved', 'escalated')
RETURNING *;
