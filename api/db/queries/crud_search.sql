-- name: SearchRecords :many
SELECT * FROM (
SELECT 'employees'::text AS type, id, reference, reference::text AS title, site_id AS site_id, NULL::text AS status FROM employees WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'form_submissions'::text AS type, id, reference, reference::text AS title, site_id AS site_id, status AS status FROM form_submissions WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR reference::text ILIKE '%' || sqlc.arg(term)::text || '%'
UNION ALL
SELECT 'actions'::text AS type, id, reference, description::text AS title, site_id AS site_id, status AS status FROM actions WHERE reference ILIKE '%' || sqlc.arg(term)::text || '%' OR description::text ILIKE '%' || sqlc.arg(term)::text || '%'
) hits
WHERE sqlc.arg(all_sites)::boolean OR site_id IS NULL OR site_id = ANY(sqlc.arg(site_ids)::uuid[])
ORDER BY id DESC
LIMIT sqlc.arg(row_limit);
