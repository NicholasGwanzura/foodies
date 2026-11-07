-- name: GetFloatIssue :one
SELECT * FROM float_issues WHERE id = $1;

-- name: ListFloatIssue :many
SELECT * FROM float_issues
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(till_id)::uuid IS NULL OR till_id = sqlc.narg(till_id)::uuid)
  AND (sqlc.narg(cashier_employee_id)::uuid IS NULL OR cashier_employee_id = sqlc.narg(cashier_employee_id)::uuid)
  AND (sqlc.narg(business_date)::date IS NULL OR business_date = sqlc.narg(business_date)::date)
  AND (sqlc.narg(from_date)::date IS NULL OR business_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR business_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateFloatIssue :one
INSERT INTO float_issues (reference, site_id, till_id, cashier_employee_id, business_date, shift, amounts, issued_at, returned_amounts, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(till_id), sqlc.arg(cashier_employee_id), sqlc.arg(business_date), sqlc.arg(shift), sqlc.arg(amounts), sqlc.arg(issued_at), sqlc.narg(returned_amounts), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetFloatIssueStatus :one
UPDATE float_issues
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActAcknowledgeFloatIssue :one
UPDATE float_issues
SET acknowledged_at = now(), status = 'acknowledged', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('issued')
RETURNING *;

-- name: ActReturnFloatIssue :one
UPDATE float_issues
SET returned_amounts = COALESCE(sqlc.narg(returned_amounts), returned_amounts), returned_at = now(), status = 'returned', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('issued', 'acknowledged', 'exception')
RETURNING *;
