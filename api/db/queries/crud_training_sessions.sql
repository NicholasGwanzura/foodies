-- name: GetTrainingSession :one
SELECT * FROM training_sessions WHERE id = $1;

-- name: ListTrainingSession :many
SELECT * FROM training_sessions
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(course_id)::uuid IS NULL OR course_id = sqlc.narg(course_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateTrainingSession :one
INSERT INTO training_sessions (reference, course_id, site_id, venue, starts_at, ends_at, shift, trainer_employee_id, external_trainer, capacity, cost_amount, cost_currency, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(course_id), sqlc.narg(site_id), sqlc.narg(venue), sqlc.arg(starts_at), sqlc.narg(ends_at), sqlc.narg(shift), sqlc.narg(trainer_employee_id), sqlc.narg(external_trainer), sqlc.narg(capacity), sqlc.narg(cost_amount), sqlc.narg(cost_currency), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateTrainingSession :one
UPDATE training_sessions
SET course_id = COALESCE(sqlc.narg(course_id), course_id), site_id = COALESCE(sqlc.narg(site_id), site_id), venue = COALESCE(sqlc.narg(venue), venue), starts_at = COALESCE(sqlc.narg(starts_at), starts_at), ends_at = COALESCE(sqlc.narg(ends_at), ends_at), shift = COALESCE(sqlc.narg(shift), shift), trainer_employee_id = COALESCE(sqlc.narg(trainer_employee_id), trainer_employee_id), external_trainer = COALESCE(sqlc.narg(external_trainer), external_trainer), capacity = COALESCE(sqlc.narg(capacity), capacity), cost_amount = COALESCE(sqlc.narg(cost_amount), cost_amount), cost_currency = COALESCE(sqlc.narg(cost_currency), cost_currency), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetTrainingSessionStatus :one
UPDATE training_sessions
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListTrainingParticipants :many
SELECT * FROM training_participants WHERE training_session_id = $1 ORDER BY id;

-- name: InsertTrainingParticipants :one
INSERT INTO training_participants (training_session_id, employee_id, attendance, checked_in_at, created_by)
VALUES (sqlc.arg(training_session_id), sqlc.arg(employee_id), sqlc.narg(attendance), sqlc.narg(checked_in_at), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteTrainingParticipants :exec
DELETE FROM training_participants WHERE training_session_id = $1;

-- name: ActCloseTrainingSession :one
UPDATE training_sessions
SET status = 'held', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('planned')
RETURNING *;
