-- name: GetReportDeadlineNotice :one
SELECT * FROM report_deadline_notices WHERE id = $1;

-- name: ListReportDeadlineNotice :many
SELECT * FROM report_deadline_notices
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(from_date)::date IS NULL OR period_start::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR period_start::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateReportDeadlineNotice :one
INSERT INTO report_deadline_notices (report_definition_id, period_start, new_due_at, notice_attachment_id, created_by, updated_by)
VALUES (sqlc.arg(report_definition_id), sqlc.arg(period_start), sqlc.arg(new_due_at), sqlc.narg(notice_attachment_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;
