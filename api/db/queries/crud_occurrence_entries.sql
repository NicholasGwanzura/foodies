-- name: GetOccurrenceEntry :one
SELECT * FROM occurrence_entries WHERE id = $1;

-- name: ListOccurrenceEntry :many
SELECT * FROM occurrence_entries
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%' OR description ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(open_item)::boolean IS NULL OR open_item = sqlc.narg(open_item)::boolean)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateOccurrenceEntry :one
INSERT INTO occurrence_entries (reference, site_id, post, occurred_at, officer_employee_id, description, action_taken, persons_involved, open_item, incident_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.narg(post), sqlc.arg(occurred_at), sqlc.arg(officer_employee_id), sqlc.arg(description), sqlc.narg(action_taken), sqlc.narg(persons_involved), COALESCE(sqlc.narg(open_item)::boolean, false), sqlc.narg(incident_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;
