-- name: GetMaintenancePlan :one
SELECT * FROM maintenance_plans WHERE id = $1;

-- name: ListMaintenancePlan :many
SELECT * FROM maintenance_plans
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(asset_id)::uuid IS NULL OR asset_id = sqlc.narg(asset_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR next_due_on::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR next_due_on::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateMaintenancePlan :one
INSERT INTO maintenance_plans (asset_id, task, frequency_rule_key, assigned_role_id, next_due_on, form_template_id, created_by, updated_by)
VALUES (sqlc.arg(asset_id), sqlc.arg(task), sqlc.arg(frequency_rule_key), sqlc.narg(assigned_role_id), sqlc.narg(next_due_on), sqlc.narg(form_template_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateMaintenancePlan :one
UPDATE maintenance_plans
SET asset_id = COALESCE(sqlc.narg(asset_id), asset_id), task = COALESCE(sqlc.narg(task), task), frequency_rule_key = COALESCE(sqlc.narg(frequency_rule_key), frequency_rule_key), assigned_role_id = COALESCE(sqlc.narg(assigned_role_id), assigned_role_id), next_due_on = COALESCE(sqlc.narg(next_due_on), next_due_on), form_template_id = COALESCE(sqlc.narg(form_template_id), form_template_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetMaintenancePlanStatus :one
UPDATE maintenance_plans
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
