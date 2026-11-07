-- name: GetExternalVisit :one
SELECT * FROM external_visits WHERE id = $1;

-- name: ListExternalVisit :many
SELECT * FROM external_visits
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR requested_on::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR requested_on::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateExternalVisit :one
INSERT INTO external_visits (reference, organisation, group_leader, contact, site_id, requested_on, attendee_count, includes_minors, special_requirements, risk_notes, disclaimer_attachment_id, nda_required, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(organisation), sqlc.arg(group_leader), sqlc.narg(contact), sqlc.arg(site_id), sqlc.arg(requested_on), sqlc.arg(attendee_count), COALESCE(sqlc.narg(includes_minors)::boolean, false), sqlc.narg(special_requirements), sqlc.narg(risk_notes), sqlc.narg(disclaimer_attachment_id), COALESCE(sqlc.narg(nda_required)::boolean, false), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateExternalVisit :one
UPDATE external_visits
SET organisation = COALESCE(sqlc.narg(organisation), organisation), group_leader = COALESCE(sqlc.narg(group_leader), group_leader), contact = COALESCE(sqlc.narg(contact), contact), site_id = COALESCE(sqlc.narg(site_id), site_id), requested_on = COALESCE(sqlc.narg(requested_on), requested_on), attendee_count = COALESCE(sqlc.narg(attendee_count), attendee_count), includes_minors = COALESCE(sqlc.narg(includes_minors), includes_minors), special_requirements = COALESCE(sqlc.narg(special_requirements), special_requirements), risk_notes = COALESCE(sqlc.narg(risk_notes), risk_notes), disclaimer_attachment_id = COALESCE(sqlc.narg(disclaimer_attachment_id), disclaimer_attachment_id), nda_required = COALESCE(sqlc.narg(nda_required), nda_required), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetExternalVisitStatus :one
UPDATE external_visits
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitExternalVisit :one
UPDATE external_visits
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('requested')
RETURNING *;

-- name: ActCloseExternalVisit :one
UPDATE external_visits
SET closed_at = now(), status = 'closed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('requested', 'screening', 'pending_approval', 'approved', 'held')
RETURNING *;
