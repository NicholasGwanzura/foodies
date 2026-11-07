-- +goose Up

CREATE TABLE items (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  code text NOT NULL,
  "name" text NOT NULL,
  category text NOT NULL,
  stock_class text NOT NULL CHECK (stock_class IN ('dry', 'chilled', 'frozen', 'packaging', 'chemical', 'consumable')),
  uom text NOT NULL,
  critical boolean NOT NULL DEFAULT false,
  rotation text NOT NULL CHECK (rotation IN ('fifo', 'fefo')),
  shelf_life_days integer,
  gaap_ref text,
  active boolean NOT NULL DEFAULT true,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE (code)
);

CREATE TABLE stock_balances (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  site_id uuid,
  item_id uuid,
  quantity numeric(18,4),
  uom text,
  held_quantity numeric(18,4),
  average_daily_usage numeric(18,4),
  days_of_cover numeric(18,4),
  as_of timestamptz,
  source_run_id uuid,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE stock_requisitions (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  requesting_department_id uuid,
  requester_employee_id uuid NOT NULL,
  issued_by uuid,
  issued_at timestamptz,
  acknowledged_at timestamptz,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'submitted', 'approved', 'issued', 'acknowledged', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE stock_requisition_lines (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  stock_requisition_id uuid NOT NULL,
  item_id uuid NOT NULL,
  requested_quantity numeric(18,4) NOT NULL,
  issued_quantity numeric(18,4),
  uom text NOT NULL,
  batch_lot text,
  bin_card_ref text,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE stock_transfers (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  from_site_id uuid NOT NULL,
  to_site_id uuid NOT NULL,
  requested_by uuid,
  vehicle text,
  dispatched_at timestamptz,
  dispatch_security_employee_id uuid,
  arrived_at timestamptz,
  receiving_security_employee_id uuid,
  received_by uuid,
  tba_followup_due timestamptz,
  invoice_attachment_id uuid,
  status text NOT NULL DEFAULT 'requested' CHECK (status IN ('requested', 'accepted', 'rejected', 'dispatched', 'received', 'discrepancy', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE stock_transfer_lines (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  stock_transfer_id uuid NOT NULL,
  item_id uuid NOT NULL,
  requested_quantity numeric(18,4) NOT NULL,
  approved_quantity numeric(18,4),
  dispatched_quantity numeric(18,4),
  received_quantity numeric(18,4),
  uom text NOT NULL,
  batch_lot text,
  condition text,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE stock_returns (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  stock_transfer_id uuid,
  from_site_id uuid NOT NULL,
  to_site_id uuid NOT NULL,
  item_id uuid NOT NULL,
  quantity numeric(18,4) NOT NULL,
  "reason" text NOT NULL,
  followup_due timestamptz,
  credit_note_reference text,
  closure_due timestamptz,
  closed_at timestamptz,
  status text NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'followed_up', 'credited', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE stocktakes (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  kind text NOT NULL CHECK (kind IN ('shift_open', 'shift_close', 'close_of_business', 'cycle', 'month_end', 'year_end')),
  count_date date NOT NULL,
  shift text,
  finance_instruction_ref text,
  alternative_control text,
  system_capture_ref text,
  status text NOT NULL DEFAULT 'planned' CHECK (status IN ('planned', 'counting', 'submitted', 'reviewed', 'approved', 'posted', 'skipped')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE stocktake_lines (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  stocktake_id uuid NOT NULL,
  item_id uuid NOT NULL,
  location_id uuid,
  sheet_number text,
  counter_employee_id uuid,
  count_1 numeric(18,4),
  count_2 numeric(18,4),
  recount numeric(18,4),
  drs_quantity numeric(18,4),
  gaap_quantity numeric(18,4),
  bin_card_quantity numeric(18,4),
  variance_quantity numeric(18,4),
  variance_value_amount numeric(18,4),
  variance_value_currency text,
  obsolete boolean NOT NULL DEFAULT false,
  "comment" text,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE stock_variances (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  stocktake_id uuid NOT NULL,
  item_id uuid NOT NULL,
  variance_quantity numeric(18,4) NOT NULL,
  variance_value_amount numeric(18,4),
  variance_value_currency text,
  documents_checked text[],
  root_cause text,
  classification text CHECK (classification IN ('timing', 'capture', 'actual_loss')),
  responsible_employee_id uuid,
  recovery text,
  status text NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'investigating', 'resolved', 'escalated', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE stock_adjustments (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  item_id uuid NOT NULL,
  quantity numeric(18,4) NOT NULL,
  value_amount numeric(18,4),
  value_currency text,
  reason_tag text NOT NULL CHECK (reason_tag IN ('wastage', 'damage', 'off_cuts', 'donation', 'expired_inventory', 'variance', 'other')),
  root_cause text,
  stock_variance_id uuid,
  supporting_attachment_id uuid,
  system_adjustment_ref text,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'posted', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE waste_records (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  occurred_at timestamptz NOT NULL,
  item_id uuid,
  product_id uuid,
  batch_id uuid,
  quantity numeric(18,4) NOT NULL,
  uom text NOT NULL,
  category text NOT NULL CHECK (category IN ('trimming', 'cooking_loss', 'spoilage', 'expiry', 'damage', 'over_production', 'reject', 'over_portioning', 'contamination', 'other')),
  cause text,
  value_amount numeric(18,4),
  value_currency text,
  photo_attachment_id uuid,
  disposition text,
  status text NOT NULL DEFAULT 'recorded' CHECK (status IN ('recorded', 'reviewed', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE disposals (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  item_id uuid,
  batch_id uuid,
  hold_id uuid,
  quantity numeric(18,4) NOT NULL,
  value_amount numeric(18,4),
  value_currency text,
  "reason" text NOT NULL,
  "method" text NOT NULL,
  location text,
  witness_employee_id uuid NOT NULL,
  stock_adjustment_id uuid,
  disposed_at timestamptz,
  evidence_attachment_ids uuid[],
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'disposed', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE recipes (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  code text NOT NULL,
  "name" text NOT NULL,
  product_id uuid,
  owner_role_id uuid,
  current_version_id uuid,
  active boolean NOT NULL DEFAULT true,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE (code)
);

CREATE TABLE recipe_versions (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  recipe_id uuid NOT NULL,
  "version" integer,
  batch_size numeric(18,4) NOT NULL,
  uom text NOT NULL,
  standard_output numeric(18,4) NOT NULL,
  standard_yield_percent numeric(18,4) NOT NULL,
  normal_waste_percent numeric(18,4),
  method_document_id uuid,
  effective_from date,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'active', 'superseded')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE recipe_ingredients (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  recipe_version_id uuid NOT NULL,
  item_id uuid NOT NULL,
  quantity numeric(18,4) NOT NULL,
  uom text NOT NULL,
  allergen_tags text[],
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE kitchen_orders (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  cycle text NOT NULL CHECK (cycle IN ('morning', 'evening')),
  order_date date NOT NULL,
  supervisor_employee_id uuid NOT NULL,
  deadline_at timestamptz,
  submitted_at timestamptz,
  late boolean NOT NULL DEFAULT false,
  late_reason text,
  expected_submission_at timestamptz,
  acknowledged_by uuid,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'submitted', 'late_pending', 'confirmed', 'planned', 'dispatched', 'received', 'cancelled')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE kitchen_order_lines (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  kitchen_order_id uuid NOT NULL,
  product_id uuid NOT NULL,
  quantity numeric(18,4) NOT NULL,
  uom text NOT NULL,
  stock_on_hand numeric(18,4),
  par_level numeric(18,4),
  expected_sales numeric(18,4),
  unusual_flag boolean NOT NULL DEFAULT false,
  confirmed_quantity numeric(18,4),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE production_plans (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  plan_date date NOT NULL,
  cycle text NOT NULL CHECK (cycle IN ('morning', 'evening', 'day')),
  line_clearance_submission_id uuid,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'confirmed', 'in_progress', 'completed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE production_plan_lines (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  production_plan_id uuid NOT NULL,
  product_id uuid NOT NULL,
  production_line_id uuid,
  planned_quantity numeric(18,4) NOT NULL,
  uom text NOT NULL,
  material_shortage text,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE batches (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  batch_number text NOT NULL,
  product_id uuid NOT NULL,
  recipe_version_id uuid,
  site_id uuid NOT NULL,
  production_line_id uuid,
  production_plan_id uuid,
  production_date date NOT NULL,
  shift text,
  operator_employee_id uuid NOT NULL,
  input_quantity numeric(18,4),
  planned_output numeric(18,4),
  actual_output numeric(18,4),
  uom text NOT NULL,
  actual_yield_percent numeric(18,4),
  standard_yield_percent numeric(18,4),
  yield_variance_percent numeric(18,4),
  qc_status text CHECK (qc_status IN ('pending', 'released', 'hold', 'rework', 'rejected', 'disposed')),
  expiry_date date,
  holding_condition text,
  gaap_transaction_ref text,
  status text NOT NULL DEFAULT 'in_progress' CHECK (status IN ('in_progress', 'completed', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE batch_ingredients (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  batch_id uuid NOT NULL,
  item_id uuid NOT NULL,
  supplier_lot text NOT NULL,
  goods_receipt_line_id uuid,
  planned_quantity numeric(18,4),
  issued_quantity numeric(18,4),
  actual_quantity numeric(18,4),
  uom text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE dispatch_notes (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  from_site_id uuid NOT NULL,
  to_site_id uuid NOT NULL,
  kitchen_order_id uuid,
  vehicle text NOT NULL,
  driver_employee_id uuid NOT NULL,
  loading_temperature_c numeric(18,4),
  dispatched_at timestamptz,
  delivery_temperature_c numeric(18,4),
  delivered_at timestamptz,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'picked', 'verified', 'security_cleared', 'in_transit', 'delivered', 'discrepancy', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE dispatch_note_lines (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  dispatch_note_id uuid NOT NULL,
  product_id uuid NOT NULL,
  batch_id uuid NOT NULL,
  ordered_quantity numeric(18,4),
  picked_quantity numeric(18,4) NOT NULL,
  delivered_quantity numeric(18,4),
  expiry_date date,
  condition text,
  raw_or_rte text CHECK (raw_or_rte IN ('raw', 'rte')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE products (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  code text NOT NULL,
  "name" text NOT NULL,
  category text NOT NULL,
  description text,
  active boolean NOT NULL DEFAULT true,
  branch_site_ids uuid[],
  recipe_id uuid,
  menu_version text,
  allergens text[],
  allergens_verified_by uuid,
  portion_standard text,
  storage_and_shelf_life text,
  image_attachment_id uuid,
  card jsonb,
  faq jsonb,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE (code)
);

CREATE TABLE product_changes (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  product_id uuid NOT NULL,
  change_type text NOT NULL CHECK (change_type IN ('new', 'recipe', 'price', 'menu', 'withdrawal', 'seasonal')),
  description text NOT NULL,
  effective_on date NOT NULL,
  pos_change_reference text,
  affected_role_ids uuid[],
  training_request_id uuid,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'effective', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE cost_sheets (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  recipe_version_id uuid NOT NULL,
  currency text NOT NULL,
  ingredient_cost numeric(18,4) NOT NULL,
  labour_cost numeric(18,4),
  overhead_cost numeric(18,4),
  cost_per_unit numeric(18,4),
  selling_price numeric(18,4),
  gross_margin_percent numeric(18,4),
  review_trigger text,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'superseded')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE npd_projects (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  "name" text NOT NULL,
  business_case text NOT NULL,
  current_gate text CHECK (current_gate IN ('G0', 'G1', 'G2', 'G3', 'G4', 'G5')),
  owner_employee_id uuid,
  product_id uuid,
  launched_on date,
  status text NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'on_hold', 'launched', 'cancelled')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE INDEX ON stock_balances (site_id);
CREATE INDEX ON stock_requisitions (site_id);
CREATE INDEX ON stock_requisitions (status);
CREATE INDEX ON stock_requisition_lines (stock_requisition_id);
CREATE INDEX ON stock_transfers (status);
CREATE INDEX ON stock_transfer_lines (stock_transfer_id);
CREATE INDEX ON stock_returns (status);
CREATE INDEX ON stocktakes (site_id);
CREATE INDEX ON stocktakes (status);
CREATE INDEX ON stocktake_lines (stocktake_id);
CREATE INDEX ON stock_variances (status);
CREATE INDEX ON stock_adjustments (site_id);
CREATE INDEX ON stock_adjustments (status);
CREATE INDEX ON waste_records (site_id);
CREATE INDEX ON waste_records (status);
CREATE INDEX ON disposals (site_id);
CREATE INDEX ON disposals (status);
CREATE INDEX ON recipe_versions (status);
CREATE INDEX ON recipe_versions (recipe_id);
CREATE INDEX ON recipe_ingredients (recipe_version_id);
CREATE INDEX ON kitchen_orders (site_id);
CREATE INDEX ON kitchen_orders (status);
CREATE INDEX ON kitchen_order_lines (kitchen_order_id);
CREATE INDEX ON production_plans (site_id);
CREATE INDEX ON production_plans (status);
CREATE INDEX ON production_plan_lines (production_plan_id);
CREATE INDEX ON batches (site_id);
CREATE INDEX ON batches (status);
CREATE INDEX ON batch_ingredients (batch_id);
CREATE INDEX ON dispatch_notes (status);
CREATE INDEX ON dispatch_note_lines (dispatch_note_id);
CREATE INDEX ON product_changes (status);
CREATE INDEX ON cost_sheets (status);
CREATE INDEX ON npd_projects (status);

ALTER TABLE batch_ingredients ADD FOREIGN KEY (goods_receipt_line_id) REFERENCES goods_receipt_lines (id) ON DELETE CASCADE;
ALTER TABLE batch_ingredients ADD FOREIGN KEY (item_id) REFERENCES items (id) ON DELETE CASCADE;
ALTER TABLE batches ADD FOREIGN KEY (operator_employee_id) REFERENCES employees (id);
ALTER TABLE batches ADD FOREIGN KEY (product_id) REFERENCES products (id);
ALTER TABLE batches ADD FOREIGN KEY (production_line_id) REFERENCES production_lines (id);
ALTER TABLE batches ADD FOREIGN KEY (production_plan_id) REFERENCES production_plans (id);
ALTER TABLE batches ADD FOREIGN KEY (recipe_version_id) REFERENCES recipe_versions (id);
ALTER TABLE batches ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE cost_sheets ADD FOREIGN KEY (currency) REFERENCES currencies (code);
ALTER TABLE cost_sheets ADD FOREIGN KEY (recipe_version_id) REFERENCES recipe_versions (id);
ALTER TABLE delivery_discrepancies ADD FOREIGN KEY (item_id) REFERENCES items (id);
ALTER TABLE dispatch_note_lines ADD FOREIGN KEY (batch_id) REFERENCES batches (id) ON DELETE CASCADE;
ALTER TABLE dispatch_note_lines ADD FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE;
ALTER TABLE dispatch_notes ADD FOREIGN KEY (driver_employee_id) REFERENCES employees (id);
ALTER TABLE dispatch_notes ADD FOREIGN KEY (from_site_id) REFERENCES sites (id);
ALTER TABLE dispatch_notes ADD FOREIGN KEY (kitchen_order_id) REFERENCES kitchen_orders (id);
ALTER TABLE dispatch_notes ADD FOREIGN KEY (to_site_id) REFERENCES sites (id);
ALTER TABLE disposals ADD FOREIGN KEY (batch_id) REFERENCES batches (id);
ALTER TABLE disposals ADD FOREIGN KEY (item_id) REFERENCES items (id);
ALTER TABLE disposals ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE disposals ADD FOREIGN KEY (stock_adjustment_id) REFERENCES stock_adjustments (id);
ALTER TABLE disposals ADD FOREIGN KEY (value_currency) REFERENCES currencies (code);
ALTER TABLE disposals ADD FOREIGN KEY (witness_employee_id) REFERENCES employees (id);
ALTER TABLE goods_receipt_lines ADD FOREIGN KEY (item_id) REFERENCES items (id) ON DELETE CASCADE;
ALTER TABLE kitchen_order_lines ADD FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE;
ALTER TABLE kitchen_orders ADD FOREIGN KEY (acknowledged_by) REFERENCES users (id);
ALTER TABLE kitchen_orders ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE kitchen_orders ADD FOREIGN KEY (supervisor_employee_id) REFERENCES employees (id);
ALTER TABLE npd_projects ADD FOREIGN KEY (owner_employee_id) REFERENCES employees (id);
ALTER TABLE npd_projects ADD FOREIGN KEY (product_id) REFERENCES products (id);
ALTER TABLE product_changes ADD FOREIGN KEY (product_id) REFERENCES products (id);
ALTER TABLE production_plan_lines ADD FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE;
ALTER TABLE production_plan_lines ADD FOREIGN KEY (production_line_id) REFERENCES production_lines (id) ON DELETE CASCADE;
ALTER TABLE production_plans ADD FOREIGN KEY (line_clearance_submission_id) REFERENCES form_submissions (id);
ALTER TABLE production_plans ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE products ADD FOREIGN KEY (allergens_verified_by) REFERENCES users (id);
ALTER TABLE products ADD FOREIGN KEY (image_attachment_id) REFERENCES attachments (id);
ALTER TABLE products ADD FOREIGN KEY (recipe_id) REFERENCES recipes (id);
ALTER TABLE purchase_order_lines ADD FOREIGN KEY (item_id) REFERENCES items (id) ON DELETE CASCADE;
ALTER TABLE purchase_requisition_lines ADD FOREIGN KEY (item_id) REFERENCES items (id) ON DELETE CASCADE;
ALTER TABLE recipe_ingredients ADD FOREIGN KEY (item_id) REFERENCES items (id) ON DELETE CASCADE;
ALTER TABLE recipes ADD FOREIGN KEY (current_version_id) REFERENCES recipe_versions (id);
ALTER TABLE recipes ADD FOREIGN KEY (owner_role_id) REFERENCES roles (id);
ALTER TABLE recipes ADD FOREIGN KEY (product_id) REFERENCES products (id);
ALTER TABLE stock_adjustments ADD FOREIGN KEY (item_id) REFERENCES items (id);
ALTER TABLE stock_adjustments ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE stock_adjustments ADD FOREIGN KEY (stock_variance_id) REFERENCES stock_variances (id);
ALTER TABLE stock_adjustments ADD FOREIGN KEY (supporting_attachment_id) REFERENCES attachments (id);
ALTER TABLE stock_adjustments ADD FOREIGN KEY (value_currency) REFERENCES currencies (code);
ALTER TABLE stock_balances ADD FOREIGN KEY (item_id) REFERENCES items (id);
ALTER TABLE stock_balances ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE stock_balances ADD FOREIGN KEY (source_run_id) REFERENCES sync_runs (id);
ALTER TABLE stock_requisition_lines ADD FOREIGN KEY (item_id) REFERENCES items (id) ON DELETE CASCADE;
ALTER TABLE stock_requisitions ADD FOREIGN KEY (issued_by) REFERENCES users (id);
ALTER TABLE stock_requisitions ADD FOREIGN KEY (requester_employee_id) REFERENCES employees (id);
ALTER TABLE stock_requisitions ADD FOREIGN KEY (requesting_department_id) REFERENCES departments (id);
ALTER TABLE stock_requisitions ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE stock_returns ADD FOREIGN KEY (from_site_id) REFERENCES sites (id);
ALTER TABLE stock_returns ADD FOREIGN KEY (item_id) REFERENCES items (id);
ALTER TABLE stock_returns ADD FOREIGN KEY (stock_transfer_id) REFERENCES stock_transfers (id);
ALTER TABLE stock_returns ADD FOREIGN KEY (to_site_id) REFERENCES sites (id);
ALTER TABLE stock_transfer_lines ADD FOREIGN KEY (item_id) REFERENCES items (id) ON DELETE CASCADE;
ALTER TABLE stock_transfers ADD FOREIGN KEY (dispatch_security_employee_id) REFERENCES employees (id);
ALTER TABLE stock_transfers ADD FOREIGN KEY (from_site_id) REFERENCES sites (id);
ALTER TABLE stock_transfers ADD FOREIGN KEY (invoice_attachment_id) REFERENCES attachments (id);
ALTER TABLE stock_transfers ADD FOREIGN KEY (received_by) REFERENCES users (id);
ALTER TABLE stock_transfers ADD FOREIGN KEY (receiving_security_employee_id) REFERENCES employees (id);
ALTER TABLE stock_transfers ADD FOREIGN KEY (requested_by) REFERENCES users (id);
ALTER TABLE stock_transfers ADD FOREIGN KEY (to_site_id) REFERENCES sites (id);
ALTER TABLE stock_variances ADD FOREIGN KEY (item_id) REFERENCES items (id);
ALTER TABLE stock_variances ADD FOREIGN KEY (responsible_employee_id) REFERENCES employees (id);
ALTER TABLE stock_variances ADD FOREIGN KEY (stocktake_id) REFERENCES stocktakes (id);
ALTER TABLE stock_variances ADD FOREIGN KEY (variance_value_currency) REFERENCES currencies (code);
ALTER TABLE stocktake_lines ADD FOREIGN KEY (counter_employee_id) REFERENCES employees (id) ON DELETE CASCADE;
ALTER TABLE stocktake_lines ADD FOREIGN KEY (item_id) REFERENCES items (id) ON DELETE CASCADE;
ALTER TABLE stocktake_lines ADD FOREIGN KEY (location_id) REFERENCES storage_locations (id) ON DELETE CASCADE;
ALTER TABLE stocktake_lines ADD FOREIGN KEY (variance_value_currency) REFERENCES currencies (code);
ALTER TABLE stocktakes ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE waste_records ADD FOREIGN KEY (batch_id) REFERENCES batches (id);
ALTER TABLE waste_records ADD FOREIGN KEY (item_id) REFERENCES items (id);
ALTER TABLE waste_records ADD FOREIGN KEY (photo_attachment_id) REFERENCES attachments (id);
ALTER TABLE waste_records ADD FOREIGN KEY (product_id) REFERENCES products (id);
ALTER TABLE waste_records ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE waste_records ADD FOREIGN KEY (value_currency) REFERENCES currencies (code);

-- +goose Down

DROP TABLE IF EXISTS npd_projects CASCADE;
DROP TABLE IF EXISTS cost_sheets CASCADE;
DROP TABLE IF EXISTS product_changes CASCADE;
DROP TABLE IF EXISTS products CASCADE;
DROP TABLE IF EXISTS dispatch_note_lines CASCADE;
DROP TABLE IF EXISTS dispatch_notes CASCADE;
DROP TABLE IF EXISTS batch_ingredients CASCADE;
DROP TABLE IF EXISTS batches CASCADE;
DROP TABLE IF EXISTS production_plan_lines CASCADE;
DROP TABLE IF EXISTS production_plans CASCADE;
DROP TABLE IF EXISTS kitchen_order_lines CASCADE;
DROP TABLE IF EXISTS kitchen_orders CASCADE;
DROP TABLE IF EXISTS recipe_ingredients CASCADE;
DROP TABLE IF EXISTS recipe_versions CASCADE;
DROP TABLE IF EXISTS recipes CASCADE;
DROP TABLE IF EXISTS disposals CASCADE;
DROP TABLE IF EXISTS waste_records CASCADE;
DROP TABLE IF EXISTS stock_adjustments CASCADE;
DROP TABLE IF EXISTS stock_variances CASCADE;
DROP TABLE IF EXISTS stocktake_lines CASCADE;
DROP TABLE IF EXISTS stocktakes CASCADE;
DROP TABLE IF EXISTS stock_returns CASCADE;
DROP TABLE IF EXISTS stock_transfer_lines CASCADE;
DROP TABLE IF EXISTS stock_transfers CASCADE;
DROP TABLE IF EXISTS stock_requisition_lines CASCADE;
DROP TABLE IF EXISTS stock_requisitions CASCADE;
DROP TABLE IF EXISTS stock_balances CASCADE;
DROP TABLE IF EXISTS items CASCADE;
