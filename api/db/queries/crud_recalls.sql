-- name: GetRecall :one
SELECT * FROM recalls WHERE id = $1;

-- name: ListRecall :many
SELECT * FROM recalls
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%' OR description ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(kind)::text IS NULL OR kind = sqlc.narg(kind)::text)
  AND (sqlc.narg(product_id)::uuid IS NULL OR product_id = sqlc.narg(product_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateRecall :one
INSERT INTO recalls (reference, kind, trigger, description, product_id, batch_ids, site_id, reporter_employee_id, hazard, severity_rating, likelihood, customer_exposure, distribution_summary, regulatory_notification_required, effectiveness_summary, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(kind), sqlc.arg(trigger), sqlc.arg(description), sqlc.narg(product_id), sqlc.narg(batch_ids), sqlc.narg(site_id), sqlc.narg(reporter_employee_id), sqlc.narg(hazard), sqlc.narg(severity_rating), sqlc.narg(likelihood), sqlc.narg(customer_exposure), sqlc.narg(distribution_summary), COALESCE(sqlc.narg(regulatory_notification_required)::boolean, false), sqlc.narg(effectiveness_summary), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateRecall :one
UPDATE recalls
SET kind = COALESCE(sqlc.narg(kind), kind), trigger = COALESCE(sqlc.narg(trigger), trigger), description = COALESCE(sqlc.narg(description), description), product_id = COALESCE(sqlc.narg(product_id), product_id), batch_ids = COALESCE(sqlc.narg(batch_ids), batch_ids), site_id = COALESCE(sqlc.narg(site_id), site_id), reporter_employee_id = COALESCE(sqlc.narg(reporter_employee_id), reporter_employee_id), hazard = COALESCE(sqlc.narg(hazard), hazard), severity_rating = COALESCE(sqlc.narg(severity_rating), severity_rating), likelihood = COALESCE(sqlc.narg(likelihood), likelihood), customer_exposure = COALESCE(sqlc.narg(customer_exposure), customer_exposure), distribution_summary = COALESCE(sqlc.narg(distribution_summary), distribution_summary), regulatory_notification_required = COALESCE(sqlc.narg(regulatory_notification_required), regulatory_notification_required), effectiveness_summary = COALESCE(sqlc.narg(effectiveness_summary), effectiveness_summary), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetRecallStatus :one
UPDATE recalls
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActAssessRecall :one
UPDATE recalls
SET hazard = COALESCE(sqlc.narg(hazard), hazard), severity_rating = COALESCE(sqlc.narg(severity_rating), severity_rating), likelihood = COALESCE(sqlc.narg(likelihood), likelihood), status = 'assessed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('held')
RETURNING *;

-- name: ActDecideRecall :one
UPDATE recalls
SET decision = COALESCE(sqlc.narg(decision), decision), decided_at = now(), status = 'decided', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('initiated', 'held', 'assessed')
RETURNING *;

-- name: ActCloseRecall :one
UPDATE recalls
SET effectiveness_summary = COALESCE(sqlc.narg(effectiveness_summary), effectiveness_summary), closed_at = now(), status = 'closed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('initiated', 'held', 'assessed', 'decided', 'retrieving', 'reconciling')
RETURNING *;
