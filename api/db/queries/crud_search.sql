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
UNION ALL
SELECT 'stock_requisitions'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM stock_requisitions WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'stock_transfers'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM stock_transfers WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'stock_returns'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM stock_returns WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'stocktakes'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM stocktakes WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'stock_variances'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM stock_variances WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'stock_adjustments'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM stock_adjustments WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'waste_records'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM waste_records WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'disposals'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM disposals WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'kitchen_orders'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM kitchen_orders WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'production_plans'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM production_plans WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'dispatch_notes'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM dispatch_notes WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'product_changes'::text AS type, id, reference, description::text AS title, NULL::uuid AS site_id, status AS status FROM product_changes WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR description::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'cost_sheets'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM cost_sheets WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'npd_projects'::text AS type, id, reference, "name"::text AS title, NULL::uuid AS site_id, status AS status FROM npd_projects WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR "name"::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'ncrs'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM ncrs WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'holds'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM holds WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'dispositions'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM dispositions WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'recalls'::text AS type, id, reference, description::text AS title, site_id AS site_id, status AS status FROM recalls WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR description::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'incidents'::text AS type, id, reference, description::text AS title, site_id AS site_id, status AS status FROM incidents WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR description::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'complaints'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM complaints WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'service_recoveries'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM service_recoveries WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'payment_incidents'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM payment_incidents WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'pos_outages'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM pos_outages WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'equipment_faults'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM equipment_faults WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'work_orders'::text AS type, id, reference, description::text AS title, site_id AS site_id, status AS status FROM work_orders WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR description::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'reports'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM reports WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'management_meetings'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM management_meetings WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'compliance_obligations'::text AS type, id, reference, title::text AS title, site_id AS site_id, status AS status FROM compliance_obligations WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR title::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'capex_requests'::text AS type, id, reference, description::text AS title, site_id AS site_id, status AS status FROM capex_requests WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR description::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'asset_transfers'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM asset_transfers WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'asset_verifications'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM asset_verifications WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'asset_disposals'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM asset_disposals WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'key_duplication_requests'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM key_duplication_requests WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'key_audits'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM key_audits WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'lost_property_items'::text AS type, id, reference, description::text AS title, site_id AS site_id, status AS status FROM lost_property_items WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR description::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'document_requests'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM document_requests WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'training_sessions'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM training_sessions WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'training_requests'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM training_requests WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'leave_requests'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM leave_requests WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'exit_clearances'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM exit_clearances WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'rosters'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM rosters WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'uniform_issues'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM uniform_issues WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'recruitment_requisitions'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM recruitment_requisitions WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'disciplinary_cases'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM disciplinary_cases WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'grievances'::text AS type, id, reference, subject::text AS title, NULL::uuid AS site_id, status AS status FROM grievances WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR subject::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'store_openings'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM store_openings WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'occurrence_entries'::text AS type, id, reference, description::text AS title, site_id AS site_id, NULL::text AS status FROM occurrence_entries WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR description::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'patrol_rounds'::text AS type, id, reference, reference::text AS title, NULL::uuid AS site_id, status AS status FROM patrol_rounds WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'cctv_outages'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM cctv_outages WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'footage_requests'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM footage_requests WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'external_visits'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM external_visits WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
) hits
WHERE sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[])
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);
