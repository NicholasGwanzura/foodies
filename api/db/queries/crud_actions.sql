-- name: GetAction :one
SELECT * FROM actions WHERE id = $1;

-- name: ListAction :many
SELECT * FROM actions
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%' OR description ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(source_type)::text IS NULL OR source_type = sqlc.narg(source_type)::text)
  AND (sqlc.narg(source_id)::uuid IS NULL OR source_id = sqlc.narg(source_id)::uuid)
  AND (sqlc.narg(owner_employee_id)::uuid IS NULL OR owner_employee_id = sqlc.narg(owner_employee_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR due_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR due_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateAction :one
INSERT INTO actions (reference, source_type, source_id, kind, description, root_cause, owner_employee_id, site_id, due_date, priority, percent_complete, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.narg(source_type), sqlc.narg(source_id), sqlc.arg(kind), sqlc.arg(description), sqlc.narg(root_cause), sqlc.arg(owner_employee_id), sqlc.narg(site_id), sqlc.arg(due_date), sqlc.narg(priority), sqlc.narg(percent_complete), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateAction :one
UPDATE actions
SET source_type = COALESCE(sqlc.narg(source_type), source_type), source_id = COALESCE(sqlc.narg(source_id), source_id), kind = COALESCE(sqlc.narg(kind), kind), description = COALESCE(sqlc.narg(description), description), root_cause = COALESCE(sqlc.narg(root_cause), root_cause), owner_employee_id = COALESCE(sqlc.narg(owner_employee_id), owner_employee_id), site_id = COALESCE(sqlc.narg(site_id), site_id), due_date = COALESCE(sqlc.narg(due_date), due_date), priority = COALESCE(sqlc.narg(priority), priority), percent_complete = COALESCE(sqlc.narg(percent_complete), percent_complete), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetActionStatus :one
UPDATE actions
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActCompleteAction :one
UPDATE actions
SET status = 'pending_verification', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('open', 'in_progress')
RETURNING *;

-- name: ActVerifyAction :one
UPDATE actions
SET effectiveness_result = COALESCE(sqlc.narg(effectiveness_result), effectiveness_result), verified_by = sqlc.narg(updated_by), status = 'closed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('pending_verification')
RETURNING *;

-- name: ActReopenAction :one
UPDATE actions
SET status = 'open', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('pending_verification', 'closed')
RETURNING *;
