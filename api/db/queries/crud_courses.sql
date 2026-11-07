-- name: GetCourse :one
SELECT * FROM courses WHERE id = $1;

-- name: ListCourse :many
SELECT * FROM courses
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR code ILIKE '%' || sqlc.narg(q)::text || '%' OR title ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(category)::text IS NULL OR category = sqlc.narg(category)::text)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateCourse :one
INSERT INTO courses (code, title, category, objective, document_ids, pass_mark, requires_practical, refresher_rule_key, material_attachment_ids, created_by, updated_by)
VALUES (sqlc.arg(code), sqlc.arg(title), sqlc.arg(category), sqlc.narg(objective), sqlc.narg(document_ids), sqlc.narg(pass_mark), COALESCE(sqlc.narg(requires_practical)::boolean, false), sqlc.narg(refresher_rule_key), sqlc.narg(material_attachment_ids), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateCourse :one
UPDATE courses
SET code = COALESCE(sqlc.narg(code), code), title = COALESCE(sqlc.narg(title), title), category = COALESCE(sqlc.narg(category), category), objective = COALESCE(sqlc.narg(objective), objective), document_ids = COALESCE(sqlc.narg(document_ids), document_ids), pass_mark = COALESCE(sqlc.narg(pass_mark), pass_mark), requires_practical = COALESCE(sqlc.narg(requires_practical), requires_practical), refresher_rule_key = COALESCE(sqlc.narg(refresher_rule_key), refresher_rule_key), material_attachment_ids = COALESCE(sqlc.narg(material_attachment_ids), material_attachment_ids), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
