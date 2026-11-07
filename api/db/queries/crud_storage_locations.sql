-- name: GetStorageLocation :one
SELECT * FROM storage_locations WHERE id = $1;

-- name: ListStorageLocation :many
SELECT * FROM storage_locations
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR code ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateStorageLocation :one
INSERT INTO storage_locations (site_id, code, "zone", custodian_employee_id, active, created_by, updated_by)
VALUES (sqlc.arg(site_id), sqlc.arg(code), sqlc.arg(zone), sqlc.narg(custodian_employee_id), COALESCE(sqlc.narg(active)::boolean, true), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateStorageLocation :one
UPDATE storage_locations
SET site_id = COALESCE(sqlc.narg(site_id), site_id), code = COALESCE(sqlc.narg(code), code), "zone" = COALESCE(sqlc.narg(zone), "zone"), custodian_employee_id = COALESCE(sqlc.narg(custodian_employee_id), custodian_employee_id), active = COALESCE(sqlc.narg(active), active), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
