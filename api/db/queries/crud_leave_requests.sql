-- name: GetLeaveRequest :one
SELECT * FROM leave_requests WHERE id = $1;

-- name: ListLeaveRequest :many
SELECT * FROM leave_requests
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(employee_id)::uuid IS NULL OR employee_id = sqlc.narg(employee_id)::uuid)
  AND (sqlc.narg(leave_type_id)::uuid IS NULL OR leave_type_id = sqlc.narg(leave_type_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR start_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR start_date::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateLeaveRequest :one
INSERT INTO leave_requests (reference, employee_id, leave_type_id, start_date, end_date, "reason", document_attachment_id, coverage_plan, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(employee_id), sqlc.arg(leave_type_id), sqlc.arg(start_date), sqlc.arg(end_date), sqlc.narg(reason), sqlc.narg(document_attachment_id), sqlc.narg(coverage_plan), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateLeaveRequest :one
UPDATE leave_requests
SET employee_id = COALESCE(sqlc.narg(employee_id), employee_id), leave_type_id = COALESCE(sqlc.narg(leave_type_id), leave_type_id), start_date = COALESCE(sqlc.narg(start_date), start_date), end_date = COALESCE(sqlc.narg(end_date), end_date), "reason" = COALESCE(sqlc.narg(reason), "reason"), document_attachment_id = COALESCE(sqlc.narg(document_attachment_id), document_attachment_id), coverage_plan = COALESCE(sqlc.narg(coverage_plan), coverage_plan), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetLeaveRequestStatus :one
UPDATE leave_requests
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitLeaveRequest :one
UPDATE leave_requests
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActCancelLeaveRequest :one
UPDATE leave_requests
SET status = 'cancelled', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'pending_approval', 'approved', 'taken')
RETURNING *;
