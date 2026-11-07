-- +goose Up

CREATE TABLE report_definitions (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  "key" text NOT NULL,
  "name" text NOT NULL,
  frequency text NOT NULL CHECK (frequency IN ('daily', 'weekly', 'bi_weekly', 'monthly', 'ad_hoc')),
  due_rule_key text NOT NULL,
  department_id uuid,
  submitter_role_id uuid,
  reviewer_role_id uuid,
  template_id uuid,
  shop_pattern text CHECK (shop_pattern IN ('closes_2200', 'open_24h')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE reports (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  report_definition_id uuid NOT NULL,
  period_start date NOT NULL,
  period_end date NOT NULL,
  site_id uuid,
  department_id uuid,
  due_at timestamptz,
  submitted_at timestamptz,
  submitted_by uuid,
  deputy_reason text,
  form_submission_id uuid,
  late boolean NOT NULL DEFAULT false,
  delay_reason text,
  escalated_at timestamptz,
  status text NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'draft', 'submitted', 'reviewed', 'late', 'missing')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE report_deadline_notices (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  report_definition_id uuid NOT NULL,
  period_start date NOT NULL,
  new_due_at timestamptz NOT NULL,
  notice_attachment_id uuid,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE management_meetings (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  meeting_at timestamptz NOT NULL,
  agenda text,
  attendee_employee_ids uuid[],
  minutes text,
  decisions text,
  status text NOT NULL DEFAULT 'planned' CHECK (status IN ('planned', 'held')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE compliance_obligations (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  kind text NOT NULL CHECK (kind IN ('tax', 'statutory', 'certificate', 'contract', 'licence', 'insurance', 'calibration', 'medical')),
  title text NOT NULL,
  authority text,
  site_id uuid,
  subject_type text,
  subject_id uuid,
  owner_employee_id uuid,
  due_on date NOT NULL,
  recurrence_rule_key text,
  reminder_days integer,
  completed_at timestamptz,
  evidence_attachment_id uuid,
  status text NOT NULL DEFAULT 'upcoming' CHECK (status IN ('upcoming', 'due', 'overdue', 'completed', 'cancelled')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE INDEX ON reports (site_id);
CREATE INDEX ON reports (status);
CREATE INDEX ON management_meetings (status);
CREATE INDEX ON compliance_obligations (site_id);
CREATE INDEX ON compliance_obligations (status);
CREATE INDEX ON compliance_obligations (subject_type, subject_id);

ALTER TABLE compliance_obligations ADD FOREIGN KEY (evidence_attachment_id) REFERENCES attachments (id);
ALTER TABLE compliance_obligations ADD FOREIGN KEY (owner_employee_id) REFERENCES employees (id);
ALTER TABLE compliance_obligations ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE report_deadline_notices ADD FOREIGN KEY (notice_attachment_id) REFERENCES attachments (id);
ALTER TABLE report_deadline_notices ADD FOREIGN KEY (report_definition_id) REFERENCES report_definitions (id);
ALTER TABLE report_definitions ADD FOREIGN KEY (department_id) REFERENCES departments (id);
ALTER TABLE report_definitions ADD FOREIGN KEY (reviewer_role_id) REFERENCES roles (id);
ALTER TABLE report_definitions ADD FOREIGN KEY (submitter_role_id) REFERENCES roles (id);
ALTER TABLE report_definitions ADD FOREIGN KEY (template_id) REFERENCES form_templates (id);
ALTER TABLE reports ADD FOREIGN KEY (department_id) REFERENCES departments (id);
ALTER TABLE reports ADD FOREIGN KEY (form_submission_id) REFERENCES form_submissions (id);
ALTER TABLE reports ADD FOREIGN KEY (report_definition_id) REFERENCES report_definitions (id);
ALTER TABLE reports ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE reports ADD FOREIGN KEY (submitted_by) REFERENCES users (id);

-- +goose Down

DROP TABLE IF EXISTS compliance_obligations CASCADE;
DROP TABLE IF EXISTS management_meetings CASCADE;
DROP TABLE IF EXISTS report_deadline_notices CASCADE;
DROP TABLE IF EXISTS reports CASCADE;
DROP TABLE IF EXISTS report_definitions CASCADE;
