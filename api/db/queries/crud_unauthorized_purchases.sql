-- name: GetUnauthorizedPurchase :one
SELECT * FROM unauthorized_purchases WHERE id = $1;

-- name: ListUnauthorizedPurchase :many
SELECT * FROM unauthorized_purchases
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR purchased_on::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR purchased_on::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateUnauthorizedPurchase :one
INSERT INTO unauthorized_purchases (reference, site_id, originator_employee_id, supplier_name, amount_amount, amount_currency, purchased_on, "reason", operations_review, finance_treatment, corrective_action_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(originator_employee_id), sqlc.narg(supplier_name), sqlc.arg(amount_amount), sqlc.arg(amount_currency), sqlc.arg(purchased_on), sqlc.arg(reason), sqlc.narg(operations_review), sqlc.narg(finance_treatment), sqlc.narg(corrective_action_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateUnauthorizedPurchase :one
UPDATE unauthorized_purchases
SET site_id = COALESCE(sqlc.narg(site_id), site_id), originator_employee_id = COALESCE(sqlc.narg(originator_employee_id), originator_employee_id), supplier_name = COALESCE(sqlc.narg(supplier_name), supplier_name), amount_amount = COALESCE(sqlc.narg(amount_amount), amount_amount), amount_currency = COALESCE(sqlc.narg(amount_currency), amount_currency), purchased_on = COALESCE(sqlc.narg(purchased_on), purchased_on), "reason" = COALESCE(sqlc.narg(reason), "reason"), operations_review = COALESCE(sqlc.narg(operations_review), operations_review), finance_treatment = COALESCE(sqlc.narg(finance_treatment), finance_treatment), corrective_action_id = COALESCE(sqlc.narg(corrective_action_id), corrective_action_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetUnauthorizedPurchaseStatus :one
UPDATE unauthorized_purchases
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActCloseUnauthorizedPurchase :one
UPDATE unauthorized_purchases
SET status = 'closed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('reported', 'under_review')
RETURNING *;
