-- name: GetCertificate :one
SELECT * FROM certificates WHERE id = $1;

-- name: ListCertificate :many
SELECT * FROM certificates
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(employee_id)::uuid IS NULL OR employee_id = sqlc.narg(employee_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR issued_on::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR issued_on::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateCertificate :one
INSERT INTO certificates (employee_id, course_name, provider, certificate_number, issued_on, expires_on, attachment_id, created_by, updated_by)
VALUES (sqlc.arg(employee_id), sqlc.arg(course_name), sqlc.narg(provider), sqlc.narg(certificate_number), sqlc.narg(issued_on), sqlc.narg(expires_on), sqlc.narg(attachment_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateCertificate :one
UPDATE certificates
SET employee_id = COALESCE(sqlc.narg(employee_id), employee_id), course_name = COALESCE(sqlc.narg(course_name), course_name), provider = COALESCE(sqlc.narg(provider), provider), certificate_number = COALESCE(sqlc.narg(certificate_number), certificate_number), issued_on = COALESCE(sqlc.narg(issued_on), issued_on), expires_on = COALESCE(sqlc.narg(expires_on), expires_on), attachment_id = COALESCE(sqlc.narg(attachment_id), attachment_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
