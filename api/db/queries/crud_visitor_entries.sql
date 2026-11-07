-- name: GetVisitorEntry :one
SELECT * FROM visitor_entries WHERE id = $1;

-- name: ListVisitorEntry :many
SELECT * FROM visitor_entries
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(kind)::text IS NULL OR kind = sqlc.narg(kind)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateVisitorEntry :one
INSERT INTO visitor_entries (site_id, kind, full_name, organisation, contact, id_reference, purpose, host_employee_id, work_order_id, work_area, induction_done, badge, arrived_at, created_by, updated_by)
VALUES (sqlc.arg(site_id), sqlc.arg(kind), sqlc.arg(full_name), sqlc.narg(organisation), sqlc.narg(contact), sqlc.narg(id_reference), sqlc.arg(purpose), sqlc.narg(host_employee_id), sqlc.narg(work_order_id), sqlc.narg(work_area), COALESCE(sqlc.narg(induction_done)::boolean, false), sqlc.narg(badge), sqlc.arg(arrived_at), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateVisitorEntry :one
UPDATE visitor_entries
SET site_id = COALESCE(sqlc.narg(site_id), site_id), kind = COALESCE(sqlc.narg(kind), kind), full_name = COALESCE(sqlc.narg(full_name), full_name), organisation = COALESCE(sqlc.narg(organisation), organisation), contact = COALESCE(sqlc.narg(contact), contact), id_reference = COALESCE(sqlc.narg(id_reference), id_reference), purpose = COALESCE(sqlc.narg(purpose), purpose), host_employee_id = COALESCE(sqlc.narg(host_employee_id), host_employee_id), work_order_id = COALESCE(sqlc.narg(work_order_id), work_order_id), work_area = COALESCE(sqlc.narg(work_area), work_area), induction_done = COALESCE(sqlc.narg(induction_done), induction_done), badge = COALESCE(sqlc.narg(badge), badge), arrived_at = COALESCE(sqlc.narg(arrived_at), arrived_at), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetVisitorEntryStatus :one
UPDATE visitor_entries
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSignOutVisitorEntry :one
UPDATE visitor_entries
SET status = 'departed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('on_site')
RETURNING *;
