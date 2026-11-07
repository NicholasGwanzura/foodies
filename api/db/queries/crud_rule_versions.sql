-- name: GetRuleVersion :one
SELECT * FROM rule_versions WHERE id = $1;

-- name: ListRuleVersion :many
SELECT * FROM rule_versions
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateRuleVersion :one
INSERT INTO rule_versions (rule_definition_id, scope_type, scope_id, "value", effective_from, effective_to, source_document_id, change_reason, created_by, updated_by)
VALUES (sqlc.arg(rule_definition_id), sqlc.arg(scope_type), sqlc.narg(scope_id), sqlc.arg(value), sqlc.arg(effective_from), sqlc.narg(effective_to), sqlc.narg(source_document_id), sqlc.arg(change_reason), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetRuleVersionStatus :one
UPDATE rule_versions
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitRuleVersion :one
UPDATE rule_versions
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
