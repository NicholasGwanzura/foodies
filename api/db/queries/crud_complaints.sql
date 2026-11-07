-- name: GetComplaint :one
SELECT * FROM complaints WHERE id = $1;

-- name: ListComplaint :many
SELECT * FROM complaints
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(category)::text IS NULL OR category = sqlc.narg(category)::text)
  AND (sqlc.narg(severity_level)::integer IS NULL OR severity_level = sqlc.narg(severity_level)::integer)
  AND (sqlc.narg(owner_employee_id)::uuid IS NULL OR owner_employee_id = sqlc.narg(owner_employee_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateComplaint :one
INSERT INTO complaints (reference, channel, site_id, received_at, customer_name, customer_contact, category, severity_level, pos_transaction_ref, product_id, batch_id, owner_employee_id, due_at, food_safety_details, evidence_location, ncr_id, resolution, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(channel), sqlc.arg(site_id), sqlc.arg(received_at), sqlc.narg(customer_name), sqlc.narg(customer_contact), sqlc.arg(category), sqlc.arg(severity_level), sqlc.narg(pos_transaction_ref), sqlc.narg(product_id), sqlc.narg(batch_id), sqlc.narg(owner_employee_id), sqlc.narg(due_at), sqlc.narg(food_safety_details), sqlc.narg(evidence_location), sqlc.narg(ncr_id), sqlc.narg(resolution), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateComplaint :one
UPDATE complaints
SET channel = COALESCE(sqlc.narg(channel), channel), site_id = COALESCE(sqlc.narg(site_id), site_id), received_at = COALESCE(sqlc.narg(received_at), received_at), customer_name = COALESCE(sqlc.narg(customer_name), customer_name), customer_contact = COALESCE(sqlc.narg(customer_contact), customer_contact), category = COALESCE(sqlc.narg(category), category), severity_level = COALESCE(sqlc.narg(severity_level), severity_level), pos_transaction_ref = COALESCE(sqlc.narg(pos_transaction_ref), pos_transaction_ref), product_id = COALESCE(sqlc.narg(product_id), product_id), batch_id = COALESCE(sqlc.narg(batch_id), batch_id), owner_employee_id = COALESCE(sqlc.narg(owner_employee_id), owner_employee_id), due_at = COALESCE(sqlc.narg(due_at), due_at), food_safety_details = COALESCE(sqlc.narg(food_safety_details), food_safety_details), evidence_location = COALESCE(sqlc.narg(evidence_location), evidence_location), ncr_id = COALESCE(sqlc.narg(ncr_id), ncr_id), resolution = COALESCE(sqlc.narg(resolution), resolution), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetComplaintStatus :one
UPDATE complaints
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActTriageComplaint :one
UPDATE complaints
SET severity_level = COALESCE(sqlc.narg(severity_level), severity_level), owner_employee_id = COALESCE(sqlc.narg(owner_employee_id), owner_employee_id), status = 'triaged', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('new')
RETURNING *;

-- name: ActResolveComplaint :one
UPDATE complaints
SET resolution = COALESCE(sqlc.narg(resolution), resolution), status = 'resolved', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('new', 'triaged', 'investigating', 'awaiting_customer')
RETURNING *;

-- name: ActCloseComplaint :one
UPDATE complaints
SET closed_at = now(), status = 'closed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('new', 'triaged', 'investigating', 'awaiting_customer', 'resolved')
RETURNING *;
