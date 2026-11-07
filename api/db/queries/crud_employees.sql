-- name: GetEmployee :one
SELECT * FROM employees WHERE id = $1;

-- name: ListEmployee :many
SELECT * FROM employees
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%' OR employee_number ILIKE '%' || sqlc.narg(q)::text || '%' OR last_name ILIKE '%' || sqlc.narg(q)::text || '%' OR first_name ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(department_id)::uuid IS NULL OR department_id = sqlc.narg(department_id)::uuid)
  AND (sqlc.narg(employment_status)::text IS NULL OR employment_status = sqlc.narg(employment_status)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR start_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR start_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateEmployee :one
INSERT INTO employees (reference, employee_number, first_name, last_name, job_title, department_id, site_id, manager_employee_id, start_date, contract_type, probation_end, contract_end, phone_e164, email, belina_ref, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(employee_number), sqlc.arg(first_name), sqlc.arg(last_name), sqlc.arg(job_title), sqlc.arg(department_id), sqlc.arg(site_id), sqlc.narg(manager_employee_id), sqlc.arg(start_date), sqlc.narg(contract_type), sqlc.narg(probation_end), sqlc.narg(contract_end), sqlc.narg(phone_e164), sqlc.narg(email), sqlc.narg(belina_ref), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateEmployee :one
UPDATE employees
SET employee_number = COALESCE(sqlc.narg(employee_number), employee_number), first_name = COALESCE(sqlc.narg(first_name), first_name), last_name = COALESCE(sqlc.narg(last_name), last_name), job_title = COALESCE(sqlc.narg(job_title), job_title), department_id = COALESCE(sqlc.narg(department_id), department_id), site_id = COALESCE(sqlc.narg(site_id), site_id), manager_employee_id = COALESCE(sqlc.narg(manager_employee_id), manager_employee_id), start_date = COALESCE(sqlc.narg(start_date), start_date), contract_type = COALESCE(sqlc.narg(contract_type), contract_type), probation_end = COALESCE(sqlc.narg(probation_end), probation_end), contract_end = COALESCE(sqlc.narg(contract_end), contract_end), phone_e164 = COALESCE(sqlc.narg(phone_e164), phone_e164), email = COALESCE(sqlc.narg(email), email), belina_ref = COALESCE(sqlc.narg(belina_ref), belina_ref), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
