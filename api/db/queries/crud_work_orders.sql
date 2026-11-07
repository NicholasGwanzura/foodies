-- name: GetWorkOrder :one
SELECT * FROM work_orders WHERE id = $1;

-- name: ListWorkOrder :many
SELECT * FROM work_orders
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%' OR description ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(asset_id)::uuid IS NULL OR asset_id = sqlc.narg(asset_id)::uuid)
  AND (sqlc.narg(kind)::text IS NULL OR kind = sqlc.narg(kind)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateWorkOrder :one
INSERT INTO work_orders (reference, asset_id, site_id, kind, equipment_fault_id, maintenance_plan_id, priority, description, assigned_employee_id, contractor, repair_eta, parts, service_report_attachment_id, invoice_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(asset_id), sqlc.narg(site_id), sqlc.arg(kind), sqlc.narg(equipment_fault_id), sqlc.narg(maintenance_plan_id), sqlc.arg(priority), sqlc.arg(description), sqlc.narg(assigned_employee_id), sqlc.narg(contractor), sqlc.narg(repair_eta), sqlc.narg(parts), sqlc.narg(service_report_attachment_id), sqlc.narg(invoice_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateWorkOrder :one
UPDATE work_orders
SET asset_id = COALESCE(sqlc.narg(asset_id), asset_id), kind = COALESCE(sqlc.narg(kind), kind), equipment_fault_id = COALESCE(sqlc.narg(equipment_fault_id), equipment_fault_id), maintenance_plan_id = COALESCE(sqlc.narg(maintenance_plan_id), maintenance_plan_id), priority = COALESCE(sqlc.narg(priority), priority), description = COALESCE(sqlc.narg(description), description), assigned_employee_id = COALESCE(sqlc.narg(assigned_employee_id), assigned_employee_id), contractor = COALESCE(sqlc.narg(contractor), contractor), repair_eta = COALESCE(sqlc.narg(repair_eta), repair_eta), parts = COALESCE(sqlc.narg(parts), parts), service_report_attachment_id = COALESCE(sqlc.narg(service_report_attachment_id), service_report_attachment_id), invoice_id = COALESCE(sqlc.narg(invoice_id), invoice_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetWorkOrderStatus :one
UPDATE work_orders
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActAssignWorkOrder :one
UPDATE work_orders
SET assigned_employee_id = COALESCE(sqlc.narg(assigned_employee_id), assigned_employee_id), status = 'assigned', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('open')
RETURNING *;

-- name: ActStartWorkOrder :one
UPDATE work_orders
SET status = 'in_progress', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('assigned')
RETURNING *;

-- name: ActCompleteWorkOrder :one
UPDATE work_orders
SET service_report_attachment_id = COALESCE(sqlc.narg(service_report_attachment_id), service_report_attachment_id), status = 'closed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('open', 'assigned', 'in_progress', 'pending_clearance')
RETURNING *;

-- name: ActClearWorkOrder :one
UPDATE work_orders
SET status = 'pending_clearance', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('in_progress')
RETURNING *;
