-- name: GetReport :one
SELECT * FROM reports WHERE id = $1;

-- name: ListReport :many
SELECT * FROM reports
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(report_definition_id)::uuid IS NULL OR report_definition_id = sqlc.narg(report_definition_id)::uuid)
  AND (sqlc.narg(department_id)::uuid IS NULL OR department_id = sqlc.narg(department_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR period_start::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR period_start::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: UpdateReport :one
UPDATE reports
SET report_definition_id = COALESCE(sqlc.narg(report_definition_id), report_definition_id), period_start = COALESCE(sqlc.narg(period_start), period_start), period_end = COALESCE(sqlc.narg(period_end), period_end), site_id = COALESCE(sqlc.narg(site_id), site_id), department_id = COALESCE(sqlc.narg(department_id), department_id), deputy_reason = COALESCE(sqlc.narg(deputy_reason), deputy_reason), form_submission_id = COALESCE(sqlc.narg(form_submission_id), form_submission_id), delay_reason = COALESCE(sqlc.narg(delay_reason), delay_reason), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetReportStatus :one
UPDATE reports
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitReport :one
UPDATE reports
SET submitted_at = now(), submitted_by = sqlc.narg(updated_by), status = 'submitted', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActReviewReport :one
UPDATE reports
SET status = 'reviewed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('pending', 'submitted', 'reviewed')
RETURNING *;

-- name: ActExplainLateReport :one
UPDATE reports
SET delay_reason = COALESCE(sqlc.narg(delay_reason), delay_reason), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
