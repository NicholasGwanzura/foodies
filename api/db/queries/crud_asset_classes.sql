-- name: GetAssetClass :one
SELECT * FROM asset_classes WHERE id = $1;

-- name: ListAssetClass :many
SELECT * FROM asset_classes
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR code ILIKE '%' || sqlc.narg(q)::text || '%' OR "name" ILIKE '%' || sqlc.narg(q)::text || '%')
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateAssetClass :one
INSERT INTO asset_classes (code, "name", depreciation_rate_percent, useful_life_years, capitalise, created_by, updated_by)
VALUES (sqlc.arg(code), sqlc.arg(name), sqlc.narg(depreciation_rate_percent), sqlc.narg(useful_life_years), COALESCE(sqlc.narg(capitalise)::boolean, false), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateAssetClass :one
UPDATE asset_classes
SET code = COALESCE(sqlc.narg(code), code), "name" = COALESCE(sqlc.narg(name), "name"), depreciation_rate_percent = COALESCE(sqlc.narg(depreciation_rate_percent), depreciation_rate_percent), useful_life_years = COALESCE(sqlc.narg(useful_life_years), useful_life_years), capitalise = COALESCE(sqlc.narg(capitalise), capitalise), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
