-- +goose Up

CREATE TABLE suppliers (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  legal_name text NOT NULL,
  supplier_type text NOT NULL CHECK (supplier_type IN ('local', 'foreign')),
  sap_code text,
  categories text[],
  currency text,
  address text,
  contact text,
  tin text,
  vat_number text,
  tax_clearance_expiry date,
  bank_details_verified boolean NOT NULL DEFAULT false,
  bank_details_masked text,
  payment_terms text,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_verification', 'active', 'suspended', 'blocked')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE supplier_amendments (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  supplier_id uuid NOT NULL,
  field text NOT NULL,
  old_value text,
  new_value text NOT NULL,
  evidence_attachment_id uuid,
  verified_by uuid,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE purchase_requisitions (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  department_id uuid NOT NULL,
  cost_centre_id uuid NOT NULL,
  requester_employee_id uuid NOT NULL,
  purpose text NOT NULL,
  requirement_type text NOT NULL CHECK (requirement_type IN ('trading', 'departmental', 'capex', 'emergency')),
  required_date date NOT NULL,
  budget_code text,
  yield_estimate text,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'submitted', 'returned', 'validated', 'pending_approval', 'approved', 'rejected', 'sourcing', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE purchase_requisition_lines (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  purchase_requisition_id uuid NOT NULL,
  item_id uuid,
  description text NOT NULL,
  specification text,
  quantity_on_hand numeric(18,4),
  quantity_required numeric(18,4) NOT NULL,
  uom text NOT NULL,
  estimated_unit_price_amount numeric(18,4),
  estimated_unit_price_currency text,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE quotations (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  purchase_requisition_id uuid NOT NULL,
  supplier_id uuid NOT NULL,
  quote_number text NOT NULL,
  quote_date date NOT NULL,
  valid_until date,
  total_amount numeric(18,4) NOT NULL,
  total_currency text NOT NULL,
  charges_amount numeric(18,4),
  charges_currency text,
  lead_time_days integer,
  payment_terms text,
  specification_compliant boolean NOT NULL DEFAULT false,
  selected boolean NOT NULL DEFAULT false,
  selection_reason text,
  attachment_id uuid,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE purchase_orders (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  purchase_requisition_id uuid NOT NULL,
  supplier_id uuid NOT NULL,
  order_date date,
  currency text NOT NULL,
  total_amount numeric(18,4),
  total_currency text,
  delivery_site_id uuid NOT NULL,
  delivery_date date NOT NULL,
  payment_terms text,
  special_conditions text,
  sap_po_ref text,
  confirmed_at timestamptz,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'issued', 'part_received', 'received', 'closed', 'cancelled')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE purchase_order_lines (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  purchase_order_id uuid NOT NULL,
  item_id uuid,
  description text NOT NULL,
  quantity numeric(18,4) NOT NULL,
  uom text NOT NULL,
  unit_price numeric(18,4) NOT NULL,
  line_total numeric(18,4),
  received_quantity numeric(18,4),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE goods_receipts (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  purchase_order_id uuid,
  supplier_id uuid NOT NULL,
  site_id uuid NOT NULL,
  delivery_note_number text,
  supplier_invoice_number text,
  received_at timestamptz NOT NULL,
  receiver_employee_id uuid NOT NULL,
  delivery_note_attachment_id uuid,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'inspected', 'accepted', 'part_rejected', 'rejected', 'posted')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE goods_receipt_lines (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  goods_receipt_id uuid NOT NULL,
  item_id uuid NOT NULL,
  ordered_quantity numeric(18,4),
  received_quantity numeric(18,4) NOT NULL,
  uom text NOT NULL,
  batch_lot text,
  expiry_date date,
  temperature_c numeric(18,4),
  thermometer_id text,
  pack_condition text,
  outcome text NOT NULL CHECK (outcome IN ('accepted', 'rejected', 'quarantined')),
  rejection_reason text,
  hold_id uuid,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE delivery_discrepancies (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  goods_receipt_id uuid NOT NULL,
  item_id uuid,
  discrepancy_type text NOT NULL CHECK (discrepancy_type IN ('short', 'wrong_item', 'damage', 'quality', 'expiry', 'price', 'specification')),
  description text NOT NULL,
  supplier_notified_at timestamptz,
  resolution text CHECK (resolution IN ('credit', 'replacement', 'accepted', 'returned')),
  resolution_reference text,
  closed_at timestamptz,
  status text NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'supplier_notified', 'resolved', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE supplier_invoices (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  supplier_id uuid NOT NULL,
  invoice_number text NOT NULL,
  invoice_date date NOT NULL,
  total_amount numeric(18,4) NOT NULL,
  total_currency text NOT NULL,
  vat_amount numeric(18,4),
  vat_currency text,
  due_date date,
  purchase_order_id uuid,
  goods_receipt_ids uuid[],
  invoice_received_at timestamptz NOT NULL,
  forwarded_to_finance_at timestamptz,
  qr_verified boolean NOT NULL DEFAULT false,
  match_status text CHECK (match_status IN ('unmatched', 'matched', 'exception')),
  duplicate_of_id uuid,
  attachment_id uuid,
  paper_original_location text,
  sap_doc_ref text,
  status text NOT NULL DEFAULT 'received' CHECK (status IN ('received', 'matched', 'on_hold', 'approved_exception', 'posted', 'paid', 'void')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE supplier_invoice_lines (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  supplier_invoice_id uuid NOT NULL,
  description text NOT NULL,
  quantity numeric(18,4) NOT NULL,
  unit_price numeric(18,4) NOT NULL,
  line_total numeric(18,4),
  po_line_id uuid,
  match_result text CHECK (match_result IN ('matched', 'quantity_mismatch', 'price_mismatch', 'unmatched')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE payment_requisitions (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  supplier_id uuid NOT NULL,
  invoice_ids uuid[] NOT NULL,
  total_amount numeric(18,4),
  total_currency text,
  due_date date,
  priority text CHECK (priority IN ('statutory', 'normal', 'low')),
  payment_type text NOT NULL CHECK (payment_type IN ('transfer', 'advance', 'deposit', 'cash')),
  withholding_tax_amount numeric(18,4),
  withholding_tax_currency text,
  tax_clearance_valid boolean NOT NULL DEFAULT false,
  "reason" text,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'scheduled', 'paid', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE payment_batches (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  run_date date NOT NULL,
  currency text NOT NULL,
  requisition_ids uuid[] NOT NULL,
  total_amount numeric(18,4),
  total_currency text,
  bank_reference text,
  proof_attachment_id uuid,
  executed_at timestamptz,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'executed', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE emergency_procurements (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  event text NOT NULL,
  why_unforeseeable text NOT NULL,
  continuity_impact text,
  supplier_id uuid,
  amount_amount numeric(18,4) NOT NULL,
  amount_currency text NOT NULL,
  purchased_at timestamptz,
  purchase_requisition_id uuid,
  post_review text,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'rejected', 'reviewed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE unauthorized_purchases (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  originator_employee_id uuid NOT NULL,
  supplier_name text,
  amount_amount numeric(18,4) NOT NULL,
  amount_currency text NOT NULL,
  purchased_on date NOT NULL,
  "reason" text NOT NULL,
  operations_review text,
  finance_treatment text,
  corrective_action_id uuid,
  status text NOT NULL DEFAULT 'reported' CHECK (status IN ('reported', 'under_review', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE budgets (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  financial_year integer NOT NULL,
  department_id uuid NOT NULL,
  "version" integer,
  submission_due date,
  total_amount numeric(18,4),
  total_currency text,
  board_approval_attachment_id uuid,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'submitted', 'consolidated', 'pending_approval', 'approved', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE budget_lines (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  budget_id uuid NOT NULL,
  account_code text NOT NULL,
  description text,
  "month" integer NOT NULL,
  amount numeric(18,4) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE budget_variances (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  budget_id uuid NOT NULL,
  "month" date NOT NULL,
  account_code text NOT NULL,
  budget_amount numeric(18,4) NOT NULL,
  actual_amount numeric(18,4) NOT NULL,
  variance_percent numeric(18,4),
  explanation text,
  escalated_at timestamptz,
  escalation_due date,
  status text NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'explained', 'escalated', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE budget_reallocations (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  budget_id uuid NOT NULL,
  from_account text,
  to_account text,
  amount_amount numeric(18,4) NOT NULL,
  amount_currency text NOT NULL,
  "reason" text NOT NULL,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE INDEX ON suppliers (status);
CREATE INDEX ON supplier_amendments (status);
CREATE INDEX ON supplier_amendments (supplier_id);
CREATE INDEX ON purchase_requisitions (site_id);
CREATE INDEX ON purchase_requisitions (status);
CREATE INDEX ON purchase_requisition_lines (purchase_requisition_id);
CREATE INDEX ON quotations (purchase_requisition_id);
CREATE INDEX ON purchase_orders (status);
CREATE INDEX ON purchase_order_lines (purchase_order_id);
CREATE INDEX ON goods_receipts (site_id);
CREATE INDEX ON goods_receipts (status);
CREATE INDEX ON goods_receipt_lines (goods_receipt_id);
CREATE INDEX ON delivery_discrepancies (status);
CREATE INDEX ON supplier_invoices (status);
CREATE INDEX ON supplier_invoice_lines (supplier_invoice_id);
CREATE INDEX ON payment_requisitions (status);
CREATE INDEX ON payment_batches (status);
CREATE INDEX ON emergency_procurements (site_id);
CREATE INDEX ON emergency_procurements (status);
CREATE INDEX ON unauthorized_purchases (site_id);
CREATE INDEX ON unauthorized_purchases (status);
CREATE INDEX ON budgets (status);
CREATE INDEX ON budget_lines (budget_id);
CREATE INDEX ON budget_variances (status);
CREATE INDEX ON budget_variances (budget_id);
CREATE INDEX ON budget_reallocations (status);

ALTER TABLE budget_reallocations ADD FOREIGN KEY (amount_currency) REFERENCES currencies (code);
ALTER TABLE budget_reallocations ADD FOREIGN KEY (budget_id) REFERENCES budgets (id);
ALTER TABLE budgets ADD FOREIGN KEY (board_approval_attachment_id) REFERENCES attachments (id);
ALTER TABLE budgets ADD FOREIGN KEY (department_id) REFERENCES departments (id);
ALTER TABLE budgets ADD FOREIGN KEY (total_currency) REFERENCES currencies (code);
ALTER TABLE delivery_discrepancies ADD FOREIGN KEY (goods_receipt_id) REFERENCES goods_receipts (id);
ALTER TABLE emergency_procurements ADD FOREIGN KEY (amount_currency) REFERENCES currencies (code);
ALTER TABLE emergency_procurements ADD FOREIGN KEY (purchase_requisition_id) REFERENCES purchase_requisitions (id);
ALTER TABLE emergency_procurements ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE emergency_procurements ADD FOREIGN KEY (supplier_id) REFERENCES suppliers (id);
ALTER TABLE goods_receipts ADD FOREIGN KEY (delivery_note_attachment_id) REFERENCES attachments (id);
ALTER TABLE goods_receipts ADD FOREIGN KEY (purchase_order_id) REFERENCES purchase_orders (id);
ALTER TABLE goods_receipts ADD FOREIGN KEY (receiver_employee_id) REFERENCES employees (id);
ALTER TABLE goods_receipts ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE goods_receipts ADD FOREIGN KEY (supplier_id) REFERENCES suppliers (id);
ALTER TABLE payment_batches ADD FOREIGN KEY (currency) REFERENCES currencies (code);
ALTER TABLE payment_batches ADD FOREIGN KEY (proof_attachment_id) REFERENCES attachments (id);
ALTER TABLE payment_batches ADD FOREIGN KEY (total_currency) REFERENCES currencies (code);
ALTER TABLE payment_requisitions ADD FOREIGN KEY (supplier_id) REFERENCES suppliers (id);
ALTER TABLE payment_requisitions ADD FOREIGN KEY (total_currency) REFERENCES currencies (code);
ALTER TABLE payment_requisitions ADD FOREIGN KEY (withholding_tax_currency) REFERENCES currencies (code);
ALTER TABLE purchase_orders ADD FOREIGN KEY (currency) REFERENCES currencies (code);
ALTER TABLE purchase_orders ADD FOREIGN KEY (delivery_site_id) REFERENCES sites (id);
ALTER TABLE purchase_orders ADD FOREIGN KEY (purchase_requisition_id) REFERENCES purchase_requisitions (id);
ALTER TABLE purchase_orders ADD FOREIGN KEY (supplier_id) REFERENCES suppliers (id);
ALTER TABLE purchase_orders ADD FOREIGN KEY (total_currency) REFERENCES currencies (code);
ALTER TABLE purchase_requisition_lines ADD FOREIGN KEY (estimated_unit_price_currency) REFERENCES currencies (code);
ALTER TABLE purchase_requisitions ADD FOREIGN KEY (cost_centre_id) REFERENCES cost_centres (id);
ALTER TABLE purchase_requisitions ADD FOREIGN KEY (department_id) REFERENCES departments (id);
ALTER TABLE purchase_requisitions ADD FOREIGN KEY (requester_employee_id) REFERENCES employees (id);
ALTER TABLE purchase_requisitions ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE quotations ADD FOREIGN KEY (attachment_id) REFERENCES attachments (id);
ALTER TABLE quotations ADD FOREIGN KEY (charges_currency) REFERENCES currencies (code);
ALTER TABLE quotations ADD FOREIGN KEY (supplier_id) REFERENCES suppliers (id);
ALTER TABLE quotations ADD FOREIGN KEY (total_currency) REFERENCES currencies (code);
ALTER TABLE supplier_amendments ADD FOREIGN KEY (evidence_attachment_id) REFERENCES attachments (id);
ALTER TABLE supplier_amendments ADD FOREIGN KEY (verified_by) REFERENCES users (id);
ALTER TABLE supplier_invoice_lines ADD FOREIGN KEY (po_line_id) REFERENCES purchase_order_lines (id) ON DELETE CASCADE;
ALTER TABLE supplier_invoices ADD FOREIGN KEY (attachment_id) REFERENCES attachments (id);
ALTER TABLE supplier_invoices ADD FOREIGN KEY (duplicate_of_id) REFERENCES supplier_invoices (id);
ALTER TABLE supplier_invoices ADD FOREIGN KEY (purchase_order_id) REFERENCES purchase_orders (id);
ALTER TABLE supplier_invoices ADD FOREIGN KEY (supplier_id) REFERENCES suppliers (id);
ALTER TABLE supplier_invoices ADD FOREIGN KEY (total_currency) REFERENCES currencies (code);
ALTER TABLE supplier_invoices ADD FOREIGN KEY (vat_currency) REFERENCES currencies (code);
ALTER TABLE suppliers ADD FOREIGN KEY (currency) REFERENCES currencies (code);
ALTER TABLE unauthorized_purchases ADD FOREIGN KEY (amount_currency) REFERENCES currencies (code);
ALTER TABLE unauthorized_purchases ADD FOREIGN KEY (corrective_action_id) REFERENCES actions (id);
ALTER TABLE unauthorized_purchases ADD FOREIGN KEY (originator_employee_id) REFERENCES employees (id);
ALTER TABLE unauthorized_purchases ADD FOREIGN KEY (site_id) REFERENCES sites (id);
CREATE UNIQUE INDEX supplier_invoice_unique ON supplier_invoices (supplier_id, invoice_number);

-- +goose Down

DROP TABLE IF EXISTS budget_reallocations CASCADE;
DROP TABLE IF EXISTS budget_variances CASCADE;
DROP TABLE IF EXISTS budget_lines CASCADE;
DROP TABLE IF EXISTS budgets CASCADE;
DROP TABLE IF EXISTS unauthorized_purchases CASCADE;
DROP TABLE IF EXISTS emergency_procurements CASCADE;
DROP TABLE IF EXISTS payment_batches CASCADE;
DROP TABLE IF EXISTS payment_requisitions CASCADE;
DROP TABLE IF EXISTS supplier_invoice_lines CASCADE;
DROP TABLE IF EXISTS supplier_invoices CASCADE;
DROP TABLE IF EXISTS delivery_discrepancies CASCADE;
DROP TABLE IF EXISTS goods_receipt_lines CASCADE;
DROP TABLE IF EXISTS goods_receipts CASCADE;
DROP TABLE IF EXISTS purchase_order_lines CASCADE;
DROP TABLE IF EXISTS purchase_orders CASCADE;
DROP TABLE IF EXISTS quotations CASCADE;
DROP TABLE IF EXISTS purchase_requisition_lines CASCADE;
DROP TABLE IF EXISTS purchase_requisitions CASCADE;
DROP TABLE IF EXISTS supplier_amendments CASCADE;
DROP TABLE IF EXISTS suppliers CASCADE;
