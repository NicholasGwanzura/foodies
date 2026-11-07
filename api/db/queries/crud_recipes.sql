-- name: GetRecipe :one
SELECT * FROM recipes WHERE id = $1;

-- name: ListRecipe :many
SELECT * FROM recipes
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR code ILIKE '%' || sqlc.narg(q)::text || '%' OR "name" ILIKE '%' || sqlc.narg(q)::text || '%')
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateRecipe :one
INSERT INTO recipes (code, "name", product_id, owner_role_id, active, created_by, updated_by)
VALUES (sqlc.arg(code), sqlc.arg(name), sqlc.narg(product_id), sqlc.narg(owner_role_id), COALESCE(sqlc.narg(active)::boolean, true), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateRecipe :one
UPDATE recipes
SET code = COALESCE(sqlc.narg(code), code), "name" = COALESCE(sqlc.narg(name), "name"), product_id = COALESCE(sqlc.narg(product_id), product_id), owner_role_id = COALESCE(sqlc.narg(owner_role_id), owner_role_id), active = COALESCE(sqlc.narg(active), active), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
