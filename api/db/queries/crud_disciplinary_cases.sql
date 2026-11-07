-- name: GetDisciplinaryCase :one
SELECT * FROM disciplinary_cases WHERE id = $1;

-- name: ListDisciplinaryCase :many
SELECT * FROM disciplinary_cases
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR hearing_on::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR hearing_on::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateDisciplinaryCase :one
INSERT INTO disciplinary_cases (reference, employee_id, allegation, source_type, source_id, stage, hearing_on, outcome, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(employee_id), sqlc.arg(allegation), sqlc.narg(source_type), sqlc.narg(source_id), sqlc.narg(stage), sqlc.narg(hearing_on), sqlc.narg(outcome), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateDisciplinaryCase :one
UPDATE disciplinary_cases
SET employee_id = COALESCE(sqlc.narg(employee_id), employee_id), allegation = COALESCE(sqlc.narg(allegation), allegation), source_type = COALESCE(sqlc.narg(source_type), source_type), source_id = COALESCE(sqlc.narg(source_id), source_id), stage = COALESCE(sqlc.narg(stage), stage), hearing_on = COALESCE(sqlc.narg(hearing_on), hearing_on), outcome = COALESCE(sqlc.narg(outcome), outcome), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetDisciplinaryCaseStatus :one
UPDATE disciplinary_cases
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;
