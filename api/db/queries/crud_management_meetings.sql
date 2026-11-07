-- name: GetManagementMeeting :one
SELECT * FROM management_meetings WHERE id = $1;

-- name: ListManagementMeeting :many
SELECT * FROM management_meetings
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateManagementMeeting :one
INSERT INTO management_meetings (reference, meeting_at, agenda, attendee_employee_ids, minutes, decisions, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(meeting_at), sqlc.narg(agenda), sqlc.narg(attendee_employee_ids), sqlc.narg(minutes), sqlc.narg(decisions), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateManagementMeeting :one
UPDATE management_meetings
SET meeting_at = COALESCE(sqlc.narg(meeting_at), meeting_at), agenda = COALESCE(sqlc.narg(agenda), agenda), attendee_employee_ids = COALESCE(sqlc.narg(attendee_employee_ids), attendee_employee_ids), minutes = COALESCE(sqlc.narg(minutes), minutes), decisions = COALESCE(sqlc.narg(decisions), decisions), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetManagementMeetingStatus :one
UPDATE management_meetings
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
