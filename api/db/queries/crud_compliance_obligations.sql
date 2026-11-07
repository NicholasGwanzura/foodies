-- name: GetComplianceObligation :one
SELECT * FROM compliance_obligations WHERE id = $1;

-- name: ListComplianceObligation :many
SELECT * FROM compliance_obligations
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%' OR title ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(kind)::text IS NULL OR kind = sqlc.narg(kind)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR due_on::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR due_on::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateComplianceObligation :one
INSERT INTO compliance_obligations (reference, kind, title, authority, site_id, subject_type, subject_id, owner_employee_id, due_on, recurrence_rule_key, reminder_days, evidence_attachment_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(kind), sqlc.arg(title), sqlc.narg(authority), sqlc.narg(site_id), sqlc.narg(subject_type), sqlc.narg(subject_id), sqlc.narg(owner_employee_id), sqlc.arg(due_on), sqlc.narg(recurrence_rule_key), sqlc.narg(reminder_days), sqlc.narg(evidence_attachment_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateComplianceObligation :one
UPDATE compliance_obligations
SET kind = COALESCE(sqlc.narg(kind), kind), title = COALESCE(sqlc.narg(title), title), authority = COALESCE(sqlc.narg(authority), authority), site_id = COALESCE(sqlc.narg(site_id), site_id), subject_type = COALESCE(sqlc.narg(subject_type), subject_type), subject_id = COALESCE(sqlc.narg(subject_id), subject_id), owner_employee_id = COALESCE(sqlc.narg(owner_employee_id), owner_employee_id), due_on = COALESCE(sqlc.narg(due_on), due_on), recurrence_rule_key = COALESCE(sqlc.narg(recurrence_rule_key), recurrence_rule_key), reminder_days = COALESCE(sqlc.narg(reminder_days), reminder_days), evidence_attachment_id = COALESCE(sqlc.narg(evidence_attachment_id), evidence_attachment_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetComplianceObligationStatus :one
UPDATE compliance_obligations
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActCompleteComplianceObligation :one
UPDATE compliance_obligations
SET evidence_attachment_id = COALESCE(sqlc.narg(evidence_attachment_id), evidence_attachment_id), completed_at = now(), status = 'completed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('upcoming', 'due', 'overdue')
RETURNING *;
