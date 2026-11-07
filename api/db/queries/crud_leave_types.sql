-- name: GetLeaveType :one
SELECT * FROM leave_types WHERE id = $1;

-- name: ListLeaveType :many
SELECT * FROM leave_types
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR code ILIKE '%' || sqlc.narg(q)::text || '%' OR "name" ILIKE '%' || sqlc.narg(q)::text || '%')
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateLeaveType :one
INSERT INTO leave_types (code, "name", accrual_rule_key, requires_document, restricted, active, created_by, updated_by)
VALUES (sqlc.arg(code), sqlc.arg(name), sqlc.narg(accrual_rule_key), COALESCE(sqlc.narg(requires_document)::boolean, false), COALESCE(sqlc.narg(restricted)::boolean, false), COALESCE(sqlc.narg(active)::boolean, true), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateLeaveType :one
UPDATE leave_types
SET code = COALESCE(sqlc.narg(code), code), "name" = COALESCE(sqlc.narg(name), "name"), accrual_rule_key = COALESCE(sqlc.narg(accrual_rule_key), accrual_rule_key), requires_document = COALESCE(sqlc.narg(requires_document), requires_document), restricted = COALESCE(sqlc.narg(restricted), restricted), active = COALESCE(sqlc.narg(active), active), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
