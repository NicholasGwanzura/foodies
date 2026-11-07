-- +goose Up

CREATE TABLE float_certificates (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  brand text,
  location text,
  change_float_amount numeric(18,4) NOT NULL,
  change_float_currency text NOT NULL,
  petty_cash_amount numeric(18,4) NOT NULL,
  petty_cash_currency text NOT NULL,
  effective_date date NOT NULL,
  certificate_attachment_id uuid,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'active', 'superseded', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE float_issues (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  till_id uuid NOT NULL,
  cashier_employee_id uuid NOT NULL,
  business_date date NOT NULL,
  shift text NOT NULL,
  amounts jsonb NOT NULL,
  issued_at timestamptz NOT NULL,
  acknowledged_at timestamptz,
  returned_amounts jsonb,
  returned_at timestamptz,
  status text NOT NULL DEFAULT 'issued' CHECK (status IN ('issued', 'acknowledged', 'returned', 'exception')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE float_exceptions (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  "reason" text NOT NULL,
  amount_amount numeric(18,4) NOT NULL,
  amount_currency text NOT NULL,
  transaction_date date NOT NULL,
  reimbursement_due date,
  reimbursed_at timestamptz,
  receipt_attachment_ids uuid[],
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'reimbursed', 'overdue', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE till_flushes (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  till_id uuid NOT NULL,
  cashier_employee_id uuid NOT NULL,
  supervisor_employee_id uuid NOT NULL,
  security_witness_employee_id uuid,
  business_date date NOT NULL,
  shift text NOT NULL,
  threshold_reached_at timestamptz,
  flushed_at timestamptz NOT NULL,
  amount_amount numeric(18,4) NOT NULL,
  amount_currency text NOT NULL,
  denominations jsonb,
  safe_id uuid NOT NULL,
  safe_reference text,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_signatures', 'closed', 'exception')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE cash_ups (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  till_id uuid NOT NULL,
  cashier_employee_id uuid NOT NULL,
  supervisor_employee_id uuid NOT NULL,
  witness_employee_id uuid,
  business_date date NOT NULL,
  shift text NOT NULL,
  kind text NOT NULL CHECK (kind IN ('shift', 'close_of_day')),
  x_read_id uuid,
  z_report_id uuid,
  fx_rate_id uuid,
  fx_rate_value numeric(18,4),
  variance_base_amount numeric(18,4),
  variance_base_currency text,
  safe_id uuid,
  safe_deposit_seal text,
  closed_at timestamptz,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'counted', 'pending_signatures', 'pending_review', 'approved', 'returned', 'exception')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE cash_up_tenders (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  cash_up_id uuid NOT NULL,
  currency text NOT NULL,
  tender text NOT NULL CHECK (tender IN ('cash', 'card', 'mobile_money', 'other')),
  expected numeric(18,4) NOT NULL,
  actual numeric(18,4) NOT NULL,
  variance numeric(18,4),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE cash_variances (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  cash_up_id uuid NOT NULL,
  site_id uuid,
  cashier_employee_id uuid,
  variance_type text CHECK (variance_type IN ('shortage', 'overage')),
  amount_amount numeric(18,4),
  amount_currency text,
  explanation text,
  recovery_status text CHECK (recovery_status IN ('none', 'paid', 'not_applicable')),
  recovered_at timestamptz,
  recovery_reference text,
  occurrence_week integer,
  occurrence_month integer,
  amount_band text,
  hr_referred_at timestamptz,
  disciplinary_case_id uuid,
  status text NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'recovered', 'under_investigation', 'referred', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE cit_providers (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  "name" text NOT NULL,
  contract_reference text,
  insurance_limit_amount numeric(18,4),
  insurance_limit_currency text,
  active boolean NOT NULL DEFAULT true,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE cit_personnel (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  cit_provider_id uuid NOT NULL,
  full_name text NOT NULL,
  id_number text NOT NULL,
  active boolean NOT NULL DEFAULT true,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE cit_transits (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  provider_id uuid NOT NULL,
  transit_date date NOT NULL,
  total_amount numeric(18,4) NOT NULL,
  total_currency text NOT NULL,
  deposit_slip_number text,
  slip_copy_dispatched boolean NOT NULL DEFAULT false,
  slip_copy_retained boolean NOT NULL DEFAULT false,
  identity_verified_person_id uuid,
  collected_at timestamptz,
  cit_receipt_number text,
  armoured_vehicle_confirmed boolean NOT NULL DEFAULT false,
  vault_received_at timestamptz,
  vault_amount_amount numeric(18,4),
  vault_amount_currency text,
  received_by uuid,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'sealed', 'identity_verified', 'handed_over', 'in_transit', 'received', 'reconciled', 'discrepancy', 'cancelled')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE cit_bags (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  cit_transit_id uuid NOT NULL,
  bag_number text NOT NULL,
  seal_number text NOT NULL,
  amount_amount numeric(18,4) NOT NULL,
  amount_currency text NOT NULL,
  seal_condition_on_receipt text CHECK (seal_condition_on_receipt IN ('intact', 'damaged', 'mismatch')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE bank_deposits (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  trading_date date NOT NULL,
  route text NOT NULL CHECK (route IN ('cit', 'direct_bank')),
  takings_amount numeric(18,4) NOT NULL,
  takings_currency text NOT NULL,
  deposit_amount_amount numeric(18,4) NOT NULL,
  deposit_amount_currency text NOT NULL,
  bank text,
  deposited_at timestamptz,
  slip_number text,
  finance_instruction_ref text,
  proof_attachment_id uuid,
  proof_received_at timestamptz,
  delay_reason text,
  cit_transit_id uuid,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'banked', 'proof_submitted', 'verified', 'exception')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE cit_reconciliations (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  reconciliation_date date NOT NULL,
  site_id uuid NOT NULL,
  dispatched_amount numeric(18,4) NOT NULL,
  dispatched_currency text NOT NULL,
  received_amount numeric(18,4),
  received_currency text,
  deposited_amount numeric(18,4),
  deposited_currency text,
  variance_amount numeric(18,4),
  variance_currency text,
  investigation_owner_employee_id uuid,
  status text NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'reconciled', 'unresolved')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE surprise_counts (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  counted_at timestamptz NOT NULL,
  count_team uuid[] NOT NULL,
  cashier_employee_id uuid,
  pos_balance_amount numeric(18,4),
  pos_balance_currency text,
  physical_cash_amount numeric(18,4),
  physical_cash_currency text,
  variance_amount numeric(18,4),
  variance_currency text,
  findings text,
  corrective_action_id uuid,
  status text NOT NULL DEFAULT 'planned' CHECK (status IN ('planned', 'in_progress', 'completed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE petty_cash_accounts (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  site_id uuid NOT NULL,
  currency text NOT NULL,
  approved_limit_amount numeric(18,4),
  approved_limit_currency text,
  custodian_employee_id uuid,
  safe_id uuid,
  book_balance_amount numeric(18,4),
  book_balance_currency text,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE petty_cash_vouchers (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  account_id uuid NOT NULL,
  site_id uuid,
  supervisor_employee_id uuid NOT NULL,
  "reason" text NOT NULL,
  amount_amount numeric(18,4) NOT NULL,
  amount_currency text NOT NULL,
  expense_category text NOT NULL,
  recipient text NOT NULL,
  verified_by uuid,
  issued_amount_amount numeric(18,4),
  issued_amount_currency text,
  balance_before_amount numeric(18,4),
  balance_before_currency text,
  balance_after_amount numeric(18,4),
  balance_after_currency text,
  issued_at timestamptz,
  receipt_attachment_id uuid,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'issued', 'receipted', 'reimbursed', 'rejected', 'cancelled')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE petty_cash_handovers (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  account_id uuid NOT NULL,
  outgoing_employee_id uuid NOT NULL,
  incoming_employee_id uuid NOT NULL,
  book_balance_amount numeric(18,4) NOT NULL,
  book_balance_currency text NOT NULL,
  physical_balance_amount numeric(18,4) NOT NULL,
  physical_balance_currency text NOT NULL,
  variance_amount numeric(18,4),
  variance_currency text,
  handed_over_at timestamptz NOT NULL,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_signatures', 'closed', 'exception')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE petty_cash_reconciliations (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  account_id uuid NOT NULL,
  kind text NOT NULL CHECK (kind IN ('end_of_day', 'spot_check')),
  business_date date NOT NULL,
  opening_balance_amount numeric(18,4) NOT NULL,
  opening_balance_currency text NOT NULL,
  receipts_amount numeric(18,4) NOT NULL,
  receipts_currency text NOT NULL,
  disbursements_amount numeric(18,4) NOT NULL,
  disbursements_currency text NOT NULL,
  expected_closing_amount numeric(18,4),
  expected_closing_currency text,
  physical_closing_amount numeric(18,4) NOT NULL,
  physical_closing_currency text NOT NULL,
  variance_amount numeric(18,4),
  variance_currency text,
  explanation text,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'submitted', 'reviewed', 'exception')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE petty_cash_reimbursements (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  account_id uuid NOT NULL,
  voucher_ids uuid[] NOT NULL,
  total_amount numeric(18,4),
  total_currency text,
  submitted_at timestamptz,
  paid_at timestamptz,
  payment_reference text,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'submitted', 'approved', 'paid', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE INDEX ON float_certificates (site_id);
CREATE INDEX ON float_certificates (status);
CREATE INDEX ON float_issues (site_id);
CREATE INDEX ON float_issues (status);
CREATE INDEX ON float_issues (business_date);
CREATE INDEX ON float_exceptions (site_id);
CREATE INDEX ON float_exceptions (status);
CREATE INDEX ON till_flushes (site_id);
CREATE INDEX ON till_flushes (status);
CREATE INDEX ON till_flushes (business_date);
CREATE INDEX ON cash_ups (site_id);
CREATE INDEX ON cash_ups (status);
CREATE INDEX ON cash_ups (business_date);
CREATE INDEX ON cash_up_tenders (cash_up_id);
CREATE INDEX ON cash_variances (site_id);
CREATE INDEX ON cash_variances (status);
CREATE INDEX ON cit_personnel (cit_provider_id);
CREATE INDEX ON cit_transits (site_id);
CREATE INDEX ON cit_transits (status);
CREATE INDEX ON cit_bags (cit_transit_id);
CREATE INDEX ON bank_deposits (site_id);
CREATE INDEX ON bank_deposits (status);
CREATE INDEX ON cit_reconciliations (site_id);
CREATE INDEX ON cit_reconciliations (status);
CREATE INDEX ON surprise_counts (site_id);
CREATE INDEX ON surprise_counts (status);
CREATE INDEX ON petty_cash_accounts (site_id);
CREATE INDEX ON petty_cash_vouchers (site_id);
CREATE INDEX ON petty_cash_vouchers (status);
CREATE INDEX ON petty_cash_handovers (status);
CREATE INDEX ON petty_cash_reconciliations (status);
CREATE INDEX ON petty_cash_reconciliations (business_date);
CREATE INDEX ON petty_cash_reimbursements (status);

ALTER TABLE bank_deposits ADD FOREIGN KEY (cit_transit_id) REFERENCES cit_transits (id);
ALTER TABLE bank_deposits ADD FOREIGN KEY (deposit_amount_currency) REFERENCES currencies (code);
ALTER TABLE bank_deposits ADD FOREIGN KEY (proof_attachment_id) REFERENCES attachments (id);
ALTER TABLE bank_deposits ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE bank_deposits ADD FOREIGN KEY (takings_currency) REFERENCES currencies (code);
ALTER TABLE cash_up_tenders ADD FOREIGN KEY (currency) REFERENCES currencies (code) ON DELETE CASCADE;
ALTER TABLE cash_ups ADD FOREIGN KEY (cashier_employee_id) REFERENCES employees (id);
ALTER TABLE cash_ups ADD FOREIGN KEY (fx_rate_id) REFERENCES fx_rates (id);
ALTER TABLE cash_ups ADD FOREIGN KEY (safe_id) REFERENCES safes (id);
ALTER TABLE cash_ups ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE cash_ups ADD FOREIGN KEY (supervisor_employee_id) REFERENCES employees (id);
ALTER TABLE cash_ups ADD FOREIGN KEY (till_id) REFERENCES tills (id);
ALTER TABLE cash_ups ADD FOREIGN KEY (variance_base_currency) REFERENCES currencies (code);
ALTER TABLE cash_ups ADD FOREIGN KEY (witness_employee_id) REFERENCES employees (id);
ALTER TABLE cash_ups ADD FOREIGN KEY (x_read_id) REFERENCES z_reports (id);
ALTER TABLE cash_ups ADD FOREIGN KEY (z_report_id) REFERENCES z_reports (id);
ALTER TABLE cash_variances ADD FOREIGN KEY (amount_currency) REFERENCES currencies (code);
ALTER TABLE cash_variances ADD FOREIGN KEY (cash_up_id) REFERENCES cash_ups (id);
ALTER TABLE cash_variances ADD FOREIGN KEY (cashier_employee_id) REFERENCES employees (id);
ALTER TABLE cash_variances ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE cit_bags ADD FOREIGN KEY (amount_currency) REFERENCES currencies (code);
ALTER TABLE cit_providers ADD FOREIGN KEY (insurance_limit_currency) REFERENCES currencies (code);
ALTER TABLE cit_reconciliations ADD FOREIGN KEY (deposited_currency) REFERENCES currencies (code);
ALTER TABLE cit_reconciliations ADD FOREIGN KEY (dispatched_currency) REFERENCES currencies (code);
ALTER TABLE cit_reconciliations ADD FOREIGN KEY (investigation_owner_employee_id) REFERENCES employees (id);
ALTER TABLE cit_reconciliations ADD FOREIGN KEY (received_currency) REFERENCES currencies (code);
ALTER TABLE cit_reconciliations ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE cit_reconciliations ADD FOREIGN KEY (variance_currency) REFERENCES currencies (code);
ALTER TABLE cit_transits ADD FOREIGN KEY (identity_verified_person_id) REFERENCES cit_personnel (id);
ALTER TABLE cit_transits ADD FOREIGN KEY (provider_id) REFERENCES cit_providers (id);
ALTER TABLE cit_transits ADD FOREIGN KEY (received_by) REFERENCES users (id);
ALTER TABLE cit_transits ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE cit_transits ADD FOREIGN KEY (total_currency) REFERENCES currencies (code);
ALTER TABLE cit_transits ADD FOREIGN KEY (vault_amount_currency) REFERENCES currencies (code);
ALTER TABLE float_certificates ADD FOREIGN KEY (certificate_attachment_id) REFERENCES attachments (id);
ALTER TABLE float_certificates ADD FOREIGN KEY (change_float_currency) REFERENCES currencies (code);
ALTER TABLE float_certificates ADD FOREIGN KEY (petty_cash_currency) REFERENCES currencies (code);
ALTER TABLE float_certificates ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE float_exceptions ADD FOREIGN KEY (amount_currency) REFERENCES currencies (code);
ALTER TABLE float_exceptions ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE float_issues ADD FOREIGN KEY (cashier_employee_id) REFERENCES employees (id);
ALTER TABLE float_issues ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE float_issues ADD FOREIGN KEY (till_id) REFERENCES tills (id);
ALTER TABLE petty_cash_accounts ADD FOREIGN KEY (approved_limit_currency) REFERENCES currencies (code);
ALTER TABLE petty_cash_accounts ADD FOREIGN KEY (book_balance_currency) REFERENCES currencies (code);
ALTER TABLE petty_cash_accounts ADD FOREIGN KEY (currency) REFERENCES currencies (code);
ALTER TABLE petty_cash_accounts ADD FOREIGN KEY (custodian_employee_id) REFERENCES employees (id);
ALTER TABLE petty_cash_accounts ADD FOREIGN KEY (safe_id) REFERENCES safes (id);
ALTER TABLE petty_cash_accounts ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE petty_cash_handovers ADD FOREIGN KEY (account_id) REFERENCES petty_cash_accounts (id);
ALTER TABLE petty_cash_handovers ADD FOREIGN KEY (book_balance_currency) REFERENCES currencies (code);
ALTER TABLE petty_cash_handovers ADD FOREIGN KEY (incoming_employee_id) REFERENCES employees (id);
ALTER TABLE petty_cash_handovers ADD FOREIGN KEY (outgoing_employee_id) REFERENCES employees (id);
ALTER TABLE petty_cash_handovers ADD FOREIGN KEY (physical_balance_currency) REFERENCES currencies (code);
ALTER TABLE petty_cash_handovers ADD FOREIGN KEY (variance_currency) REFERENCES currencies (code);
ALTER TABLE petty_cash_reconciliations ADD FOREIGN KEY (account_id) REFERENCES petty_cash_accounts (id);
ALTER TABLE petty_cash_reconciliations ADD FOREIGN KEY (disbursements_currency) REFERENCES currencies (code);
ALTER TABLE petty_cash_reconciliations ADD FOREIGN KEY (expected_closing_currency) REFERENCES currencies (code);
ALTER TABLE petty_cash_reconciliations ADD FOREIGN KEY (opening_balance_currency) REFERENCES currencies (code);
ALTER TABLE petty_cash_reconciliations ADD FOREIGN KEY (physical_closing_currency) REFERENCES currencies (code);
ALTER TABLE petty_cash_reconciliations ADD FOREIGN KEY (receipts_currency) REFERENCES currencies (code);
ALTER TABLE petty_cash_reconciliations ADD FOREIGN KEY (variance_currency) REFERENCES currencies (code);
ALTER TABLE petty_cash_reimbursements ADD FOREIGN KEY (account_id) REFERENCES petty_cash_accounts (id);
ALTER TABLE petty_cash_reimbursements ADD FOREIGN KEY (total_currency) REFERENCES currencies (code);
ALTER TABLE petty_cash_vouchers ADD FOREIGN KEY (account_id) REFERENCES petty_cash_accounts (id);
ALTER TABLE petty_cash_vouchers ADD FOREIGN KEY (amount_currency) REFERENCES currencies (code);
ALTER TABLE petty_cash_vouchers ADD FOREIGN KEY (balance_after_currency) REFERENCES currencies (code);
ALTER TABLE petty_cash_vouchers ADD FOREIGN KEY (balance_before_currency) REFERENCES currencies (code);
ALTER TABLE petty_cash_vouchers ADD FOREIGN KEY (issued_amount_currency) REFERENCES currencies (code);
ALTER TABLE petty_cash_vouchers ADD FOREIGN KEY (receipt_attachment_id) REFERENCES attachments (id);
ALTER TABLE petty_cash_vouchers ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE petty_cash_vouchers ADD FOREIGN KEY (supervisor_employee_id) REFERENCES employees (id);
ALTER TABLE petty_cash_vouchers ADD FOREIGN KEY (verified_by) REFERENCES users (id);
ALTER TABLE surprise_counts ADD FOREIGN KEY (cashier_employee_id) REFERENCES employees (id);
ALTER TABLE surprise_counts ADD FOREIGN KEY (corrective_action_id) REFERENCES actions (id);
ALTER TABLE surprise_counts ADD FOREIGN KEY (physical_cash_currency) REFERENCES currencies (code);
ALTER TABLE surprise_counts ADD FOREIGN KEY (pos_balance_currency) REFERENCES currencies (code);
ALTER TABLE surprise_counts ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE surprise_counts ADD FOREIGN KEY (variance_currency) REFERENCES currencies (code);
ALTER TABLE till_flushes ADD FOREIGN KEY (amount_currency) REFERENCES currencies (code);
ALTER TABLE till_flushes ADD FOREIGN KEY (cashier_employee_id) REFERENCES employees (id);
ALTER TABLE till_flushes ADD FOREIGN KEY (safe_id) REFERENCES safes (id);
ALTER TABLE till_flushes ADD FOREIGN KEY (security_witness_employee_id) REFERENCES employees (id);
ALTER TABLE till_flushes ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE till_flushes ADD FOREIGN KEY (supervisor_employee_id) REFERENCES employees (id);
ALTER TABLE till_flushes ADD FOREIGN KEY (till_id) REFERENCES tills (id);
CREATE UNIQUE INDEX cit_bags_seal_unique ON cit_bags (seal_number);

-- +goose Down

DROP TABLE IF EXISTS petty_cash_reimbursements CASCADE;
DROP TABLE IF EXISTS petty_cash_reconciliations CASCADE;
DROP TABLE IF EXISTS petty_cash_handovers CASCADE;
DROP TABLE IF EXISTS petty_cash_vouchers CASCADE;
DROP TABLE IF EXISTS petty_cash_accounts CASCADE;
DROP TABLE IF EXISTS surprise_counts CASCADE;
DROP TABLE IF EXISTS cit_reconciliations CASCADE;
DROP TABLE IF EXISTS bank_deposits CASCADE;
DROP TABLE IF EXISTS cit_bags CASCADE;
DROP TABLE IF EXISTS cit_transits CASCADE;
DROP TABLE IF EXISTS cit_personnel CASCADE;
DROP TABLE IF EXISTS cit_providers CASCADE;
DROP TABLE IF EXISTS cash_variances CASCADE;
DROP TABLE IF EXISTS cash_up_tenders CASCADE;
DROP TABLE IF EXISTS cash_ups CASCADE;
DROP TABLE IF EXISTS till_flushes CASCADE;
DROP TABLE IF EXISTS float_exceptions CASCADE;
DROP TABLE IF EXISTS float_issues CASCADE;
DROP TABLE IF EXISTS float_certificates CASCADE;
