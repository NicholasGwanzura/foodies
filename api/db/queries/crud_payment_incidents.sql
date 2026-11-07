-- name: GetPaymentIncident :one
SELECT * FROM payment_incidents WHERE id = $1;

-- name: ListPaymentIncident :many
SELECT * FROM payment_incidents
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(error_type)::text IS NULL OR error_type = sqlc.narg(error_type)::text)
  AND (sqlc.narg(cashier_employee_id)::uuid IS NULL OR cashier_employee_id = sqlc.narg(cashier_employee_id)::uuid)
  AND (sqlc.narg(from_date)::date IS NULL OR created_at::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR created_at::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreatePaymentIncident :one
INSERT INTO payment_incidents (reference, site_id, occurred_at, cashier_employee_id, error_type, channel, provider, terminal_id, merchant_id, transaction_reference, pos_transaction_ref, intended_amount, intended_currency, actual_amount, actual_currency, card_last_four, customer_contact, correction, correction_reference, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(occurred_at), sqlc.arg(cashier_employee_id), sqlc.arg(error_type), sqlc.arg(channel), sqlc.narg(provider), sqlc.narg(terminal_id), sqlc.narg(merchant_id), sqlc.narg(transaction_reference), sqlc.narg(pos_transaction_ref), sqlc.arg(intended_amount), sqlc.arg(intended_currency), sqlc.arg(actual_amount), sqlc.arg(actual_currency), sqlc.narg(card_last_four), sqlc.narg(customer_contact), sqlc.narg(correction), sqlc.narg(correction_reference), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdatePaymentIncident :one
UPDATE payment_incidents
SET site_id = COALESCE(sqlc.narg(site_id), site_id), occurred_at = COALESCE(sqlc.narg(occurred_at), occurred_at), cashier_employee_id = COALESCE(sqlc.narg(cashier_employee_id), cashier_employee_id), error_type = COALESCE(sqlc.narg(error_type), error_type), channel = COALESCE(sqlc.narg(channel), channel), provider = COALESCE(sqlc.narg(provider), provider), terminal_id = COALESCE(sqlc.narg(terminal_id), terminal_id), merchant_id = COALESCE(sqlc.narg(merchant_id), merchant_id), transaction_reference = COALESCE(sqlc.narg(transaction_reference), transaction_reference), pos_transaction_ref = COALESCE(sqlc.narg(pos_transaction_ref), pos_transaction_ref), intended_amount = COALESCE(sqlc.narg(intended_amount), intended_amount), intended_currency = COALESCE(sqlc.narg(intended_currency), intended_currency), actual_amount = COALESCE(sqlc.narg(actual_amount), actual_amount), actual_currency = COALESCE(sqlc.narg(actual_currency), actual_currency), card_last_four = COALESCE(sqlc.narg(card_last_four), card_last_four), customer_contact = COALESCE(sqlc.narg(customer_contact), customer_contact), correction = COALESCE(sqlc.narg(correction), correction), correction_reference = COALESCE(sqlc.narg(correction_reference), correction_reference), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetPaymentIncidentStatus :one
UPDATE payment_incidents
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ListPaymentIncidentUpdates :many
SELECT * FROM payment_incident_updates WHERE payment_incident_id = $1 ORDER BY id;

-- name: InsertPaymentIncidentUpdates :one
INSERT INTO payment_incident_updates (payment_incident_id, contacted_at, "method", content, confirmed_status, next_update_on, created_by)
VALUES (sqlc.arg(payment_incident_id), sqlc.arg(contacted_at), sqlc.arg(method), sqlc.arg(content), sqlc.narg(confirmed_status), sqlc.narg(next_update_on), sqlc.narg(created_by))
RETURNING *;

-- name: DeletePaymentIncidentUpdates :exec
DELETE FROM payment_incident_updates WHERE payment_incident_id = $1;

-- name: ActVerifyStatusPaymentIncident :one
UPDATE payment_incidents
SET provider_status = COALESCE(sqlc.narg(provider_status), provider_status), status = 'verified', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('open')
RETURNING *;

-- name: ActProposeCorrectionPaymentIncident :one
UPDATE payment_incidents
SET correction = COALESCE(sqlc.narg(correction), correction), status = 'correction_pending', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('verified')
RETURNING *;

-- name: ActReconcilePaymentIncident :one
UPDATE payment_incidents
SET status = 'reconciling', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('correction_approved')
RETURNING *;

-- name: ActClosePaymentIncident :one
UPDATE payment_incidents
SET closed_at = now(), status = 'closed', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('open', 'verified', 'correction_pending', 'correction_approved', 'reconciling')
RETURNING *;
