-- name: GetCctvOutage :one
SELECT * FROM cctv_outages WHERE id = $1;

-- name: ListCctvOutage :many
SELECT * FROM cctv_outages
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

-- name: CreateCctvOutage :one
INSERT INTO cctv_outages (reference, site_id, camera_ids, started_at, cause, alternative_measures, ict_case, test_outcome, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(camera_ids), sqlc.arg(started_at), sqlc.narg(cause), sqlc.arg(alternative_measures), sqlc.narg(ict_case), sqlc.narg(test_outcome), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateCctvOutage :one
UPDATE cctv_outages
SET site_id = COALESCE(sqlc.narg(site_id), site_id), camera_ids = COALESCE(sqlc.narg(camera_ids), camera_ids), started_at = COALESCE(sqlc.narg(started_at), started_at), cause = COALESCE(sqlc.narg(cause), cause), alternative_measures = COALESCE(sqlc.narg(alternative_measures), alternative_measures), ict_case = COALESCE(sqlc.narg(ict_case), ict_case), test_outcome = COALESCE(sqlc.narg(test_outcome), test_outcome), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetCctvOutageStatus :one
UPDATE cctv_outages
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActRestoreCctvOutage :one
UPDATE cctv_outages
SET ended_at = COALESCE(sqlc.narg(ended_at), ended_at), test_outcome = COALESCE(sqlc.narg(test_outcome), test_outcome), status = 'restored', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('open')
RETURNING *;
