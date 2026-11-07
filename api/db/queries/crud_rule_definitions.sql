-- name: GetRuleDefinition :one
SELECT * FROM rule_definitions WHERE id = $1;

-- name: ListRuleDefinition :many
SELECT * FROM rule_definitions
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR "name" ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(module)::text IS NULL OR module = sqlc.narg(module)::text)
  AND (sqlc.narg(blocking)::boolean IS NULL OR blocking = sqlc.narg(blocking)::boolean)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: UpdateRuleDefinition :one
UPDATE rule_definitions
SET "key" = COALESCE(sqlc.narg(key), "key"), "name" = COALESCE(sqlc.narg(name), "name"), module = COALESCE(sqlc.narg(module), module), value_type = COALESCE(sqlc.narg(value_type), value_type), scope_level = COALESCE(sqlc.narg(scope_level), scope_level), owner_role_id = COALESCE(sqlc.narg(owner_role_id), owner_role_id), source_clause = COALESCE(sqlc.narg(source_clause), source_clause), decision_id = COALESCE(sqlc.narg(decision_id), decision_id), blocking = COALESCE(sqlc.narg(blocking), blocking), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
