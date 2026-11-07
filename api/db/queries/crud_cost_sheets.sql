-- name: GetCostSheet :one
SELECT * FROM cost_sheets WHERE id = $1;

-- name: ListCostSheet :many
SELECT * FROM cost_sheets
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateCostSheet :one
INSERT INTO cost_sheets (reference, recipe_version_id, currency, ingredient_cost, labour_cost, overhead_cost, selling_price, review_trigger, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(recipe_version_id), sqlc.arg(currency), sqlc.arg(ingredient_cost), sqlc.narg(labour_cost), sqlc.narg(overhead_cost), sqlc.narg(selling_price), sqlc.narg(review_trigger), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateCostSheet :one
UPDATE cost_sheets
SET recipe_version_id = COALESCE(sqlc.narg(recipe_version_id), recipe_version_id), currency = COALESCE(sqlc.narg(currency), currency), ingredient_cost = COALESCE(sqlc.narg(ingredient_cost), ingredient_cost), labour_cost = COALESCE(sqlc.narg(labour_cost), labour_cost), overhead_cost = COALESCE(sqlc.narg(overhead_cost), overhead_cost), selling_price = COALESCE(sqlc.narg(selling_price), selling_price), review_trigger = COALESCE(sqlc.narg(review_trigger), review_trigger), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetCostSheetStatus :one
UPDATE cost_sheets
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitCostSheet :one
UPDATE cost_sheets
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
