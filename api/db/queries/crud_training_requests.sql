-- name: GetTrainingRequest :one
SELECT * FROM training_requests WHERE id = $1;

-- name: ListTrainingRequest :many
SELECT * FROM training_requests
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR proposed_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR proposed_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateTrainingRequest :one
INSERT INTO training_requests (reference, site_id, topic, "reason", target_group, estimated_cost_amount, estimated_cost_currency, proposed_date, source_type, source_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.narg(site_id), sqlc.arg(topic), sqlc.arg(reason), sqlc.narg(target_group), sqlc.narg(estimated_cost_amount), sqlc.narg(estimated_cost_currency), sqlc.narg(proposed_date), sqlc.narg(source_type), sqlc.narg(source_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateTrainingRequest :one
UPDATE training_requests
SET site_id = COALESCE(sqlc.narg(site_id), site_id), topic = COALESCE(sqlc.narg(topic), topic), "reason" = COALESCE(sqlc.narg(reason), "reason"), target_group = COALESCE(sqlc.narg(target_group), target_group), estimated_cost_amount = COALESCE(sqlc.narg(estimated_cost_amount), estimated_cost_amount), estimated_cost_currency = COALESCE(sqlc.narg(estimated_cost_currency), estimated_cost_currency), proposed_date = COALESCE(sqlc.narg(proposed_date), proposed_date), source_type = COALESCE(sqlc.narg(source_type), source_type), source_id = COALESCE(sqlc.narg(source_id), source_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetTrainingRequestStatus :one
UPDATE training_requests
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitTrainingRequest :one
UPDATE training_requests
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
