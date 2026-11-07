-- name: GetZReport :one
SELECT * FROM z_reports WHERE id = $1;

-- name: ListZReport :many
SELECT * FROM z_reports
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(till_id)::uuid IS NULL OR till_id = sqlc.narg(till_id)::uuid)
  AND (sqlc.narg(business_date)::date IS NULL OR business_date = sqlc.narg(business_date)::date)
  AND (sqlc.narg(report_type)::text IS NULL OR report_type = sqlc.narg(report_type)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR business_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR business_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateZReport :one
INSERT INTO z_reports (site_id, till_id, business_date, report_type, report_time, total_amount, total_currency, fiscal_status, attachment_id, created_by, updated_by)
VALUES (sqlc.arg(site_id), sqlc.arg(till_id), sqlc.arg(business_date), sqlc.arg(report_type), sqlc.arg(report_time), sqlc.arg(total_amount), sqlc.arg(total_currency), sqlc.narg(fiscal_status), sqlc.narg(attachment_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;
