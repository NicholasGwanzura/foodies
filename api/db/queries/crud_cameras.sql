-- name: GetCamera :one
SELECT * FROM cameras WHERE id = $1;

-- name: ListCamera :many
SELECT * FROM cameras
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR label ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateCamera :one
INSERT INTO cameras (site_id, label, location, critical_area, recorder, "status", created_by, updated_by)
VALUES (sqlc.arg(site_id), sqlc.arg(label), sqlc.arg(location), COALESCE(sqlc.narg(critical_area)::boolean, false), sqlc.narg(recorder), sqlc.narg(status), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateCamera :one
UPDATE cameras
SET site_id = COALESCE(sqlc.narg(site_id), site_id), label = COALESCE(sqlc.narg(label), label), location = COALESCE(sqlc.narg(location), location), critical_area = COALESCE(sqlc.narg(critical_area), critical_area), recorder = COALESCE(sqlc.narg(recorder), recorder), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
