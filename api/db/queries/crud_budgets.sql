-- name: GetBudget :one
SELECT * FROM budgets WHERE id = $1;

-- name: ListBudget :many
SELECT * FROM budgets
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(financial_year)::integer IS NULL OR financial_year = sqlc.narg(financial_year)::integer)
  AND (sqlc.narg(department_id)::uuid IS NULL OR department_id = sqlc.narg(department_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR submission_due::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR submission_due::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateBudget :one
INSERT INTO budgets (reference, financial_year, department_id, submission_due, board_approval_attachment_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(financial_year), sqlc.arg(department_id), sqlc.narg(submission_due), sqlc.narg(board_approval_attachment_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateBudget :one
UPDATE budgets
SET financial_year = COALESCE(sqlc.narg(financial_year), financial_year), department_id = COALESCE(sqlc.narg(department_id), department_id), submission_due = COALESCE(sqlc.narg(submission_due), submission_due), board_approval_attachment_id = COALESCE(sqlc.narg(board_approval_attachment_id), board_approval_attachment_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetBudgetStatus :one
UPDATE budgets
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListBudgetLines :many
SELECT * FROM budget_lines WHERE budget_id = $1 ORDER BY id;

-- name: InsertBudgetLines :one
INSERT INTO budget_lines (budget_id, account_code, description, "month", amount, created_by)
VALUES (sqlc.arg(budget_id), sqlc.arg(account_code), sqlc.narg(description), sqlc.arg(month), sqlc.arg(amount), sqlc.narg(created_by))
RETURNING *;

-- name: DeleteBudgetLines :exec
DELETE FROM budget_lines WHERE budget_id = $1;

-- name: ActSubmitBudget :one
UPDATE budgets
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
