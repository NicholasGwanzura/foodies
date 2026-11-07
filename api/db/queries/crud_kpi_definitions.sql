-- name: GetKpiDefinition :one
SELECT * FROM kpi_definitions WHERE id = $1;

-- name: ListKpiDefinition :many
SELECT * FROM kpi_definitions
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR "name" ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(module)::text IS NULL OR module = sqlc.narg(module)::text)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateKpiDefinition :one
INSERT INTO kpi_definitions ("key", "name", module, formula, target_rule_key, unit, drilldown, created_by, updated_by)
VALUES (sqlc.arg(key), sqlc.arg(name), sqlc.arg(module), sqlc.narg(formula), sqlc.narg(target_rule_key), sqlc.arg(unit), sqlc.narg(drilldown), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateKpiDefinition :one
UPDATE kpi_definitions
SET "key" = COALESCE(sqlc.narg(key), "key"), "name" = COALESCE(sqlc.narg(name), "name"), module = COALESCE(sqlc.narg(module), module), formula = COALESCE(sqlc.narg(formula), formula), target_rule_key = COALESCE(sqlc.narg(target_rule_key), target_rule_key), unit = COALESCE(sqlc.narg(unit), unit), drilldown = COALESCE(sqlc.narg(drilldown), drilldown), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
