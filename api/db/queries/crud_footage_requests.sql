-- name: GetFootageRequest :one
SELECT * FROM footage_requests WHERE id = $1;

-- name: ListFootageRequest :many
SELECT * FROM footage_requests
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR hold_until::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR hold_until::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateFootageRequest :one
INSERT INTO footage_requests (reference, requester_employee_id, site_id, purpose, camera_ids, window_start, window_end, incident_id, hold_until, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(requester_employee_id), sqlc.arg(site_id), sqlc.arg(purpose), sqlc.arg(camera_ids), sqlc.arg(window_start), sqlc.arg(window_end), sqlc.narg(incident_id), sqlc.narg(hold_until), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateFootageRequest :one
UPDATE footage_requests
SET requester_employee_id = COALESCE(sqlc.narg(requester_employee_id), requester_employee_id), site_id = COALESCE(sqlc.narg(site_id), site_id), purpose = COALESCE(sqlc.narg(purpose), purpose), camera_ids = COALESCE(sqlc.narg(camera_ids), camera_ids), window_start = COALESCE(sqlc.narg(window_start), window_start), window_end = COALESCE(sqlc.narg(window_end), window_end), incident_id = COALESCE(sqlc.narg(incident_id), incident_id), hold_until = COALESCE(sqlc.narg(hold_until), hold_until), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetFootageRequestStatus :one
UPDATE footage_requests
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitFootageRequest :one
UPDATE footage_requests
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActRecordExportFootageRequest :one
UPDATE footage_requests
SET export_filename = COALESCE(sqlc.narg(export_filename), export_filename), export_location = COALESCE(sqlc.narg(export_location), export_location), file_hash = COALESCE(sqlc.narg(file_hash), file_hash), status = 'exported', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('approved', 'preserved')
RETURNING *;

-- name: ActAddCustodyEntryFootageRequest :one
UPDATE footage_requests
SET updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActCloseFootageRequest :one
UPDATE footage_requests
SET closed_at = now(), status = 'closed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'pending_approval', 'approved', 'preserved', 'exported')
RETURNING *;
