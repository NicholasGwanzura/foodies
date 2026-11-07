-- name: GetRecallRetrieval :one
SELECT * FROM recall_retrievals WHERE id = $1;

-- name: ListRecallRetrieval :many
SELECT * FROM recall_retrievals
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateRecallRetrieval :one
INSERT INTO recall_retrievals (recall_id, location_site_id, location_name, batch_id, quantity_dispatched, quantity_returned, returned_at, condition, quarantine_location_id, created_by, updated_by)
VALUES (sqlc.arg(recall_id), sqlc.narg(location_site_id), sqlc.narg(location_name), sqlc.narg(batch_id), sqlc.narg(quantity_dispatched), sqlc.narg(quantity_returned), sqlc.narg(returned_at), sqlc.narg(condition), sqlc.narg(quarantine_location_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateRecallRetrieval :one
UPDATE recall_retrievals
SET location_site_id = COALESCE(sqlc.narg(location_site_id), location_site_id), location_name = COALESCE(sqlc.narg(location_name), location_name), batch_id = COALESCE(sqlc.narg(batch_id), batch_id), quantity_dispatched = COALESCE(sqlc.narg(quantity_dispatched), quantity_dispatched), quantity_returned = COALESCE(sqlc.narg(quantity_returned), quantity_returned), returned_at = COALESCE(sqlc.narg(returned_at), returned_at), condition = COALESCE(sqlc.narg(condition), condition), quarantine_location_id = COALESCE(sqlc.narg(quarantine_location_id), quarantine_location_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetRecallRetrievalStatus :one
UPDATE recall_retrievals
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
