-- name: GetEquipmentFault :one
SELECT * FROM equipment_faults WHERE id = $1;

-- name: ListEquipmentFault :many
SELECT * FROM equipment_faults
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(asset_id)::uuid IS NULL OR asset_id = sqlc.narg(asset_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateEquipmentFault :one
INSERT INTO equipment_faults (reference, asset_id, site_id, reported_at, symptoms, severity_level, trading_impact, safety_impact, out_of_service, downtime_started_at, temporary_control, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(asset_id), sqlc.narg(site_id), sqlc.arg(reported_at), sqlc.arg(symptoms), sqlc.arg(severity_level), COALESCE(sqlc.narg(trading_impact)::boolean, false), COALESCE(sqlc.narg(safety_impact)::boolean, false), COALESCE(sqlc.narg(out_of_service)::boolean, false), sqlc.narg(downtime_started_at), sqlc.narg(temporary_control), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateEquipmentFault :one
UPDATE equipment_faults
SET asset_id = COALESCE(sqlc.narg(asset_id), asset_id), reported_at = COALESCE(sqlc.narg(reported_at), reported_at), symptoms = COALESCE(sqlc.narg(symptoms), symptoms), severity_level = COALESCE(sqlc.narg(severity_level), severity_level), trading_impact = COALESCE(sqlc.narg(trading_impact), trading_impact), safety_impact = COALESCE(sqlc.narg(safety_impact), safety_impact), out_of_service = COALESCE(sqlc.narg(out_of_service), out_of_service), downtime_started_at = COALESCE(sqlc.narg(downtime_started_at), downtime_started_at), temporary_control = COALESCE(sqlc.narg(temporary_control), temporary_control), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetEquipmentFaultStatus :one
UPDATE equipment_faults
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActConvertEquipmentFault :one
UPDATE equipment_faults
SET updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
