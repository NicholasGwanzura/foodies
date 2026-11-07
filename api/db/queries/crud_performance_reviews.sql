-- name: GetPerformanceReview :one
SELECT * FROM performance_reviews WHERE id = $1;

-- name: ListPerformanceReview :many
SELECT * FROM performance_reviews
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreatePerformanceReview :one
INSERT INTO performance_reviews (employee_id, reviewer_employee_id, period, objectives, rating, comments, improvement_plan, created_by, updated_by)
VALUES (sqlc.arg(employee_id), sqlc.arg(reviewer_employee_id), sqlc.arg(period), sqlc.narg(objectives), sqlc.narg(rating), sqlc.narg(comments), sqlc.narg(improvement_plan), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdatePerformanceReview :one
UPDATE performance_reviews
SET employee_id = COALESCE(sqlc.narg(employee_id), employee_id), reviewer_employee_id = COALESCE(sqlc.narg(reviewer_employee_id), reviewer_employee_id), period = COALESCE(sqlc.narg(period), period), objectives = COALESCE(sqlc.narg(objectives), objectives), rating = COALESCE(sqlc.narg(rating), rating), comments = COALESCE(sqlc.narg(comments), comments), improvement_plan = COALESCE(sqlc.narg(improvement_plan), improvement_plan), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetPerformanceReviewStatus :one
UPDATE performance_reviews
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActCompletePerformanceReview :one
UPDATE performance_reviews
SET status = 'completed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'in_review')
RETURNING *;
