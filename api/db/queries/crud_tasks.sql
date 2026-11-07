-- name: GetTask :one
SELECT * FROM tasks WHERE id = $1;

-- name: ListTask :many
SELECT * FROM tasks
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR title ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(overdue)::boolean IS NULL OR overdue = sqlc.narg(overdue)::boolean)
  AND (sqlc.narg(business_date)::date IS NULL OR business_date = sqlc.narg(business_date)::date)
  AND (sqlc.narg(from_date)::date IS NULL OR business_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR business_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateTask :one
INSERT INTO tasks (title, subject_type, subject_id, site_id, business_date, shift, assignee_user_id, assignee_role_id, due_at, priority, created_by, updated_by)
VALUES (sqlc.arg(title), sqlc.narg(subject_type), sqlc.narg(subject_id), sqlc.arg(site_id), sqlc.narg(business_date), sqlc.narg(shift), sqlc.narg(assignee_user_id), sqlc.narg(assignee_role_id), sqlc.narg(due_at), sqlc.narg(priority), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateTask :one
UPDATE tasks
SET title = COALESCE(sqlc.narg(title), title), subject_type = COALESCE(sqlc.narg(subject_type), subject_type), subject_id = COALESCE(sqlc.narg(subject_id), subject_id), site_id = COALESCE(sqlc.narg(site_id), site_id), business_date = COALESCE(sqlc.narg(business_date), business_date), shift = COALESCE(sqlc.narg(shift), shift), assignee_user_id = COALESCE(sqlc.narg(assignee_user_id), assignee_user_id), assignee_role_id = COALESCE(sqlc.narg(assignee_role_id), assignee_role_id), due_at = COALESCE(sqlc.narg(due_at), due_at), priority = COALESCE(sqlc.narg(priority), priority), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetTaskStatus :one
UPDATE tasks
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActClaimTask :one
UPDATE tasks
SET status = 'in_progress', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('open')
RETURNING *;

-- name: ActCompleteTask :one
UPDATE tasks
SET completed_at = now(), status = 'done', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('open', 'in_progress')
RETURNING *;

-- name: ActSkipTask :one
UPDATE tasks
SET status = 'skipped', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('open', 'in_progress')
RETURNING *;

-- name: ActReassignTask :one
UPDATE tasks
SET updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
