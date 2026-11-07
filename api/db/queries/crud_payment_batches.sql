-- name: GetPaymentBatch :one
SELECT * FROM payment_batches WHERE id = $1;

-- name: ListPaymentBatch :many
SELECT * FROM payment_batches
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR run_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR run_date::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreatePaymentBatch :one
INSERT INTO payment_batches (reference, run_date, currency, requisition_ids, bank_reference, proof_attachment_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(run_date), sqlc.arg(currency), sqlc.arg(requisition_ids), sqlc.narg(bank_reference), sqlc.narg(proof_attachment_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdatePaymentBatch :one
UPDATE payment_batches
SET run_date = COALESCE(sqlc.narg(run_date), run_date), currency = COALESCE(sqlc.narg(currency), currency), requisition_ids = COALESCE(sqlc.narg(requisition_ids), requisition_ids), bank_reference = COALESCE(sqlc.narg(bank_reference), bank_reference), proof_attachment_id = COALESCE(sqlc.narg(proof_attachment_id), proof_attachment_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetPaymentBatchStatus :one
UPDATE payment_batches
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitPaymentBatch :one
UPDATE payment_batches
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActRecordExecutionPaymentBatch :one
UPDATE payment_batches
SET bank_reference = COALESCE(sqlc.narg(bank_reference), bank_reference), proof_attachment_id = COALESCE(sqlc.narg(proof_attachment_id), proof_attachment_id), status = 'executed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('approved')
RETURNING *;
