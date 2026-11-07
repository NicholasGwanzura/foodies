-- name: GetPublicHoliday :one
SELECT * FROM public_holidays WHERE id = $1;

-- name: ListPublicHoliday :many
SELECT * FROM public_holidays
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR "name" ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(from_date)::date IS NULL OR holiday_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR holiday_date::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreatePublicHoliday :one
INSERT INTO public_holidays (holiday_date, "name", applies_to_site_id, created_by, updated_by)
VALUES (sqlc.arg(holiday_date), sqlc.arg(name), sqlc.narg(applies_to_site_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdatePublicHoliday :one
UPDATE public_holidays
SET holiday_date = COALESCE(sqlc.narg(holiday_date), holiday_date), "name" = COALESCE(sqlc.narg(name), "name"), applies_to_site_id = COALESCE(sqlc.narg(applies_to_site_id), applies_to_site_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: DeletePublicHoliday :exec
DELETE FROM public_holidays WHERE id = $1;
