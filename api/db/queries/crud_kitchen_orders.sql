-- name: GetKitchenOrder :one
SELECT * FROM kitchen_orders WHERE id = $1;

-- name: ListKitchenOrder :many
SELECT * FROM kitchen_orders
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(cycle)::text IS NULL OR cycle = sqlc.narg(cycle)::text)
  AND (sqlc.narg(order_date)::date IS NULL OR order_date = sqlc.narg(order_date)::date)
  AND (sqlc.narg(from_date)::date IS NULL OR order_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR order_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateKitchenOrder :one
INSERT INTO kitchen_orders (reference, site_id, cycle, order_date, supervisor_employee_id, late_reason, expected_submission_at, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(cycle), sqlc.arg(order_date), sqlc.arg(supervisor_employee_id), sqlc.narg(late_reason), sqlc.narg(expected_submission_at), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateKitchenOrder :one
UPDATE kitchen_orders
SET site_id = COALESCE(sqlc.narg(site_id), site_id), cycle = COALESCE(sqlc.narg(cycle), cycle), order_date = COALESCE(sqlc.narg(order_date), order_date), supervisor_employee_id = COALESCE(sqlc.narg(supervisor_employee_id), supervisor_employee_id), late_reason = COALESCE(sqlc.narg(late_reason), late_reason), expected_submission_at = COALESCE(sqlc.narg(expected_submission_at), expected_submission_at), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetKitchenOrderStatus :one
UPDATE kitchen_orders
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListKitchenOrderLines :many
SELECT * FROM kitchen_order_lines WHERE kitchen_order_id = $1 ORDER BY id;

-- name: InsertKitchenOrderLines :one
INSERT INTO kitchen_order_lines (kitchen_order_id, product_id, quantity, uom, stock_on_hand, par_level, expected_sales, unusual_flag, confirmed_quantity, created_by)
VALUES (sqlc.arg(kitchen_order_id), sqlc.arg(product_id), sqlc.arg(quantity), sqlc.arg(uom), sqlc.narg(stock_on_hand), sqlc.narg(par_level), sqlc.narg(expected_sales), COALESCE(sqlc.narg(unusual_flag)::boolean, false), sqlc.narg(confirmed_quantity), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteKitchenOrderLines :exec
DELETE FROM kitchen_order_lines WHERE kitchen_order_id = $1;

-- name: ActSubmitKitchenOrder :one
UPDATE kitchen_orders
SET submitted_at = now(), status = 'submitted', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'planned')
RETURNING *;

-- name: ActAcknowledgeLateKitchenOrder :one
UPDATE kitchen_orders
SET status = 'confirmed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('late_pending')
RETURNING *;

-- name: ActConfirmKitchenOrder :one
UPDATE kitchen_orders
SET status = 'confirmed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'submitted', 'late_pending')
RETURNING *;
