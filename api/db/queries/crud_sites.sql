-- name: GetSite :one
SELECT * FROM sites WHERE id = $1;

-- name: ListSite :many
SELECT * FROM sites
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR code ILIKE '%' || sqlc.narg(q)::text || '%' OR "name" ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(complex_id)::uuid IS NULL OR complex_id = sqlc.narg(complex_id)::uuid)
  AND (sqlc.narg(type)::text IS NULL OR "type" = sqlc.narg(type)::text)
  AND (sqlc.arg(all_sites)::boolean OR id IS NULL OR id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateSite :one
INSERT INTO sites (code, "name", "type", complex_id, brand, trading_pattern, opens_at, closes_at, stockholding, address, escalation_contacts, active, created_by, updated_by)
VALUES (sqlc.arg(code), sqlc.arg(name), sqlc.arg(type), sqlc.narg(complex_id), sqlc.narg(brand), sqlc.narg(trading_pattern), sqlc.narg(opens_at), sqlc.narg(closes_at), COALESCE(sqlc.narg(stockholding)::boolean, false), sqlc.narg(address), sqlc.narg(escalation_contacts), COALESCE(sqlc.narg(active)::boolean, true), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateSite :one
UPDATE sites
SET code = COALESCE(sqlc.narg(code), code), "name" = COALESCE(sqlc.narg(name), "name"), "type" = COALESCE(sqlc.narg(type), "type"), complex_id = COALESCE(sqlc.narg(complex_id), complex_id), brand = COALESCE(sqlc.narg(brand), brand), trading_pattern = COALESCE(sqlc.narg(trading_pattern), trading_pattern), opens_at = COALESCE(sqlc.narg(opens_at), opens_at), closes_at = COALESCE(sqlc.narg(closes_at), closes_at), stockholding = COALESCE(sqlc.narg(stockholding), stockholding), address = COALESCE(sqlc.narg(address), address), escalation_contacts = COALESCE(sqlc.narg(escalation_contacts), escalation_contacts), active = COALESCE(sqlc.narg(active), active), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
