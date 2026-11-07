-- name: GetStockReturn :one
SELECT * FROM stock_returns WHERE id = $1;

-- name: ListStockReturn :many
SELECT * FROM stock_returns
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateStockReturn :one
INSERT INTO stock_returns (reference, stock_transfer_id, from_site_id, to_site_id, item_id, quantity, "reason", credit_note_reference, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.narg(stock_transfer_id), sqlc.arg(from_site_id), sqlc.arg(to_site_id), sqlc.arg(item_id), sqlc.arg(quantity), sqlc.arg(reason), sqlc.narg(credit_note_reference), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateStockReturn :one
UPDATE stock_returns
SET stock_transfer_id = COALESCE(sqlc.narg(stock_transfer_id), stock_transfer_id), from_site_id = COALESCE(sqlc.narg(from_site_id), from_site_id), to_site_id = COALESCE(sqlc.narg(to_site_id), to_site_id), item_id = COALESCE(sqlc.narg(item_id), item_id), quantity = COALESCE(sqlc.narg(quantity), quantity), "reason" = COALESCE(sqlc.narg(reason), "reason"), credit_note_reference = COALESCE(sqlc.narg(credit_note_reference), credit_note_reference), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetStockReturnStatus :one
UPDATE stock_returns
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActCloseStockReturn :one
UPDATE stock_returns
SET closed_at = now(), status = 'closed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('open', 'followed_up', 'credited')
RETURNING *;
