-- +goose Up

CREATE TABLE user_credentials (
  user_id uuid PRIMARY KEY,
  password_hash text,
  pin_hash text,
  totp_secret text,
  totp_confirmed_at timestamptz,
  failed_logins integer NOT NULL DEFAULT 0,
  failed_pins integer NOT NULL DEFAULT 0,
  pin_locked_at timestamptz,
  updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE auth_sessions (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  user_id uuid NOT NULL,
  channel text NOT NULL,
  device_id uuid,
  ip text,
  created_at timestamptz NOT NULL DEFAULT now(),
  last_seen_at timestamptz NOT NULL DEFAULT now(),
  revoked_at timestamptz
);
CREATE INDEX ON auth_sessions (user_id);

CREATE TABLE refresh_tokens (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  session_id uuid NOT NULL REFERENCES auth_sessions (id) ON DELETE CASCADE,
  token_hash text NOT NULL UNIQUE,
  expires_at timestamptz NOT NULL,
  used_at timestamptz
);

CREATE TABLE mfa_challenges (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  user_id uuid NOT NULL,
  device_id uuid,
  channel text NOT NULL,
  expires_at timestamptz NOT NULL,
  failed_attempts integer NOT NULL DEFAULT 0,
  consumed_at timestamptz
);

CREATE TABLE idempotency_keys (
  actor_id uuid NOT NULL,
  key text NOT NULL,
  request_hash text NOT NULL,
  status integer,
  headers jsonb,
  body bytea,
  created_at timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY (actor_id, key)
);

CREATE TABLE reference_sequences (
  prefix text NOT NULL,
  site_code text NOT NULL,
  year integer NOT NULL,
  last_value bigint NOT NULL DEFAULT 0,
  PRIMARY KEY (prefix, site_code, year)
);

-- +goose StatementBegin
CREATE FUNCTION forbid_change() RETURNS trigger LANGUAGE plpgsql AS $$
BEGIN
  RAISE EXCEPTION '% is append-only', TG_TABLE_NAME;
END;
$$;
-- +goose StatementEnd

CREATE TABLE complexes (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  code text NOT NULL,
  "name" text NOT NULL,
  manager_employee_id uuid,
  active boolean NOT NULL DEFAULT true,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE (code)
);

CREATE TABLE sites (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  code text NOT NULL,
  "name" text NOT NULL,
  "type" text NOT NULL CHECK ("type" IN ('head_office', 'branch', 'central_kitchen', 'central_store', 'bakery', 'butchery')),
  complex_id uuid,
  brand text,
  trading_pattern text CHECK (trading_pattern IN ('closes_2200', 'open_24h', 'non_trading')),
  opens_at text,
  closes_at text,
  stockholding boolean NOT NULL DEFAULT false,
  address text,
  escalation_contacts jsonb,
  active boolean NOT NULL DEFAULT true,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE (code)
);

CREATE TABLE departments (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  code text NOT NULL,
  "name" text NOT NULL,
  parent_id uuid,
  head_employee_id uuid,
  active boolean NOT NULL DEFAULT true,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE (code)
);

CREATE TABLE cost_centres (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  code text NOT NULL,
  "name" text NOT NULL,
  department_id uuid,
  site_id uuid,
  sap_code text,
  active boolean NOT NULL DEFAULT true,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE (site_id, code)
);

CREATE TABLE production_lines (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  site_id uuid NOT NULL,
  "name" text NOT NULL,
  "type" text NOT NULL CHECK ("type" IN ('chicken_chips', 'pizza', 'bakery', 'beverages', 'butchery', 'other')),
  active boolean NOT NULL DEFAULT true,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE tills (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  site_id uuid NOT NULL,
  code text NOT NULL,
  pos_terminal_ref text,
  active boolean NOT NULL DEFAULT true,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE (site_id, code)
);

CREATE TABLE safes (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  site_id uuid NOT NULL,
  code text NOT NULL,
  location text,
  lock_status text CHECK (lock_status IN ('ok', 'defective')),
  active boolean NOT NULL DEFAULT true,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE (site_id, code)
);

CREATE TABLE storage_locations (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  site_id uuid NOT NULL,
  code text NOT NULL,
  "zone" text NOT NULL CHECK ("zone" IN ('dry', 'cold_room', 'freezer', 'chemical', 'packaging', 'quarantine', 'other')),
  custodian_employee_id uuid,
  active boolean NOT NULL DEFAULT true,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE (site_id, code)
);

CREATE TABLE currencies (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  code text NOT NULL,
  "name" text NOT NULL,
  symbol text,
  decimal_places integer,
  is_base boolean NOT NULL DEFAULT false,
  active boolean NOT NULL DEFAULT true,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE (code)
);

CREATE TABLE fx_rates (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  from_currency text NOT NULL,
  to_currency text NOT NULL,
  rate numeric(18,4) NOT NULL,
  rate_date date NOT NULL,
  "source" text NOT NULL CHECK ("source" IN ('rbz', 'manual')),
  change_reason text,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE public_holidays (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  holiday_date date NOT NULL,
  "name" text NOT NULL,
  applies_to_site_id uuid,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE employees (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  employee_number text NOT NULL,
  first_name text NOT NULL,
  last_name text NOT NULL,
  job_title text NOT NULL,
  department_id uuid NOT NULL,
  site_id uuid NOT NULL,
  manager_employee_id uuid,
  employment_status text CHECK (employment_status IN ('active', 'on_leave', 'suspended', 'transferred', 'resigned', 'terminated', 'retired')),
  start_date date NOT NULL,
  contract_type text CHECK (contract_type IN ('permanent', 'fixed_term', 'casual', 'other')),
  probation_end date,
  contract_end date,
  phone_e164 text,
  email text,
  belina_ref text,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE (employee_number)
);

CREATE TABLE employee_status_changes (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  employee_id uuid NOT NULL,
  "status" text NOT NULL CHECK ("status" IN ('active', 'on_leave', 'suspended', 'transferred', 'resigned', 'terminated', 'retired')),
  effective_date date NOT NULL,
  "reason" text NOT NULL,
  supporting_attachment_id uuid,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE employee_transfers (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  employee_id uuid NOT NULL,
  from_site_id uuid,
  to_site_id uuid,
  from_department_id uuid,
  to_department_id uuid,
  new_job_title text,
  new_manager_employee_id uuid,
  effective_date date NOT NULL,
  "reason" text NOT NULL,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'rejected', 'effective')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE employee_documents (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  employee_id uuid NOT NULL,
  document_type text NOT NULL,
  issue_date date,
  expiry_date date,
  attachment_id uuid NOT NULL,
  visibility text CHECK (visibility IN ('hr_only', 'hr_and_manager')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE emergency_contacts (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  employee_id uuid NOT NULL,
  contact_name text NOT NULL,
  relationship text NOT NULL,
  phone_e164 text NOT NULL,
  alternate_phone text,
  last_verified_on date,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE users (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  employee_id uuid NOT NULL,
  channels text NOT NULL CHECK (channels IN ('web', 'whatsapp', 'both')),
  mfa_required boolean NOT NULL DEFAULT false,
  last_sign_in_at timestamptz,
  credential_status text CHECK (credential_status IN ('unset', 'active', 'must_change')),
  status text NOT NULL DEFAULT 'invited' CHECK (status IN ('invited', 'active', 'locked', 'disabled')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE (employee_id)
);

CREATE TABLE roles (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  "key" text NOT NULL,
  "name" text NOT NULL,
  description text,
  permission_keys text[] NOT NULL,
  requires_mfa boolean NOT NULL DEFAULT false,
  system boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE ("key")
);

CREATE TABLE role_assignments (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  user_id uuid NOT NULL,
  role_id uuid NOT NULL,
  scope_type text NOT NULL CHECK (scope_type IN ('site', 'complex', 'all')),
  scope_id uuid,
  effective_from date NOT NULL,
  effective_to date,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE delegations (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  from_user_id uuid NOT NULL,
  to_user_id uuid NOT NULL,
  role_id uuid,
  workflow_keys text[],
  valid_from timestamptz NOT NULL,
  valid_to timestamptz NOT NULL,
  "reason" text NOT NULL,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'active', 'expired', 'revoked')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE sod_rules (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  "key" text NOT NULL,
  "name" text NOT NULL,
  action_a text NOT NULL,
  action_b text NOT NULL,
  "scope" text NOT NULL CHECK ("scope" IN ('same_record', 'same_business_date', 'same_supplier')),
  enforcement text NOT NULL CHECK (enforcement IN ('block', 'flag')),
  active boolean NOT NULL DEFAULT true,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE ("key")
);

CREATE TABLE devices (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  label text NOT NULL,
  site_id uuid,
  platform text,
  shared boolean NOT NULL DEFAULT false,
  last_seen_at timestamptz,
  status text NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'revoked')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE audit_events (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  occurred_at timestamptz,
  actor_user_id uuid,
  actor_role text,
  "action" text,
  subject_type text,
  subject_id uuid,
  site_id uuid,
  channel text CHECK (channel IN ('web', 'whatsapp', 'offline_web', 'integration', 'system')),
  device_id uuid,
  ip text,
  before jsonb,
  after jsonb,
  prev_hash text,
  hash text,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE INDEX ON cost_centres (site_id);
CREATE INDEX ON production_lines (site_id);
CREATE INDEX ON tills (site_id);
CREATE INDEX ON safes (site_id);
CREATE INDEX ON storage_locations (site_id);
CREATE INDEX ON employees (site_id);
CREATE INDEX ON employee_status_changes (employee_id);
CREATE INDEX ON employee_status_changes (employee_id);
CREATE INDEX ON employee_transfers (status);
CREATE INDEX ON employee_transfers (employee_id);
CREATE INDEX ON employee_transfers (employee_id);
CREATE INDEX ON employee_documents (employee_id);
CREATE INDEX ON employee_documents (employee_id);
CREATE INDEX ON emergency_contacts (employee_id);
CREATE INDEX ON emergency_contacts (employee_id);
CREATE INDEX ON users (status);
CREATE INDEX ON users (employee_id);
CREATE INDEX ON delegations (status);
CREATE INDEX ON devices (site_id);
CREATE INDEX ON devices (status);
CREATE INDEX ON audit_events (site_id);
CREATE INDEX ON audit_events (subject_type, subject_id);

ALTER TABLE audit_events ADD FOREIGN KEY (actor_user_id) REFERENCES users (id);
ALTER TABLE audit_events ADD FOREIGN KEY (device_id) REFERENCES devices (id);
ALTER TABLE audit_events ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE complexes ADD FOREIGN KEY (manager_employee_id) REFERENCES employees (id);
ALTER TABLE cost_centres ADD FOREIGN KEY (department_id) REFERENCES departments (id);
ALTER TABLE cost_centres ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE delegations ADD FOREIGN KEY (from_user_id) REFERENCES users (id);
ALTER TABLE delegations ADD FOREIGN KEY (role_id) REFERENCES roles (id);
ALTER TABLE delegations ADD FOREIGN KEY (to_user_id) REFERENCES users (id);
ALTER TABLE departments ADD FOREIGN KEY (head_employee_id) REFERENCES employees (id);
ALTER TABLE departments ADD FOREIGN KEY (parent_id) REFERENCES departments (id);
ALTER TABLE devices ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE employee_transfers ADD FOREIGN KEY (from_department_id) REFERENCES departments (id);
ALTER TABLE employee_transfers ADD FOREIGN KEY (from_site_id) REFERENCES sites (id);
ALTER TABLE employee_transfers ADD FOREIGN KEY (new_manager_employee_id) REFERENCES employees (id);
ALTER TABLE employee_transfers ADD FOREIGN KEY (to_department_id) REFERENCES departments (id);
ALTER TABLE employee_transfers ADD FOREIGN KEY (to_site_id) REFERENCES sites (id);
ALTER TABLE employees ADD FOREIGN KEY (department_id) REFERENCES departments (id);
ALTER TABLE employees ADD FOREIGN KEY (manager_employee_id) REFERENCES employees (id);
ALTER TABLE employees ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE fx_rates ADD FOREIGN KEY (from_currency) REFERENCES currencies (code);
ALTER TABLE fx_rates ADD FOREIGN KEY (to_currency) REFERENCES currencies (code);
ALTER TABLE production_lines ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE public_holidays ADD FOREIGN KEY (applies_to_site_id) REFERENCES sites (id);
ALTER TABLE role_assignments ADD FOREIGN KEY (role_id) REFERENCES roles (id);
ALTER TABLE role_assignments ADD FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE safes ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE sites ADD FOREIGN KEY (complex_id) REFERENCES complexes (id);
ALTER TABLE storage_locations ADD FOREIGN KEY (custodian_employee_id) REFERENCES employees (id);
ALTER TABLE storage_locations ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE tills ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE users ADD FOREIGN KEY (employee_id) REFERENCES employees (id);

ALTER TABLE audit_events ADD COLUMN seq bigint GENERATED ALWAYS AS IDENTITY;
CREATE UNIQUE INDEX ON audit_events (seq);
CREATE TRIGGER audit_events_append_only BEFORE UPDATE OR DELETE ON audit_events
  FOR EACH ROW EXECUTE FUNCTION forbid_change();

-- +goose Down

DROP FUNCTION IF EXISTS forbid_change CASCADE;
DROP TABLE IF EXISTS reference_sequences, idempotency_keys, mfa_challenges, refresh_tokens, auth_sessions, user_credentials CASCADE;

DROP TABLE IF EXISTS audit_events CASCADE;
DROP TABLE IF EXISTS devices CASCADE;
DROP TABLE IF EXISTS sod_rules CASCADE;
DROP TABLE IF EXISTS delegations CASCADE;
DROP TABLE IF EXISTS role_assignments CASCADE;
DROP TABLE IF EXISTS roles CASCADE;
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS emergency_contacts CASCADE;
DROP TABLE IF EXISTS employee_documents CASCADE;
DROP TABLE IF EXISTS employee_transfers CASCADE;
DROP TABLE IF EXISTS employee_status_changes CASCADE;
DROP TABLE IF EXISTS employees CASCADE;
DROP TABLE IF EXISTS public_holidays CASCADE;
DROP TABLE IF EXISTS fx_rates CASCADE;
DROP TABLE IF EXISTS currencies CASCADE;
DROP TABLE IF EXISTS storage_locations CASCADE;
DROP TABLE IF EXISTS safes CASCADE;
DROP TABLE IF EXISTS tills CASCADE;
DROP TABLE IF EXISTS production_lines CASCADE;
DROP TABLE IF EXISTS cost_centres CASCADE;
DROP TABLE IF EXISTS departments CASCADE;
DROP TABLE IF EXISTS sites CASCADE;
DROP TABLE IF EXISTS complexes CASCADE;
