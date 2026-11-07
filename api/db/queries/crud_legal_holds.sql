-- name: GetLegalHold :one
SELECT * FROM legal_holds WHERE id = $1;

-- name: ListLegalHold :many
SELECT * FROM legal_holds
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateLegalHold :one
INSERT INTO legal_holds (subject_type, subject_id, "reason", owner_user_id, created_by, updated_by)
VALUES (sqlc.arg(subject_type), sqlc.arg(subject_id), sqlc.arg(reason), sqlc.narg(owner_user_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetLegalHoldStatus :one
UPDATE legal_holds
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActReleaseLegalHold :one
UPDATE legal_holds
SET "reason" = COALESCE(sqlc.narg(reason), "reason"), released_at = now(), status = 'released', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('active')
RETURNING *;
