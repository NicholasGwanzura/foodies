-- name: GetPosting :one
SELECT * FROM postings WHERE id = $1;

-- name: ListPosting :many
SELECT * FROM postings
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(target_system)::text IS NULL OR target_system = sqlc.narg(target_system)::text)
  AND (sqlc.narg(subject_id)::uuid IS NULL OR subject_id = sqlc.narg(subject_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: SetPostingStatus :one
UPDATE postings
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActRetryPosting :one
UPDATE postings
SET updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActMarkPostedPosting :one
UPDATE postings
SET external_ref = COALESCE(sqlc.narg(external_ref), external_ref), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
