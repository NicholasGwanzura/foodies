-- name: GetPatrolRoute :one
SELECT * FROM patrol_routes WHERE id = $1;

-- name: ListPatrolRoute :many
SELECT * FROM patrol_routes
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR "name" ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreatePatrolRoute :one
INSERT INTO patrol_routes (site_id, "name", frequency_rule_key, active, created_by, updated_by)
VALUES (sqlc.arg(site_id), sqlc.arg(name), sqlc.narg(frequency_rule_key), COALESCE(sqlc.narg(active)::boolean, true), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdatePatrolRoute :one
UPDATE patrol_routes
SET site_id = COALESCE(sqlc.narg(site_id), site_id), "name" = COALESCE(sqlc.narg(name), "name"), frequency_rule_key = COALESCE(sqlc.narg(frequency_rule_key), frequency_rule_key), active = COALESCE(sqlc.narg(active), active), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: ListPatrolCheckpoints :many
SELECT * FROM patrol_checkpoints WHERE patrol_route_id = $1 ORDER BY id;

-- name: InsertPatrolCheckpoints :one
INSERT INTO patrol_checkpoints (patrol_route_id, sequence, "name", risk_category, qr_code, blind_spot, created_by)
VALUES (sqlc.arg(patrol_route_id), sqlc.arg(sequence), sqlc.arg(name), sqlc.narg(risk_category), sqlc.narg(qr_code), COALESCE(sqlc.narg(blind_spot)::boolean, false), sqlc.narg(created_by))
RETURNING *;

-- name: DeletePatrolCheckpoints :exec
DELETE FROM patrol_checkpoints WHERE patrol_route_id = $1;
