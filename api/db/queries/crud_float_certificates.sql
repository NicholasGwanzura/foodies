-- name: GetFloatCertificate :one
SELECT * FROM float_certificates WHERE id = $1;

-- name: ListFloatCertificate :many
SELECT * FROM float_certificates
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(from_date)::date IS NULL OR effective_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR effective_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateFloatCertificate :one
INSERT INTO float_certificates (reference, site_id, brand, location, change_float_amount, change_float_currency, petty_cash_amount, petty_cash_currency, effective_date, certificate_attachment_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.narg(brand), sqlc.narg(location), sqlc.arg(change_float_amount), sqlc.arg(change_float_currency), sqlc.arg(petty_cash_amount), sqlc.arg(petty_cash_currency), sqlc.arg(effective_date), sqlc.narg(certificate_attachment_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: SetFloatCertificateStatus :one
UPDATE float_certificates
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitFloatCertificate :one
UPDATE float_certificates
SET status = 'pending_approval', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft')
RETURNING *;
