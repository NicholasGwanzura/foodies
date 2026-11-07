-- name: GetWorkflowActivation :one
SELECT * FROM workflow_activations WHERE id = $1;

-- name: ListWorkflowActivation :many
SELECT * FROM workflow_activations
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(workflow_key)::text IS NULL OR workflow_key = sqlc.narg(workflow_key)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: SetWorkflowActivationStatus :one
UPDATE workflow_activations
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActActivateWorkflowActivation :one
UPDATE workflow_activations
SET activated_at = now(), activated_by = sqlc.narg(updated_by), status = 'live', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('not_ready', 'ready')
RETURNING *;

-- name: ActSuspendWorkflowActivation :one
UPDATE workflow_activations
SET status = 'suspended', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('not_ready', 'ready', 'live')
RETURNING *;
