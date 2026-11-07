-- name: GetAttendanceRecord :one
SELECT * FROM attendance_records WHERE id = $1;

-- name: ListAttendanceRecord :many
SELECT * FROM attendance_records
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(employee_id)::uuid IS NULL OR employee_id = sqlc.narg(employee_id)::uuid)
  AND (sqlc.narg(work_date)::date IS NULL OR work_date = sqlc.narg(work_date)::date)
  AND (sqlc.narg(from_date)::date IS NULL OR work_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR work_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateAttendanceRecord :one
INSERT INTO attendance_records (employee_id, site_id, work_date, clock_in, clock_out, "status", "reason", created_by, updated_by)
VALUES (sqlc.arg(employee_id), sqlc.arg(site_id), sqlc.arg(work_date), sqlc.narg(clock_in), sqlc.narg(clock_out), sqlc.arg(status), sqlc.narg(reason), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateAttendanceRecord :one
UPDATE attendance_records
SET employee_id = COALESCE(sqlc.narg(employee_id), employee_id), site_id = COALESCE(sqlc.narg(site_id), site_id), work_date = COALESCE(sqlc.narg(work_date), work_date), clock_in = COALESCE(sqlc.narg(clock_in), clock_in), clock_out = COALESCE(sqlc.narg(clock_out), clock_out), "reason" = COALESCE(sqlc.narg(reason), "reason"), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
