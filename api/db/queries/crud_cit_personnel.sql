-- name: GetCitPerson :one
SELECT * FROM cit_personnel WHERE id = $1;

-- name: ListCitPerson :many
SELECT * FROM cit_personnel
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateCitPerson :one
INSERT INTO cit_personnel (cit_provider_id, full_name, id_number, active, created_by, updated_by)
VALUES (sqlc.arg(cit_provider_id), sqlc.arg(full_name), sqlc.arg(id_number), COALESCE(sqlc.narg(active)::boolean, true), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateCitPerson :one
UPDATE cit_personnel
SET full_name = COALESCE(sqlc.narg(full_name), full_name), id_number = COALESCE(sqlc.narg(id_number), id_number), active = COALESCE(sqlc.narg(active), active), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
