-- name: GetIncident :one
SELECT * FROM incidents WHERE id = $1;

-- name: ListIncident :many
SELECT * FROM incidents
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%' OR description ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(incident_type)::text IS NULL OR incident_type = sqlc.narg(incident_type)::text)
  AND (sqlc.narg(severity_level)::integer IS NULL OR severity_level = sqlc.narg(severity_level)::integer)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateIncident :one
INSERT INTO incidents (reference, site_id, incident_type, severity_level, occurred_at, discovered_at, description, persons_involved, immediate_action, details, reporter_employee_id, statutory_notification_required, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(incident_type), sqlc.arg(severity_level), sqlc.arg(occurred_at), sqlc.arg(discovered_at), sqlc.arg(description), sqlc.narg(persons_involved), sqlc.narg(immediate_action), sqlc.narg(details), sqlc.narg(reporter_employee_id), COALESCE(sqlc.narg(statutory_notification_required)::boolean, false), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateIncident :one
UPDATE incidents
SET site_id = COALESCE(sqlc.narg(site_id), site_id), incident_type = COALESCE(sqlc.narg(incident_type), incident_type), severity_level = COALESCE(sqlc.narg(severity_level), severity_level), occurred_at = COALESCE(sqlc.narg(occurred_at), occurred_at), discovered_at = COALESCE(sqlc.narg(discovered_at), discovered_at), description = COALESCE(sqlc.narg(description), description), persons_involved = COALESCE(sqlc.narg(persons_involved), persons_involved), immediate_action = COALESCE(sqlc.narg(immediate_action), immediate_action), details = COALESCE(sqlc.narg(details), details), reporter_employee_id = COALESCE(sqlc.narg(reporter_employee_id), reporter_employee_id), statutory_notification_required = COALESCE(sqlc.narg(statutory_notification_required), statutory_notification_required), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetIncidentStatus :one
UPDATE incidents
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActCloseIncident :one
UPDATE incidents
SET closed_at = now(), status = 'closed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('reported', 'acknowledged', 'investigating', 'action_pending')
RETURNING *;
