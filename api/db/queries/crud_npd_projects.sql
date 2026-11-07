-- name: GetNpdProject :one
SELECT * FROM npd_projects WHERE id = $1;

-- name: ListNpdProject :many
SELECT * FROM npd_projects
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%' OR "name" ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR launched_on::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR launched_on::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateNpdProject :one
INSERT INTO npd_projects (reference, "name", business_case, owner_employee_id, product_id, launched_on, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(name), sqlc.arg(business_case), sqlc.narg(owner_employee_id), sqlc.narg(product_id), sqlc.narg(launched_on), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateNpdProject :one
UPDATE npd_projects
SET "name" = COALESCE(sqlc.narg(name), "name"), business_case = COALESCE(sqlc.narg(business_case), business_case), owner_employee_id = COALESCE(sqlc.narg(owner_employee_id), owner_employee_id), product_id = COALESCE(sqlc.narg(product_id), product_id), launched_on = COALESCE(sqlc.narg(launched_on), launched_on), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetNpdProjectStatus :one
UPDATE npd_projects
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActAdvanceGateNpdProject :one
UPDATE npd_projects
SET updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
