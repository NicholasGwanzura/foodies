-- name: GetSafe :one
SELECT * FROM safes WHERE id = $1;

-- name: ListSafe :many
SELECT * FROM safes
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR code ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateSafe :one
INSERT INTO safes (site_id, code, location, lock_status, active, created_by, updated_by)
VALUES (sqlc.arg(site_id), sqlc.arg(code), sqlc.narg(location), sqlc.narg(lock_status), COALESCE(sqlc.narg(active)::boolean, true), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateSafe :one
UPDATE safes
SET site_id = COALESCE(sqlc.narg(site_id), site_id), code = COALESCE(sqlc.narg(code), code), location = COALESCE(sqlc.narg(location), location), lock_status = COALESCE(sqlc.narg(lock_status), lock_status), active = COALESCE(sqlc.narg(active), active), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
