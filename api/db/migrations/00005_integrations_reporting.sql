-- +goose Up

CREATE TABLE sync_conflicts (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  batch_id text,
  device_id uuid,
  site_id uuid,
  client_record_id uuid,
  subject_type text,
  subject_id uuid,
  "reason" text,
  client_payload jsonb,
  resolved_at timestamptz,
  status text NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'resolved')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE whatsapp_identities (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  employee_id uuid NOT NULL,
  phone_e164 text NOT NULL,
  verified_at timestamptz,
  pin_set_at timestamptz,
  last_active_at timestamptz,
  status text NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'verified', 'revoked')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE whatsapp_templates (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  "name" text NOT NULL,
  language text NOT NULL,
  category text NOT NULL CHECK (category IN ('utility', 'authentication', 'marketing')),
  body text NOT NULL,
  provider_status text,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE whatsapp_messages (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  direction text CHECK (direction IN ('inbound', 'outbound')),
  identity_id uuid,
  provider_message_id text,
  message_type text,
  summary text,
  subject_type text,
  subject_id uuid,
  sent_at timestamptz,
  delivery_status text,
  cost_category text,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE integrations (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  system text NOT NULL CHECK (system IN ('gaap_pos', 'gaap_unity', 'sap_b1', 'belina', 'bank', 'cctv')),
  mode text NOT NULL CHECK (mode IN ('api', 'file', 'manual')),
  config jsonb,
  last_success_at timestamptz,
  last_error text,
  status text NOT NULL DEFAULT 'disabled' CHECK (status IN ('disabled', 'active', 'degraded')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE sync_runs (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  integration_id uuid NOT NULL,
  direction text CHECK (direction IN ('inbound', 'outbound')),
  started_at timestamptz,
  finished_at timestamptz,
  records_ok integer,
  records_failed integer,
  error_summary text,
  status text NOT NULL DEFAULT 'running' CHECK (status IN ('running', 'succeeded', 'partial', 'failed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE data_imports (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  kind text NOT NULL CHECK (kind IN ('pos_sales', 'z_report', 'stock_balance', 'bank_statement', 'payroll_list', 'sales_by_batch', 'other')),
  attachment_id uuid NOT NULL,
  site_id uuid,
  business_date date,
  rows_ok integer,
  rows_failed integer,
  error_report_attachment_id uuid,
  status text NOT NULL DEFAULT 'uploaded' CHECK (status IN ('uploaded', 'validating', 'applied', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE postings (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  target_system text CHECK (target_system IN ('gaap_unity', 'sap_b1', 'belina', 'gaap_pos')),
  posting_type text,
  subject_type text,
  subject_id uuid,
  site_id uuid,
  payload jsonb,
  external_ref text,
  attempts integer,
  last_error text,
  posted_at timestamptz,
  status text NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'sent', 'posted', 'failed', 'manual_required', 'cancelled')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE pos_sales_summaries (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  site_id uuid,
  till_id uuid,
  business_date date,
  shift text,
  cashier_employee_id uuid,
  currency text,
  tender text CHECK (tender IN ('cash', 'card', 'mobile_money', 'other')),
  gross numeric(18,4),
  voids numeric(18,4),
  refunds numeric(18,4),
  net numeric(18,4),
  transaction_count integer,
  source_run_id uuid,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE z_reports (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  site_id uuid NOT NULL,
  till_id uuid NOT NULL,
  business_date date NOT NULL,
  report_type text NOT NULL CHECK (report_type IN ('x_read', 'z_report')),
  report_time timestamptz NOT NULL,
  total_amount numeric(18,4) NOT NULL,
  total_currency text NOT NULL,
  fiscal_status text CHECK (fiscal_status IN ('ok', 'failed', 'not_fiscalised')),
  attachment_id uuid,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE kpi_definitions (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  "key" text NOT NULL,
  "name" text NOT NULL,
  module text NOT NULL,
  formula text,
  target_rule_key text,
  unit text NOT NULL CHECK (unit IN ('count', 'percent', 'money', 'duration')),
  drilldown text,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE ("key")
);

CREATE TABLE export_jobs (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  kind text NOT NULL CHECK (kind IN ('pdf', 'xlsx', 'csv')),
  "source" text NOT NULL,
  filters jsonb,
  file_attachment_id uuid,
  finished_at timestamptz,
  status text NOT NULL DEFAULT 'queued' CHECK (status IN ('queued', 'running', 'ready', 'failed', 'expired')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE retention_classes (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  "key" text NOT NULL,
  "name" text NOT NULL,
  period_value integer,
  period_unit text CHECK (period_unit IN ('months', 'years')),
  trigger text NOT NULL CHECK (trigger IN ('record_date', 'closed', 'employment_end', 'product_end')),
  source_clause text,
  decision_id text,
  disposition text CHECK (disposition IN ('archive', 'destroy', 'review')),
  status text NOT NULL DEFAULT 'unset' CHECK (status IN ('unset', 'draft', 'active', 'superseded')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE ("key")
);

CREATE TABLE legal_holds (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  subject_type text NOT NULL,
  subject_id uuid NOT NULL,
  "reason" text NOT NULL,
  owner_user_id uuid,
  released_at timestamptz,
  status text NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'released')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE disposition_batches (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  retention_class_id uuid NOT NULL,
  cutoff_date date NOT NULL,
  record_count integer,
  completed_at timestamptz,
  status text NOT NULL DEFAULT 'proposed' CHECK (status IN ('proposed', 'pending_approval', 'approved', 'completed', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE INDEX ON sync_conflicts (site_id);
CREATE INDEX ON sync_conflicts (status);
CREATE INDEX ON sync_conflicts (subject_type, subject_id);
CREATE INDEX ON whatsapp_identities (status);
CREATE INDEX ON whatsapp_identities (employee_id);
CREATE INDEX ON whatsapp_messages (subject_type, subject_id);
CREATE INDEX ON integrations (status);
CREATE INDEX ON sync_runs (status);
CREATE INDEX ON sync_runs (integration_id);
CREATE INDEX ON data_imports (site_id);
CREATE INDEX ON data_imports (status);
CREATE INDEX ON data_imports (business_date);
CREATE INDEX ON postings (site_id);
CREATE INDEX ON postings (status);
CREATE INDEX ON postings (subject_type, subject_id);
CREATE INDEX ON pos_sales_summaries (site_id);
CREATE INDEX ON pos_sales_summaries (business_date);
CREATE INDEX ON z_reports (site_id);
CREATE INDEX ON z_reports (business_date);
CREATE INDEX ON export_jobs (status);
CREATE INDEX ON retention_classes (status);
CREATE INDEX ON legal_holds (status);
CREATE INDEX ON legal_holds (subject_type, subject_id);
CREATE INDEX ON disposition_batches (status);

ALTER TABLE data_imports ADD FOREIGN KEY (attachment_id) REFERENCES attachments (id);
ALTER TABLE data_imports ADD FOREIGN KEY (error_report_attachment_id) REFERENCES attachments (id);
ALTER TABLE data_imports ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE disposition_batches ADD FOREIGN KEY (retention_class_id) REFERENCES retention_classes (id);
ALTER TABLE export_jobs ADD FOREIGN KEY (file_attachment_id) REFERENCES attachments (id);
ALTER TABLE form_template_versions ADD FOREIGN KEY (retention_class_id) REFERENCES retention_classes (id);
ALTER TABLE legal_holds ADD FOREIGN KEY (owner_user_id) REFERENCES users (id);
ALTER TABLE pos_sales_summaries ADD FOREIGN KEY (cashier_employee_id) REFERENCES employees (id);
ALTER TABLE pos_sales_summaries ADD FOREIGN KEY (currency) REFERENCES currencies (code);
ALTER TABLE pos_sales_summaries ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE pos_sales_summaries ADD FOREIGN KEY (source_run_id) REFERENCES sync_runs (id);
ALTER TABLE pos_sales_summaries ADD FOREIGN KEY (till_id) REFERENCES tills (id);
ALTER TABLE postings ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE sync_conflicts ADD FOREIGN KEY (device_id) REFERENCES devices (id);
ALTER TABLE sync_conflicts ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE whatsapp_identities ADD FOREIGN KEY (employee_id) REFERENCES employees (id);
ALTER TABLE whatsapp_messages ADD FOREIGN KEY (identity_id) REFERENCES whatsapp_identities (id);
ALTER TABLE z_reports ADD FOREIGN KEY (attachment_id) REFERENCES attachments (id);
ALTER TABLE z_reports ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE z_reports ADD FOREIGN KEY (till_id) REFERENCES tills (id);
ALTER TABLE z_reports ADD FOREIGN KEY (total_currency) REFERENCES currencies (code);

-- +goose Down

DROP TABLE IF EXISTS disposition_batches CASCADE;
DROP TABLE IF EXISTS legal_holds CASCADE;
DROP TABLE IF EXISTS retention_classes CASCADE;
DROP TABLE IF EXISTS export_jobs CASCADE;
DROP TABLE IF EXISTS kpi_definitions CASCADE;
DROP TABLE IF EXISTS z_reports CASCADE;
DROP TABLE IF EXISTS pos_sales_summaries CASCADE;
DROP TABLE IF EXISTS postings CASCADE;
DROP TABLE IF EXISTS data_imports CASCADE;
DROP TABLE IF EXISTS sync_runs CASCADE;
DROP TABLE IF EXISTS integrations CASCADE;
DROP TABLE IF EXISTS whatsapp_messages CASCADE;
DROP TABLE IF EXISTS whatsapp_templates CASCADE;
DROP TABLE IF EXISTS whatsapp_identities CASCADE;
DROP TABLE IF EXISTS sync_conflicts CASCADE;
