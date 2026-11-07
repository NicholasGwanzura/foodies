-- name: GetItem :one
SELECT * FROM items WHERE id = $1;

-- name: ListItem :many
SELECT * FROM items
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR code ILIKE '%' || sqlc.narg(q)::text || '%' OR "name" ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(category)::text IS NULL OR category = sqlc.narg(category)::text)
  AND (sqlc.narg(stock_class)::text IS NULL OR stock_class = sqlc.narg(stock_class)::text)
  AND (sqlc.narg(critical)::boolean IS NULL OR critical = sqlc.narg(critical)::boolean)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateItem :one
INSERT INTO items (code, "name", category, stock_class, uom, critical, rotation, shelf_life_days, gaap_ref, active, created_by, updated_by)
VALUES (sqlc.arg(code), sqlc.arg(name), sqlc.arg(category), sqlc.arg(stock_class), sqlc.arg(uom), COALESCE(sqlc.narg(critical)::boolean, false), sqlc.arg(rotation), sqlc.narg(shelf_life_days), sqlc.narg(gaap_ref), COALESCE(sqlc.narg(active)::boolean, true), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateItem :one
UPDATE items
SET code = COALESCE(sqlc.narg(code), code), "name" = COALESCE(sqlc.narg(name), "name"), category = COALESCE(sqlc.narg(category), category), stock_class = COALESCE(sqlc.narg(stock_class), stock_class), uom = COALESCE(sqlc.narg(uom), uom), critical = COALESCE(sqlc.narg(critical), critical), rotation = COALESCE(sqlc.narg(rotation), rotation), shelf_life_days = COALESCE(sqlc.narg(shelf_life_days), shelf_life_days), gaap_ref = COALESCE(sqlc.narg(gaap_ref), gaap_ref), active = COALESCE(sqlc.narg(active), active), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
