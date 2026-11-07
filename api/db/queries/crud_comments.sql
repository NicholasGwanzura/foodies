-- name: GetComment :one
SELECT * FROM comments WHERE id = $1;

-- name: ListComment :many
SELECT * FROM comments
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(subject_type)::text IS NULL OR subject_type = sqlc.narg(subject_type)::text)
  AND (sqlc.narg(subject_id)::uuid IS NULL OR subject_id = sqlc.narg(subject_id)::uuid)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateComment :one
INSERT INTO comments (subject_type, subject_id, body, created_by, updated_by)
VALUES (sqlc.arg(subject_type), sqlc.arg(subject_id), sqlc.arg(body), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;
