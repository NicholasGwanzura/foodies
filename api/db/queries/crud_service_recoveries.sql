-- name: GetServiceRecovery :one
SELECT * FROM service_recoveries WHERE id = $1;

-- name: ListServiceRecovery :many
SELECT * FROM service_recoveries
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateServiceRecovery :one
INSERT INTO service_recoveries (reference, complaint_id, remedy_type, value_amount, value_currency, "reason", customer_agreed, pos_reference, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(complaint_id), sqlc.arg(remedy_type), sqlc.narg(value_amount), sqlc.narg(value_currency), sqlc.arg(reason), COALESCE(sqlc.narg(customer_agreed)::boolean, false), sqlc.narg(pos_reference), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetServiceRecoveryStatus :one
UPDATE service_recoveries
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitServiceRecovery :one
UPDATE service_recoveries
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActExecuteServiceRecovery :one
UPDATE service_recoveries
SET pos_reference = COALESCE(sqlc.narg(pos_reference), pos_reference), status = 'executed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'pending_approval', 'approved')
RETURNING *;
