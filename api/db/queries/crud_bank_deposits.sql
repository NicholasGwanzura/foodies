-- name: GetBankDeposit :one
SELECT * FROM bank_deposits WHERE id = $1;

-- name: ListBankDeposit :many
SELECT * FROM bank_deposits
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (sqlc.narg(updated_since)::timestamptz IS NULL OR updated_at > sqlc.narg(updated_since))
  AND (sqlc.narg(q)::text IS NULL OR reference ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.narg(site_id)::uuid IS NULL OR site_id = sqlc.narg(site_id)::uuid)
  AND (sqlc.narg(status)::text[] IS NULL OR status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(trading_date)::date IS NULL OR trading_date = sqlc.narg(trading_date)::date)
  AND (sqlc.narg(route)::text IS NULL OR route = sqlc.narg(route)::text)
  AND (sqlc.narg(from_date)::date IS NULL OR trading_date::date >= sqlc.narg(from_date)::date)
  AND (sqlc.narg(to_date)::date IS NULL OR trading_date::date <= sqlc.narg(to_date)::date)
  AND (sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[]))
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);

-- name: CreateBankDeposit :one
INSERT INTO bank_deposits (reference, site_id, trading_date, route, takings_amount, takings_currency, deposit_amount_amount, deposit_amount_currency, bank, deposited_at, slip_number, finance_instruction_ref, proof_attachment_id, delay_reason, cit_transit_id, created_by, updated_by)
VALUES (sqlc.arg(reference), sqlc.arg(site_id), sqlc.arg(trading_date), sqlc.arg(route), sqlc.arg(takings_amount), sqlc.arg(takings_currency), sqlc.arg(deposit_amount_amount), sqlc.arg(deposit_amount_currency), sqlc.narg(bank), sqlc.narg(deposited_at), sqlc.narg(slip_number), sqlc.narg(finance_instruction_ref), sqlc.narg(proof_attachment_id), sqlc.narg(delay_reason), sqlc.narg(cit_transit_id), sqlc.narg(created_by), sqlc.narg(created_by))
RETURNING *;

-- name: UpdateBankDeposit :one
UPDATE bank_deposits
SET site_id = COALESCE(sqlc.narg(site_id), site_id), trading_date = COALESCE(sqlc.narg(trading_date), trading_date), route = COALESCE(sqlc.narg(route), route), takings_amount = COALESCE(sqlc.narg(takings_amount), takings_amount), takings_currency = COALESCE(sqlc.narg(takings_currency), takings_currency), deposit_amount_amount = COALESCE(sqlc.narg(deposit_amount_amount), deposit_amount_amount), deposit_amount_currency = COALESCE(sqlc.narg(deposit_amount_currency), deposit_amount_currency), bank = COALESCE(sqlc.narg(bank), bank), deposited_at = COALESCE(sqlc.narg(deposited_at), deposited_at), slip_number = COALESCE(sqlc.narg(slip_number), slip_number), finance_instruction_ref = COALESCE(sqlc.narg(finance_instruction_ref), finance_instruction_ref), proof_attachment_id = COALESCE(sqlc.narg(proof_attachment_id), proof_attachment_id), delay_reason = COALESCE(sqlc.narg(delay_reason), delay_reason), cit_transit_id = COALESCE(sqlc.narg(cit_transit_id), cit_transit_id), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND row_version = sqlc.arg(row_version)
RETURNING *;

-- name: SetBankDepositStatus :one
UPDATE bank_deposits
SET status = sqlc.arg(status), updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: ActSubmitProofBankDeposit :one
UPDATE bank_deposits
SET proof_attachment_id = COALESCE(sqlc.narg(proof_attachment_id), proof_attachment_id), status = 'proof_submitted', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('banked')
RETURNING *;

-- name: ActVerifyBankDeposit :one
UPDATE bank_deposits
SET status = 'verified', updated_at = now(), updated_by = sqlc.narg(updated_by), row_version = row_version + 1
WHERE id = sqlc.arg(id) AND status IN ('draft', 'banked', 'proof_submitted')
RETURNING *;
