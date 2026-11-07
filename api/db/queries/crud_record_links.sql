-- name: GetRecordLink :one
SELECT * FROM record_links WHERE id = $1;

-- name: ListRecordLink :many
SELECT * FROM record_links
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(from_id)::uuid IS NULL OR from_id = sqlc.narg(from_id)::uuid)
  AND (sqlc.narg(to_id)::uuid IS NULL OR to_id = sqlc.narg(to_id)::uuid)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateRecordLink :one
INSERT INTO record_links (from_type, from_id, to_type, to_id, relation, created_by, updated_by)
VALUES (sqlc.arg(from_type), sqlc.arg(from_id), sqlc.arg(to_type), sqlc.arg(to_id), sqlc.arg(relation), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteRecordLink :exec
DELETE FROM record_links WHERE id = $1;
