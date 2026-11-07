-- name: GetHold :one
SELECT * FROM holds WHERE id = $1;

-- name: ListHold :many
SELECT * FROM holds
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(batch_id)::uuid IS NULL OR batch_id = sqlc.narg(batch_id)::uuid)
  AND (sqlc.narg(product_id)::uuid IS NULL OR product_id = sqlc.narg(product_id)::uuid)
  AND (sqlc.narg(item_id)::uuid IS NULL OR item_id = sqlc.narg(item_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateHold :one
INSERT INTO holds (reference, site_id, product_id, item_id, batch_id, batch_lot, quantity, uom, "reason", quarantine_location_id, owner_employee_id, ncr_id, recall_id, label_confirmed, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.narg(product_id), sqlc.narg(item_id), sqlc.narg(batch_id), sqlc.narg(batch_lot), sqlc.arg(quantity), sqlc.arg(uom), sqlc.arg(reason), sqlc.narg(quarantine_location_id), sqlc.arg(owner_employee_id), sqlc.narg(ncr_id), sqlc.narg(recall_id), COALESCE(sqlc.narg(label_confirmed)::boolean, false), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateHold :one
UPDATE holds
SET site_id = COALESCE(sqlc.narg(site_id), site_id), product_id = COALESCE(sqlc.narg(product_id), product_id), item_id = COALESCE(sqlc.narg(item_id), item_id), batch_id = COALESCE(sqlc.narg(batch_id), batch_id), batch_lot = COALESCE(sqlc.narg(batch_lot), batch_lot), quantity = COALESCE(sqlc.narg(quantity), quantity), uom = COALESCE(sqlc.narg(uom), uom), "reason" = COALESCE(sqlc.narg(reason), "reason"), quarantine_location_id = COALESCE(sqlc.narg(quarantine_location_id), quarantine_location_id), owner_employee_id = COALESCE(sqlc.narg(owner_employee_id), owner_employee_id), ncr_id = COALESCE(sqlc.narg(ncr_id), ncr_id), recall_id = COALESCE(sqlc.narg(recall_id), recall_id), label_confirmed = COALESCE(sqlc.narg(label_confirmed), label_confirmed), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetHoldStatus :one
UPDATE holds
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActReleaseHold :one
UPDATE holds
SET released_at = now(), status = 'released', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('on_hold', 'pending_disposition')
RETURNING *;
