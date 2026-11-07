-- name: GetSupplier :one
SELECT * FROM suppliers WHERE id = $1;

-- name: ListSupplier :many
SELECT * FROM suppliers
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(supplier_type)::text IS NULL OR supplier_type = sqlc.narg(supplier_type)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR tax_clearance_expiry::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR tax_clearance_expiry::date <= sqlc.narg(to_date)::date)
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateSupplier :one
INSERT INTO suppliers (reference, legal_name, supplier_type, sap_code, categories, currency, address, contact, tin, vat_number, tax_clearance_expiry, payment_terms, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(legal_name), sqlc.arg(supplier_type), sqlc.narg(sap_code), sqlc.narg(categories), sqlc.narg(currency), sqlc.narg(address), sqlc.narg(contact), sqlc.narg(tin), sqlc.narg(vat_number), sqlc.narg(tax_clearance_expiry), sqlc.narg(payment_terms), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateSupplier :one
UPDATE suppliers
SET legal_name = COALESCE(sqlc.narg(legal_name), legal_name), supplier_type = COALESCE(sqlc.narg(supplier_type), supplier_type), sap_code = COALESCE(sqlc.narg(sap_code), sap_code), categories = COALESCE(sqlc.narg(categories), categories), currency = COALESCE(sqlc.narg(currency), currency), address = COALESCE(sqlc.narg(address), address), contact = COALESCE(sqlc.narg(contact), contact), tin = COALESCE(sqlc.narg(tin), tin), vat_number = COALESCE(sqlc.narg(vat_number), vat_number), tax_clearance_expiry = COALESCE(sqlc.narg(tax_clearance_expiry), tax_clearance_expiry), payment_terms = COALESCE(sqlc.narg(payment_terms), payment_terms), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetSupplierStatus :one
UPDATE suppliers
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitSupplier :one
UPDATE suppliers
SET status = 'pending_verification', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;

-- name: ActSuspendSupplier :one
UPDATE suppliers
SET status = 'suspended', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'pending_verification', 'active', 'blocked')
RETURNING *;

-- name: ActReinstateSupplier :one
UPDATE suppliers
SET status = 'active', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'pending_verification')
RETURNING *;
