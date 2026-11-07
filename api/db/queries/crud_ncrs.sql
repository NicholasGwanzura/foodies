-- name: GetNcr :one
SELECT * FROM ncrs WHERE id = $1;

-- name: ListNcr :many
SELECT * FROM ncrs
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(severity_level)::integer IS NULL OR severity_level = sqlc.narg(severity_level)::integer)
  AND (sqlc.narg(product_id)::uuid IS NULL OR product_id = sqlc.narg(product_id)::uuid)
  AND (sqlc.narg(batch_id)::uuid IS NULL OR batch_id = sqlc.narg(batch_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateNcr :one
INSERT INTO ncrs (reference, site_id, occurred_at, product_id, item_id, batch_id, quantity, location, deviation, severity_level, containment, production_stopped, root_cause, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(occurred_at), sqlc.narg(product_id), sqlc.narg(item_id), sqlc.narg(batch_id), sqlc.narg(quantity), sqlc.narg(location), sqlc.arg(deviation), sqlc.arg(severity_level), sqlc.narg(containment), COALESCE(sqlc.narg(production_stopped)::boolean, false), sqlc.narg(root_cause), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateNcr :one
UPDATE ncrs
SET site_id = COALESCE(sqlc.narg(site_id), site_id), occurred_at = COALESCE(sqlc.narg(occurred_at), occurred_at), product_id = COALESCE(sqlc.narg(product_id), product_id), item_id = COALESCE(sqlc.narg(item_id), item_id), batch_id = COALESCE(sqlc.narg(batch_id), batch_id), quantity = COALESCE(sqlc.narg(quantity), quantity), location = COALESCE(sqlc.narg(location), location), deviation = COALESCE(sqlc.narg(deviation), deviation), severity_level = COALESCE(sqlc.narg(severity_level), severity_level), containment = COALESCE(sqlc.narg(containment), containment), production_stopped = COALESCE(sqlc.narg(production_stopped), production_stopped), root_cause = COALESCE(sqlc.narg(root_cause), root_cause), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetNcrStatus :one
UPDATE ncrs
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActContainNcr :one
UPDATE ncrs
SET containment = COALESCE(sqlc.narg(containment), containment), status = 'contained', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('open')
RETURNING *;

-- name: ActVerifyCloseNcr :one
UPDATE ncrs
SET status = 'closed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('pending_verification')
RETURNING *;
