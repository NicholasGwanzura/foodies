-- name: GetPosSalesSummary :one
SELECT * FROM pos_sales_summaries WHERE id = $1;

-- name: ListPosSalesSummary :many
SELECT * FROM pos_sales_summaries
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(till_id)::uuid IS NULL OR till_id = sqlc.narg(till_id)::uuid)
  AND (sqlc.narg(business_date)::date IS NULL OR business_date = sqlc.narg(business_date)::date)
  AND (sqlc.narg(cashier_employee_id)::uuid IS NULL OR cashier_employee_id = sqlc.narg(cashier_employee_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR business_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR business_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);
