-- +goose Up

CREATE TABLE notification_preferences (
  user_id uuid PRIMARY KEY,
  categories jsonb NOT NULL DEFAULT '{}',
  updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE schedules (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  title text NOT NULL,
  workflow_key text,
  form_template_id uuid,
  frequency_rule_key text,
  site_ids uuid[],
  assignee_role_id uuid NOT NULL,
  due_offset_minutes integer,
  reminder_minutes integer,
  escalate_after_minutes integer,
  escalate_to_role_id uuid,
  status text NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'paused')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE tasks (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  title text NOT NULL,
  schedule_id uuid,
  subject_type text,
  subject_id uuid,
  site_id uuid NOT NULL,
  business_date date,
  shift text,
  assignee_user_id uuid,
  assignee_role_id uuid,
  due_at timestamptz,
  priority text CHECK (priority IN ('low', 'normal', 'high', 'critical')),
  overdue boolean NOT NULL DEFAULT false,
  completed_at timestamptz,
  skip_reason text,
  status text NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'in_progress', 'done', 'skipped', 'cancelled')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE form_templates (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  "key" text NOT NULL,
  title text NOT NULL,
  kind text NOT NULL CHECK (kind IN ('checklist', 'register', 'log', 'inspection', 'report', 'acknowledgement')),
  module text NOT NULL,
  area text,
  source_sop text,
  owner_role_id uuid,
  current_version_id uuid,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'active', 'retired')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE ("key")
);

CREATE TABLE form_template_versions (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  form_template_id uuid NOT NULL,
  "version" integer,
  definition jsonb NOT NULL,
  offline_allowed boolean NOT NULL DEFAULT false,
  channels text NOT NULL CHECK (channels IN ('web', 'whatsapp', 'both')),
  whatsapp_flow_id text,
  retention_class_id uuid,
  required_rule_keys text[],
  effective_from timestamptz,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'published', 'retired')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE form_submissions (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  template_version_id uuid NOT NULL,
  site_id uuid NOT NULL,
  business_date date NOT NULL,
  shift text,
  task_id uuid,
  answers jsonb NOT NULL,
  "result" text CHECK ("result" IN ('pass', 'fail', 'na')),
  captured_at timestamptz NOT NULL,
  channel text CHECK (channel IN ('web', 'whatsapp', 'offline_web')),
  client_record_id uuid,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'submitted', 'pending_signature', 'pending_review', 'closed', 'returned')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE escalation_policies (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  domain text NOT NULL CHECK (domain IN ('quality', 'customer', 'security', 'cash', 'she', 'equipment', 'general')),
  severity_level integer NOT NULL,
  label text NOT NULL,
  after_hours boolean NOT NULL DEFAULT false,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'active', 'superseded')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE escalation_policy_steps (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  escalation_policy_id uuid NOT NULL,
  sequence integer NOT NULL,
  notify_role_id uuid NOT NULL,
  channels text[] NOT NULL,
  ack_within_minutes integer,
  next_on_timeout boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE escalations (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  subject_type text,
  subject_id uuid,
  domain text,
  severity_level integer,
  site_id uuid,
  current_step integer,
  triggered_at timestamptz,
  closed_at timestamptz,
  status text NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'acknowledged', 'breached', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE escalation_recipients (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  escalation_id uuid NOT NULL,
  step integer,
  user_id uuid,
  channel text,
  notified_at timestamptz,
  acknowledged_at timestamptz,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE actions (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  source_type text,
  source_id uuid,
  kind text NOT NULL CHECK (kind IN ('corrective', 'preventive', 'general', 'report', 'audit')),
  description text NOT NULL,
  root_cause text,
  owner_employee_id uuid NOT NULL,
  site_id uuid,
  due_date date NOT NULL,
  priority text CHECK (priority IN ('low', 'normal', 'high', 'critical')),
  percent_complete integer,
  effectiveness_result text,
  verified_by uuid,
  closed_at timestamptz,
  status text NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'in_progress', 'pending_verification', 'closed', 'cancelled')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE notifications (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  recipient_user_id uuid,
  category text,
  title text,
  body text,
  channel text CHECK (channel IN ('in_app', 'push', 'email', 'whatsapp')),
  subject_type text,
  subject_id uuid,
  sent_at timestamptz,
  delivered_at timestamptz,
  read_at timestamptz,
  failure_reason text,
  status text NOT NULL DEFAULT 'queued' CHECK (status IN ('queued', 'sent', 'delivered', 'read', 'failed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE INDEX ON schedules (status);
CREATE INDEX ON tasks (site_id);
CREATE INDEX ON tasks (status);
CREATE INDEX ON tasks (business_date);
CREATE INDEX ON tasks (subject_type, subject_id);
CREATE INDEX ON form_templates (status);
CREATE INDEX ON form_template_versions (status);
CREATE INDEX ON form_template_versions (form_template_id);
CREATE INDEX ON form_submissions (site_id);
CREATE INDEX ON form_submissions (status);
CREATE INDEX ON form_submissions (business_date);
CREATE INDEX ON escalation_policies (status);
CREATE INDEX ON escalation_policy_steps (escalation_policy_id);
CREATE INDEX ON escalations (site_id);
CREATE INDEX ON escalations (status);
CREATE INDEX ON escalations (subject_type, subject_id);
CREATE INDEX ON escalation_recipients (escalation_id);
CREATE INDEX ON actions (site_id);
CREATE INDEX ON actions (status);
CREATE INDEX ON notifications (status);
CREATE INDEX ON notifications (subject_type, subject_id);

ALTER TABLE actions ADD FOREIGN KEY (owner_employee_id) REFERENCES employees (id);
ALTER TABLE actions ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE actions ADD FOREIGN KEY (verified_by) REFERENCES users (id);
ALTER TABLE escalation_policy_steps ADD FOREIGN KEY (notify_role_id) REFERENCES roles (id) ON DELETE CASCADE;
ALTER TABLE escalation_recipients ADD FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;
ALTER TABLE escalations ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE form_submissions ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE form_submissions ADD FOREIGN KEY (task_id) REFERENCES tasks (id);
ALTER TABLE form_submissions ADD FOREIGN KEY (template_version_id) REFERENCES form_template_versions (id);
ALTER TABLE form_templates ADD FOREIGN KEY (current_version_id) REFERENCES form_template_versions (id);
ALTER TABLE form_templates ADD FOREIGN KEY (owner_role_id) REFERENCES roles (id);
ALTER TABLE notifications ADD FOREIGN KEY (recipient_user_id) REFERENCES users (id);
ALTER TABLE schedules ADD FOREIGN KEY (assignee_role_id) REFERENCES roles (id);
ALTER TABLE schedules ADD FOREIGN KEY (escalate_to_role_id) REFERENCES roles (id);
ALTER TABLE schedules ADD FOREIGN KEY (form_template_id) REFERENCES form_templates (id);
ALTER TABLE tasks ADD FOREIGN KEY (assignee_role_id) REFERENCES roles (id);
ALTER TABLE tasks ADD FOREIGN KEY (assignee_user_id) REFERENCES users (id);
ALTER TABLE tasks ADD FOREIGN KEY (schedule_id) REFERENCES schedules (id);
ALTER TABLE tasks ADD FOREIGN KEY (site_id) REFERENCES sites (id);

CREATE UNIQUE INDEX tasks_once_per_day ON tasks (schedule_id, site_id, business_date) WHERE schedule_id IS NOT NULL;
CREATE INDEX ON tasks (assignee_user_id);
CREATE INDEX ON notifications (recipient_user_id, read_at);
ALTER TABLE tasks ADD COLUMN reminded_at timestamptz;

-- +goose Down
DROP TABLE IF EXISTS notification_preferences;
DROP TABLE IF EXISTS notifications CASCADE;
DROP TABLE IF EXISTS actions CASCADE;
DROP TABLE IF EXISTS escalation_recipients CASCADE;
DROP TABLE IF EXISTS escalations CASCADE;
DROP TABLE IF EXISTS escalation_policy_steps CASCADE;
DROP TABLE IF EXISTS escalation_policies CASCADE;
DROP TABLE IF EXISTS form_submissions CASCADE;
DROP TABLE IF EXISTS form_template_versions CASCADE;
DROP TABLE IF EXISTS form_templates CASCADE;
DROP TABLE IF EXISTS tasks CASCADE;
DROP TABLE IF EXISTS schedules CASCADE;
