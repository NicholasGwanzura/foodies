-- name: GetDisposal :one
SELECT * FROM disposals WHERE id = $1;

-- name: ListDisposal :many
SELECT * FROM disposals
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(hold_id)::uuid IS NULL OR hold_id = sqlc.narg(hold_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateDisposal :one
INSERT INTO disposals (reference, site_id, item_id, batch_id, hold_id, quantity, value_amount, value_currency, "reason", "method", location, witness_employee_id, stock_adjustment_id, evidence_attachment_ids, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.narg(item_id), sqlc.narg(batch_id), sqlc.narg(hold_id), sqlc.arg(quantity), sqlc.narg(value_amount), sqlc.narg(value_currency), sqlc.arg(reason), sqlc.arg(method), sqlc.narg(location), sqlc.arg(witness_employee_id), sqlc.narg(stock_adjustment_id), sqlc.narg(evidence_attachment_ids), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateDisposal :one
UPDATE disposals
SET site_id = COALESCE(sqlc.narg(site_id), site_id), item_id = COALESCE(sqlc.narg(item_id), item_id), batch_id = COALESCE(sqlc.narg(batch_id), batch_id), hold_id = COALESCE(sqlc.narg(hold_id), hold_id), quantity = COALESCE(sqlc.narg(quantity), quantity), value_amount = COALESCE(sqlc.narg(value_amount), value_amount), value_currency = COALESCE(sqlc.narg(value_currency), value_currency), "reason" = COALESCE(sqlc.narg(reason), "reason"), "method" = COALESCE(sqlc.narg(method), "method"), location = COALESCE(sqlc.narg(location), location), witness_employee_id = COALESCE(sqlc.narg(witness_employee_id), witness_employee_id), stock_adjustment_id = COALESCE(sqlc.narg(stock_adjustment_id), stock_adjustment_id), evidence_attachment_ids = COALESCE(sqlc.narg(evidence_attachment_ids), evidence_attachment_ids), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetDisposalStatus :one
UPDATE disposals
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitDisposal :one
UPDATE disposals
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActRecordDisposalDisposal :one
UPDATE disposals
SET disposed_at = COALESCE(sqlc.narg(disposed_at), disposed_at), status = 'disposed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('approved')
RETURNING *;
