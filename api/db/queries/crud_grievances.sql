-- name: GetGrievance :one
SELECT * FROM grievances WHERE id = $1;

-- name: ListGrievance :many
SELECT * FROM grievances
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateGrievance :one
INSERT INTO grievances (reference, employee_id, subject, investigator_employee_id, outcome, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(employee_id), sqlc.arg(subject), sqlc.narg(investigator_employee_id), sqlc.narg(outcome), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateGrievance :one
UPDATE grievances
SET employee_id = COALESCE(sqlc.narg(employee_id), employee_id), subject = COALESCE(sqlc.narg(subject), subject), investigator_employee_id = COALESCE(sqlc.narg(investigator_employee_id), investigator_employee_id), outcome = COALESCE(sqlc.narg(outcome), outcome), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetGrievanceStatus :one
UPDATE grievances
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
