-- name: GetWhatsappMessage :one
SELECT * FROM whatsapp_messages WHERE id = $1;

-- name: ListWhatsappMessage :many
SELECT * FROM whatsapp_messages
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(identity_id)::uuid IS NULL OR identity_id = sqlc.narg(identity_id)::uuid)
  AND (sqlc.narg(direction)::text IS NULL OR direction = sqlc.narg(direction)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);
