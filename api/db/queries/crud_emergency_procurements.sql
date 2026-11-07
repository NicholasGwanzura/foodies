-- name: GetEmergencyProcurement :one
SELECT * FROM emergency_procurements WHERE id = $1;

-- name: ListEmergencyProcurement :many
SELECT * FROM emergency_procurements
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

-- name: CreateEmergencyProcurement :one
INSERT INTO emergency_procurements (reference, site_id, event, why_unforeseeable, continuity_impact, supplier_id, amount_amount, amount_currency, purchased_at, purchase_requisition_id, post_review, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(event), sqlc.arg(why_unforeseeable), sqlc.narg(continuity_impact), sqlc.narg(supplier_id), sqlc.arg(amount_amount), sqlc.arg(amount_currency), sqlc.narg(purchased_at), sqlc.narg(purchase_requisition_id), sqlc.narg(post_review), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateEmergencyProcurement :one
UPDATE emergency_procurements
SET site_id = COALESCE(sqlc.narg(site_id), site_id), event = COALESCE(sqlc.narg(event), event), why_unforeseeable = COALESCE(sqlc.narg(why_unforeseeable), why_unforeseeable), continuity_impact = COALESCE(sqlc.narg(continuity_impact), continuity_impact), supplier_id = COALESCE(sqlc.narg(supplier_id), supplier_id), amount_amount = COALESCE(sqlc.narg(amount_amount), amount_amount), amount_currency = COALESCE(sqlc.narg(amount_currency), amount_currency), purchased_at = COALESCE(sqlc.narg(purchased_at), purchased_at), purchase_requisition_id = COALESCE(sqlc.narg(purchase_requisition_id), purchase_requisition_id), post_review = COALESCE(sqlc.narg(post_review), post_review), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetEmergencyProcurementStatus :one
UPDATE emergency_procurements
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitEmergencyProcurement :one
UPDATE emergency_procurements
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActReviewEmergencyProcurement :one
UPDATE emergency_procurements
SET post_review = COALESCE(sqlc.narg(post_review), post_review), status = 'reviewed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('pending_approval', 'reviewed')
RETURNING *;
