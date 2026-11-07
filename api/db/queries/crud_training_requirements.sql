-- name: GetTrainingRequirement :one
SELECT * FROM training_requirements WHERE id = $1;

-- name: ListTrainingRequirement :many
SELECT * FROM training_requirements
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateTrainingRequirement :one
INSERT INTO training_requirements (role_id, course_id, department_id, mandatory, refresher_months, created_by, updated_by)
VALUES (sqlc.arg(role_id), sqlc.arg(course_id), sqlc.narg(department_id), COALESCE(sqlc.narg(mandatory)::boolean, false), sqlc.narg(refresher_months), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateTrainingRequirement :one
UPDATE training_requirements
SET role_id = COALESCE(sqlc.narg(role_id), role_id), course_id = COALESCE(sqlc.narg(course_id), course_id), department_id = COALESCE(sqlc.narg(department_id), department_id), mandatory = COALESCE(sqlc.narg(mandatory), mandatory), refresher_months = COALESCE(sqlc.narg(refresher_months), refresher_months), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: DeleteTrainingRequirement :exec
DELETE FROM training_requirements WHERE id = $1;
