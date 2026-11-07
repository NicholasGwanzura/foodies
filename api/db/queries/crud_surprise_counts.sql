-- name: GetSurpriseCount :one
SELECT * FROM surprise_counts WHERE id = $1;

-- name: ListSurpriseCount :many
SELECT * FROM surprise_counts
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateSurpriseCount :one
INSERT INTO surprise_counts (reference, site_id, counted_at, count_team, cashier_employee_id, pos_balance_amount, pos_balance_currency, physical_cash_amount, physical_cash_currency, findings, corrective_action_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(counted_at), sqlc.arg(count_team), sqlc.narg(cashier_employee_id), sqlc.narg(pos_balance_amount), sqlc.narg(pos_balance_currency), sqlc.narg(physical_cash_amount), sqlc.narg(physical_cash_currency), sqlc.narg(findings), sqlc.narg(corrective_action_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateSurpriseCount :one
UPDATE surprise_counts
SET site_id = COALESCE(sqlc.narg(site_id), site_id), counted_at = COALESCE(sqlc.narg(counted_at), counted_at), count_team = COALESCE(sqlc.narg(count_team), count_team), cashier_employee_id = COALESCE(sqlc.narg(cashier_employee_id), cashier_employee_id), pos_balance_amount = COALESCE(sqlc.narg(pos_balance_amount), pos_balance_amount), pos_balance_currency = COALESCE(sqlc.narg(pos_balance_currency), pos_balance_currency), physical_cash_amount = COALESCE(sqlc.narg(physical_cash_amount), physical_cash_amount), physical_cash_currency = COALESCE(sqlc.narg(physical_cash_currency), physical_cash_currency), findings = COALESCE(sqlc.narg(findings), findings), corrective_action_id = COALESCE(sqlc.narg(corrective_action_id), corrective_action_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetSurpriseCountStatus :one
UPDATE surprise_counts
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActCompleteSurpriseCount :one
UPDATE surprise_counts
SET status = 'completed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('planned', 'in_progress')
RETURNING *;
