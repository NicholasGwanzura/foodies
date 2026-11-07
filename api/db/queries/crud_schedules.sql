-- name: GetSchedule :one
SELECT * FROM schedules WHERE id = $1;

-- name: ListSchedule :many
SELECT * FROM schedules
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR title ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateSchedule :one
INSERT INTO schedules (title, workflow_key, form_template_id, frequency_rule_key, site_ids, assignee_role_id, due_offset_minutes, reminder_minutes, escalate_after_minutes, escalate_to_role_id, created_by, updated_by)
VALUES (sqlc.arg(title), sqlc.narg(workflow_key), sqlc.narg(form_template_id), sqlc.narg(frequency_rule_key), sqlc.narg(site_ids), sqlc.arg(assignee_role_id), sqlc.narg(due_offset_minutes), sqlc.narg(reminder_minutes), sqlc.narg(escalate_after_minutes), sqlc.narg(escalate_to_role_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateSchedule :one
UPDATE schedules
SET title = COALESCE(sqlc.narg(title), title), workflow_key = COALESCE(sqlc.narg(workflow_key), workflow_key), form_template_id = COALESCE(sqlc.narg(form_template_id), form_template_id), frequency_rule_key = COALESCE(sqlc.narg(frequency_rule_key), frequency_rule_key), site_ids = COALESCE(sqlc.narg(site_ids), site_ids), assignee_role_id = COALESCE(sqlc.narg(assignee_role_id), assignee_role_id), due_offset_minutes = COALESCE(sqlc.narg(due_offset_minutes), due_offset_minutes), reminder_minutes = COALESCE(sqlc.narg(reminder_minutes), reminder_minutes), escalate_after_minutes = COALESCE(sqlc.narg(escalate_after_minutes), escalate_after_minutes), escalate_to_role_id = COALESCE(sqlc.narg(escalate_to_role_id), escalate_to_role_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetScheduleStatus :one
UPDATE schedules
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActPauseSchedule :one
UPDATE schedules
SET status = 'paused', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('active')
RETURNING *;

-- name: ActResumeSchedule :one
UPDATE schedules
SET status = 'active', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
