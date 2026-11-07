-- name: GetRecipeVersion :one
SELECT * FROM recipe_versions WHERE id = $1;

-- name: ListRecipeVersion :many
SELECT * FROM recipe_versions
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR effective_from::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR effective_from::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateRecipeVersion :one
INSERT INTO recipe_versions (recipe_id, batch_size, uom, standard_output, standard_yield_percent, normal_waste_percent, method_document_id, effective_from, created_by, updated_by)
VALUES (sqlc.arg(recipe_id), sqlc.arg(batch_size), sqlc.arg(uom), sqlc.arg(standard_output), sqlc.arg(standard_yield_percent), sqlc.narg(normal_waste_percent), sqlc.narg(method_document_id), sqlc.narg(effective_from), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateRecipeVersion :one
UPDATE recipe_versions
SET batch_size = COALESCE(sqlc.narg(batch_size), batch_size), uom = COALESCE(sqlc.narg(uom), uom), standard_output = COALESCE(sqlc.narg(standard_output), standard_output), standard_yield_percent = COALESCE(sqlc.narg(standard_yield_percent), standard_yield_percent), normal_waste_percent = COALESCE(sqlc.narg(normal_waste_percent), normal_waste_percent), method_document_id = COALESCE(sqlc.narg(method_document_id), method_document_id), effective_from = COALESCE(sqlc.narg(effective_from), effective_from), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetRecipeVersionStatus :one
UPDATE recipe_versions
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListRecipeIngredients :many
SELECT * FROM recipe_ingredients WHERE recipe_version_id = $1 ORDER BY id;

-- name: InsertRecipeIngredients :one
INSERT INTO recipe_ingredients (recipe_version_id, item_id, quantity, uom, allergen_tags, created_by)
VALUES (sqlc.arg(recipe_version_id), sqlc.arg(item_id), sqlc.arg(quantity), sqlc.arg(uom), sqlc.narg(allergen_tags), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteRecipeIngredients :exec
DELETE FROM recipe_ingredients WHERE recipe_version_id = $1;

-- name: ActSubmitRecipeVersion :one
UPDATE recipe_versions
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
