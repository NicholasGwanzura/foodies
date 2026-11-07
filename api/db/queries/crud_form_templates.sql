-- name: GetFormTemplate :one
SELECT * FROM form_templates WHERE id = $1;

-- name: ListFormTemplate :many
SELECT * FROM form_templates
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR title ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(module)::text IS NULL OR module = sqlc.narg(module)::text)
  AND (sqlc.narg(kind)::text IS NULL OR kind = sqlc.narg(kind)::text)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateFormTemplate :one
INSERT INTO form_templates ("key", title, kind, module, area, source_sop, owner_role_id, created_by, updated_by)
VALUES (sqlc.arg(key), sqlc.arg(title), sqlc.arg(kind), sqlc.arg(module), sqlc.narg(area), sqlc.narg(source_sop), sqlc.narg(owner_role_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateFormTemplate :one
UPDATE form_templates
SET "key" = COALESCE(sqlc.narg(key), "key"), title = COALESCE(sqlc.narg(title), title), kind = COALESCE(sqlc.narg(kind), kind), module = COALESCE(sqlc.narg(module), module), area = COALESCE(sqlc.narg(area), area), source_sop = COALESCE(sqlc.narg(source_sop), source_sop), owner_role_id = COALESCE(sqlc.narg(owner_role_id), owner_role_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetFormTemplateStatus :one
UPDATE form_templates
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
