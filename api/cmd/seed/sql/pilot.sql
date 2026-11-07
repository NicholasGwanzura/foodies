-- Pilot branch: the money trail is live at Samora Machel, and every other site
-- stays not ready, so both sides of readiness can be seen. Real values wait for
-- their owners' approval. The two with no stated value are made up.
INSERT INTO rule_versions (rule_definition_id, scope_type, value, effective_from, change_reason, status)
SELECT d.id, 'group', v.value, now() - interval '1 day', 'Demo value', 'active'
FROM rule_definitions d
JOIN LATERAL (
  SELECT COALESCE(
    (SELECT x.value FROM rule_versions x WHERE x.rule_definition_id = d.id ORDER BY x.id LIMIT 1),
    '"Friday"'::jsonb) AS value
) v ON true
WHERE d."key" IN (SELECT rule_key FROM workflow_requirements WHERE workflow_key LIKE 'cash.%')
  AND NOT EXISTS (SELECT 1 FROM rule_versions x WHERE x.rule_definition_id = d.id AND x.status = 'active');

INSERT INTO workflow_activations (workflow_key, site_id, status, missing_rule_keys, activated_at)
SELECT DISTINCT r.workflow_key, s.id, 'live', '{}'::text[], now()
FROM workflow_requirements r CROSS JOIN sites s
WHERE s.code = 'HRE01' AND r.workflow_key LIKE 'cash.%'
ON CONFLICT (workflow_key, site_id) DO UPDATE SET status = 'live', missing_rule_keys = '{}';

INSERT INTO fx_rates (from_currency, to_currency, rate, rate_date, source)
SELECT 'ZWG', 'USD', 0.0370, DATE '2026-11-01', 'manual'
WHERE NOT EXISTS (SELECT 1 FROM fx_rates);

INSERT INTO petty_cash_accounts (site_id, currency, approved_limit_amount, approved_limit_currency, book_balance_amount, book_balance_currency, custodian_employee_id)
SELECT s.id, 'USD', 1000, 'USD', 500, 'USD', e.id FROM sites s, employees e
WHERE s.code = 'HRE01' AND e.employee_number = 'E1009'
  AND NOT EXISTS (SELECT 1 FROM petty_cash_accounts);

INSERT INTO cit_providers (name, contract_reference, active)
SELECT 'Secure Cash Couriers', 'CIT-2026-01', true WHERE NOT EXISTS (SELECT 1 FROM cit_providers);

INSERT INTO cit_personnel (cit_provider_id, full_name, id_number, active)
SELECT p.id, 'Daniel Mapfumo', '63-123456A42', true FROM cit_providers p
WHERE NOT EXISTS (SELECT 1 FROM cit_personnel);

INSERT INTO items (code, name, category, stock_class, uom, critical, rotation, active) VALUES
  ('CHK-WHOLE', 'Whole chicken', 'Protein', 'frozen', 'kg', true, 'fefo', true),
  ('OIL-20L', 'Cooking oil 20L', 'Dry goods', 'dry', 'each', true, 'fifo', true),
  ('FLOUR-25', 'Bread flour 25kg', 'Dry goods', 'dry', 'each', false, 'fifo', true),
  ('BOX-M', 'Takeaway box, medium', 'Packaging', 'packaging', 'each', false, 'fifo', true)
ON CONFLICT (code) DO NOTHING;

INSERT INTO products (code, name, category, active) VALUES
  ('PIE-CHK', 'Chicken pie', 'Bakery', true),
  ('MEAL-2PC', 'Two-piece meal', 'Chicken and chips', true)
ON CONFLICT (code) DO NOTHING;

INSERT INTO suppliers (reference, legal_name, supplier_type, sap_code, tin, tax_clearance_expiry, status)
SELECT 'SUP-GRP-2026-000001', 'Mazoe Poultry Supplies', 'local', 'LOSUP0001', '2000123456', DATE '2027-06-30', 'active'
WHERE NOT EXISTS (SELECT 1 FROM suppliers);
