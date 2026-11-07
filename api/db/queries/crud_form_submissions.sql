-- name: GetFormSubmission :one
SELECT * FROM form_submissions WHERE id = $1;

-- name: ListFormSubmission :many
SELECT * FROM form_submissions
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(template_version_id)::uuid IS NULL OR template_version_id = sqlc.narg(template_version_id)::uuid)
  AND (sqlc.narg(business_date)::date IS NULL OR business_date = sqlc.narg(business_date)::date)
  AND (sqlc.narg(result)::text IS NULL OR "result" = sqlc.narg(result)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR business_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR business_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateFormSubmission :one
INSERT INTO form_submissions (reference, template_version_id, site_id, business_date, shift, task_id, answers, captured_at, client_record_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(template_version_id), sqlc.arg(site_id), sqlc.arg(business_date), sqlc.narg(shift), sqlc.narg(task_id), sqlc.arg(answers), sqlc.arg(captured_at), sqlc.narg(client_record_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateFormSubmission :one
UPDATE form_submissions
SET template_version_id = COALESCE(sqlc.narg(template_version_id), template_version_id), site_id = COALESCE(sqlc.narg(site_id), site_id), business_date = COALESCE(sqlc.narg(business_date), business_date), shift = COALESCE(sqlc.narg(shift), shift), task_id = COALESCE(sqlc.narg(task_id), task_id), answers = COALESCE(sqlc.narg(answers), answers), captured_at = COALESCE(sqlc.narg(captured_at), captured_at), client_record_id = COALESCE(sqlc.narg(client_record_id), client_record_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetFormSubmissionStatus :one
UPDATE form_submissions
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitFormSubmission :one
UPDATE form_submissions
SET status = 'pending_review', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'returned')
RETURNING *;

-- name: ActReviewFormSubmission :one
UPDATE form_submissions
SET status = 'closed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'submitted', 'pending_signature', 'pending_review')
RETURNING *;

-- name: ActReturnFormSubmission :one
UPDATE form_submissions
SET status = 'returned', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'submitted', 'pending_signature', 'pending_review')
RETURNING *;
