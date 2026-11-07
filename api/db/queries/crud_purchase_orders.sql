-- name: GetPurchaseOrder :one
SELECT * FROM purchase_orders WHERE id = $1;

-- name: ListPurchaseOrder :many
SELECT * FROM purchase_orders
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(supplier_id)::uuid IS NULL OR supplier_id = sqlc.narg(supplier_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR order_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR order_date::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreatePurchaseOrder :one
INSERT INTO purchase_orders (reference, purchase_requisition_id, supplier_id, currency, delivery_site_id, delivery_date, payment_terms, special_conditions, confirmed_at, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(purchase_requisition_id), sqlc.arg(supplier_id), sqlc.arg(currency), sqlc.arg(delivery_site_id), sqlc.arg(delivery_date), sqlc.narg(payment_terms), sqlc.narg(special_conditions), sqlc.narg(confirmed_at), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdatePurchaseOrder :one
UPDATE purchase_orders
SET purchase_requisition_id = COALESCE(sqlc.narg(purchase_requisition_id), purchase_requisition_id), supplier_id = COALESCE(sqlc.narg(supplier_id), supplier_id), currency = COALESCE(sqlc.narg(currency), currency), delivery_site_id = COALESCE(sqlc.narg(delivery_site_id), delivery_site_id), delivery_date = COALESCE(sqlc.narg(delivery_date), delivery_date), payment_terms = COALESCE(sqlc.narg(payment_terms), payment_terms), special_conditions = COALESCE(sqlc.narg(special_conditions), special_conditions), confirmed_at = COALESCE(sqlc.narg(confirmed_at), confirmed_at), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetPurchaseOrderStatus :one
UPDATE purchase_orders
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListPurchaseOrderLines :many
SELECT * FROM purchase_order_lines WHERE purchase_order_id = $1 ORDER BY id;

-- name: InsertPurchaseOrderLines :one
INSERT INTO purchase_order_lines (purchase_order_id, item_id, description, quantity, uom, unit_price, line_total, received_quantity, created_by)
VALUES (sqlc.arg(purchase_order_id), sqlc.narg(item_id), sqlc.arg(description), sqlc.arg(quantity), sqlc.arg(uom), sqlc.arg(unit_price), sqlc.narg(line_total), sqlc.narg(received_quantity), sqlc.narg(created_by))
RETURNING *;

-- name: DeletePurchaseOrderLines :exec
DELETE FROM purchase_order_lines WHERE purchase_order_id = $1;

-- name: ActSubmitPurchaseOrder :one
UPDATE purchase_orders
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActIssuePurchaseOrder :one
UPDATE purchase_orders
SET status = 'issued', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'pending_approval', 'approved')
RETURNING *;

-- name: ActCancelPurchaseOrder :one
UPDATE purchase_orders
SET status = 'cancelled', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'pending_approval', 'approved', 'issued', 'part_received', 'received')
RETURNING *;
