-- name: GetReportDefinition :one
SELECT * FROM report_definitions WHERE id = $1;

-- name: ListReportDefinition :many
SELECT * FROM report_definitions
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR "name" ILIKE '%' || sqlc.narg(q)::text || '%')
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateReportDefinition :one
INSERT INTO report_definitions ("key", "name", frequency, due_rule_key, department_id, submitter_role_id, reviewer_role_id, template_id, shop_pattern, created_by, updated_by)
VALUES (sqlc.arg(key), sqlc.arg(name), sqlc.arg(frequency), sqlc.arg(due_rule_key), sqlc.narg(department_id), sqlc.narg(submitter_role_id), sqlc.narg(reviewer_role_id), sqlc.narg(template_id), sqlc.narg(shop_pattern), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateReportDefinition :one
UPDATE report_definitions
SET "key" = COALESCE(sqlc.narg(key), "key"), "name" = COALESCE(sqlc.narg(name), "name"), frequency = COALESCE(sqlc.narg(frequency), frequency), due_rule_key = COALESCE(sqlc.narg(due_rule_key), due_rule_key), department_id = COALESCE(sqlc.narg(department_id), department_id), submitter_role_id = COALESCE(sqlc.narg(submitter_role_id), submitter_role_id), reviewer_role_id = COALESCE(sqlc.narg(reviewer_role_id), reviewer_role_id), template_id = COALESCE(sqlc.narg(template_id), template_id), shop_pattern = COALESCE(sqlc.narg(shop_pattern), shop_pattern), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
