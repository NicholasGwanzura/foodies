-- name: GetReceiptBook :one
SELECT * FROM receipt_books WHERE id = $1;

-- name: ListReceiptBook :many
SELECT * FROM receipt_books
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateReceiptBook :one
INSERT INTO receipt_books (site_id, book_number, first_number, last_number, custodian_employee_id, storage_location, created_by, updated_by)
VALUES (sqlc.arg(site_id), sqlc.arg(book_number), sqlc.arg(first_number), sqlc.arg(last_number), sqlc.narg(custodian_employee_id), sqlc.narg(storage_location), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateReceiptBook :one
UPDATE receipt_books
SET site_id = COALESCE(sqlc.narg(site_id), site_id), book_number = COALESCE(sqlc.narg(book_number), book_number), first_number = COALESCE(sqlc.narg(first_number), first_number), last_number = COALESCE(sqlc.narg(last_number), last_number), custodian_employee_id = COALESCE(sqlc.narg(custodian_employee_id), custodian_employee_id), storage_location = COALESCE(sqlc.narg(storage_location), storage_location), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetReceiptBookStatus :one
UPDATE receipt_books
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActIssueReceiptBook :one
UPDATE receipt_books
SET issued_at = now(), status = 'issued', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('in_storage')
RETURNING *;

-- name: ActReturnReceiptBook :one
UPDATE receipt_books
SET returned_at = now(), status = 'returned', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('in_storage', 'issued', 'missing')
RETURNING *;
