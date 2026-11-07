-- +goose Up

CREATE TABLE store_openings (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  business_date date NOT NULL,
  scheduled_at timestamptz,
  status_colour text CHECK (status_colour IN ('red', 'amber', 'green')),
  critical_controls jsonb,
  amber_issue text,
  rationale text,
  corrective_action text,
  decided_by uuid,
  decided_at timestamptz,
  actual_opened_at timestamptz,
  status text NOT NULL DEFAULT 'in_progress' CHECK (status IN ('in_progress', 'no_go', 'management_decision', 'go', 'opened')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE occurrence_entries (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  post text,
  occurred_at timestamptz NOT NULL,
  officer_employee_id uuid NOT NULL,
  description text NOT NULL,
  action_taken text,
  persons_involved text,
  open_item boolean NOT NULL DEFAULT false,
  incident_id uuid,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE patrol_routes (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  site_id uuid NOT NULL,
  "name" text NOT NULL,
  frequency_rule_key text,
  active boolean NOT NULL DEFAULT true,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE patrol_checkpoints (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  patrol_route_id uuid NOT NULL,
  sequence integer NOT NULL,
  "name" text NOT NULL,
  risk_category text,
  qr_code text,
  blind_spot boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE patrol_rounds (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  patrol_route_id uuid NOT NULL,
  officer_employee_id uuid NOT NULL,
  started_at timestamptz NOT NULL,
  finished_at timestamptz,
  missed_checkpoints integer,
  status text NOT NULL DEFAULT 'in_progress' CHECK (status IN ('in_progress', 'completed', 'incomplete')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE patrol_visits (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  patrol_round_id uuid NOT NULL,
  checkpoint_id uuid NOT NULL,
  visited_at timestamptz NOT NULL,
  observation text,
  attachment_id uuid,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE cameras (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  site_id uuid NOT NULL,
  label text NOT NULL,
  location text NOT NULL,
  critical_area boolean NOT NULL DEFAULT false,
  recorder text,
  "status" text CHECK ("status" IN ('ok', 'fault', 'offline')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE cctv_outages (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  camera_ids uuid[] NOT NULL,
  started_at timestamptz NOT NULL,
  ended_at timestamptz,
  cause text,
  alternative_measures text NOT NULL,
  ict_case text,
  test_outcome text,
  status text NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'restored', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE footage_requests (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  requester_employee_id uuid NOT NULL,
  site_id uuid NOT NULL,
  purpose text NOT NULL,
  camera_ids uuid[] NOT NULL,
  window_start timestamptz NOT NULL,
  window_end timestamptz NOT NULL,
  incident_id uuid,
  hold_until date,
  export_filename text,
  export_location text,
  file_hash text,
  exported_by uuid,
  closed_at timestamptz,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'preserved', 'exported', 'closed', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE external_visits (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  organisation text NOT NULL,
  group_leader text NOT NULL,
  contact text,
  site_id uuid NOT NULL,
  requested_on date NOT NULL,
  attendee_count integer NOT NULL,
  includes_minors boolean NOT NULL DEFAULT false,
  special_requirements text,
  risk_notes text,
  disclaimer_attachment_id uuid,
  nda_required boolean NOT NULL DEFAULT false,
  closed_at timestamptz,
  status text NOT NULL DEFAULT 'requested' CHECK (status IN ('requested', 'screening', 'pending_approval', 'approved', 'held', 'cancelled', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE INDEX ON store_openings (site_id);
CREATE INDEX ON store_openings (status);
CREATE INDEX ON store_openings (business_date);
CREATE INDEX ON occurrence_entries (site_id);
CREATE INDEX ON patrol_routes (site_id);
CREATE INDEX ON patrol_checkpoints (patrol_route_id);
CREATE INDEX ON patrol_rounds (status);
CREATE INDEX ON patrol_visits (patrol_round_id);
CREATE INDEX ON cameras (site_id);
CREATE INDEX ON cctv_outages (site_id);
CREATE INDEX ON cctv_outages (status);
CREATE INDEX ON footage_requests (site_id);
CREATE INDEX ON footage_requests (status);
CREATE INDEX ON external_visits (site_id);
CREATE INDEX ON external_visits (status);

ALTER TABLE cameras ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE cctv_outages ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE external_visits ADD FOREIGN KEY (disclaimer_attachment_id) REFERENCES attachments (id);
ALTER TABLE external_visits ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE footage_requests ADD FOREIGN KEY (exported_by) REFERENCES users (id);
ALTER TABLE footage_requests ADD FOREIGN KEY (incident_id) REFERENCES incidents (id);
ALTER TABLE footage_requests ADD FOREIGN KEY (requester_employee_id) REFERENCES employees (id);
ALTER TABLE footage_requests ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE occurrence_entries ADD FOREIGN KEY (incident_id) REFERENCES incidents (id);
ALTER TABLE occurrence_entries ADD FOREIGN KEY (officer_employee_id) REFERENCES employees (id);
ALTER TABLE occurrence_entries ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE patrol_rounds ADD FOREIGN KEY (officer_employee_id) REFERENCES employees (id);
ALTER TABLE patrol_rounds ADD FOREIGN KEY (patrol_route_id) REFERENCES patrol_routes (id);
ALTER TABLE patrol_routes ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE patrol_visits ADD FOREIGN KEY (attachment_id) REFERENCES attachments (id) ON DELETE CASCADE;
ALTER TABLE patrol_visits ADD FOREIGN KEY (checkpoint_id) REFERENCES patrol_checkpoints (id) ON DELETE CASCADE;
ALTER TABLE store_openings ADD FOREIGN KEY (decided_by) REFERENCES users (id);
ALTER TABLE store_openings ADD FOREIGN KEY (site_id) REFERENCES sites (id);

-- +goose Down

DROP TABLE IF EXISTS external_visits CASCADE;
DROP TABLE IF EXISTS footage_requests CASCADE;
DROP TABLE IF EXISTS cctv_outages CASCADE;
DROP TABLE IF EXISTS cameras CASCADE;
DROP TABLE IF EXISTS patrol_visits CASCADE;
DROP TABLE IF EXISTS patrol_rounds CASCADE;
DROP TABLE IF EXISTS patrol_checkpoints CASCADE;
DROP TABLE IF EXISTS patrol_routes CASCADE;
DROP TABLE IF EXISTS occurrence_entries CASCADE;
DROP TABLE IF EXISTS store_openings CASCADE;
