-- name: GetFloatException :one
SELECT * FROM float_exceptions WHERE id = $1;

-- name: ListFloatException :many
SELECT * FROM float_exceptions
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR transaction_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR transaction_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateFloatException :one
INSERT INTO float_exceptions (reference, site_id, "reason", amount_amount, amount_currency, transaction_date, receipt_attachment_ids, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(reason), sqlc.arg(amount_amount), sqlc.arg(amount_currency), sqlc.arg(transaction_date), sqlc.narg(receipt_attachment_ids), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetFloatExceptionStatus :one
UPDATE float_exceptions
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitFloatException :one
UPDATE float_exceptions
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActReimburseFloatException :one
UPDATE float_exceptions
SET receipt_attachment_ids = COALESCE(sqlc.narg(receipt_attachment_ids), receipt_attachment_ids), reimbursed_at = now(), status = 'reimbursed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'pending_approval', 'approved')
RETURNING *;
