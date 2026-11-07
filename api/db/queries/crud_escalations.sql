-- name: GetEscalation :one
SELECT * FROM escalations WHERE id = $1;

-- name: ListEscalation :many
SELECT * FROM escalations
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(domain)::text IS NULL OR domain = sqlc.narg(domain)::text)
  AND (sqlc.narg(severity_level)::integer IS NULL OR severity_level = sqlc.narg(severity_level)::integer)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: SetEscalationStatus :one
UPDATE escalations
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListEscalationRecipients :many
SELECT * FROM escalation_recipients WHERE escalation_id = $1 ORDER BY id;

-- name: InsertEscalationRecipients :one
INSERT INTO escalation_recipients (escalation_id, step, user_id, channel, notified_at, acknowledged_at, created_by)
VALUES (sqlc.arg(escalation_id), sqlc.narg(step), sqlc.narg(user_id), sqlc.narg(channel), sqlc.narg(notified_at), sqlc.narg(acknowledged_at), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteEscalationRecipients :exec
DELETE FROM escalation_recipients WHERE escalation_id = $1;

-- name: ActAcknowledgeEscalation :one
UPDATE escalations
SET status = 'acknowledged', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('open')
RETURNING *;

-- name: ActAdvanceEscalation :one
UPDATE escalations
SET updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
