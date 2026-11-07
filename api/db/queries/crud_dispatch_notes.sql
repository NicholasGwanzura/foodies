-- name: GetDispatchNote :one
SELECT * FROM dispatch_notes WHERE id = $1;

-- name: ListDispatchNote :many
SELECT * FROM dispatch_notes
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(to_site_id)::uuid IS NULL OR to_site_id = sqlc.narg(to_site_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateDispatchNote :one
INSERT INTO dispatch_notes (reference, from_site_id, to_site_id, kitchen_order_id, vehicle, driver_employee_id, loading_temperature_c, delivery_temperature_c, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(from_site_id), sqlc.arg(to_site_id), sqlc.narg(kitchen_order_id), sqlc.arg(vehicle), sqlc.arg(driver_employee_id), sqlc.narg(loading_temperature_c), sqlc.narg(delivery_temperature_c), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateDispatchNote :one
UPDATE dispatch_notes
SET from_site_id = COALESCE(sqlc.narg(from_site_id), from_site_id), to_site_id = COALESCE(sqlc.narg(to_site_id), to_site_id), kitchen_order_id = COALESCE(sqlc.narg(kitchen_order_id), kitchen_order_id), vehicle = COALESCE(sqlc.narg(vehicle), vehicle), driver_employee_id = COALESCE(sqlc.narg(driver_employee_id), driver_employee_id), loading_temperature_c = COALESCE(sqlc.narg(loading_temperature_c), loading_temperature_c), delivery_temperature_c = COALESCE(sqlc.narg(delivery_temperature_c), delivery_temperature_c), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetDispatchNoteStatus :one
UPDATE dispatch_notes
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListDispatchNoteLines :many
SELECT * FROM dispatch_note_lines WHERE dispatch_note_id = $1 ORDER BY id;

-- name: InsertDispatchNoteLines :one
INSERT INTO dispatch_note_lines (dispatch_note_id, product_id, batch_id, ordered_quantity, picked_quantity, delivered_quantity, expiry_date, condition, raw_or_rte, created_by)
VALUES (sqlc.arg(dispatch_note_id), sqlc.arg(product_id), sqlc.arg(batch_id), sqlc.narg(ordered_quantity), sqlc.arg(picked_quantity), sqlc.narg(delivered_quantity), sqlc.narg(expiry_date), sqlc.narg(condition), sqlc.narg(raw_or_rte), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteDispatchNoteLines :exec
DELETE FROM dispatch_note_lines WHERE dispatch_note_id = $1;

-- name: ActVerifyDispatchNote :one
UPDATE dispatch_notes
SET status = 'verified', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'picked')
RETURNING *;

-- name: ActSecurityClearDispatchNote :one
UPDATE dispatch_notes
SET status = 'security_cleared', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('verified')
RETURNING *;

-- name: ActConfirmDeliveryDispatchNote :one
UPDATE dispatch_notes
SET status = 'delivered', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('security_cleared', 'in_transit')
RETURNING *;

-- name: ActRaiseDiscrepancyDispatchNote :one
UPDATE dispatch_notes
SET status = 'discrepancy', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('security_cleared', 'in_transit', 'delivered')
RETURNING *;
