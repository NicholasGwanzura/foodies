-- +goose Up

CREATE TABLE asset_classes (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  code text NOT NULL,
  "name" text NOT NULL,
  depreciation_rate_percent numeric(18,4),
  useful_life_years integer,
  capitalise boolean NOT NULL DEFAULT false,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE (code)
);

CREATE TABLE assets (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  tag text NOT NULL,
  description text NOT NULL,
  asset_class_id uuid NOT NULL,
  serial_number text,
  make_model text,
  acquisition_date date,
  supplier_id uuid,
  purchase_order_id uuid,
  goods_receipt_id uuid,
  invoice_id uuid,
  cost_amount numeric(18,4),
  cost_currency text,
  site_id uuid NOT NULL,
  department_id uuid,
  cost_centre_id uuid,
  custodian_employee_id uuid,
  condition text CHECK (condition IN ('good', 'fair', 'poor', 'out_of_service')),
  useful_life_years integer,
  residual_value_amount numeric(18,4),
  residual_value_currency text,
  critical boolean NOT NULL DEFAULT false,
  authorised_operator_ids uuid[],
  sap_asset_ref text,
  status text NOT NULL DEFAULT 'wip' CHECK (status IN ('wip', 'active', 'out_of_service', 'disposed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE capex_requests (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  department_id uuid NOT NULL,
  description text NOT NULL,
  justification text NOT NULL,
  amount_amount numeric(18,4) NOT NULL,
  amount_currency text NOT NULL,
  budgeted boolean NOT NULL DEFAULT false,
  risk_assessment text,
  classification text CHECK (classification IN ('capex', 'opex')),
  purchase_requisition_id uuid,
  asset_id uuid,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'rejected', 'ordered', 'capitalised')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE asset_transfers (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  asset_id uuid NOT NULL,
  to_site_id uuid NOT NULL,
  to_custodian_employee_id uuid,
  to_cost_centre_id uuid,
  "reason" text NOT NULL,
  received_at timestamptz,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'received', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE asset_verifications (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  verification_year integer NOT NULL,
  completed_at timestamptz,
  status text NOT NULL DEFAULT 'planned' CHECK (status IN ('planned', 'in_progress', 'submitted', 'reviewed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE asset_verification_lines (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  asset_verification_id uuid NOT NULL,
  asset_id uuid NOT NULL,
  found boolean NOT NULL DEFAULT false,
  condition text,
  location_correct boolean NOT NULL DEFAULT false,
  note text,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE asset_disposals (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  asset_id uuid NOT NULL,
  "method" text NOT NULL CHECK ("method" IN ('sale', 'auction', 'scrap', 'write_off', 'insurance_claim')),
  "reason" text NOT NULL,
  proceeds_amount numeric(18,4),
  proceeds_currency text,
  removal_confirmed_at timestamptz,
  invoice_reference text,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'removed', 'completed', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE key_items (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  site_id uuid NOT NULL,
  label text NOT NULL,
  key_type text NOT NULL CHECK (key_type IN ('store', 'storeroom', 'office_safe', 'back_door', 'alarm_remote', 'other')),
  areas_secured text,
  primary_custodian_employee_id uuid,
  backup_location text,
  master_location text,
  status text NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'in_use', 'lost', 'faulty', 'retired')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE key_movements (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  key_item_id uuid NOT NULL,
  holder_employee_id uuid NOT NULL,
  purpose text NOT NULL,
  issued_at timestamptz NOT NULL,
  returned_at timestamptz,
  condition text,
  movement_kind text CHECK (movement_kind IN ('daily', 'emergency', 'backup_collection')),
  status text NOT NULL DEFAULT 'out' CHECK (status IN ('out', 'returned', 'overdue')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE key_duplication_requests (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  key_type text NOT NULL,
  "reason" text NOT NULL,
  quantity integer NOT NULL,
  risk_assessment text,
  memo_attachment_id uuid,
  locksmith text,
  cut_on date,
  issued_key_ids uuid[],
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'cut', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE key_audits (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  kind text NOT NULL CHECK (kind IN ('monthly', 'annual')),
  period text NOT NULL,
  auditor_employee_id uuid NOT NULL,
  status text NOT NULL DEFAULT 'planned' CHECK (status IN ('planned', 'in_progress', 'completed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE key_audit_lines (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  key_audit_id uuid NOT NULL,
  key_item_id uuid NOT NULL,
  expected_custodian_employee_id uuid,
  verified boolean NOT NULL DEFAULT false,
  discrepancy text,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE visitor_entries (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  site_id uuid NOT NULL,
  kind text NOT NULL CHECK (kind IN ('visitor', 'contractor')),
  full_name text NOT NULL,
  organisation text,
  contact text,
  id_reference text,
  purpose text NOT NULL,
  host_employee_id uuid,
  work_order_id uuid,
  work_area text,
  induction_done boolean NOT NULL DEFAULT false,
  badge text,
  arrived_at timestamptz NOT NULL,
  departed_at timestamptz,
  status text NOT NULL DEFAULT 'on_site' CHECK (status IN ('on_site', 'departed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE lost_property_items (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  item_type text NOT NULL CHECK (item_type IN ('bank_card', 'personal_item')),
  description text NOT NULL,
  found_at timestamptz NOT NULL,
  found_location text NOT NULL,
  finder_employee_id uuid NOT NULL,
  custodian_employee_id uuid NOT NULL,
  storage_location text,
  issuer_if_visible text,
  name_if_visible text,
  disposition_due date,
  contact_attempts jsonb,
  claimant_name text,
  claimant_evidence text,
  released_at timestamptz,
  disposition text CHECK (disposition IN ('returned', 'bank_surrendered', 'destroyed', 'unclaimed_disposed')),
  status text NOT NULL DEFAULT 'in_custody' CHECK (status IN ('in_custody', 'returned', 'disposed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE documents (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  document_number text NOT NULL,
  title text NOT NULL,
  document_type text NOT NULL CHECK (document_type IN ('policy', 'sop', 'work_instruction', 'form', 'register', 'manual', 'guideline', 'objective')),
  department_id uuid NOT NULL,
  owner_employee_id uuid,
  current_revision_id uuid,
  issue_date date,
  next_review_on date,
  storage_location text,
  applicable_site_ids uuid[],
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'in_review', 'approved', 'issued', 'obsolete')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE document_revisions (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  document_id uuid NOT NULL,
  revision text NOT NULL,
  revision_date date,
  summary_of_change text NOT NULL,
  file_attachment_id uuid NOT NULL,
  prepared_by uuid,
  reviewed_by uuid,
  temporary boolean NOT NULL DEFAULT false,
  expires_on date,
  requires_acknowledgement boolean NOT NULL DEFAULT false,
  training_required boolean NOT NULL DEFAULT false,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'in_review', 'qms_review', 'pending_approval', 'approved', 'issued', 'withdrawn')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE document_requests (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  raised_by_employee_id uuid NOT NULL,
  department_id uuid NOT NULL,
  process text NOT NULL,
  gap text NOT NULL,
  trigger text NOT NULL CHECK (trigger IN ('new_process', 'audit_finding', 'nonconformance', 'corrective_action', 'legal_change', 'customer_change', 'improvement', 'other')),
  proposed_type text,
  target_on date,
  extension_reason text,
  document_id uuid,
  status text NOT NULL DEFAULT 'raised' CHECK (status IN ('raised', 'classified', 'drafting', 'in_review', 'approved', 'closed', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE document_acknowledgements (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  document_revision_id uuid NOT NULL,
  employee_id uuid NOT NULL,
  due_on date,
  acknowledged_at timestamptz,
  signature_id uuid,
  status text NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'acknowledged', 'overdue')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE courses (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  code text NOT NULL,
  title text NOT NULL,
  category text NOT NULL CHECK (category IN ('induction', 'food_safety', 'she', 'security', 'operational', 'customer_service', 'supervisory', 'product_knowledge', 'other')),
  objective text,
  document_ids uuid[],
  pass_mark integer,
  requires_practical boolean NOT NULL DEFAULT false,
  refresher_rule_key text,
  material_attachment_ids uuid[],
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE (code)
);

CREATE TABLE training_requirements (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  role_id uuid NOT NULL,
  course_id uuid NOT NULL,
  department_id uuid,
  mandatory boolean NOT NULL DEFAULT false,
  refresher_months integer,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE training_sessions (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  course_id uuid NOT NULL,
  site_id uuid,
  venue text,
  starts_at timestamptz NOT NULL,
  ends_at timestamptz,
  shift text,
  trainer_employee_id uuid,
  external_trainer text,
  capacity integer,
  cost_amount numeric(18,4),
  cost_currency text,
  status text NOT NULL DEFAULT 'planned' CHECK (status IN ('planned', 'held', 'cancelled')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE training_participants (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  training_session_id uuid NOT NULL,
  employee_id uuid NOT NULL,
  attendance text CHECK (attendance IN ('invited', 'attended', 'absent', 'excused')),
  checked_in_at timestamptz,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE training_assessments (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  employee_id uuid NOT NULL,
  course_id uuid NOT NULL,
  training_session_id uuid,
  kind text NOT NULL CHECK (kind IN ('knowledge', 'practical')),
  score integer,
  attempt integer,
  "result" text NOT NULL CHECK ("result" IN ('pass', 'fail', 'not_yet_competent')),
  assessor_employee_id uuid,
  assessed_on date NOT NULL,
  evidence_attachment_id uuid,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE competencies (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  employee_id uuid NOT NULL,
  course_id uuid NOT NULL,
  "status" text NOT NULL CHECK ("status" IN ('not_started', 'in_progress', 'competent', 'expired', 'restricted')),
  achieved_on date,
  expires_on date,
  assessor_employee_id uuid,
  restriction text,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE certificates (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  employee_id uuid NOT NULL,
  course_name text NOT NULL,
  provider text,
  certificate_number text,
  issued_on date,
  expires_on date,
  attachment_id uuid,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE training_requests (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid,
  topic text NOT NULL,
  "reason" text NOT NULL,
  target_group text,
  estimated_cost_amount numeric(18,4),
  estimated_cost_currency text,
  proposed_date date,
  source_type text,
  source_id uuid,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'scheduled', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE leave_types (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  code text NOT NULL,
  "name" text NOT NULL,
  accrual_rule_key text,
  requires_document boolean NOT NULL DEFAULT false,
  restricted boolean NOT NULL DEFAULT false,
  active boolean NOT NULL DEFAULT true,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1,
  UNIQUE (code)
);

CREATE TABLE leave_balances (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  employee_id uuid,
  leave_type_id uuid,
  period text,
  opening numeric(18,4),
  accrued numeric(18,4),
  taken numeric(18,4),
  adjusted numeric(18,4),
  closing numeric(18,4),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE leave_requests (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  employee_id uuid NOT NULL,
  leave_type_id uuid NOT NULL,
  start_date date NOT NULL,
  end_date date NOT NULL,
  working_days numeric(18,4),
  "reason" text,
  document_attachment_id uuid,
  coverage_gap boolean NOT NULL DEFAULT false,
  coverage_plan text,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'rejected', 'cancelled', 'taken')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE exit_clearances (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  employee_id uuid NOT NULL,
  exit_type text NOT NULL CHECK (exit_type IN ('resignation', 'termination', 'retirement', 'other')),
  notice_date date,
  last_working_day date NOT NULL,
  reason_category text,
  hr_officer_employee_id uuid,
  access_revoked_at timestamptz,
  payroll_handover_at timestamptz,
  exit_interview_notes text,
  status text NOT NULL DEFAULT 'initiated' CHECK (status IN ('initiated', 'in_clearance', 'cleared', 'completed', 'cancelled')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE exit_clearance_sections (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  exit_clearance_id uuid NOT NULL,
  section text NOT NULL CHECK (section IN ('stores', 'finance', 'ict', 'security', 'uniform', 'department', 'keys')),
  "status" text CHECK ("status" IN ('pending', 'cleared', 'outstanding')),
  outstanding_items text,
  signed_by uuid,
  signed_at timestamptz,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE rosters (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  site_id uuid NOT NULL,
  period_start date NOT NULL,
  period_end date NOT NULL,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'published')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE roster_shifts (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  roster_id uuid NOT NULL,
  employee_id uuid NOT NULL,
  shift_date date NOT NULL,
  starts_at text NOT NULL,
  ends_at text NOT NULL,
  "role" text,
  swapped_with_employee_id uuid,
  swap_reason text,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE attendance_records (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  employee_id uuid NOT NULL,
  site_id uuid NOT NULL,
  work_date date NOT NULL,
  clock_in timestamptz,
  clock_out timestamptz,
  "status" text NOT NULL CHECK ("status" IN ('present', 'late', 'absent', 'no_show', 'leave')),
  "reason" text,
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE uniform_issues (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  employee_id uuid NOT NULL,
  kind text NOT NULL CHECK (kind IN ('initial', 'six_month', 'early_replacement', 'return')),
  items jsonb NOT NULL,
  sizes jsonb,
  issued_on date NOT NULL,
  next_replacement_due date,
  "reason" text,
  old_pair_returned boolean NOT NULL DEFAULT false,
  site_id uuid,
  status text NOT NULL DEFAULT 'issued' CHECK (status IN ('issued', 'returned', 'cancelled')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE recruitment_requisitions (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  department_id uuid NOT NULL,
  site_id uuid,
  job_title text NOT NULL,
  headcount integer NOT NULL,
  "reason" text NOT NULL,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'pending_approval', 'approved', 'open', 'filled', 'cancelled')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE candidates (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  recruitment_requisition_id uuid NOT NULL,
  full_name text NOT NULL,
  contact text,
  cv_attachment_id uuid,
  assessment text,
  offer_status text CHECK (offer_status IN ('none', 'issued', 'accepted', 'declined')),
  status text NOT NULL DEFAULT 'applied' CHECK (status IN ('applied', 'shortlisted', 'interviewed', 'offered', 'hired', 'rejected')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE performance_reviews (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  employee_id uuid NOT NULL,
  reviewer_employee_id uuid NOT NULL,
  period text NOT NULL,
  objectives jsonb,
  rating text,
  comments text,
  improvement_plan text,
  status text NOT NULL DEFAULT 'draft' CHECK (status IN ('draft', 'in_review', 'completed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE disciplinary_cases (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  employee_id uuid NOT NULL,
  allegation text NOT NULL,
  source_type text,
  source_id uuid,
  stage text,
  hearing_on date,
  outcome text,
  closed_at timestamptz,
  status text NOT NULL DEFAULT 'registered' CHECK (status IN ('registered', 'investigating', 'hearing', 'decided', 'appeal', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE TABLE grievances (
  id uuid PRIMARY KEY DEFAULT uuidv7(),
  reference text NOT NULL UNIQUE,
  employee_id uuid NOT NULL,
  subject text NOT NULL,
  investigator_employee_id uuid,
  outcome text,
  closed_at timestamptz,
  status text NOT NULL DEFAULT 'received' CHECK (status IN ('received', 'investigating', 'closed')),
  created_at timestamptz NOT NULL DEFAULT now(),
  updated_at timestamptz NOT NULL DEFAULT now(),
  created_by uuid,
  updated_by uuid,
  row_version integer NOT NULL DEFAULT 1
);

CREATE INDEX ON assets (site_id);
CREATE INDEX ON assets (status);
CREATE INDEX ON capex_requests (site_id);
CREATE INDEX ON capex_requests (status);
CREATE INDEX ON asset_transfers (status);
CREATE INDEX ON asset_verifications (site_id);
CREATE INDEX ON asset_verifications (status);
CREATE INDEX ON asset_verification_lines (asset_verification_id);
CREATE INDEX ON asset_disposals (status);
CREATE INDEX ON key_items (site_id);
CREATE INDEX ON key_items (status);
CREATE INDEX ON key_movements (status);
CREATE INDEX ON key_duplication_requests (site_id);
CREATE INDEX ON key_duplication_requests (status);
CREATE INDEX ON key_audits (site_id);
CREATE INDEX ON key_audits (status);
CREATE INDEX ON key_audit_lines (key_audit_id);
CREATE INDEX ON visitor_entries (site_id);
CREATE INDEX ON visitor_entries (status);
CREATE INDEX ON lost_property_items (site_id);
CREATE INDEX ON lost_property_items (status);
CREATE INDEX ON documents (status);
CREATE INDEX ON document_revisions (status);
CREATE INDEX ON document_revisions (document_id);
CREATE INDEX ON document_requests (status);
CREATE INDEX ON document_acknowledgements (status);
CREATE INDEX ON document_acknowledgements (employee_id);
CREATE INDEX ON training_sessions (site_id);
CREATE INDEX ON training_sessions (status);
CREATE INDEX ON training_participants (employee_id);
CREATE INDEX ON training_participants (training_session_id);
CREATE INDEX ON training_assessments (employee_id);
CREATE INDEX ON competencies (employee_id);
CREATE INDEX ON certificates (employee_id);
CREATE INDEX ON training_requests (site_id);
CREATE INDEX ON training_requests (status);
CREATE INDEX ON leave_balances (employee_id);
CREATE INDEX ON leave_requests (status);
CREATE INDEX ON leave_requests (employee_id);
CREATE INDEX ON exit_clearances (status);
CREATE INDEX ON exit_clearances (employee_id);
CREATE INDEX ON exit_clearance_sections (exit_clearance_id);
CREATE INDEX ON rosters (site_id);
CREATE INDEX ON rosters (status);
CREATE INDEX ON roster_shifts (employee_id);
CREATE INDEX ON roster_shifts (roster_id);
CREATE INDEX ON attendance_records (site_id);
CREATE INDEX ON attendance_records (employee_id);
CREATE INDEX ON uniform_issues (site_id);
CREATE INDEX ON uniform_issues (status);
CREATE INDEX ON uniform_issues (employee_id);
CREATE INDEX ON recruitment_requisitions (site_id);
CREATE INDEX ON recruitment_requisitions (status);
CREATE INDEX ON candidates (status);
CREATE INDEX ON performance_reviews (status);
CREATE INDEX ON performance_reviews (employee_id);
CREATE INDEX ON disciplinary_cases (status);
CREATE INDEX ON disciplinary_cases (employee_id);
CREATE INDEX ON grievances (status);
CREATE INDEX ON grievances (employee_id);

ALTER TABLE approval_policies ADD FOREIGN KEY (source_document_id) REFERENCES documents (id);
ALTER TABLE asset_disposals ADD FOREIGN KEY (asset_id) REFERENCES assets (id);
ALTER TABLE asset_disposals ADD FOREIGN KEY (proceeds_currency) REFERENCES currencies (code);
ALTER TABLE asset_transfers ADD FOREIGN KEY (asset_id) REFERENCES assets (id);
ALTER TABLE asset_transfers ADD FOREIGN KEY (to_cost_centre_id) REFERENCES cost_centres (id);
ALTER TABLE asset_transfers ADD FOREIGN KEY (to_custodian_employee_id) REFERENCES employees (id);
ALTER TABLE asset_transfers ADD FOREIGN KEY (to_site_id) REFERENCES sites (id);
ALTER TABLE asset_verification_lines ADD FOREIGN KEY (asset_id) REFERENCES assets (id) ON DELETE CASCADE;
ALTER TABLE asset_verifications ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE assets ADD FOREIGN KEY (asset_class_id) REFERENCES asset_classes (id);
ALTER TABLE assets ADD FOREIGN KEY (cost_centre_id) REFERENCES cost_centres (id);
ALTER TABLE assets ADD FOREIGN KEY (cost_currency) REFERENCES currencies (code);
ALTER TABLE assets ADD FOREIGN KEY (custodian_employee_id) REFERENCES employees (id);
ALTER TABLE assets ADD FOREIGN KEY (department_id) REFERENCES departments (id);
ALTER TABLE assets ADD FOREIGN KEY (goods_receipt_id) REFERENCES goods_receipts (id);
ALTER TABLE assets ADD FOREIGN KEY (invoice_id) REFERENCES supplier_invoices (id);
ALTER TABLE assets ADD FOREIGN KEY (purchase_order_id) REFERENCES purchase_orders (id);
ALTER TABLE assets ADD FOREIGN KEY (residual_value_currency) REFERENCES currencies (code);
ALTER TABLE assets ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE assets ADD FOREIGN KEY (supplier_id) REFERENCES suppliers (id);
ALTER TABLE attendance_records ADD FOREIGN KEY (employee_id) REFERENCES employees (id);
ALTER TABLE attendance_records ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE candidates ADD FOREIGN KEY (cv_attachment_id) REFERENCES attachments (id);
ALTER TABLE candidates ADD FOREIGN KEY (recruitment_requisition_id) REFERENCES recruitment_requisitions (id);
ALTER TABLE capex_requests ADD FOREIGN KEY (amount_currency) REFERENCES currencies (code);
ALTER TABLE capex_requests ADD FOREIGN KEY (asset_id) REFERENCES assets (id);
ALTER TABLE capex_requests ADD FOREIGN KEY (department_id) REFERENCES departments (id);
ALTER TABLE capex_requests ADD FOREIGN KEY (purchase_requisition_id) REFERENCES purchase_requisitions (id);
ALTER TABLE capex_requests ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE cash_variances ADD FOREIGN KEY (disciplinary_case_id) REFERENCES disciplinary_cases (id);
ALTER TABLE certificates ADD FOREIGN KEY (attachment_id) REFERENCES attachments (id);
ALTER TABLE certificates ADD FOREIGN KEY (employee_id) REFERENCES employees (id);
ALTER TABLE competencies ADD FOREIGN KEY (assessor_employee_id) REFERENCES employees (id);
ALTER TABLE competencies ADD FOREIGN KEY (course_id) REFERENCES courses (id);
ALTER TABLE competencies ADD FOREIGN KEY (employee_id) REFERENCES employees (id);
ALTER TABLE disciplinary_cases ADD FOREIGN KEY (employee_id) REFERENCES employees (id);
ALTER TABLE document_acknowledgements ADD FOREIGN KEY (document_revision_id) REFERENCES document_revisions (id);
ALTER TABLE document_acknowledgements ADD FOREIGN KEY (employee_id) REFERENCES employees (id);
ALTER TABLE document_acknowledgements ADD FOREIGN KEY (signature_id) REFERENCES signature_requests (id);
ALTER TABLE document_requests ADD FOREIGN KEY (department_id) REFERENCES departments (id);
ALTER TABLE document_requests ADD FOREIGN KEY (document_id) REFERENCES documents (id);
ALTER TABLE document_requests ADD FOREIGN KEY (raised_by_employee_id) REFERENCES employees (id);
ALTER TABLE document_revisions ADD FOREIGN KEY (file_attachment_id) REFERENCES attachments (id);
ALTER TABLE document_revisions ADD FOREIGN KEY (prepared_by) REFERENCES users (id);
ALTER TABLE document_revisions ADD FOREIGN KEY (reviewed_by) REFERENCES users (id);
ALTER TABLE documents ADD FOREIGN KEY (current_revision_id) REFERENCES document_revisions (id);
ALTER TABLE documents ADD FOREIGN KEY (department_id) REFERENCES departments (id);
ALTER TABLE documents ADD FOREIGN KEY (owner_employee_id) REFERENCES employees (id);
ALTER TABLE equipment_faults ADD FOREIGN KEY (asset_id) REFERENCES assets (id);
ALTER TABLE exit_clearance_sections ADD FOREIGN KEY (signed_by) REFERENCES users (id) ON DELETE CASCADE;
ALTER TABLE exit_clearances ADD FOREIGN KEY (employee_id) REFERENCES employees (id);
ALTER TABLE exit_clearances ADD FOREIGN KEY (hr_officer_employee_id) REFERENCES employees (id);
ALTER TABLE grievances ADD FOREIGN KEY (employee_id) REFERENCES employees (id);
ALTER TABLE grievances ADD FOREIGN KEY (investigator_employee_id) REFERENCES employees (id);
ALTER TABLE key_audit_lines ADD FOREIGN KEY (expected_custodian_employee_id) REFERENCES employees (id) ON DELETE CASCADE;
ALTER TABLE key_audit_lines ADD FOREIGN KEY (key_item_id) REFERENCES key_items (id) ON DELETE CASCADE;
ALTER TABLE key_audits ADD FOREIGN KEY (auditor_employee_id) REFERENCES employees (id);
ALTER TABLE key_audits ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE key_duplication_requests ADD FOREIGN KEY (memo_attachment_id) REFERENCES attachments (id);
ALTER TABLE key_duplication_requests ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE key_items ADD FOREIGN KEY (primary_custodian_employee_id) REFERENCES employees (id);
ALTER TABLE key_items ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE key_movements ADD FOREIGN KEY (holder_employee_id) REFERENCES employees (id);
ALTER TABLE key_movements ADD FOREIGN KEY (key_item_id) REFERENCES key_items (id);
ALTER TABLE leave_balances ADD FOREIGN KEY (employee_id) REFERENCES employees (id);
ALTER TABLE leave_balances ADD FOREIGN KEY (leave_type_id) REFERENCES leave_types (id);
ALTER TABLE leave_requests ADD FOREIGN KEY (document_attachment_id) REFERENCES attachments (id);
ALTER TABLE leave_requests ADD FOREIGN KEY (employee_id) REFERENCES employees (id);
ALTER TABLE leave_requests ADD FOREIGN KEY (leave_type_id) REFERENCES leave_types (id);
ALTER TABLE lost_property_items ADD FOREIGN KEY (custodian_employee_id) REFERENCES employees (id);
ALTER TABLE lost_property_items ADD FOREIGN KEY (finder_employee_id) REFERENCES employees (id);
ALTER TABLE lost_property_items ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE maintenance_plans ADD FOREIGN KEY (asset_id) REFERENCES assets (id);
ALTER TABLE performance_reviews ADD FOREIGN KEY (employee_id) REFERENCES employees (id);
ALTER TABLE performance_reviews ADD FOREIGN KEY (reviewer_employee_id) REFERENCES employees (id);
ALTER TABLE product_changes ADD FOREIGN KEY (training_request_id) REFERENCES training_requests (id);
ALTER TABLE recipe_versions ADD FOREIGN KEY (method_document_id) REFERENCES documents (id);
ALTER TABLE recruitment_requisitions ADD FOREIGN KEY (department_id) REFERENCES departments (id);
ALTER TABLE recruitment_requisitions ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE roster_shifts ADD FOREIGN KEY (employee_id) REFERENCES employees (id) ON DELETE CASCADE;
ALTER TABLE roster_shifts ADD FOREIGN KEY (swapped_with_employee_id) REFERENCES employees (id) ON DELETE CASCADE;
ALTER TABLE rosters ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE rule_versions ADD FOREIGN KEY (source_document_id) REFERENCES documents (id);
ALTER TABLE training_assessments ADD FOREIGN KEY (assessor_employee_id) REFERENCES employees (id);
ALTER TABLE training_assessments ADD FOREIGN KEY (course_id) REFERENCES courses (id);
ALTER TABLE training_assessments ADD FOREIGN KEY (employee_id) REFERENCES employees (id);
ALTER TABLE training_assessments ADD FOREIGN KEY (evidence_attachment_id) REFERENCES attachments (id);
ALTER TABLE training_assessments ADD FOREIGN KEY (training_session_id) REFERENCES training_sessions (id);
ALTER TABLE training_participants ADD FOREIGN KEY (employee_id) REFERENCES employees (id) ON DELETE CASCADE;
ALTER TABLE training_requests ADD FOREIGN KEY (estimated_cost_currency) REFERENCES currencies (code);
ALTER TABLE training_requests ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE training_requirements ADD FOREIGN KEY (course_id) REFERENCES courses (id);
ALTER TABLE training_requirements ADD FOREIGN KEY (department_id) REFERENCES departments (id);
ALTER TABLE training_requirements ADD FOREIGN KEY (role_id) REFERENCES roles (id);
ALTER TABLE training_sessions ADD FOREIGN KEY (cost_currency) REFERENCES currencies (code);
ALTER TABLE training_sessions ADD FOREIGN KEY (course_id) REFERENCES courses (id);
ALTER TABLE training_sessions ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE training_sessions ADD FOREIGN KEY (trainer_employee_id) REFERENCES employees (id);
ALTER TABLE uniform_issues ADD FOREIGN KEY (employee_id) REFERENCES employees (id);
ALTER TABLE uniform_issues ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE visitor_entries ADD FOREIGN KEY (host_employee_id) REFERENCES employees (id);
ALTER TABLE visitor_entries ADD FOREIGN KEY (site_id) REFERENCES sites (id);
ALTER TABLE visitor_entries ADD FOREIGN KEY (work_order_id) REFERENCES work_orders (id);
ALTER TABLE work_orders ADD FOREIGN KEY (asset_id) REFERENCES assets (id);

-- +goose Down

DROP TABLE IF EXISTS grievances CASCADE;
DROP TABLE IF EXISTS disciplinary_cases CASCADE;
DROP TABLE IF EXISTS performance_reviews CASCADE;
DROP TABLE IF EXISTS candidates CASCADE;
DROP TABLE IF EXISTS recruitment_requisitions CASCADE;
DROP TABLE IF EXISTS uniform_issues CASCADE;
DROP TABLE IF EXISTS attendance_records CASCADE;
DROP TABLE IF EXISTS roster_shifts CASCADE;
DROP TABLE IF EXISTS rosters CASCADE;
DROP TABLE IF EXISTS exit_clearance_sections CASCADE;
DROP TABLE IF EXISTS exit_clearances CASCADE;
DROP TABLE IF EXISTS leave_requests CASCADE;
DROP TABLE IF EXISTS leave_balances CASCADE;
DROP TABLE IF EXISTS leave_types CASCADE;
DROP TABLE IF EXISTS training_requests CASCADE;
DROP TABLE IF EXISTS certificates CASCADE;
DROP TABLE IF EXISTS competencies CASCADE;
DROP TABLE IF EXISTS training_assessments CASCADE;
DROP TABLE IF EXISTS training_participants CASCADE;
DROP TABLE IF EXISTS training_sessions CASCADE;
DROP TABLE IF EXISTS training_requirements CASCADE;
DROP TABLE IF EXISTS courses CASCADE;
DROP TABLE IF EXISTS document_acknowledgements CASCADE;
DROP TABLE IF EXISTS document_requests CASCADE;
DROP TABLE IF EXISTS document_revisions CASCADE;
DROP TABLE IF EXISTS documents CASCADE;
DROP TABLE IF EXISTS lost_property_items CASCADE;
DROP TABLE IF EXISTS visitor_entries CASCADE;
DROP TABLE IF EXISTS key_audit_lines CASCADE;
DROP TABLE IF EXISTS key_audits CASCADE;
DROP TABLE IF EXISTS key_duplication_requests CASCADE;
DROP TABLE IF EXISTS key_movements CASCADE;
DROP TABLE IF EXISTS key_items CASCADE;
DROP TABLE IF EXISTS asset_disposals CASCADE;
DROP TABLE IF EXISTS asset_verification_lines CASCADE;
DROP TABLE IF EXISTS asset_verifications CASCADE;
DROP TABLE IF EXISTS asset_transfers CASCADE;
DROP TABLE IF EXISTS capex_requests CASCADE;
DROP TABLE IF EXISTS assets CASCADE;
DROP TABLE IF EXISTS asset_classes CASCADE;
