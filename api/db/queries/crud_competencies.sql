-- name: GetCompetency :one
SELECT * FROM competencies WHERE id = $1;

-- name: ListCompetency :many
SELECT * FROM competencies
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(employee_id)::uuid IS NULL OR employee_id = sqlc.narg(employee_id)::uuid)
  AND (sqlc.narg(course_id)::uuid IS NULL OR course_id = sqlc.narg(course_id)::uuid)
  AND (sqlc.narg(status)::text IS NULL OR "status" = sqlc.narg(status)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR achieved_on::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR achieved_on::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateCompetency :one
INSERT INTO competencies (employee_id, course_id, "status", achieved_on, expires_on, assessor_employee_id, restriction, created_by, updated_by)
VALUES (sqlc.arg(employee_id), sqlc.arg(course_id), sqlc.arg(status), sqlc.narg(achieved_on), sqlc.narg(expires_on), sqlc.narg(assessor_employee_id), sqlc.narg(restriction), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateCompetency :one
UPDATE competencies
SET employee_id = COALESCE(sqlc.narg(employee_id), employee_id), course_id = COALESCE(sqlc.narg(course_id), course_id), achieved_on = COALESCE(sqlc.narg(achieved_on), achieved_on), expires_on = COALESCE(sqlc.narg(expires_on), expires_on), assessor_employee_id = COALESCE(sqlc.narg(assessor_employee_id), assessor_employee_id), restriction = COALESCE(sqlc.narg(restriction), restriction), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
