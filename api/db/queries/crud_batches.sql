-- name: GetBatch :one
SELECT * FROM batches WHERE id = $1;

-- name: ListBatch :many
SELECT * FROM batches
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(product_id)::uuid IS NULL OR product_id = sqlc.narg(product_id)::uuid)
  AND (sqlc.narg(production_date)::date IS NULL OR production_date = sqlc.narg(production_date)::date)
  AND (sqlc.narg(qc_status)::text IS NULL OR qc_status = sqlc.narg(qc_status)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR production_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR production_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateBatch :one
INSERT INTO batches (batch_number, product_id, recipe_version_id, site_id, production_line_id, production_plan_id, production_date, shift, operator_employee_id, input_quantity, planned_output, actual_output, uom, expiry_date, holding_condition, created_by, updated_by)
VALUES (sqlc.arg(batch_number), sqlc.arg(product_id), sqlc.narg(recipe_version_id), sqlc.arg(site_id), sqlc.narg(production_line_id), sqlc.narg(production_plan_id), sqlc.arg(production_date), sqlc.narg(shift), sqlc.arg(operator_employee_id), sqlc.narg(input_quantity), sqlc.narg(planned_output), sqlc.narg(actual_output), sqlc.arg(uom), sqlc.narg(expiry_date), sqlc.narg(holding_condition), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateBatch :one
UPDATE batches
SET batch_number = COALESCE(sqlc.narg(batch_number), batch_number), product_id = COALESCE(sqlc.narg(product_id), product_id), recipe_version_id = COALESCE(sqlc.narg(recipe_version_id), recipe_version_id), site_id = COALESCE(sqlc.narg(site_id), site_id), production_line_id = COALESCE(sqlc.narg(production_line_id), production_line_id), production_plan_id = COALESCE(sqlc.narg(production_plan_id), production_plan_id), production_date = COALESCE(sqlc.narg(production_date), production_date), shift = COALESCE(sqlc.narg(shift), shift), operator_employee_id = COALESCE(sqlc.narg(operator_employee_id), operator_employee_id), input_quantity = COALESCE(sqlc.narg(input_quantity), input_quantity), planned_output = COALESCE(sqlc.narg(planned_output), planned_output), actual_output = COALESCE(sqlc.narg(actual_output), actual_output), uom = COALESCE(sqlc.narg(uom), uom), expiry_date = COALESCE(sqlc.narg(expiry_date), expiry_date), holding_condition = COALESCE(sqlc.narg(holding_condition), holding_condition), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetBatchStatus :one
UPDATE batches
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListBatchIngredients :many
SELECT * FROM batch_ingredients WHERE batch_id = $1 ORDER BY id;

-- name: InsertBatchIngredients :one
INSERT INTO batch_ingredients (batch_id, item_id, supplier_lot, goods_receipt_line_id, planned_quantity, issued_quantity, actual_quantity, uom, created_by)
VALUES (sqlc.arg(batch_id), sqlc.arg(item_id), sqlc.arg(supplier_lot), sqlc.narg(goods_receipt_line_id), sqlc.narg(planned_quantity), sqlc.narg(issued_quantity), sqlc.narg(actual_quantity), sqlc.arg(uom), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteBatchIngredients :exec
DELETE FROM batch_ingredients WHERE batch_id = $1;

-- name: ActCompleteBatch :one
UPDATE batches
SET status = 'completed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('in_progress')
RETURNING *;

-- name: ActReleaseBatch :one
UPDATE batches
SET status = 'closed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('completed')
RETURNING *;
