-- name: GetPatrolRound :one
SELECT * FROM patrol_rounds WHERE id = $1;

-- name: ListPatrolRound :many
SELECT * FROM patrol_rounds
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(patrol_route_id)::uuid IS NULL OR patrol_route_id = sqlc.narg(patrol_route_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreatePatrolRound :one
INSERT INTO patrol_rounds (reference, patrol_route_id, officer_employee_id, started_at, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(patrol_route_id), sqlc.arg(officer_employee_id), sqlc.arg(started_at), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdatePatrolRound :one
UPDATE patrol_rounds
SET patrol_route_id = COALESCE(sqlc.narg(patrol_route_id), patrol_route_id), officer_employee_id = COALESCE(sqlc.narg(officer_employee_id), officer_employee_id), started_at = COALESCE(sqlc.narg(started_at), started_at), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetPatrolRoundStatus :one
UPDATE patrol_rounds
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListPatrolVisits :many
SELECT * FROM patrol_visits WHERE patrol_round_id = $1 ORDER BY id;

-- name: InsertPatrolVisits :one
INSERT INTO patrol_visits (patrol_round_id, checkpoint_id, visited_at, observation, attachment_id, created_by)
VALUES (sqlc.arg(patrol_round_id), sqlc.arg(checkpoint_id), sqlc.arg(visited_at), sqlc.narg(observation), sqlc.narg(attachment_id), sqlc.narg(created_by))
RETURNING *;

-- name: DeletePatrolVisits :exec
DELETE FROM patrol_visits WHERE patrol_round_id = $1;

-- name: ActCompletePatrolRound :one
UPDATE patrol_rounds
SET status = 'completed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('in_progress')
RETURNING *;
