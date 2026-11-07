-- name: GetUniformIssue :one
SELECT * FROM uniform_issues WHERE id = $1;

-- name: ListUniformIssue :many
SELECT * FROM uniform_issues
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(employee_id)::uuid IS NULL OR employee_id = sqlc.narg(employee_id)::uuid)
  AND (sqlc.narg(kind)::text IS NULL OR kind = sqlc.narg(kind)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR issued_on::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR issued_on::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateUniformIssue :one
INSERT INTO uniform_issues (reference, employee_id, kind, items, sizes, issued_on, "reason", old_pair_returned, site_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(employee_id), sqlc.arg(kind), sqlc.arg(items), sqlc.narg(sizes), sqlc.arg(issued_on), sqlc.narg(reason), COALESCE(sqlc.narg(old_pair_returned)::boolean, false), sqlc.narg(site_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateUniformIssue :one
UPDATE uniform_issues
SET employee_id = COALESCE(sqlc.narg(employee_id), employee_id), kind = COALESCE(sqlc.narg(kind), kind), items = COALESCE(sqlc.narg(items), items), sizes = COALESCE(sqlc.narg(sizes), sizes), issued_on = COALESCE(sqlc.narg(issued_on), issued_on), "reason" = COALESCE(sqlc.narg(reason), "reason"), old_pair_returned = COALESCE(sqlc.narg(old_pair_returned), old_pair_returned), site_id = COALESCE(sqlc.narg(site_id), site_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetUniformIssueStatus :one
UPDATE uniform_issues
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
