-- name: GetRoster :one
SELECT * FROM rosters WHERE id = $1;

-- name: ListRoster :many
SELECT * FROM rosters
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR period_start::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR period_start::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateRoster :one
INSERT INTO rosters (reference, site_id, period_start, period_end, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(period_start), sqlc.arg(period_end), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateRoster :one
UPDATE rosters
SET site_id = COALESCE(sqlc.narg(site_id), site_id), period_start = COALESCE(sqlc.narg(period_start), period_start), period_end = COALESCE(sqlc.narg(period_end), period_end), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetRosterStatus :one
UPDATE rosters
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListRosterShifts :many
SELECT * FROM roster_shifts WHERE roster_id = $1 ORDER BY id;

-- name: InsertRosterShifts :one
INSERT INTO roster_shifts (roster_id, employee_id, shift_date, starts_at, ends_at, "role", swapped_with_employee_id, swap_reason, created_by)
VALUES (sqlc.arg(roster_id), sqlc.arg(employee_id), sqlc.arg(shift_date), sqlc.arg(starts_at), sqlc.arg(ends_at), sqlc.narg(role), sqlc.narg(swapped_with_employee_id), sqlc.narg(swap_reason), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteRosterShifts :exec
DELETE FROM roster_shifts WHERE roster_id = $1;

-- name: ActSubmitRoster :one
UPDATE rosters
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
