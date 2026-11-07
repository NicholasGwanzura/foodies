-- name: GetEmergencyContact :one
SELECT * FROM emergency_contacts WHERE id = $1;

-- name: ListEmergencyContact :many
SELECT * FROM emergency_contacts
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(from_date)::date IS NULL OR last_verified_on::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR last_verified_on::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateEmergencyContact :one
INSERT INTO emergency_contacts (employee_id, contact_name, relationship, phone_e164, alternate_phone, last_verified_on, created_by, updated_by)
VALUES (sqlc.arg(employee_id), sqlc.arg(contact_name), sqlc.arg(relationship), sqlc.arg(phone_e164), sqlc.narg(alternate_phone), sqlc.narg(last_verified_on), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateEmergencyContact :one
UPDATE emergency_contacts
SET contact_name = COALESCE(sqlc.narg(contact_name), contact_name), relationship = COALESCE(sqlc.narg(relationship), relationship), phone_e164 = COALESCE(sqlc.narg(phone_e164), phone_e164), alternate_phone = COALESCE(sqlc.narg(alternate_phone), alternate_phone), last_verified_on = COALESCE(sqlc.narg(last_verified_on), last_verified_on), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: DeleteEmergencyContact :exec
DELETE FROM emergency_contacts WHERE id = $1;
