-- name: GetStocktake :one
SELECT * FROM stocktakes WHERE id = $1;

-- name: ListStocktake :many
SELECT * FROM stocktakes
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(kind)::text IS NULL OR kind = sqlc.narg(kind)::text)
  AND (sqlc.narg(count_date)::date IS NULL OR count_date = sqlc.narg(count_date)::date)
  AND (sqlc.narg(from_date)::date IS NULL OR count_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR count_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateStocktake :one
INSERT INTO stocktakes (reference, site_id, kind, count_date, shift, finance_instruction_ref, alternative_control, system_capture_ref, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(kind), sqlc.arg(count_date), sqlc.narg(shift), sqlc.narg(finance_instruction_ref), sqlc.narg(alternative_control), sqlc.narg(system_capture_ref), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateStocktake :one
UPDATE stocktakes
SET site_id = COALESCE(sqlc.narg(site_id), site_id), kind = COALESCE(sqlc.narg(kind), kind), count_date = COALESCE(sqlc.narg(count_date), count_date), shift = COALESCE(sqlc.narg(shift), shift), finance_instruction_ref = COALESCE(sqlc.narg(finance_instruction_ref), finance_instruction_ref), alternative_control = COALESCE(sqlc.narg(alternative_control), alternative_control), system_capture_ref = COALESCE(sqlc.narg(system_capture_ref), system_capture_ref), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetStocktakeStatus :one
UPDATE stocktakes
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListStocktakeLines :many
SELECT * FROM stocktake_lines WHERE stocktake_id = $1 ORDER BY id;

-- name: InsertStocktakeLines :one
INSERT INTO stocktake_lines (stocktake_id, item_id, location_id, sheet_number, counter_employee_id, count_1, count_2, recount, drs_quantity, gaap_quantity, bin_card_quantity, variance_quantity, variance_value_amount, variance_value_currency, obsolete, "comment", created_by)
VALUES (sqlc.arg(stocktake_id), sqlc.arg(item_id), sqlc.narg(location_id), sqlc.narg(sheet_number), sqlc.narg(counter_employee_id), sqlc.narg(count_1), sqlc.narg(count_2), sqlc.narg(recount), sqlc.narg(drs_quantity), sqlc.narg(gaap_quantity), sqlc.narg(bin_card_quantity), sqlc.narg(variance_quantity), sqlc.narg(variance_value_amount), sqlc.narg(variance_value_currency), COALESCE(sqlc.narg(obsolete)::boolean, false), sqlc.narg(comment), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteStocktakeLines :exec
DELETE FROM stocktake_lines WHERE stocktake_id = $1;

-- name: ActSubmitStocktake :one
UPDATE stocktakes
SET status = 'submitted', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('planned')
RETURNING *;

-- name: ActApproveStocktake :one
UPDATE stocktakes
SET status = 'approved', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('submitted', 'reviewed')
RETURNING *;

-- name: ActRequestSkipStocktake :one
UPDATE stocktakes
SET alternative_control = COALESCE(sqlc.narg(alternative_control), alternative_control), status = 'skipped', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('planned')
RETURNING *;
