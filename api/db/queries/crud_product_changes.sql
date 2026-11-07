-- name: GetProductChange :one
SELECT * FROM product_changes WHERE id = $1;

-- name: ListProductChange :many
SELECT * FROM product_changes
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%' OR description ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR effective_on::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR effective_on::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateProductChange :one
INSERT INTO product_changes (reference, product_id, change_type, description, effective_on, pos_change_reference, affected_role_ids, training_request_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(product_id), sqlc.arg(change_type), sqlc.arg(description), sqlc.arg(effective_on), sqlc.narg(pos_change_reference), sqlc.narg(affected_role_ids), sqlc.narg(training_request_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateProductChange :one
UPDATE product_changes
SET product_id = COALESCE(sqlc.narg(product_id), product_id), change_type = COALESCE(sqlc.narg(change_type), change_type), description = COALESCE(sqlc.narg(description), description), effective_on = COALESCE(sqlc.narg(effective_on), effective_on), pos_change_reference = COALESCE(sqlc.narg(pos_change_reference), pos_change_reference), affected_role_ids = COALESCE(sqlc.narg(affected_role_ids), affected_role_ids), training_request_id = COALESCE(sqlc.narg(training_request_id), training_request_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetProductChangeStatus :one
UPDATE product_changes
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitProductChange :one
UPDATE product_changes
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
