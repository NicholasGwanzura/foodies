-- name: GetApprovalRequest :one
SELECT * FROM approval_requests WHERE id = $1;

-- name: ListApprovalRequest :many
SELECT * FROM approval_requests
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(workflow_key)::text IS NULL OR workflow_key = sqlc.narg(workflow_key)::text)
  AND (sqlc.narg(subject_type)::text IS NULL OR subject_type = sqlc.narg(subject_type)::text)
  AND (sqlc.narg(subject_id)::uuid IS NULL OR subject_id = sqlc.narg(subject_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: SetApprovalRequestStatus :one
UPDATE approval_requests
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListApprovalRequestSteps :many
SELECT * FROM approval_request_steps WHERE approval_request_id = $1 ORDER BY id;

-- name: InsertApprovalRequestSteps :one
INSERT INTO approval_request_steps (approval_request_id, sequence, role_id, decided_by, decision, decided_at, channel, "comment", signature_id, due_at, created_by)
VALUES (sqlc.arg(approval_request_id), sqlc.narg(sequence), sqlc.narg(role_id), sqlc.narg(decided_by), sqlc.narg(decision), sqlc.narg(decided_at), sqlc.narg(channel), sqlc.narg(comment), sqlc.narg(signature_id), sqlc.narg(due_at), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteApprovalRequestSteps :exec
DELETE FROM approval_request_steps WHERE approval_request_id = $1;

-- name: ActApproveApprovalRequest :one
UPDATE approval_requests
SET status = 'approved', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('pending')
RETURNING *;

-- name: ActRejectApprovalRequest :one
UPDATE approval_requests
SET status = 'rejected', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('pending', 'approved', 'returned')
RETURNING *;

-- name: ActReturnApprovalRequest :one
UPDATE approval_requests
SET status = 'returned', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('pending', 'approved')
RETURNING *;

-- name: ActReassignApprovalRequest :one
UPDATE approval_requests
SET updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
