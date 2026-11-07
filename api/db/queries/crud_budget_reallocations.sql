-- name: GetBudgetReallocation :one
SELECT * FROM budget_reallocations WHERE id = $1;

-- name: ListBudgetReallocation :many
SELECT * FROM budget_reallocations
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateBudgetReallocation :one
INSERT INTO budget_reallocations (reference, budget_id, from_account, to_account, amount_amount, amount_currency, "reason", created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(budget_id), sqlc.narg(from_account), sqlc.narg(to_account), sqlc.arg(amount_amount), sqlc.arg(amount_currency), sqlc.arg(reason), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetBudgetReallocationStatus :one
UPDATE budget_reallocations
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitBudgetReallocation :one
UPDATE budget_reallocations
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
