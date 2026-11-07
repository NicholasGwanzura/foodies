-- name: GetStockRequisition :one
SELECT * FROM stock_requisitions WHERE id = $1;

-- name: ListStockRequisition :many
SELECT * FROM stock_requisitions
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateStockRequisition :one
INSERT INTO stock_requisitions (reference, site_id, requesting_department_id, requester_employee_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.narg(requesting_department_id), sqlc.arg(requester_employee_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateStockRequisition :one
UPDATE stock_requisitions
SET site_id = COALESCE(sqlc.narg(site_id), site_id), requesting_department_id = COALESCE(sqlc.narg(requesting_department_id), requesting_department_id), requester_employee_id = COALESCE(sqlc.narg(requester_employee_id), requester_employee_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetStockRequisitionStatus :one
UPDATE stock_requisitions
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListStockRequisitionLines :many
SELECT * FROM stock_requisition_lines WHERE stock_requisition_id = $1 ORDER BY id;

-- name: InsertStockRequisitionLines :one
INSERT INTO stock_requisition_lines (stock_requisition_id, item_id, requested_quantity, issued_quantity, uom, batch_lot, bin_card_ref, created_by)
VALUES (sqlc.arg(stock_requisition_id), sqlc.arg(item_id), sqlc.arg(requested_quantity), sqlc.narg(issued_quantity), sqlc.arg(uom), sqlc.narg(batch_lot), sqlc.narg(bin_card_ref), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteStockRequisitionLines :exec
DELETE FROM stock_requisition_lines WHERE stock_requisition_id = $1;

-- name: ActSubmitStockRequisition :one
UPDATE stock_requisitions
SET status = 'submitted', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActIssueStockRequisition :one
UPDATE stock_requisitions
SET issued_at = now(), issued_by = sqlc.narg(updated_by), status = 'issued', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'submitted', 'approved')
RETURNING *;

-- name: ActAcknowledgeStockRequisition :one
UPDATE stock_requisitions
SET acknowledged_at = now(), status = 'acknowledged', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'submitted', 'approved', 'issued')
RETURNING *;
