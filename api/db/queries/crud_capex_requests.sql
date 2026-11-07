-- name: GetCapexRequest :one
SELECT * FROM capex_requests WHERE id = $1;

-- name: ListCapexRequest :many
SELECT * FROM capex_requests
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%' OR description ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateCapexRequest :one
INSERT INTO capex_requests (reference, site_id, department_id, description, justification, amount_amount, amount_currency, budgeted, risk_assessment, purchase_requisition_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(department_id), sqlc.arg(description), sqlc.arg(justification), sqlc.arg(amount_amount), sqlc.arg(amount_currency), COALESCE(sqlc.narg(budgeted)::boolean, false), sqlc.narg(risk_assessment), sqlc.narg(purchase_requisition_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateCapexRequest :one
UPDATE capex_requests
SET site_id = COALESCE(sqlc.narg(site_id), site_id), department_id = COALESCE(sqlc.narg(department_id), department_id), description = COALESCE(sqlc.narg(description), description), justification = COALESCE(sqlc.narg(justification), justification), amount_amount = COALESCE(sqlc.narg(amount_amount), amount_amount), amount_currency = COALESCE(sqlc.narg(amount_currency), amount_currency), budgeted = COALESCE(sqlc.narg(budgeted), budgeted), risk_assessment = COALESCE(sqlc.narg(risk_assessment), risk_assessment), purchase_requisition_id = COALESCE(sqlc.narg(purchase_requisition_id), purchase_requisition_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetCapexRequestStatus :one
UPDATE capex_requests
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitCapexRequest :one
UPDATE capex_requests
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
