-- name: GetStoreOpening :one
SELECT * FROM store_openings WHERE id = $1;

-- name: ListStoreOpening :many
SELECT * FROM store_openings
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(business_date)::date IS NULL OR business_date = sqlc.narg(business_date)::date)
  AND (sqlc.narg(from_date)::date IS NULL OR business_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR business_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateStoreOpening :one
INSERT INTO store_openings (reference, site_id, business_date, amber_issue, rationale, corrective_action, actual_opened_at, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(business_date), sqlc.narg(amber_issue), sqlc.narg(rationale), sqlc.narg(corrective_action), sqlc.narg(actual_opened_at), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateStoreOpening :one
UPDATE store_openings
SET site_id = COALESCE(sqlc.narg(site_id), site_id), business_date = COALESCE(sqlc.narg(business_date), business_date), amber_issue = COALESCE(sqlc.narg(amber_issue), amber_issue), rationale = COALESCE(sqlc.narg(rationale), rationale), corrective_action = COALESCE(sqlc.narg(corrective_action), corrective_action), actual_opened_at = COALESCE(sqlc.narg(actual_opened_at), actual_opened_at), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetStoreOpeningStatus :one
UPDATE store_openings
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActDecideStoreOpening :one
UPDATE store_openings
SET rationale = COALESCE(sqlc.narg(rationale), rationale), decided_at = now(), decided_by = sqlc.narg(updated_by), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActRecordOpenedStoreOpening :one
UPDATE store_openings
SET actual_opened_at = COALESCE(sqlc.narg(actual_opened_at), actual_opened_at), status = 'opened', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('go')
RETURNING *;
