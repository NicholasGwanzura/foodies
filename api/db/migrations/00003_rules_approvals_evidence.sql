-- +goose Up

-- Which rules must have an active value before a workflow can go live. The data
-- model keeps this on templates (required_rule_keys); dedicated workflows use this.
CREATE TABLE workflow_requirements (
  workflow_key text NOT NULL,
  rule_key text NOT NULL,
  PRIMARY KEY (workflow_key, rule_key)
);

CREATE TABLE rule_definitions (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  "key" text NOT NULL,
  "name" text NOT NULL,
  module text NOT NULL,
  value_type text NOT NULL CHECK (value_type IN ('number', 'money', 'duration', 'time', 'percent', 'enum', 'schedule', 'json')),
  scope_level text NOT NULL CHECK (scope_level IN ('group', 'complex', 'site', 'item', 'asset')),
  owner_role_id uuid,
  source_clause text,
  decision_id text,
  blocking boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE ("key")
);

CREATE TABLE rule_versions (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  rule_definition_id uuid NOT NULL,
  scope_type text NOT NULL CHECK (scope_type IN ('group', 'complex', 'site', 'item', 'asset')),
  scope_id uuid,
  "value" jsonb NOT NULL,
  effective_from timestamptz NOT NULL,
  effective_to timestamptz,
  source_document_id uuid,
  change_reason text NOT NULL,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'active', 'superseded', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE workflow_activations (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  workflow_key text NOT NULL,
  site_id uuid NOT NULL,
  missing_rule_keys text[],
  activated_at timestamptz,
  activated_by uuid,
  status text NOT NULL DEFAULT 'not_ready' CHECK (status IN ('not_ready', 'ready', 'live', 'suspended')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE (workflow_key, site_id)
);

CREATE TABLE approval_policies (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  workflow_key text NOT NULL,
  "name" text NOT NULL,
  "version" integer,
  effective_from timestamptz,
  source_document_id uuid,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'active', 'superseded')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE approval_policy_steps (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  approval_policy_id uuid NOT NULL,
  sequence integer NOT NULL,
  role_id uuid NOT NULL,
  mode text NOT NULL CHECK (mode IN ('any', 'all')),
  condition jsonb,
  sla_minutes integer,
  escalate_to_role_id uuid,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE approval_requests (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  subject_type text,
  subject_id uuid,
  workflow_key text,
  policy_id uuid,
  site_id uuid,
  amount_amount numeric(18,4),
  amount_currency text,
  current_step integer,
  summary text,
  decided_at timestamptz,
  status text NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'approved', 'rejected', 'returned', 'cancelled')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE approval_request_steps (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  approval_request_id uuid NOT NULL,
  sequence integer,
  role_id uuid,
  decided_by uuid,
  decision text CHECK (decision IN ('pending', 'approved', 'rejected', 'returned')),
  decided_at timestamptz,
  channel text CHECK (channel IN ('web', 'whatsapp')),
  "comment" text,
  signature_id uuid,
  due_at timestamptz,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  assigned_user_id uuid
);

CREATE TABLE signature_requests (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  subject_type text NOT NULL,
  subject_id uuid NOT NULL,
  meaning text NOT NULL CHECK (meaning IN ('prepared', 'counted', 'verified', 'witnessed', 'approved', 'received', 'handed_over', 'countersigned', 'acknowledged')),
  signer_kind text NOT NULL CHECK (signer_kind IN ('employee', 'external')),
  signer_employee_id uuid,
  signer_role_id uuid,
  external_name text,
  external_organisation text,
  external_id_ref text,
  content_hash text,
  signed_at timestamptz,
  signed_channel text CHECK (signed_channel IN ('web', 'whatsapp', 'offline_web')),
  device_id uuid,
  attested_by uuid,
  proof_attachment_id uuid,
  status text NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'signed', 'declined', 'expired')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE attachments (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  filename text NOT NULL,
  content_type text NOT NULL,
  size_bytes integer NOT NULL,
  sha256 text NOT NULL,
  subject_type text,
  subject_id uuid,
  classification text CHECK (classification IN ('normal', 'restricted')),
  kind text CHECK (kind IN ('photo', 'scan', 'document', 'report', 'other')),
  paper_original_location text,
  captured_at timestamptz,
  upload_url text,
  download_url text,
  status text NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'available', 'quarantined')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE comments (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  subject_type text NOT NULL,
  subject_id uuid NOT NULL,
  body text NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE record_links (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  from_type text NOT NULL,
  from_id uuid NOT NULL,
  to_type text NOT NULL,
  to_id uuid NOT NULL,
  relation text NOT NULL CHECK (relation IN ('caused_by', 'related_to', 'escalated_to', 'launched', 'resolved_by')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE INDEX ON rule_versions (status);
CREATE INDEX ON rule_versions (rule_definition_id);
CREATE INDEX ON workflow_activations (site_id);
CREATE INDEX ON workflow_activations (status);
CREATE INDEX ON approval_policies (status);
CREATE INDEX ON approval_policy_steps (approval_policy_id);
CREATE INDEX ON approval_requests (site_id);
CREATE INDEX ON approval_requests (status);
CREATE INDEX ON approval_requests (subject_type, subject_id);
CREATE INDEX ON approval_request_steps (approval_request_id);
CREATE INDEX ON signature_requests (status);
CREATE INDEX ON signature_requests (subject_type, subject_id);
CREATE INDEX ON attachments (status);
CREATE INDEX ON attachments (subject_type, subject_id);
CREATE INDEX ON comments (subject_type, subject_id);

ALTER TABLE approval_policy_steps ADD FOREIGN KEY (escalate_to_role_id) REFERENCES roles (id) ON DELETE CASCADE;
ALTER TABLE approval_policy_steps ADD FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE;
ALTER TABLE approval_request_steps ADD FOREIGN KEY (decided_by) REFERENCES users (id) ON DELETE CASCADE;
ALTER TABLE approval_request_steps ADD FOREIGN KEY (role_id) REFERENCES roles (id) ON DELETE CASCADE;
ALTER TABLE approval_request_steps ADD FOREIGN KEY (signature_id) REFERENCES signature_requests (id) ON DELETE CASCADE;
ALTER TABLE approval_requests ADD FOREIGN KEY (amount_currency) REFERENCES currencies (code);
ALTER TABLE approval_requests ADD FOREIGN KEY (policy_id) REFERENCES approval_policies (id);
ALTER TABLE approval_requests ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE employee_documents ADD FOREIGN KEY (attachment_id) REFERENCES attachments (id);
ALTER TABLE employee_status_changes ADD FOREIGN KEY (supporting_attachment_id) REFERENCES attachments (id);
ALTER TABLE rule_definitions ADD FOREIGN KEY (owner_role_id) REFERENCES roles (id);
ALTER TABLE signature_requests ADD FOREIGN KEY (attested_by) REFERENCES users (id);
ALTER TABLE signature_requests ADD FOREIGN KEY (device_id) REFERENCES devices (id);
ALTER TABLE signature_requests ADD FOREIGN KEY (proof_attachment_id) REFERENCES attachments (id);
ALTER TABLE signature_requests ADD FOREIGN KEY (signer_employee_id) REFERENCES employees (id);
ALTER TABLE signature_requests ADD FOREIGN KEY (signer_role_id) REFERENCES roles (id);
ALTER TABLE workflow_activations ADD FOREIGN KEY (activated_by) REFERENCES users (id);
ALTER TABLE workflow_activations ADD FOREIGN KEY (site_id) REFERENCES sites (id);

-- +goose Down
DROP TABLE IF EXISTS workflow_requirements;
DROP TABLE IF EXISTS record_links CASCADE;
DROP TABLE IF EXISTS comments CASCADE;
DROP TABLE IF EXISTS attachments CASCADE;
DROP TABLE IF EXISTS signature_requests CASCADE;
DROP TABLE IF EXISTS approval_request_steps CASCADE;
DROP TABLE IF EXISTS approval_requests CASCADE;
DROP TABLE IF EXISTS approval_policy_steps CASCADE;
DROP TABLE IF EXISTS approval_policies CASCADE;
DROP TABLE IF EXISTS workflow_activations CASCADE;
DROP TABLE IF EXISTS rule_versions CASCADE;
DROP TABLE IF EXISTS rule_definitions CASCADE;
