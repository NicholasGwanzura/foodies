-- name: GetExitClearance :one
SELECT * FROM exit_clearances WHERE id = $1;

-- name: ListExitClearance :many
SELECT * FROM exit_clearances
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(employee_id)::uuid IS NULL OR employee_id = sqlc.narg(employee_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR notice_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR notice_date::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateExitClearance :one
INSERT INTO exit_clearances (reference, employee_id, exit_type, notice_date, last_working_day, reason_category, hr_officer_employee_id, exit_interview_notes, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(employee_id), sqlc.arg(exit_type), sqlc.narg(notice_date), sqlc.arg(last_working_day), sqlc.narg(reason_category), sqlc.narg(hr_officer_employee_id), sqlc.narg(exit_interview_notes), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateExitClearance :one
UPDATE exit_clearances
SET employee_id = COALESCE(sqlc.narg(employee_id), employee_id), exit_type = COALESCE(sqlc.narg(exit_type), exit_type), notice_date = COALESCE(sqlc.narg(notice_date), notice_date), last_working_day = COALESCE(sqlc.narg(last_working_day), last_working_day), reason_category = COALESCE(sqlc.narg(reason_category), reason_category), hr_officer_employee_id = COALESCE(sqlc.narg(hr_officer_employee_id), hr_officer_employee_id), exit_interview_notes = COALESCE(sqlc.narg(exit_interview_notes), exit_interview_notes), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetExitClearanceStatus :one
UPDATE exit_clearances
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListExitClearanceSections :many
SELECT * FROM exit_clearance_sections WHERE exit_clearance_id = $1 ORDER BY id;

-- name: InsertExitClearanceSections :one
INSERT INTO exit_clearance_sections (exit_clearance_id, section, "status", outstanding_items, signed_by, signed_at, created_by)
VALUES (sqlc.arg(exit_clearance_id), sqlc.arg(section), sqlc.narg(status), sqlc.narg(outstanding_items), sqlc.narg(signed_by), sqlc.narg(signed_at), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteExitClearanceSections :exec
DELETE FROM exit_clearance_sections WHERE exit_clearance_id = $1;

-- name: ActSignSectionExitClearance :one
UPDATE exit_clearances
SET updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActCompleteExitClearance :one
UPDATE exit_clearances
SET status = 'completed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('initiated', 'in_clearance', 'cleared')
RETURNING *;
