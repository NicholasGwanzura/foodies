-- name: GetPosOutage :one
SELECT * FROM pos_outages WHERE id = $1;

-- name: ListPosOutage :many
SELECT * FROM pos_outages
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreatePosOutage :one
INSERT INTO pos_outages (reference, site_id, started_at, cause, affected_systems, ict_case, restoration_estimate, reconciliation_note, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(started_at), sqlc.arg(cause), sqlc.narg(affected_systems), sqlc.narg(ict_case), sqlc.narg(restoration_estimate), sqlc.narg(reconciliation_note), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdatePosOutage :one
UPDATE pos_outages
SET site_id = COALESCE(sqlc.narg(site_id), site_id), started_at = COALESCE(sqlc.narg(started_at), started_at), cause = COALESCE(sqlc.narg(cause), cause), affected_systems = COALESCE(sqlc.narg(affected_systems), affected_systems), ict_case = COALESCE(sqlc.narg(ict_case), ict_case), restoration_estimate = COALESCE(sqlc.narg(restoration_estimate), restoration_estimate), reconciliation_note = COALESCE(sqlc.narg(reconciliation_note), reconciliation_note), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetPosOutageStatus :one
UPDATE pos_outages
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitPosOutage :one
UPDATE pos_outages
SET status = 'pending_authorisation', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('declared')
RETURNING *;

-- name: ActRestorePosOutage :one
UPDATE pos_outages
SET restored_at = COALESCE(sqlc.narg(restored_at), restored_at), status = 'restored', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('declared', 'pending_authorisation', 'authorised')
RETURNING *;

-- name: ActClosePosOutage :one
UPDATE pos_outages
SET reconciliation_note = COALESCE(sqlc.narg(reconciliation_note), reconciliation_note), status = 'closed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('declared', 'pending_authorisation', 'authorised', 'restored')
RETURNING *;
