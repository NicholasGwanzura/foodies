-- name: GetRecruitmentRequisition :one
SELECT * FROM recruitment_requisitions WHERE id = $1;

-- name: ListRecruitmentRequisition :many
SELECT * FROM recruitment_requisitions
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateRecruitmentRequisition :one
INSERT INTO recruitment_requisitions (reference, department_id, site_id, job_title, headcount, "reason", created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(department_id), sqlc.narg(site_id), sqlc.arg(job_title), sqlc.arg(headcount), sqlc.arg(reason), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateRecruitmentRequisition :one
UPDATE recruitment_requisitions
SET department_id = COALESCE(sqlc.narg(department_id), department_id), site_id = COALESCE(sqlc.narg(site_id), site_id), job_title = COALESCE(sqlc.narg(job_title), job_title), headcount = COALESCE(sqlc.narg(headcount), headcount), "reason" = COALESCE(sqlc.narg(reason), "reason"), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetRecruitmentRequisitionStatus :one
UPDATE recruitment_requisitions
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitRecruitmentRequisition :one
UPDATE recruitment_requisitions
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'open')
RETURNING *;
