-- name: GetSodRule :one
SELECT * FROM sod_rules WHERE id = $1;

-- name: ListSodRule :many
SELECT * FROM sod_rules
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR "name" ILIKE '%' || sqlc.narg(q)::text || '%')
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateSodRule :one
INSERT INTO sod_rules ("key", "name", action_a, action_b, "scope", enforcement, active, created_by, updated_by)
VALUES (sqlc.arg(key), sqlc.arg(name), sqlc.arg(action_a), sqlc.arg(action_b), sqlc.arg(scope), sqlc.arg(enforcement), COALESCE(sqlc.narg(active)::boolean, true), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateSodRule :one
UPDATE sod_rules
SET "key" = COALESCE(sqlc.narg(key), "key"), "name" = COALESCE(sqlc.narg(name), "name"), action_a = COALESCE(sqlc.narg(action_a), action_a), action_b = COALESCE(sqlc.narg(action_b), action_b), "scope" = COALESCE(sqlc.narg(scope), "scope"), enforcement = COALESCE(sqlc.narg(enforcement), enforcement), active = COALESCE(sqlc.narg(active), active), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
