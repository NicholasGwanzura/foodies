-- name: GetApprovalPolicy :one
SELECT * FROM approval_policies WHERE id = $1;

-- name: ListApprovalPolicy :many
SELECT * FROM approval_policies
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR "name" ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(workflow_key)::text IS NULL OR workflow_key = sqlc.narg(workflow_key)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateApprovalPolicy :one
INSERT INTO approval_policies (workflow_key, "name", effective_from, source_document_id, created_by, updated_by)
VALUES (sqlc.arg(workflow_key), sqlc.arg(name), sqlc.narg(effective_from), sqlc.narg(source_document_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateApprovalPolicy :one
UPDATE approval_policies
SET workflow_key = COALESCE(sqlc.narg(workflow_key), workflow_key), "name" = COALESCE(sqlc.narg(name), "name"), effective_from = COALESCE(sqlc.narg(effective_from), effective_from), source_document_id = COALESCE(sqlc.narg(source_document_id), source_document_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetApprovalPolicyStatus :one
UPDATE approval_policies
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListApprovalPolicySteps :many
SELECT * FROM approval_policy_steps WHERE approval_policy_id = $1 ORDER BY id;

-- name: InsertApprovalPolicySteps :one
INSERT INTO approval_policy_steps (approval_policy_id, sequence, role_id, mode, condition, sla_minutes, escalate_to_role_id, created_by)
VALUES (sqlc.arg(approval_policy_id), sqlc.arg(sequence), sqlc.arg(role_id), sqlc.arg(mode), sqlc.narg(condition), sqlc.narg(sla_minutes), sqlc.narg(escalate_to_role_id), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteApprovalPolicySteps :exec
DELETE FROM approval_policy_steps WHERE approval_policy_id = $1;

-- name: ActSubmitApprovalPolicy :one
UPDATE approval_policies
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
