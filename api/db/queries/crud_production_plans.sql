-- name: GetProductionPlan :one
SELECT * FROM production_plans WHERE id = $1;

-- name: ListProductionPlan :many
SELECT * FROM production_plans
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(plan_date)::date IS NULL OR plan_date = sqlc.narg(plan_date)::date)
  AND (sqlc.narg(from_date)::date IS NULL OR plan_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR plan_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateProductionPlan :one
INSERT INTO production_plans (reference, site_id, plan_date, cycle, line_clearance_submission_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(plan_date), sqlc.arg(cycle), sqlc.narg(line_clearance_submission_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateProductionPlan :one
UPDATE production_plans
SET site_id = COALESCE(sqlc.narg(site_id), site_id), plan_date = COALESCE(sqlc.narg(plan_date), plan_date), cycle = COALESCE(sqlc.narg(cycle), cycle), line_clearance_submission_id = COALESCE(sqlc.narg(line_clearance_submission_id), line_clearance_submission_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetProductionPlanStatus :one
UPDATE production_plans
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListProductionPlanLines :many
SELECT * FROM production_plan_lines WHERE production_plan_id = $1 ORDER BY id;

-- name: InsertProductionPlanLines :one
INSERT INTO production_plan_lines (production_plan_id, product_id, production_line_id, planned_quantity, uom, material_shortage, created_by)
VALUES (sqlc.arg(production_plan_id), sqlc.arg(product_id), sqlc.narg(production_line_id), sqlc.arg(planned_quantity), sqlc.arg(uom), sqlc.narg(material_shortage), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteProductionPlanLines :exec
DELETE FROM production_plan_lines WHERE production_plan_id = $1;

-- name: ActConfirmProductionPlan :one
UPDATE production_plans
SET status = 'confirmed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
