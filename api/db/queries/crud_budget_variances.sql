-- name: GetBudgetVariance :one
SELECT * FROM budget_variances WHERE id = $1;

-- name: ListBudgetVariance :many
SELECT * FROM budget_variances
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR "month"::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR "month"::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: UpdateBudgetVariance :one
UPDATE budget_variances
SET "month" = COALESCE(sqlc.narg(month), "month"), account_code = COALESCE(sqlc.narg(account_code), account_code), budget_amount = COALESCE(sqlc.narg(budget_amount), budget_amount), actual_amount = COALESCE(sqlc.narg(actual_amount), actual_amount), explanation = COALESCE(sqlc.narg(explanation), explanation), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetBudgetVarianceStatus :one
UPDATE budget_variances
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
