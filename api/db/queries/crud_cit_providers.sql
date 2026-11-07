-- name: GetCitProvider :one
SELECT * FROM cit_providers WHERE id = $1;

-- name: ListCitProvider :many
SELECT * FROM cit_providers
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR "name" ILIKE '%' || sqlc.narg(q)::text || '%')
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateCitProvider :one
INSERT INTO cit_providers ("name", contract_reference, insurance_limit_amount, insurance_limit_currency, active, created_by, updated_by)
VALUES (sqlc.arg(name), sqlc.narg(contract_reference), sqlc.narg(insurance_limit_amount), sqlc.narg(insurance_limit_currency), COALESCE(sqlc.narg(active)::boolean, true), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateCitProvider :one
UPDATE cit_providers
SET "name" = COALESCE(sqlc.narg(name), "name"), contract_reference = COALESCE(sqlc.narg(contract_reference), contract_reference), insurance_limit_amount = COALESCE(sqlc.narg(insurance_limit_amount), insurance_limit_amount), insurance_limit_currency = COALESCE(sqlc.narg(insurance_limit_currency), insurance_limit_currency), active = COALESCE(sqlc.narg(active), active), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;
