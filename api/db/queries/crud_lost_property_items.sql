-- name: GetLostPropertyItem :one
SELECT * FROM lost_property_items WHERE id = $1;

-- name: ListLostPropertyItem :many
SELECT * FROM lost_property_items
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%' OR description ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(item_type)::text IS NULL OR item_type = sqlc.narg(item_type)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR disposition_due::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR disposition_due::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateLostPropertyItem :one
INSERT INTO lost_property_items (reference, site_id, item_type, description, found_at, found_location, finder_employee_id, custodian_employee_id, storage_location, issuer_if_visible, name_if_visible, contact_attempts, claimant_name, claimant_evidence, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(item_type), sqlc.arg(description), sqlc.arg(found_at), sqlc.arg(found_location), sqlc.arg(finder_employee_id), sqlc.arg(custodian_employee_id), sqlc.narg(storage_location), sqlc.narg(issuer_if_visible), sqlc.narg(name_if_visible), sqlc.narg(contact_attempts), sqlc.narg(claimant_name), sqlc.narg(claimant_evidence), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateLostPropertyItem :one
UPDATE lost_property_items
SET site_id = COALESCE(sqlc.narg(site_id), site_id), item_type = COALESCE(sqlc.narg(item_type), item_type), description = COALESCE(sqlc.narg(description), description), found_at = COALESCE(sqlc.narg(found_at), found_at), found_location = COALESCE(sqlc.narg(found_location), found_location), finder_employee_id = COALESCE(sqlc.narg(finder_employee_id), finder_employee_id), custodian_employee_id = COALESCE(sqlc.narg(custodian_employee_id), custodian_employee_id), storage_location = COALESCE(sqlc.narg(storage_location), storage_location), issuer_if_visible = COALESCE(sqlc.narg(issuer_if_visible), issuer_if_visible), name_if_visible = COALESCE(sqlc.narg(name_if_visible), name_if_visible), contact_attempts = COALESCE(sqlc.narg(contact_attempts), contact_attempts), claimant_name = COALESCE(sqlc.narg(claimant_name), claimant_name), claimant_evidence = COALESCE(sqlc.narg(claimant_evidence), claimant_evidence), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetLostPropertyItemStatus :one
UPDATE lost_property_items
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActReturnLostPropertyItem :one
UPDATE lost_property_items
SET claimant_name = COALESCE(sqlc.narg(claimant_name), claimant_name), claimant_evidence = COALESCE(sqlc.narg(claimant_evidence), claimant_evidence), status = 'returned', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('in_custody', 'disposed')
RETURNING *;

-- name: ActDisposeLostPropertyItem :one
UPDATE lost_property_items
SET disposition = COALESCE(sqlc.narg(disposition), disposition), status = 'disposed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('in_custody')
RETURNING *;
