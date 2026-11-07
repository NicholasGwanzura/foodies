-- name: GetRetentionClass :one
SELECT * FROM retention_classes WHERE id = $1;

-- name: ListRetentionClass :many
SELECT * FROM retention_classes
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR "name" ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(decision_id)::text IS NULL OR decision_id = sqlc.narg(decision_id)::text)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateRetentionClass :one
INSERT INTO retention_classes ("key", "name", period_value, period_unit, trigger, source_clause, decision_id, disposition, created_by, updated_by)
VALUES (sqlc.arg(key), sqlc.arg(name), sqlc.narg(period_value), sqlc.narg(period_unit), sqlc.arg(trigger), sqlc.narg(source_clause), sqlc.narg(decision_id), sqlc.narg(disposition), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateRetentionClass :one
UPDATE retention_classes
SET "key" = COALESCE(sqlc.narg(key), "key"), "name" = COALESCE(sqlc.narg(name), "name"), period_value = COALESCE(sqlc.narg(period_value), period_value), period_unit = COALESCE(sqlc.narg(period_unit), period_unit), trigger = COALESCE(sqlc.narg(trigger), trigger), source_clause = COALESCE(sqlc.narg(source_clause), source_clause), decision_id = COALESCE(sqlc.narg(decision_id), decision_id), disposition = COALESCE(sqlc.narg(disposition), disposition), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetRetentionClassStatus :one
UPDATE retention_classes
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
