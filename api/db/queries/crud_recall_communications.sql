-- name: GetRecallCommunication :one
SELECT * FROM recall_communications WHERE id = $1;

-- name: ListRecallCommunication :many
SELECT * FROM recall_communications
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateRecallCommunication :one
INSERT INTO recall_communications (recall_id, audience, recipient, message, channel, created_by, updated_by)
VALUES (sqlc.arg(recall_id), sqlc.arg(audience), sqlc.arg(recipient), sqlc.arg(message), sqlc.arg(channel), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetRecallCommunicationStatus :one
UPDATE recall_communications
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitRecallCommunication :one
UPDATE recall_communications
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActMarkSentRecallCommunication :one
UPDATE recall_communications
SET status = 'sent', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'pending_approval')
RETURNING *;

-- name: ActAcknowledgeRecallCommunication :one
UPDATE recall_communications
SET acknowledged_at = now(), status = 'acknowledged', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'pending_approval', 'sent')
RETURNING *;
