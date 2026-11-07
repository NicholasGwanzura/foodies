-- name: GetDevice :one
SELECT * FROM devices WHERE id = $1;

-- name: ListDevice :many
SELECT * FROM devices
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR label ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateDevice :one
INSERT INTO devices (label, site_id, platform, shared, created_by, updated_by)
VALUES (sqlc.arg(label), sqlc.narg(site_id), sqlc.narg(platform), COALESCE(sqlc.narg(shared)::boolean, false), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateDevice :one
UPDATE devices
SET label = COALESCE(sqlc.narg(label), label), site_id = COALESCE(sqlc.narg(site_id), site_id), platform = COALESCE(sqlc.narg(platform), platform), shared = COALESCE(sqlc.narg(shared), shared), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetDeviceStatus :one
UPDATE devices
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActRevokeDevice :one
UPDATE devices
SET status = 'revoked', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('active')
RETURNING *;
