-- name: GetKeyAudit :one
SELECT * FROM key_audits WHERE id = $1;

-- name: ListKeyAudit :many
SELECT * FROM key_audits
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateKeyAudit :one
INSERT INTO key_audits (reference, site_id, kind, period, auditor_employee_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(kind), sqlc.arg(period), sqlc.arg(auditor_employee_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateKeyAudit :one
UPDATE key_audits
SET site_id = COALESCE(sqlc.narg(site_id), site_id), kind = COALESCE(sqlc.narg(kind), kind), period = COALESCE(sqlc.narg(period), period), auditor_employee_id = COALESCE(sqlc.narg(auditor_employee_id), auditor_employee_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetKeyAuditStatus :one
UPDATE key_audits
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListKeyAuditLines :many
SELECT * FROM key_audit_lines WHERE key_audit_id = $1 ORDER BY id;

-- name: InsertKeyAuditLines :one
INSERT INTO key_audit_lines (key_audit_id, key_item_id, expected_custodian_employee_id, verified, discrepancy, created_by)
VALUES (sqlc.arg(key_audit_id), sqlc.arg(key_item_id), sqlc.narg(expected_custodian_employee_id), COALESCE(sqlc.narg(verified)::boolean, false), sqlc.narg(discrepancy), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteKeyAuditLines :exec
DELETE FROM key_audit_lines WHERE key_audit_id = $1;

-- name: ActCompleteKeyAudit :one
UPDATE key_audits
SET status = 'completed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('planned', 'in_progress')
RETURNING *;
