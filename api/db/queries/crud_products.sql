-- name: GetProduct :one
SELECT * FROM products WHERE id = $1;

-- name: ListProduct :many
SELECT * FROM products
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR code ILIKE '%' || sqlc.narg(q)::text || '%' OR "name" ILIKE '%' || sqlc.narg(q)::text || '%' OR description ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(category)::text IS NULL OR category = sqlc.narg(category)::text)
  AND (sqlc.narg(active)::boolean IS NULL OR active = sqlc.narg(active)::boolean)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateProduct :one
INSERT INTO products (code, "name", category, description, active, branch_site_ids, recipe_id, menu_version, allergens, portion_standard, storage_and_shelf_life, image_attachment_id, card, faq, created_by, updated_by)
VALUES (sqlc.arg(code), sqlc.arg(name), sqlc.arg(category), sqlc.narg(description), COALESCE(sqlc.narg(active)::boolean, true), sqlc.narg(branch_site_ids), sqlc.narg(recipe_id), sqlc.narg(menu_version), sqlc.narg(allergens), sqlc.narg(portion_standard), sqlc.narg(storage_and_shelf_life), sqlc.narg(image_attachment_id), sqlc.narg(card), sqlc.narg(faq), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateProduct :one
UPDATE products
SET code = COALESCE(sqlc.narg(code), code), "name" = COALESCE(sqlc.narg(name), "name"), category = COALESCE(sqlc.narg(category), category), description = COALESCE(sqlc.narg(description), description), active = COALESCE(sqlc.narg(active), active), branch_site_ids = COALESCE(sqlc.narg(branch_site_ids), branch_site_ids), recipe_id = COALESCE(sqlc.narg(recipe_id), recipe_id), menu_version = COALESCE(sqlc.narg(menu_version), menu_version), allergens = COALESCE(sqlc.narg(allergens), allergens), portion_standard = COALESCE(sqlc.narg(portion_standard), portion_standard), storage_and_shelf_life = COALESCE(sqlc.narg(storage_and_shelf_life), storage_and_shelf_life), image_attachment_id = COALESCE(sqlc.narg(image_attachment_id), image_attachment_id), card = COALESCE(sqlc.narg(card), card), faq = COALESCE(sqlc.narg(faq), faq), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
