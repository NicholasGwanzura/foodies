-- name: GetAssetVerification :one
SELECT * FROM asset_verifications WHERE id = $1;

-- name: ListAssetVerification :many
SELECT * FROM asset_verifications
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateAssetVerification :one
INSERT INTO asset_verifications (reference, site_id, verification_year, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(verification_year), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateAssetVerification :one
UPDATE asset_verifications
SET site_id = COALESCE(sqlc.narg(site_id), site_id), verification_year = COALESCE(sqlc.narg(verification_year), verification_year), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetAssetVerificationStatus :one
UPDATE asset_verifications
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListAssetVerificationLines :many
SELECT * FROM asset_verification_lines WHERE asset_verification_id = $1 ORDER BY id;

-- name: InsertAssetVerificationLines :one
INSERT INTO asset_verification_lines (asset_verification_id, asset_id, found, condition, location_correct, note, created_by)
VALUES (sqlc.arg(asset_verification_id), sqlc.arg(asset_id), COALESCE(sqlc.narg(found)::boolean, false), sqlc.narg(condition), COALESCE(sqlc.narg(location_correct)::boolean, false), sqlc.narg(note), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteAssetVerificationLines :exec
DELETE FROM asset_verification_lines WHERE asset_verification_id = $1;

-- name: ActSubmitAssetVerification :one
UPDATE asset_verifications
SET status = 'submitted', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('planned')
RETURNING *;

-- name: ActReviewAssetVerification :one
UPDATE asset_verifications
SET status = 'reviewed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('planned', 'in_progress', 'submitted')
RETURNING *;
