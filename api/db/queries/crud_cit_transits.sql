-- name: GetCitTransit :one
SELECT * FROM cit_transits WHERE id = $1;

-- name: ListCitTransit :many
SELECT * FROM cit_transits
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(provider_id)::uuid IS NULL OR provider_id = sqlc.narg(provider_id)::uuid)
  AND (sqlc.narg(transit_date)::date IS NULL OR transit_date = sqlc.narg(transit_date)::date)
  AND (sqlc.narg(from_date)::date IS NULL OR transit_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR transit_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateCitTransit :one
INSERT INTO cit_transits (reference, site_id, provider_id, transit_date, total_amount, total_currency, slip_copy_dispatched, slip_copy_retained, cit_receipt_number, armoured_vehicle_confirmed, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(provider_id), sqlc.arg(transit_date), sqlc.arg(total_amount), sqlc.arg(total_currency), COALESCE(sqlc.narg(slip_copy_dispatched)::boolean, false), COALESCE(sqlc.narg(slip_copy_retained)::boolean, false), sqlc.narg(cit_receipt_number), COALESCE(sqlc.narg(armoured_vehicle_confirmed)::boolean, false), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateCitTransit :one
UPDATE cit_transits
SET site_id = COALESCE(sqlc.narg(site_id), site_id), provider_id = COALESCE(sqlc.narg(provider_id), provider_id), transit_date = COALESCE(sqlc.narg(transit_date), transit_date), total_amount = COALESCE(sqlc.narg(total_amount), total_amount), total_currency = COALESCE(sqlc.narg(total_currency), total_currency), slip_copy_dispatched = COALESCE(sqlc.narg(slip_copy_dispatched), slip_copy_dispatched), slip_copy_retained = COALESCE(sqlc.narg(slip_copy_retained), slip_copy_retained), cit_receipt_number = COALESCE(sqlc.narg(cit_receipt_number), cit_receipt_number), armoured_vehicle_confirmed = COALESCE(sqlc.narg(armoured_vehicle_confirmed), armoured_vehicle_confirmed), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetCitTransitStatus :one
UPDATE cit_transits
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListCitBags :many
SELECT * FROM cit_bags WHERE cit_transit_id = $1 ORDER BY id;

-- name: InsertCitBags :one
INSERT INTO cit_bags (cit_transit_id, bag_number, seal_number, amount_amount, amount_currency, seal_condition_on_receipt, created_by)
VALUES (sqlc.arg(cit_transit_id), sqlc.arg(bag_number), sqlc.arg(seal_number), sqlc.arg(amount_amount), sqlc.arg(amount_currency), sqlc.narg(seal_condition_on_receipt), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteCitBags :exec
DELETE FROM cit_bags WHERE cit_transit_id = $1;

-- name: ActSealCitTransit :one
UPDATE cit_transits
SET status = 'sealed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActVerifyIdentityCitTransit :one
UPDATE cit_transits
SET status = 'identity_verified', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('sealed')
RETURNING *;

-- name: ActHandOverCitTransit :one
UPDATE cit_transits
SET cit_receipt_number = COALESCE(sqlc.narg(cit_receipt_number), cit_receipt_number), armoured_vehicle_confirmed = COALESCE(sqlc.narg(armoured_vehicle_confirmed), armoured_vehicle_confirmed), status = 'handed_over', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('identity_verified')
RETURNING *;

-- name: ActReceiveVaultCitTransit :one
UPDATE cit_transits
SET status = 'received', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('handed_over', 'in_transit')
RETURNING *;

-- name: ActRaiseDiscrepancyCitTransit :one
UPDATE cit_transits
SET status = 'discrepancy', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('handed_over', 'in_transit', 'received')
RETURNING *;
