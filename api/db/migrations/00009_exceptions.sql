-- +goose Up

CREATE TABLE ncrs (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  occurred_at timestamptz NOT NULL,
  product_id uuid,
  item_id uuid,
  batch_id uuid,
  quantity numeric(18,4),
  location text,
  deviation text NOT NULL,
  severity_level integer NOT NULL,
  containment text,
  production_stopped boolean NOT NULL DEFAULT false,
  root_cause text,
  verified_by uuid,
  status text NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'contained', 'investigating', 'action_pending', 'pending_verification', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE holds (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  product_id uuid,
  item_id uuid,
  batch_id uuid,
  batch_lot text,
  quantity numeric(18,4) NOT NULL,
  uom text NOT NULL,
  "reason" text NOT NULL,
  quarantine_location_id uuid,
  owner_employee_id uuid NOT NULL,
  ncr_id uuid,
  recall_id uuid,
  label_confirmed boolean NOT NULL DEFAULT false,
  stop_sale_posting_id uuid,
  released_at timestamptz,
  status text NOT NULL DEFAULT 'on_hold' CHECK (status IN ('on_hold', 'pending_disposition', 'released', 'rework', 'rejected', 'disposed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE dispositions (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  hold_id uuid NOT NULL,
  disposition_type text NOT NULL CHECK (disposition_type IN ('rework', 'recovery', 'reject', 'dispose', 'supplier_return')),
  quantity numeric(18,4) NOT NULL,
  "method" text NOT NULL,
  safety_assessment text NOT NULL,
  rework_batch_id uuid,
  disposal_id uuid,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'executed', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE recalls (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  kind text NOT NULL CHECK (kind IN ('internal_hold', 'withdrawal', 'recall', 'mock')),
  trigger text NOT NULL CHECK (trigger IN ('microbiological', 'foreign_body', 'chemical', 'undeclared_allergen', 'incorrect_recipe', 'labelling', 'ccp_failure', 'packaging', 'supplier_notice', 'complaint', 'lab_result', 'regulatory')),
  description text NOT NULL,
  product_id uuid,
  batch_ids uuid[],
  site_id uuid,
  reporter_employee_id uuid,
  hazard text,
  severity_rating text CHECK (severity_rating IN ('critical', 'major', 'minor')),
  likelihood text,
  customer_exposure text,
  distribution_summary text,
  decision text CHECK (decision IN ('internal_hold', 'withdrawal', 'recall')),
  decision_rationale text,
  decided_at timestamptz,
  regulatory_notification_required boolean NOT NULL DEFAULT false,
  effectiveness_summary text,
  closed_at timestamptz,
  status text NOT NULL DEFAULT 'initiated' CHECK (status IN ('initiated', 'held', 'assessed', 'decided', 'retrieving', 'reconciling', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE recall_communications (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  recall_id uuid NOT NULL,
  audience text NOT NULL CHECK (audience IN ('internal', 'branch', 'customer', 'distributor', 'regulator')),
  recipient text NOT NULL,
  message text NOT NULL,
  channel text NOT NULL,
  approved_by uuid,
  sent_at timestamptz,
  acknowledged_at timestamptz,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'sent', 'acknowledged')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE recall_retrievals (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  recall_id uuid NOT NULL,
  location_site_id uuid,
  location_name text,
  batch_id uuid,
  quantity_dispatched numeric(18,4),
  quantity_returned numeric(18,4),
  returned_at timestamptz,
  condition text,
  quarantine_location_id uuid,
  status text NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'returned', 'no_stock', 'unaccounted')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE recall_reconciliations (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  recall_id uuid NOT NULL,
  produced_received numeric(18,4) NOT NULL,
  remaining_stock numeric(18,4) NOT NULL,
  distributed numeric(18,4) NOT NULL,
  sold numeric(18,4) NOT NULL,
  returned numeric(18,4) NOT NULL,
  disposed numeric(18,4) NOT NULL,
  variance numeric(18,4),
  explanation text,
  independent_reviewer_id uuid,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'reviewed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE incidents (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  incident_type text NOT NULL CHECK (incident_type IN ('security', 'she_injury', 'she_near_miss', 'fire', 'food_safety', 'cash_discrepancy', 'cit_discrepancy', 'key_lost', 'equipment', 'environmental', 'fraud', 'cctv', 'other')),
  severity_level integer NOT NULL,
  occurred_at timestamptz NOT NULL,
  discovered_at timestamptz NOT NULL,
  notified_at timestamptz,
  description text NOT NULL,
  persons_involved text,
  immediate_action text,
  details jsonb,
  reporter_employee_id uuid,
  statutory_notification_required boolean NOT NULL DEFAULT false,
  closed_at timestamptz,
  status text NOT NULL DEFAULT 'reported' CHECK (status IN ('reported', 'acknowledged', 'investigating', 'action_pending', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE complaints (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  channel text NOT NULL CHECK (channel IN ('in_person', 'phone', 'email', 'whatsapp', 'social', 'other')),
  site_id uuid NOT NULL,
  received_at timestamptz NOT NULL,
  customer_name text,
  customer_contact text,
  category text NOT NULL CHECK (category IN ('food_safety', 'foreign_object', 'service', 'product_quality', 'payment', 'staff', 'facility', 'other')),
  severity_level integer NOT NULL,
  pos_transaction_ref text,
  product_id uuid,
  batch_id uuid,
  owner_employee_id uuid,
  due_at timestamptz,
  food_safety_details jsonb,
  evidence_location text,
  ncr_id uuid,
  resolution text,
  closed_at timestamptz,
  status text NOT NULL DEFAULT 'new' CHECK (status IN ('new', 'triaged', 'investigating', 'awaiting_customer', 'resolved', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE service_recoveries (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  complaint_id uuid NOT NULL,
  remedy_type text NOT NULL CHECK (remedy_type IN ('remake', 'replacement', 'refund', 'voucher', 'compensation')),
  value_amount numeric(18,4),
  value_currency text,
  "reason" text NOT NULL,
  customer_agreed boolean NOT NULL DEFAULT false,
  pos_reference text,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'executed', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE payment_incidents (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  occurred_at timestamptz NOT NULL,
  cashier_employee_id uuid NOT NULL,
  error_type text NOT NULL CHECK (error_type IN ('wrong_currency', 'wrong_amount', 'duplicate_debit', 'wrong_tender', 'failed', 'pending', 'unknown', 'other')),
  channel text NOT NULL CHECK (channel IN ('card', 'mobile_money', 'other')),
  provider text,
  terminal_id text,
  merchant_id text,
  transaction_reference text,
  pos_transaction_ref text,
  intended_amount numeric(18,4) NOT NULL,
  intended_currency text NOT NULL,
  actual_amount numeric(18,4) NOT NULL,
  actual_currency text NOT NULL,
  card_last_four text,
  customer_contact text,
  provider_status text CHECK (provider_status IN ('authorised', 'declined', 'pending', 'voided', 'reversed', 'settled', 'unknown')),
  correction text,
  correction_reference text,
  settlement_reconciled boolean NOT NULL DEFAULT false,
  closed_at timestamptz,
  status text NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'verified', 'correction_pending', 'correction_approved', 'reconciling', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE payment_incident_updates (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  payment_incident_id uuid NOT NULL,
  contacted_at timestamptz NOT NULL,
  "method" text NOT NULL,
  content text NOT NULL,
  confirmed_status text,
  next_update_on date,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE pos_outages (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  started_at timestamptz NOT NULL,
  cause text NOT NULL,
  affected_systems text[],
  ict_case text,
  restoration_estimate timestamptz,
  restored_at timestamptz,
  reconciliation_note text,
  status text NOT NULL DEFAULT 'declared' CHECK (status IN ('declared', 'pending_authorisation', 'authorised', 'restored', 'reconciled', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE receipt_books (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  site_id uuid NOT NULL,
  book_number text NOT NULL,
  first_number integer NOT NULL,
  last_number integer NOT NULL,
  custodian_employee_id uuid,
  storage_location text,
  issued_to_employee_id uuid,
  issued_at timestamptz,
  returned_at timestamptz,
  status text NOT NULL DEFAULT 'in_storage' CHECK (status IN ('in_storage', 'issued', 'returned', 'missing')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE manual_receipts (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  pos_outage_id uuid NOT NULL,
  receipt_book_id uuid NOT NULL,
  receipt_number integer NOT NULL,
  issued_at timestamptz NOT NULL,
  cashier_employee_id uuid NOT NULL,
  items jsonb NOT NULL,
  amount_amount numeric(18,4) NOT NULL,
  amount_currency text NOT NULL,
  tender text NOT NULL CHECK (tender IN ('cash', 'card', 'mobile_money', 'other')),
  order_docket text,
  cancel_reason text,
  pos_transaction_ref text,
  captured_by uuid,
  status text NOT NULL DEFAULT 'issued' CHECK (status IN ('issued', 'cancelled', 'missing', 'back_captured')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE equipment_faults (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  asset_id uuid NOT NULL,
  site_id uuid,
  reported_at timestamptz NOT NULL,
  symptoms text NOT NULL,
  severity_level integer NOT NULL,
  trading_impact boolean NOT NULL DEFAULT false,
  safety_impact boolean NOT NULL DEFAULT false,
  out_of_service boolean NOT NULL DEFAULT false,
  downtime_started_at timestamptz,
  temporary_control text,
  status text NOT NULL DEFAULT 'reported' CHECK (status IN ('reported', 'assigned', 'resolved', 'cancelled')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE work_orders (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  asset_id uuid NOT NULL,
  site_id uuid,
  kind text NOT NULL CHECK (kind IN ('corrective', 'preventive')),
  equipment_fault_id uuid,
  maintenance_plan_id uuid,
  priority text NOT NULL CHECK (priority IN ('low', 'normal', 'high', 'critical')),
  description text NOT NULL,
  assigned_employee_id uuid,
  contractor text,
  repair_eta timestamptz,
  parts jsonb,
  service_report_attachment_id uuid,
  downtime_ended_at timestamptz,
  cleared_by uuid,
  invoice_id uuid,
  status text NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'assigned', 'in_progress', 'pending_clearance', 'closed', 'cancelled')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE maintenance_plans (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  asset_id uuid NOT NULL,
  task text NOT NULL,
  frequency_rule_key text NOT NULL,
  assigned_role_id uuid,
  next_due_on date,
  form_template_id uuid,
  status text NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'paused')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE INDEX ON ncrs (site_id);
CREATE INDEX ON ncrs (status);
CREATE INDEX ON holds (site_id);
CREATE INDEX ON holds (status);
CREATE INDEX ON dispositions (status);
CREATE INDEX ON dispositions (hold_id);
CREATE INDEX ON recalls (site_id);
CREATE INDEX ON recalls (status);
CREATE INDEX ON recall_communications (status);
CREATE INDEX ON recall_communications (recall_id);
CREATE INDEX ON recall_retrievals (status);
CREATE INDEX ON recall_retrievals (recall_id);
CREATE INDEX ON recall_reconciliations (status);
CREATE INDEX ON recall_reconciliations (recall_id);
CREATE INDEX ON incidents (site_id);
CREATE INDEX ON incidents (status);
CREATE INDEX ON complaints (site_id);
CREATE INDEX ON complaints (status);
CREATE INDEX ON service_recoveries (status);
CREATE INDEX ON service_recoveries (complaint_id);
CREATE INDEX ON payment_incidents (site_id);
CREATE INDEX ON payment_incidents (status);
CREATE INDEX ON payment_incident_updates (payment_incident_id);
CREATE INDEX ON pos_outages (site_id);
CREATE INDEX ON pos_outages (status);
CREATE INDEX ON receipt_books (site_id);
CREATE INDEX ON receipt_books (status);
CREATE INDEX ON manual_receipts (status);
CREATE INDEX ON manual_receipts (pos_outage_id);
CREATE INDEX ON equipment_faults (site_id);
CREATE INDEX ON equipment_faults (status);
CREATE INDEX ON work_orders (site_id);
CREATE INDEX ON work_orders (status);
CREATE INDEX ON maintenance_plans (status);

ALTER TABLE complaints ADD FOREIGN KEY (batch_id) REFERENCES batches (id);
ALTER TABLE complaints ADD FOREIGN KEY (ncr_id) REFERENCES ncrs (id);
ALTER TABLE complaints ADD FOREIGN KEY (owner_employee_id) REFERENCES employees (id);
ALTER TABLE complaints ADD FOREIGN KEY (product_id) REFERENCES products (id);
ALTER TABLE complaints ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE disposals ADD FOREIGN KEY (hold_id) REFERENCES holds (id);
ALTER TABLE dispositions ADD FOREIGN KEY (disposal_id) REFERENCES disposals (id);
ALTER TABLE dispositions ADD FOREIGN KEY (rework_batch_id) REFERENCES batches (id);
ALTER TABLE equipment_faults ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE goods_receipt_lines ADD FOREIGN KEY (hold_id) REFERENCES holds (id) ON DELETE CASCADE;
ALTER TABLE holds ADD FOREIGN KEY (batch_id) REFERENCES batches (id);
ALTER TABLE holds ADD FOREIGN KEY (item_id) REFERENCES items (id);
ALTER TABLE holds ADD FOREIGN KEY (ncr_id) REFERENCES ncrs (id);
ALTER TABLE holds ADD FOREIGN KEY (owner_employee_id) REFERENCES employees (id);
ALTER TABLE holds ADD FOREIGN KEY (product_id) REFERENCES products (id);
ALTER TABLE holds ADD FOREIGN KEY (quarantine_location_id) REFERENCES storage_locations (id);
ALTER TABLE holds ADD FOREIGN KEY (recall_id) REFERENCES recalls (id);
ALTER TABLE holds ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE holds ADD FOREIGN KEY (stop_sale_posting_id) REFERENCES postings (id);
ALTER TABLE incidents ADD FOREIGN KEY (reporter_employee_id) REFERENCES employees (id);
ALTER TABLE incidents ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE maintenance_plans ADD FOREIGN KEY (assigned_role_id) REFERENCES roles (id);
ALTER TABLE maintenance_plans ADD FOREIGN KEY (form_template_id) REFERENCES form_templates (id);
ALTER TABLE manual_receipts ADD FOREIGN KEY (amount_currency) REFERENCES currencies (code);
ALTER TABLE manual_receipts ADD FOREIGN KEY (captured_by) REFERENCES users (id);
ALTER TABLE manual_receipts ADD FOREIGN KEY (cashier_employee_id) REFERENCES employees (id);
ALTER TABLE manual_receipts ADD FOREIGN KEY (receipt_book_id) REFERENCES receipt_books (id);
ALTER TABLE ncrs ADD FOREIGN KEY (batch_id) REFERENCES batches (id);
ALTER TABLE ncrs ADD FOREIGN KEY (item_id) REFERENCES items (id);
ALTER TABLE ncrs ADD FOREIGN KEY (product_id) REFERENCES products (id);
ALTER TABLE ncrs ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE ncrs ADD FOREIGN KEY (verified_by) REFERENCES users (id);
ALTER TABLE payment_incidents ADD FOREIGN KEY (actual_currency) REFERENCES currencies (code);
ALTER TABLE payment_incidents ADD FOREIGN KEY (cashier_employee_id) REFERENCES employees (id);
ALTER TABLE payment_incidents ADD FOREIGN KEY (intended_currency) REFERENCES currencies (code);
ALTER TABLE payment_incidents ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE pos_outages ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE recall_communications ADD FOREIGN KEY (approved_by) REFERENCES users (id);
ALTER TABLE recall_reconciliations ADD FOREIGN KEY (independent_reviewer_id) REFERENCES users (id);
ALTER TABLE recall_retrievals ADD FOREIGN KEY (batch_id) REFERENCES batches (id);
ALTER TABLE recall_retrievals ADD FOREIGN KEY (location_site_id) REFERENCES sites (id);
ALTER TABLE recall_retrievals ADD FOREIGN KEY (quarantine_location_id) REFERENCES storage_locations (id);
ALTER TABLE recalls ADD FOREIGN KEY (product_id) REFERENCES products (id);
ALTER TABLE recalls ADD FOREIGN KEY (reporter_employee_id) REFERENCES employees (id);
ALTER TABLE recalls ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE receipt_books ADD FOREIGN KEY (custodian_employee_id) REFERENCES employees (id);
ALTER TABLE receipt_books ADD FOREIGN KEY (issued_to_employee_id) REFERENCES employees (id);
ALTER TABLE receipt_books ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE service_recoveries ADD FOREIGN KEY (value_currency) REFERENCES currencies (code);
ALTER TABLE work_orders ADD FOREIGN KEY (assigned_employee_id) REFERENCES employees (id);
ALTER TABLE work_orders ADD FOREIGN KEY (cleared_by) REFERENCES users (id);
ALTER TABLE work_orders ADD FOREIGN KEY (equipment_fault_id) REFERENCES equipment_faults (id);
ALTER TABLE work_orders ADD FOREIGN KEY (invoice_id) REFERENCES supplier_invoices (id);
ALTER TABLE work_orders ADD FOREIGN KEY (maintenance_plan_id) REFERENCES maintenance_plans (id);
ALTER TABLE work_orders ADD FOREIGN KEY (service_report_attachment_id) REFERENCES attachments (id);
ALTER TABLE work_orders ADD FOREIGN KEY (site_id) REFERENCES sites (id);

-- +goose Down

DROP TABLE IF EXISTS maintenance_plans CASCADE;
DROP TABLE IF EXISTS work_orders CASCADE;
DROP TABLE IF EXISTS equipment_faults CASCADE;
DROP TABLE IF EXISTS manual_receipts CASCADE;
DROP TABLE IF EXISTS receipt_books CASCADE;
DROP TABLE IF EXISTS pos_outages CASCADE;
DROP TABLE IF EXISTS payment_incident_updates CASCADE;
DROP TABLE IF EXISTS payment_incidents CASCADE;
DROP TABLE IF EXISTS service_recoveries CASCADE;
DROP TABLE IF EXISTS complaints CASCADE;
DROP TABLE IF EXISTS incidents CASCADE;
DROP TABLE IF EXISTS recall_reconciliations CASCADE;
DROP TABLE IF EXISTS recall_retrievals CASCADE;
DROP TABLE IF EXISTS recall_communications CASCADE;
DROP TABLE IF EXISTS recalls CASCADE;
DROP TABLE IF EXISTS dispositions CASCADE;
DROP TABLE IF EXISTS holds CASCADE;
DROP TABLE IF EXISTS ncrs CASCADE;
