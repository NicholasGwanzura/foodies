-- name: SearchRecords :many
SELECT * FROM (
SELECT 'employees'::text AS type, id, reference, reference::text AS title, site_id AS site_id, NULL::text AS status FROM employees WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'form_submissions'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM form_submissions WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'actions'::text AS type, id, reference, description::text AS title, site_id AS site_id, status AS status FROM actions WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR description::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'float_certificates'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM float_certificates WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'float_issues'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM float_issues WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'float_exceptions'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM float_exceptions WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'till_flushes'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM till_flushes WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'cash_ups'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM cash_ups WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'cash_variances'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM cash_variances WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'cit_transits'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM cit_transits WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'bank_deposits'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM bank_deposits WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'cit_reconciliations'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM cit_reconciliations WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'surprise_counts'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM surprise_counts WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'petty_cash_vouchers'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM petty_cash_vouchers WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'petty_cash_handovers'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM petty_cash_handovers WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'petty_cash_reconciliations'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM petty_cash_reconciliations WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'petty_cash_reimbursements'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM petty_cash_reimbursements WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'suppliers'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM suppliers WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'purchase_requisitions'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM purchase_requisitions WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'purchase_orders'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM purchase_orders WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'goods_receipts'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM goods_receipts WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'delivery_discrepancies'::text AS type, id, reference, description::text AS title, NULL::uuid AS site_id, status AS status FROM delivery_discrepancies WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR description::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'supplier_invoices'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM supplier_invoices WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'payment_requisitions'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM payment_requisitions WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'payment_batches'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM payment_batches WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'emergency_procurements'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM emergency_procurements WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'unauthorized_purchases'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM unauthorized_purchases WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'budgets'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM budgets WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'budget_reallocations'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM budget_reallocations WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
) hits
WHERE sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[])
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);
