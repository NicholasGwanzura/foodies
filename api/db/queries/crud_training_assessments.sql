-- name: GetTrainingAssessment :one
SELECT * FROM training_assessments WHERE id = $1;

-- name: ListTrainingAssessment :many
SELECT * FROM training_assessments
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(employee_id)::uuid IS NULL OR employee_id = sqlc.narg(employee_id)::uuid)
  AND (sqlc.narg(course_id)::uuid IS NULL OR course_id = sqlc.narg(course_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR assessed_on::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR assessed_on::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateTrainingAssessment :one
INSERT INTO training_assessments (employee_id, course_id, training_session_id, kind, score, attempt, "result", assessor_employee_id, assessed_on, evidence_attachment_id, created_by, updated_by)
VALUES (sqlc.arg(employee_id), sqlc.arg(course_id), sqlc.narg(training_session_id), sqlc.arg(kind), sqlc.narg(score), sqlc.narg(attempt), sqlc.arg(result), sqlc.narg(assessor_employee_id), sqlc.arg(assessed_on), sqlc.narg(evidence_attachment_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;
