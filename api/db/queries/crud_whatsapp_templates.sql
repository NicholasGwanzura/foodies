-- name: GetWhatsappTemplate :one
SELECT * FROM whatsapp_templates WHERE id = $1;

-- name: ListWhatsappTemplate :many
SELECT * FROM whatsapp_templates
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR "name" ILIKE '%' || sqlc.narg(q)::text || '%')
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateWhatsappTemplate :one
INSERT INTO whatsapp_templates ("name", language, category, body, created_by, updated_by)
VALUES (sqlc.arg(name), sqlc.arg(language), sqlc.arg(category), sqlc.arg(body), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateWhatsappTemplate :one
UPDATE whatsapp_templates
SET "name" = COALESCE(sqlc.narg(name), "name"), language = COALESCE(sqlc.narg(language), language), category = COALESCE(sqlc.narg(category), category), body = COALESCE(sqlc.narg(body), body), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
