-- name: GetPurchaseRequisition :one
SELECT * FROM purchase_requisitions WHERE id = $1;

-- name: ListPurchaseRequisition :many
SELECT * FROM purchase_requisitions
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(department_id)::uuid IS NULL OR department_id = sqlc.narg(department_id)::uuid)
  AND (sqlc.narg(requirement_type)::text IS NULL OR requirement_type = sqlc.narg(requirement_type)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR required_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR required_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreatePurchaseRequisition :one
INSERT INTO purchase_requisitions (reference, site_id, department_id, cost_centre_id, requester_employee_id, purpose, requirement_type, required_date, budget_code, yield_estimate, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(department_id), sqlc.arg(cost_centre_id), sqlc.arg(requester_employee_id), sqlc.arg(purpose), sqlc.arg(requirement_type), sqlc.arg(required_date), sqlc.narg(budget_code), sqlc.narg(yield_estimate), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdatePurchaseRequisition :one
UPDATE purchase_requisitions
SET site_id = COALESCE(sqlc.narg(site_id), site_id), department_id = COALESCE(sqlc.narg(department_id), department_id), cost_centre_id = COALESCE(sqlc.narg(cost_centre_id), cost_centre_id), requester_employee_id = COALESCE(sqlc.narg(requester_employee_id), requester_employee_id), purpose = COALESCE(sqlc.narg(purpose), purpose), requirement_type = COALESCE(sqlc.narg(requirement_type), requirement_type), required_date = COALESCE(sqlc.narg(required_date), required_date), budget_code = COALESCE(sqlc.narg(budget_code), budget_code), yield_estimate = COALESCE(sqlc.narg(yield_estimate), yield_estimate), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetPurchaseRequisitionStatus :one
UPDATE purchase_requisitions
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListPurchaseRequisitionLines :many
SELECT * FROM purchase_requisition_lines WHERE purchase_requisition_id = $1 ORDER BY id;

-- name: InsertPurchaseRequisitionLines :one
INSERT INTO purchase_requisition_lines (purchase_requisition_id, item_id, description, specification, quantity_on_hand, quantity_required, uom, estimated_unit_price_amount, estimated_unit_price_currency, created_by)
VALUES (sqlc.arg(purchase_requisition_id), sqlc.narg(item_id), sqlc.arg(description), sqlc.narg(specification), sqlc.narg(quantity_on_hand), sqlc.arg(quantity_required), sqlc.arg(uom), sqlc.narg(estimated_unit_price_amount), sqlc.narg(estimated_unit_price_currency), sqlc.narg(created_by))
RETURNING *;

-- name: DeletePurchaseRequisitionLines :exec
DELETE FROM purchase_requisition_lines WHERE purchase_requisition_id = $1;

-- name: ActSubmitPurchaseRequisition :one
UPDATE purchase_requisitions
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'returned')
RETURNING *;

-- name: ActValidatePurchaseRequisition :one
UPDATE purchase_requisitions
SET status = 'validated', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'submitted', 'returned')
RETURNING *;

-- name: ActReturnPurchaseRequisition :one
UPDATE purchase_requisitions
SET status = 'returned', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'submitted', 'validated', 'pending_approval', 'approved', 'sourcing')
RETURNING *;
