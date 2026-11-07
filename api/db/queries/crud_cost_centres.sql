-- name: GetCostCentre :one
SELECT * FROM cost_centres WHERE id = $1;

-- name: ListCostCentre :many
SELECT * FROM cost_centres
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR code ILIKE '%' || sqlc.narg(q)::text || '%' OR "name" ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateCostCentre :one
INSERT INTO cost_centres (code, "name", department_id, site_id, sap_code, active, created_by, updated_by)
VALUES (sqlc.arg(code), sqlc.arg(name), sqlc.narg(department_id), sqlc.narg(site_id), sqlc.narg(sap_code), COALESCE(sqlc.narg(active)::boolean, true), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateCostCentre :one
UPDATE cost_centres
SET code = COALESCE(sqlc.narg(code), code), "name" = COALESCE(sqlc.narg(name), "name"), department_id = COALESCE(sqlc.narg(department_id), department_id), site_id = COALESCE(sqlc.narg(site_id), site_id), sap_code = COALESCE(sqlc.narg(sap_code), sap_code), active = COALESCE(sqlc.narg(active), active), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
