-- name: GetAsset :one
SELECT * FROM assets WHERE id = $1;

-- name: ListAsset :many
SELECT * FROM assets
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR description ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(asset_class_id)::uuid IS NULL OR asset_class_id = sqlc.narg(asset_class_id)::uuid)
  AND (sqlc.narg(custodian_employee_id)::uuid IS NULL OR custodian_employee_id = sqlc.narg(custodian_employee_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR acquisition_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR acquisition_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateAsset :one
INSERT INTO assets (tag, description, asset_class_id, serial_number, make_model, acquisition_date, supplier_id, purchase_order_id, goods_receipt_id, invoice_id, cost_amount, cost_currency, site_id, department_id, cost_centre_id, custodian_employee_id, condition, useful_life_years, residual_value_amount, residual_value_currency, critical, authorised_operator_ids, sap_asset_ref, created_by, updated_by)
VALUES (sqlc.arg(tag), sqlc.arg(description), sqlc.arg(asset_class_id), sqlc.narg(serial_number), sqlc.narg(make_model), sqlc.narg(acquisition_date), sqlc.narg(supplier_id), sqlc.narg(purchase_order_id), sqlc.narg(goods_receipt_id), sqlc.narg(invoice_id), sqlc.narg(cost_amount), sqlc.narg(cost_currency), sqlc.arg(site_id), sqlc.narg(department_id), sqlc.narg(cost_centre_id), sqlc.narg(custodian_employee_id), sqlc.narg(condition), sqlc.narg(useful_life_years), sqlc.narg(residual_value_amount), sqlc.narg(residual_value_currency), COALESCE(sqlc.narg(critical)::boolean, false), sqlc.narg(authorised_operator_ids), sqlc.narg(sap_asset_ref), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateAsset :one
UPDATE assets
SET tag = COALESCE(sqlc.narg(tag), tag), description = COALESCE(sqlc.narg(description), description), asset_class_id = COALESCE(sqlc.narg(asset_class_id), asset_class_id), serial_number = COALESCE(sqlc.narg(serial_number), serial_number), make_model = COALESCE(sqlc.narg(make_model), make_model), acquisition_date = COALESCE(sqlc.narg(acquisition_date), acquisition_date), supplier_id = COALESCE(sqlc.narg(supplier_id), supplier_id), purchase_order_id = COALESCE(sqlc.narg(purchase_order_id), purchase_order_id), goods_receipt_id = COALESCE(sqlc.narg(goods_receipt_id), goods_receipt_id), invoice_id = COALESCE(sqlc.narg(invoice_id), invoice_id), cost_amount = COALESCE(sqlc.narg(cost_amount), cost_amount), cost_currency = COALESCE(sqlc.narg(cost_currency), cost_currency), site_id = COALESCE(sqlc.narg(site_id), site_id), department_id = COALESCE(sqlc.narg(department_id), department_id), cost_centre_id = COALESCE(sqlc.narg(cost_centre_id), cost_centre_id), custodian_employee_id = COALESCE(sqlc.narg(custodian_employee_id), custodian_employee_id), condition = COALESCE(sqlc.narg(condition), condition), useful_life_years = COALESCE(sqlc.narg(useful_life_years), useful_life_years), residual_value_amount = COALESCE(sqlc.narg(residual_value_amount), residual_value_amount), residual_value_currency = COALESCE(sqlc.narg(residual_value_currency), residual_value_currency), critical = COALESCE(sqlc.narg(critical), critical), authorised_operator_ids = COALESCE(sqlc.narg(authorised_operator_ids), authorised_operator_ids), sap_asset_ref = COALESCE(sqlc.narg(sap_asset_ref), sap_asset_ref), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetAssetStatus :one
UPDATE assets
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
