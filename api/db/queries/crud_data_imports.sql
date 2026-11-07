-- name: GetDataImport :one
SELECT * FROM data_imports WHERE id = $1;

-- name: ListDataImport :many
SELECT * FROM data_imports
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(kind)::text IS NULL OR kind = sqlc.narg(kind)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR business_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR business_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateDataImport :one
INSERT INTO data_imports (kind, attachment_id, site_id, business_date, created_by, updated_by)
VALUES (sqlc.arg(kind), sqlc.arg(attachment_id), sqlc.narg(site_id), sqlc.narg(business_date), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetDataImportStatus :one
UPDATE data_imports
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
