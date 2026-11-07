-- name: GetTill :one
SELECT * FROM tills WHERE id = $1;

-- name: ListTill :many
SELECT * FROM tills
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR code ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateTill :one
INSERT INTO tills (site_id, code, pos_terminal_ref, active, created_by, updated_by)
VALUES (sqlc.arg(site_id), sqlc.arg(code), sqlc.narg(pos_terminal_ref), COALESCE(sqlc.narg(active)::boolean, true), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateTill :one
UPDATE tills
SET site_id = COALESCE(sqlc.narg(site_id), site_id), code = COALESCE(sqlc.narg(code), code), pos_terminal_ref = COALESCE(sqlc.narg(pos_terminal_ref), pos_terminal_ref), active = COALESCE(sqlc.narg(active), active), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
