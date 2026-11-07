-- name: GetWhatsappIdentity :one
SELECT * FROM whatsapp_identities WHERE id = $1;

-- name: ListWhatsappIdentity :many
SELECT * FROM whatsapp_identities
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(employee_id)::uuid IS NULL OR employee_id = sqlc.narg(employee_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateWhatsappIdentity :one
INSERT INTO whatsapp_identities (employee_id, phone_e164, created_by, updated_by)
VALUES (sqlc.arg(employee_id), sqlc.arg(phone_e164), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetWhatsappIdentityStatus :one
UPDATE whatsapp_identities
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActRevokeWhatsappIdentity :one
UPDATE whatsapp_identities
SET status = 'revoked', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('pending', 'verified')
RETURNING *;

-- name: ActRebindWhatsappIdentity :one
UPDATE whatsapp_identities
SET phone_e164 = COALESCE(sqlc.narg(phone_e164), phone_e164), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
