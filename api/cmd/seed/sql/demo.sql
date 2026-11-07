-- A small demo organisation for local development and flow testing. The source
-- documents do not list sites or staff (spec N-09), so everything here is made
-- up. Safe to run repeatedly. Never load this into production.

INSERT INTO complexes (code, name, active) VALUES
  ('HRE', 'Harare Central', true),
  ('BYO', 'Bulawayo', true)
ON CONFLICT (code) DO NOTHING;

INSERT INTO sites (code, name, type, complex_id, brand, trading_pattern, opens_at, closes_at, stockholding, address, active)
SELECT v.code, v.name, v.type, c.id, 'Foodies', v.pattern, v.opens, v.closes, v.stock, v.address, true
FROM (VALUES
  ('HO',    'Head Office',      'head_office',     NULL,  'non_trading', NULL,    NULL,    false, '4 Enterprise Road, Harare'),
  ('HRE01', 'Samora Machel',    'branch',          'HRE', 'closes_2200', '07:00', '22:00', true,  '112 Samora Machel Ave, Harare'),
  ('HRE02', 'Avondale',         'branch',          'HRE', 'open_24h',    '00:00', '23:59', true,  'Avondale Shops, Harare'),
  ('BYO01', 'Fife Street',      'branch',          'BYO', 'closes_2200', '07:00', '22:00', true,  '88 Fife St, Bulawayo'),
  ('CK01',  'Central Kitchen',  'central_kitchen', NULL,  'non_trading', '05:00', '18:00', true,  '9 Factory Rd, Harare'),
  ('STR01', 'Central Store',    'central_store',   NULL,  'non_trading', '06:00', '17:00', true,  '9 Factory Rd, Harare')
) AS v(code, name, type, complex, pattern, opens, closes, stock, address)
LEFT JOIN complexes c ON c.code = v.complex
ON CONFLICT (code) DO NOTHING;

INSERT INTO departments (code, name, active) VALUES
  ('OPS', 'Operations', true), ('FIN', 'Finance and Accounts', true), ('HR', 'Human Resources', true),
  ('ICT', 'ICT', true), ('PRC', 'Procurement', true), ('QA', 'Quality and Food Safety', true),
  ('STR', 'Stores', true), ('PRD', 'Production', true), ('SEC', 'Security', true), ('MNT', 'Maintenance', true),
  ('EXE', 'Executive', true), ('SHQ', 'SHEQ', true), ('TRN', 'Training', true), ('QMS', 'QMS and Document Control', true),
  ('AUD', 'Internal Audit', true)
ON CONFLICT (code) DO NOTHING;

INSERT INTO tills (site_id, code, pos_terminal_ref, active)
SELECT s.id, t.code, s.code || '-' || t.code, true
FROM sites s CROSS JOIN (VALUES ('T1'), ('T2')) AS t(code)
WHERE s.type = 'branch'
ON CONFLICT (site_id, code) DO NOTHING;

INSERT INTO safes (site_id, code, location, lock_status, active)
SELECT id, 'SAFE1', 'Manager''s office', 'ok', true FROM sites WHERE type = 'branch'
ON CONFLICT (site_id, code) DO NOTHING;

WITH staff (number, first_name, last_name, title, dept, site, role, scope_type, scope) AS (VALUES
  ('E1001', 'Tendai',    'Moyo',     'System Administrator', 'ICT', 'HO',    'system_admin',       'all',     NULL),
  ('E1002', 'Rumbi',     'Chikomo',  'Managing Director',    'EXE', 'HO',    'md',                 'all',     NULL),
  ('E1003', 'Farai',     'Ncube',    'General Manager',      'EXE', 'HO',    'gm',                 'all',     NULL),
  ('E1004', 'Nyasha',    'Dube',     'Finance Manager',      'FIN', 'HO',    'finance_manager',    'all',     NULL),
  ('E1005', 'Tatenda',   'Sibanda',  'Accounts Clerk',       'FIN', 'HO',    'accountant',         'all',     NULL),
  ('E1006', 'Kudzai',    'Mhlanga',  'Procurement Officer',  'PRC', 'HO',    'procurement_officer','all',     NULL),
  ('E1007', 'Chipo',     'Banda',    'Operations Manager',   'OPS', 'HO',    'operations_manager', 'all',     NULL),
  ('E1008', 'Tafadzwa',  'Mutasa',   'Complex Manager',      'OPS', 'HRE01', 'complex_manager',    'complex', 'HRE'),
  ('E1009', 'Rudo',      'Nkomo',    'Branch Manager',       'OPS', 'HRE01', 'branch_manager',     'site',    'HRE01'),
  ('E1010', 'Simba',     'Gumbo',    'Shift Supervisor',     'OPS', 'HRE01', 'shift_supervisor',   'site',    'HRE01'),
  ('E1011', 'Anesu',     'Zhou',     'Cashier',              'OPS', 'HRE01', 'cashier',            'site',    'HRE01'),
  ('E1012', 'Tariro',    'Mpofu',    'Cashier',              'OPS', 'HRE02', 'cashier',            'site',    'HRE02'),
  ('E1013', 'Blessing',  'Shumba',   'Storekeeper',          'STR', 'STR01', 'storekeeper',        'site',    'STR01'),
  ('E1014', 'Ropafadzo', 'Tshuma',   'Head Chef',            'PRD', 'CK01',  'production_lead',    'site',    'CK01'),
  ('E1015', 'Munashe',   'Hove',     'Driver',               'STR', 'CK01',  'driver',             'all',     NULL),
  ('E1016', 'Vimbai',    'Marufu',   'Security Officer',     'SEC', 'HRE01', 'security_officer',   'site',    'HRE01'),
  ('E1017', 'Takudzwa',  'Chirwa',   'Quality Assurance',    'QA',  'HO',    'qa',                 'all',     NULL),
  ('E1018', 'Sekai',     'Nyoni',    'Maintenance Officer',  'MNT', 'HO',    'maintenance',        'all',     NULL),
  ('E1019', 'Tinashe',   'Ndlovu',   'HR Officer',           'HR',  'HO',    'hr_officer',         'all',     NULL),
  ('E1020', 'Yeukai',    'Madondo',  'Internal Auditor',     'AUD', 'HO',    'internal_audit',     'all',     NULL),
  ('E1021', 'Nomsa',     'Khumalo',  'Branch Manager',       'OPS', 'BYO01', 'branch_manager',     'site',    'BYO01'),
  ('E1022', 'Lovemore',  'Sithole',  'Crew Member',          'OPS', 'HRE01', 'employee',           'site',    'HRE01')
), emp AS (
  INSERT INTO employees (reference, employee_number, first_name, last_name, job_title, department_id, site_id, employment_status, start_date, contract_type, phone_e164)
  SELECT 'EMP-' || s.site || '-2026-' || lpad(substr(st.number, 2), 6, '0'), st.number, st.first_name, st.last_name, st.title, d.id, s2.id, 'active', DATE '2024-01-15', 'permanent',
         '+26377' || substr(st.number, 2) || '000'
  FROM staff st
  CROSS JOIN LATERAL (SELECT st.site) s
  JOIN departments d ON d.code = st.dept
  JOIN sites s2 ON s2.code = st.site
  ON CONFLICT (employee_number) DO NOTHING
  RETURNING id, employee_number
), usr AS (
  INSERT INTO users (employee_id, channels, status, credential_status)
  SELECT id, 'both', 'active', 'active' FROM emp
  RETURNING id, employee_id
)
INSERT INTO role_assignments (user_id, role_id, scope_type, scope_id, effective_from)
SELECT u.id, r.id, st.scope_type,
       CASE st.scope_type WHEN 'site' THEN (SELECT id FROM sites WHERE code = st.scope)
                          WHEN 'complex' THEN (SELECT id FROM complexes WHERE code = st.scope) END,
       DATE '2024-01-15'
FROM usr u
JOIN emp e ON e.id = u.employee_id
JOIN staff st ON st.number = e.employee_number
JOIN roles r ON r."key" = st.role;

UPDATE complexes c SET manager_employee_id = e.id FROM employees e
WHERE c.code = 'HRE' AND e.employee_number = 'E1008' AND c.manager_employee_id IS NULL;

-- Approval routes. The real chains wait for the Authorised Limit Matrix (spec
-- 8.3), so the demo sends everything to the General Manager in one step.
INSERT INTO approval_policies (workflow_key, name, version, status, effective_from)
SELECT t.key, 'Demo route', 1, 'active', now()
FROM (VALUES
  ('employee_transfers'),
  ('delegations'),
  ('rule_versions'),
  ('approval_policies'),
  ('form_template_versions'),
  ('disposition_batches'),
  ('float_certificates'),
  ('float_exceptions'),
  ('petty_cash_vouchers'),
  ('supplier_amendments'),
  ('purchase_requisitions'),
  ('purchase_orders'),
  ('payment_requisitions'),
  ('payment_batches'),
  ('emergency_procurements'),
  ('budgets'),
  ('budget_reallocations'),
  ('stock_adjustments'),
  ('disposals'),
  ('recipe_versions'),
  ('dispositions'),
  ('recall_communications'),
  ('service_recoveries'),
  ('capex_requests'),
  ('asset_transfers'),
  ('asset_disposals'),
  ('key_duplication_requests'),
  ('document_revisions'),
  ('training_requests'),
  ('leave_requests'),
  ('rosters'),
  ('recruitment_requisitions'),
  ('footage_requests'),
  ('product_changes'),
  ('cost_sheets'),
  ('external_visits')
) AS t(key)
WHERE NOT EXISTS (SELECT 1 FROM approval_policies p WHERE p.workflow_key = t.key);

INSERT INTO approval_policy_steps (approval_policy_id, sequence, role_id, mode)
SELECT p.id, 1, r.id, 'any' FROM approval_policies p JOIN roles r ON r."key" = 'gm'
WHERE NOT EXISTS (SELECT 1 FROM approval_policy_steps s WHERE s.approval_policy_id = p.id);

-- Escalation recipients and timers are unset rules (C-06.09). These are made up.
INSERT INTO escalation_policies (domain, severity_level, label, status)
SELECT d.domain, l.level, l.label, 'active'
FROM (VALUES ('quality'), ('customer'), ('security'), ('cash'), ('she'), ('equipment'), ('general')) AS d(domain)
CROSS JOIN (VALUES (1, 'Routine'), (2, 'Significant'), (3, 'Serious'), (4, 'Critical'), (5, 'Executive critical')) AS l(level, label)
WHERE NOT EXISTS (SELECT 1 FROM escalation_policies e WHERE e.domain = d.domain AND e.severity_level = l.level);

INSERT INTO escalation_policy_steps (escalation_policy_id, sequence, notify_role_id, channels, ack_within_minutes, next_on_timeout)
SELECT e.id, s.sequence, r.id, ARRAY['in_app'], 15, s.more
FROM escalation_policies e
CROSS JOIN (VALUES (1, 'shift_supervisor', true), (2, 'complex_manager', true), (3, 'md', false)) AS s(sequence, role, more)
JOIN roles r ON r."key" = s.role
WHERE NOT EXISTS (SELECT 1 FROM escalation_policy_steps x WHERE x.escalation_policy_id = e.id);

-- A branch opening checklist on the form engine, published, with one daily
-- schedule at Samora Machel.
INSERT INTO form_templates ("key", title, kind, module, area, status)
VALUES ('store_opening', 'Store opening checklist', 'checklist', 'store_ops', 'Store Opening', 'draft')
ON CONFLICT ("key") DO NOTHING;

INSERT INTO form_template_versions (form_template_id, version, status, channels, offline_allowed, definition)
SELECT t.id, 1, 'published', 'both', true, $json$
{"sections": [
  {"title": "Front of house", "fields": [
    {"key": "floor_clean", "label": "Floor and tables clean", "type": "pass_fail", "required": true,
     "on_fail": {"create": "action"}},
    {"key": "tills_on", "label": "Tills powered and logged in", "type": "ok_nok", "required": true},
    {"key": "float_counted", "label": "Till floats counted", "type": "pass_fail", "required": true}
  ]},
  {"title": "Kitchen", "fields": [
    {"key": "fridge_temp", "label": "Walk-in fridge temperature (C)", "type": "reading", "required": true, "max": 5,
     "on_fail": {"create": "escalation", "severity": 2, "domain": "quality"}},
    {"key": "oil_quality", "label": "Fryer oil condition", "type": "choice", "options": ["Good", "Change today", "Change now"], "required": true},
    {"key": "notes", "label": "Anything else to report", "type": "text"}
  ]}
], "require_review": false}
$json$::jsonb
FROM form_templates t
WHERE t."key" = 'store_opening'
  AND NOT EXISTS (SELECT 1 FROM form_template_versions v WHERE v.form_template_id = t.id);

UPDATE form_templates t SET status = 'active', current_version_id = v.id
FROM form_template_versions v WHERE v.form_template_id = t.id AND t.current_version_id IS NULL;

INSERT INTO schedules (title, form_template_id, site_ids, assignee_role_id, due_offset_minutes, reminder_minutes, status)
SELECT 'Store opening checklist', t.id, ARRAY[(SELECT id FROM sites WHERE code = 'HRE01')], r.id, 60, 30, 'active'
FROM form_templates t JOIN roles r ON r."key" = 'shift_supervisor'
WHERE t."key" = 'store_opening' AND NOT EXISTS (SELECT 1 FROM schedules);
