-- name: GetEscalationPolicy :one
SELECT * FROM escalation_policies WHERE id = $1;

-- name: ListEscalationPolicy :many
SELECT * FROM escalation_policies
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR label ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(domain)::text IS NULL OR domain = sqlc.narg(domain)::text)
  AND (sqlc.narg(severity_level)::integer IS NULL OR severity_level = sqlc.narg(severity_level)::integer)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateEscalationPolicy :one
INSERT INTO escalation_policies (domain, severity_level, label, after_hours, created_by, updated_by)
VALUES (sqlc.arg(domain), sqlc.arg(severity_level), sqlc.arg(label), COALESCE(sqlc.narg(after_hours)::boolean, false), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateEscalationPolicy :one
UPDATE escalation_policies
SET domain = COALESCE(sqlc.narg(domain), domain), severity_level = COALESCE(sqlc.narg(severity_level), severity_level), label = COALESCE(sqlc.narg(label), label), after_hours = COALESCE(sqlc.narg(after_hours), after_hours), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetEscalationPolicyStatus :one
UPDATE escalation_policies
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListEscalationPolicySteps :many
SELECT * FROM escalation_policy_steps WHERE escalation_policy_id = $1 ORDER BY id;

-- name: InsertEscalationPolicySteps :one
INSERT INTO escalation_policy_steps (escalation_policy_id, sequence, notify_role_id, channels, ack_within_minutes, next_on_timeout, created_by)
VALUES (sqlc.arg(escalation_policy_id), sqlc.arg(sequence), sqlc.arg(notify_role_id), sqlc.arg(channels), sqlc.narg(ack_within_minutes), COALESCE(sqlc.narg(next_on_timeout)::boolean, false), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteEscalationPolicySteps :exec
DELETE FROM escalation_policy_steps WHERE escalation_policy_id = $1;
