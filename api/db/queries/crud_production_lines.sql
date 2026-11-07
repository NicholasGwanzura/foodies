-- name: GetProductionLine :one
SELECT * FROM production_lines WHERE id = $1;

-- name: ListProductionLine :many
SELECT * FROM production_lines
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR "name" ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateProductionLine :one
INSERT INTO production_lines (site_id, "name", "type", active, created_by, updated_by)
VALUES (sqlc.arg(site_id), sqlc.arg(name), sqlc.arg(type), COALESCE(sqlc.narg(active)::boolean, true), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateProductionLine :one
UPDATE production_lines
SET site_id = COALESCE(sqlc.narg(site_id), site_id), "name" = COALESCE(sqlc.narg(name), "name"), "type" = COALESCE(sqlc.narg(type), "type"), active = COALESCE(sqlc.narg(active), active), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
