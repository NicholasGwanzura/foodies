-- name: GetFormTemplateVersion :one
SELECT * FROM form_template_versions WHERE id = $1;

-- name: ListFormTemplateVersion :many
SELECT * FROM form_template_versions
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateFormTemplateVersion :one
INSERT INTO form_template_versions (form_template_id, definition, offline_allowed, channels, retention_class_id, required_rule_keys, effective_from, created_by, updated_by)
VALUES (sqlc.arg(form_template_id), sqlc.arg(definition), COALESCE(sqlc.narg(offline_allowed)::boolean, false), sqlc.arg(channels), sqlc.narg(retention_class_id), sqlc.narg(required_rule_keys), sqlc.narg(effective_from), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateFormTemplateVersion :one
UPDATE form_template_versions
SET definition = COALESCE(sqlc.narg(definition), definition), offline_allowed = COALESCE(sqlc.narg(offline_allowed), offline_allowed), channels = COALESCE(sqlc.narg(channels), channels), retention_class_id = COALESCE(sqlc.narg(retention_class_id), retention_class_id), required_rule_keys = COALESCE(sqlc.narg(required_rule_keys), required_rule_keys), effective_from = COALESCE(sqlc.narg(effective_from), effective_from), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetFormTemplateVersionStatus :one
UPDATE form_template_versions
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitFormTemplateVersion :one
UPDATE form_template_versions
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActPublishFormTemplateVersion :one
UPDATE form_template_versions
SET status = 'published', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'pending_approval')
RETURNING *;

-- name: ActRetireFormTemplateVersion :one
UPDATE form_template_versions
SET status = 'retired', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'pending_approval', 'published')
RETURNING *;
