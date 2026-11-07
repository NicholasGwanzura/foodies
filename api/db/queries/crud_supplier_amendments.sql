-- name: GetSupplierAmendment :one
SELECT * FROM supplier_amendments WHERE id = $1;

-- name: ListSupplierAmendment :many
SELECT * FROM supplier_amendments
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateSupplierAmendment :one
INSERT INTO supplier_amendments (supplier_id, field, new_value, evidence_attachment_id, created_by, updated_by)
VALUES (sqlc.arg(supplier_id), sqlc.arg(field), sqlc.arg(new_value), sqlc.narg(evidence_attachment_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetSupplierAmendmentStatus :one
UPDATE supplier_amendments
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitSupplierAmendment :one
UPDATE supplier_amendments
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
