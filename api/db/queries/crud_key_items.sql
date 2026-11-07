-- name: GetKeyItem :one
SELECT * FROM key_items WHERE id = $1;

-- name: ListKeyItem :many
SELECT * FROM key_items
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR label ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(key_type)::text IS NULL OR key_type = sqlc.narg(key_type)::text)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateKeyItem :one
INSERT INTO key_items (site_id, label, key_type, areas_secured, primary_custodian_employee_id, backup_location, master_location, created_by, updated_by)
VALUES (sqlc.arg(site_id), sqlc.arg(label), sqlc.arg(key_type), sqlc.narg(areas_secured), sqlc.narg(primary_custodian_employee_id), sqlc.narg(backup_location), sqlc.narg(master_location), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateKeyItem :one
UPDATE key_items
SET site_id = COALESCE(sqlc.narg(site_id), site_id), label = COALESCE(sqlc.narg(label), label), key_type = COALESCE(sqlc.narg(key_type), key_type), areas_secured = COALESCE(sqlc.narg(areas_secured), areas_secured), primary_custodian_employee_id = COALESCE(sqlc.narg(primary_custodian_employee_id), primary_custodian_employee_id), backup_location = COALESCE(sqlc.narg(backup_location), backup_location), master_location = COALESCE(sqlc.narg(master_location), master_location), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetKeyItemStatus :one
UPDATE key_items
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
