-- name: UserHoldsRole :one
SELECT EXISTS (
  SELECT 1 FROM role_assignments a
  WHERE a.user_id = sqlc.arg(user_id) AND a.role_id = sqlc.arg(role_id)
    AND a.effective_from <= current_date AND (a.effective_to IS NULL OR a.effective_to >= current_date)
    AND (a.scope_type = 'all'
      OR sqlc.narg(site_id)::uuid IS NULL
      OR (a.scope_type = 'site' AND a.scope_id = sqlc.narg(site_id))
      OR (a.scope_type = 'complex' AND a.scope_id = (SELECT complex_id FROM sites WHERE id = sqlc.narg(site_id))))
)::boolean;

-- name: DelegatorsTo :many
SELECT from_user_id, role_id FROM delegations
WHERE to_user_id = $1 AND status = 'active' AND valid_from <= now() AND valid_to >= now();

-- name: UsersWithRoleAtSite :many
SELECT DISTINCT a.user_id FROM role_assignments a
WHERE a.role_id = sqlc.arg(role_id)
  AND a.effective_from <= current_date AND (a.effective_to IS NULL OR a.effective_to >= current_date)
  AND (a.scope_type = 'all'
    OR sqlc.narg(site_id)::uuid IS NULL
    OR (a.scope_type = 'site' AND a.scope_id = sqlc.narg(site_id))
    OR (a.scope_type = 'complex' AND a.scope_id = (SELECT complex_id FROM sites WHERE id = sqlc.narg(site_id))));

-- name: ActivePolicyFor :one
SELECT * FROM approval_policies WHERE workflow_key = $1 AND status = 'active'
ORDER BY effective_from DESC NULLS LAST, id DESC LIMIT 1;

-- name: PolicyVersionsFor :one
SELECT COALESCE(max(version), 0)::int FROM approval_policies WHERE workflow_key = $1;

-- name: SupersedePolicies :exec
UPDATE approval_policies SET status = 'superseded', updated_at = now(), row_version = row_version + 1
WHERE workflow_key = $1 AND status = 'active' AND id <> $2;

-- name: CreateApprovalRequest :one
INSERT INTO approval_requests (subject_type, subject_id, workflow_key, policy_id, site_id, amount_amount, amount_currency, current_step, summary, created_by, updated_by)
VALUES (@subject_type, @subject_id, @workflow_key, @policy_id, @site_id, @amount_amount, @amount_currency, 1, @summary, @created_by, @created_by)
RETURNING *;

-- name: DecideApprovalStep :exec
UPDATE approval_request_steps
SET decided_by = @decided_by, decision = @decision, decided_at = now(), channel = @channel, comment = @comment, signature_id = @signature_id,
    updated_at = now(), updated_by = @decided_by, row_version = row_version + 1
WHERE id = @id;

-- name: ReassignApprovalStep :exec
UPDATE approval_request_steps SET assigned_user_id = $2, updated_at = now(), row_version = row_version + 1 WHERE id = $1;

-- name: ClearPendingSteps :exec
UPDATE approval_request_steps SET decision = 'pending', decided_by = NULL, decided_at = NULL, row_version = row_version + 1
WHERE approval_request_id = $1 AND decision = 'pending';

-- name: AdvanceApproval :one
UPDATE approval_requests SET current_step = $2, updated_at = now(), row_version = row_version + 1 WHERE id = $1 RETURNING *;

-- name: FinishApproval :one
UPDATE approval_requests SET status = $2, decided_at = now(), updated_at = now(), row_version = row_version + 1 WHERE id = $1 RETURNING *;

-- name: ApprovalInbox :many
SELECT r.* FROM approval_requests r
JOIN approval_request_steps st ON st.approval_request_id = r.id AND st.sequence = r.current_step AND st.decision = 'pending'
WHERE r.status = 'pending'
  AND (sqlc.narg(cursor_id)::uuid IS NULL OR r.id < sqlc.narg(cursor_id))
  AND (st.assigned_user_id = sqlc.arg(user_id) OR (st.assigned_user_id IS NULL AND (
    EXISTS (SELECT 1 FROM role_assignments a WHERE a.user_id = sqlc.arg(user_id) AND a.role_id = st.role_id
      AND a.effective_from <= current_date AND (a.effective_to IS NULL OR a.effective_to >= current_date)
      AND (a.scope_type = 'all' OR r.site_id IS NULL OR (a.scope_type = 'site' AND a.scope_id = r.site_id)
        OR (a.scope_type = 'complex' AND a.scope_id = (SELECT complex_id FROM sites WHERE id = r.site_id))))
    OR EXISTS (SELECT 1 FROM delegations d WHERE d.to_user_id = sqlc.arg(user_id) AND d.status = 'active'
      AND d.valid_from <= now() AND d.valid_to >= now() AND d.role_id = st.role_id))))
ORDER BY r.id DESC LIMIT sqlc.arg(row_limit);

-- name: PendingApprovalFor :one
SELECT * FROM approval_requests WHERE subject_type = $1 AND subject_id = $2 AND status = 'pending' ORDER BY id DESC LIMIT 1;

-- name: CountPendingSignatures :one
SELECT count(*) FROM signature_requests WHERE subject_type = $1 AND subject_id = $2 AND status = 'pending';

-- name: VoidSignatures :execrows
UPDATE signature_requests SET status = 'expired', updated_at = now(), row_version = row_version + 1
WHERE subject_type = $1 AND subject_id = $2 AND status = 'signed';

-- name: RecordSignature :one
UPDATE signature_requests
SET status = 'signed', signed_at = now(), signed_channel = @signed_channel, device_id = @device_id, content_hash = @content_hash, attested_by = @attested_by,
    updated_at = now(), updated_by = @attested_by, row_version = row_version + 1
WHERE id = @id AND status = 'pending' RETURNING *;

-- name: CreateSignatureRequestRow :one
INSERT INTO signature_requests (subject_type, subject_id, meaning, signer_kind, signer_employee_id, signer_role_id,
  external_name, external_organisation, external_id_ref, proof_attachment_id, content_hash, created_by, updated_by)
VALUES (@subject_type, @subject_id, @meaning, @signer_kind, @signer_employee_id, @signer_role_id,
  @external_name, @external_organisation, @external_id_ref, @proof_attachment_id, @content_hash, @created_by, @created_by)
RETURNING *;

-- name: DeclineSignature :one
UPDATE signature_requests SET status = 'declined', updated_at = now(), updated_by = $2, row_version = row_version + 1
WHERE id = $1 AND status = 'pending' RETURNING *;

-- name: SupersedeRuleVersions :exec
UPDATE rule_versions SET status = 'superseded', effective_to = now(), updated_at = now(), row_version = row_version + 1
WHERE rule_definition_id = $1 AND scope_type = $2 AND scope_id IS NOT DISTINCT FROM $3 AND status = 'active' AND id <> $4;

-- name: ActiveRuleValue :one
SELECT v.* FROM rule_versions v JOIN rule_definitions d ON d.id = v.rule_definition_id
WHERE d."key" = sqlc.arg(key) AND v.status = 'active'
  AND v.effective_from <= now() AND (v.effective_to IS NULL OR v.effective_to > now())
  AND (v.scope_type = 'group'
    OR (v.scope_type = 'site' AND v.scope_id = sqlc.narg(site_id))
    OR (v.scope_type = 'complex' AND v.scope_id = (SELECT complex_id FROM sites WHERE id = sqlc.narg(site_id))))
ORDER BY CASE v.scope_type WHEN 'site' THEN 1 WHEN 'complex' THEN 2 ELSE 3 END LIMIT 1;

-- name: GetRuleDefinitionByKey :one
SELECT * FROM rule_definitions WHERE "key" = $1;

-- name: MissingRules :many
SELECT r.rule_key, d.name, d.blocking FROM workflow_requirements r
LEFT JOIN rule_definitions d ON d."key" = r.rule_key
WHERE r.workflow_key = sqlc.arg(workflow_key) AND COALESCE(d.blocking, true)
  AND NOT EXISTS (
    SELECT 1 FROM rule_versions v
    WHERE v.rule_definition_id = d.id AND v.status = 'active'
      AND v.effective_from <= now() AND (v.effective_to IS NULL OR v.effective_to > now())
      AND (v.scope_type = 'group'
        OR (v.scope_type = 'site' AND v.scope_id = sqlc.arg(site_id))
        OR (v.scope_type = 'complex' AND v.scope_id = (SELECT complex_id FROM sites WHERE id = sqlc.arg(site_id)))));

-- name: ListWorkflowKeys :many
SELECT DISTINCT workflow_key FROM workflow_requirements ORDER BY workflow_key;

-- name: ListAllActivations :many
SELECT * FROM workflow_activations ORDER BY workflow_key, site_id;

-- name: EnsureActivation :exec
INSERT INTO workflow_activations (workflow_key, site_id, status, missing_rule_keys)
VALUES ($1, $2, 'not_ready', '{}')
ON CONFLICT (workflow_key, site_id) DO NOTHING;

-- name: GetActivation :one
SELECT * FROM workflow_activations WHERE workflow_key = $1 AND site_id = $2;

-- name: SetActivationReadiness :exec
UPDATE workflow_activations SET missing_rule_keys = $3,
  status = CASE WHEN status IN ('live','suspended') AND cardinality($3::text[]) = 0 THEN status
                WHEN cardinality($3::text[]) = 0 THEN 'ready' ELSE 'not_ready' END,
  updated_at = now(), row_version = row_version + 1
WHERE workflow_key = $1 AND site_id = $2;

-- name: ListActiveSiteIds :many
SELECT id FROM sites WHERE active;

-- name: OwnerRoleOfRule :one
SELECT r.name FROM roles r WHERE r.id = $1;

-- name: CreateNotification :one
INSERT INTO notifications (recipient_user_id, category, title, body, channel, subject_type, subject_id, status, sent_at)
VALUES (@recipient_user_id, @category, @title, @body, @channel, @subject_type, @subject_id, 'delivered', now())
RETURNING *;

-- name: ListMyNotifications :many
SELECT * FROM notifications
WHERE recipient_user_id = sqlc.arg(user_id)
  AND (sqlc.narg(cursor_id)::uuid IS NULL OR id < sqlc.narg(cursor_id))
  AND (NOT sqlc.arg(unread_only)::boolean OR read_at IS NULL)
ORDER BY id DESC LIMIT sqlc.arg(row_limit);

-- name: GetMyNotification :one
SELECT * FROM notifications WHERE id = $1 AND recipient_user_id = $2;

-- name: ReadNotification :one
UPDATE notifications SET status = 'read', read_at = COALESCE(read_at, now()), updated_at = now(), row_version = row_version + 1
WHERE id = $1 AND recipient_user_id = $2 RETURNING *;

-- name: ReadAllNotifications :execrows
UPDATE notifications SET status = 'read', read_at = now(), updated_at = now(), row_version = row_version + 1
WHERE recipient_user_id = $1 AND read_at IS NULL;

-- name: CountUnreadNotifications :one
SELECT count(*) FROM notifications WHERE recipient_user_id = $1 AND read_at IS NULL;

-- name: GetNotificationPreferences :one
SELECT categories FROM notification_preferences WHERE user_id = $1;

-- name: SetNotificationPreferences :exec
INSERT INTO notification_preferences (user_id, categories) VALUES ($1, $2)
ON CONFLICT (user_id) DO UPDATE SET categories = EXCLUDED.categories, updated_at = now();

-- name: ListTasksFor :many
SELECT t.* FROM tasks t
WHERE (sqlc.narg(cursor_id)::uuid IS NULL OR t.id < sqlc.narg(cursor_id))
  AND (sqlc.narg(status)::text[] IS NULL OR t.status = ANY(sqlc.narg(status)::text[]))
  AND (sqlc.narg(site_id)::uuid IS NULL OR t.site_id = sqlc.narg(site_id))
  AND (sqlc.narg(business_date)::date IS NULL OR t.business_date = sqlc.narg(business_date))
  AND (sqlc.narg(overdue)::boolean IS NULL OR t.overdue = sqlc.narg(overdue))
  AND (sqlc.narg(q)::text IS NULL OR t.title ILIKE '%' || sqlc.narg(q)::text || '%')
  AND (sqlc.arg(all_sites)::boolean OR t.site_id = ANY(sqlc.arg(site_ids)::uuid[]))
  AND (sqlc.narg(assignee)::uuid IS NULL OR t.assignee_user_id = sqlc.narg(assignee)
       OR (t.assignee_user_id IS NULL AND t.assignee_role_id = ANY(sqlc.arg(role_ids)::uuid[])))
ORDER BY t.id DESC LIMIT sqlc.arg(row_limit);

-- name: UserRoleIds :many
SELECT DISTINCT role_id FROM role_assignments
WHERE user_id = $1 AND effective_from <= current_date AND (effective_to IS NULL OR effective_to >= current_date);

-- name: CreateTaskRow :one
INSERT INTO tasks (title, schedule_id, subject_type, subject_id, site_id, business_date, shift, assignee_user_id, assignee_role_id, due_at, priority, created_by, updated_by)
VALUES (@title, @schedule_id, @subject_type, @subject_id, @site_id, @business_date, @shift, @assignee_user_id, @assignee_role_id, @due_at, @priority, @created_by, @created_by)
ON CONFLICT (schedule_id, site_id, business_date) WHERE schedule_id IS NOT NULL DO NOTHING
RETURNING *;

-- name: ClaimTaskRow :one
UPDATE tasks SET status = 'in_progress', assignee_user_id = $2, updated_at = now(), updated_by = $2, row_version = row_version + 1
WHERE id = $1 AND status = 'open' RETURNING *;

-- name: CompleteTaskRow :one
UPDATE tasks SET status = 'done', completed_at = now(), overdue = false, updated_at = now(), updated_by = $2, row_version = row_version + 1
WHERE id = $1 AND status IN ('open', 'in_progress') RETURNING *;

-- name: SkipTaskRow :one
UPDATE tasks SET status = 'skipped', skip_reason = $3, updated_at = now(), updated_by = $2, row_version = row_version + 1
WHERE id = $1 AND status IN ('open', 'in_progress') RETURNING *;

-- name: ReassignTaskRow :one
UPDATE tasks SET assignee_user_id = $3, updated_at = now(), updated_by = $2, row_version = row_version + 1
WHERE id = $1 AND status IN ('open', 'in_progress') RETURNING *;

-- name: MarkOverdueTasks :many
UPDATE tasks SET overdue = true, updated_at = now(), row_version = row_version + 1
WHERE status IN ('open', 'in_progress') AND NOT overdue AND due_at < now() RETURNING *;

-- name: TasksNeedingReminder :many
SELECT t.*, s.reminder_minutes FROM tasks t JOIN schedules s ON s.id = t.schedule_id
WHERE t.status IN ('open', 'in_progress') AND t.reminded_at IS NULL AND s.reminder_minutes IS NOT NULL
  AND t.due_at - make_interval(mins => s.reminder_minutes) < now();

-- name: MarkReminded :exec
UPDATE tasks SET reminded_at = now() WHERE id = $1;

-- name: ListActiveSchedules :many
SELECT * FROM schedules WHERE status = 'active';

-- name: CreateEscalationRow :one
INSERT INTO escalations (subject_type, subject_id, domain, severity_level, site_id, current_step, triggered_at, created_by, updated_by)
VALUES (@subject_type, @subject_id, @domain, @severity_level, @site_id, 1, now(), @created_by, @created_by)
RETURNING *;

-- name: ActiveEscalationPolicy :one
SELECT * FROM escalation_policies WHERE domain = $1 AND severity_level = $2 AND status = 'active' ORDER BY id DESC LIMIT 1;

-- name: AddEscalationRecipient :one
INSERT INTO escalation_recipients (escalation_id, step, user_id, channel, notified_at)
VALUES ($1, $2, $3, $4, now()) RETURNING *;

-- name: AckEscalationRecipient :execrows
UPDATE escalation_recipients SET acknowledged_at = now(), updated_at = now()
WHERE escalation_id = $1 AND user_id = $2 AND acknowledged_at IS NULL;

-- name: SetEscalationStep :one
UPDATE escalations SET current_step = $2, status = $3, updated_at = now(), row_version = row_version + 1 WHERE id = $1 RETURNING *;

-- name: AckEscalationRow :one
UPDATE escalations SET status = 'acknowledged', updated_at = now(), updated_by = $2, row_version = row_version + 1
WHERE id = $1 AND status IN ('open', 'breached') RETURNING *;

-- name: EscalationsToCheck :many
SELECT e.*, st.ack_within_minutes, st.next_on_timeout, max(r.notified_at)::timestamptz AS last_notified
FROM escalations e
JOIN escalation_policies p ON p.domain = e.domain AND p.severity_level = e.severity_level AND p.status = 'active'
JOIN escalation_policy_steps st ON st.escalation_policy_id = p.id AND st.sequence = e.current_step
JOIN escalation_recipients r ON r.escalation_id = e.id AND r.step = e.current_step
WHERE e.status IN ('open') AND st.ack_within_minutes IS NOT NULL
GROUP BY e.id, st.ack_within_minutes, st.next_on_timeout
HAVING max(r.notified_at) + make_interval(mins => st.ack_within_minutes) < now();

-- name: CountOpenTasksFor :one
SELECT count(*) FROM tasks t
WHERE t.status IN ('open', 'in_progress')
  AND (t.assignee_user_id = sqlc.arg(user_id) OR (t.assignee_user_id IS NULL AND t.assignee_role_id = ANY(sqlc.arg(role_ids)::uuid[])
       AND (sqlc.arg(all_sites)::boolean OR t.site_id = ANY(sqlc.arg(site_ids)::uuid[]))));

-- name: CountOverdueTasksFor :one
SELECT count(*) FROM tasks t
WHERE t.status IN ('open', 'in_progress') AND t.overdue
  AND (t.assignee_user_id = sqlc.arg(user_id) OR (t.assignee_user_id IS NULL AND t.assignee_role_id = ANY(sqlc.arg(role_ids)::uuid[])
       AND (sqlc.arg(all_sites)::boolean OR t.site_id = ANY(sqlc.arg(site_ids)::uuid[]))));

-- name: CountOpenEscalations :one
SELECT count(*) FROM escalations e WHERE e.status IN ('open', 'breached')
  AND (sqlc.arg(all_sites)::boolean OR e.site_id IS NULL OR e.site_id = ANY(sqlc.arg(site_ids)::uuid[]));

-- name: LatestPublishedVersion :one
SELECT * FROM form_template_versions WHERE form_template_id = $1 AND status = 'published' ORDER BY version DESC LIMIT 1;

-- name: NextTemplateVersion :one
SELECT COALESCE(max(version), 0)::int + 1 FROM form_template_versions WHERE form_template_id = $1;

-- name: SetCurrentTemplateVersion :exec
UPDATE form_templates SET current_version_id = $2, status = 'active', updated_at = now(), row_version = row_version + 1 WHERE id = $1;

-- name: RetireOtherVersions :exec
UPDATE form_template_versions SET status = 'retired', updated_at = now(), row_version = row_version + 1
WHERE form_template_id = $1 AND status = 'published' AND id <> $2;

-- name: SetSubmissionOutcome :one
UPDATE form_submissions SET status = @status, result = @result, updated_at = now(), updated_by = @updated_by, row_version = row_version + 1
WHERE id = @id RETURNING *;

-- name: SetEmployeeStatus :exec
UPDATE employees SET employment_status = $2, updated_at = now(), row_version = row_version + 1 WHERE id = $1;

-- name: GetUserByEmployee :one
SELECT * FROM users WHERE employee_id = $1;

-- name: RevokeEmployeeIdentities :exec
UPDATE whatsapp_identities SET status = 'revoked', updated_at = now(), row_version = row_version + 1 WHERE employee_id = $1 AND status <> 'revoked';

-- name: ApplyEmployeeTransfer :exec
UPDATE employees SET site_id = COALESCE(sqlc.narg(site_id), site_id), department_id = COALESCE(sqlc.narg(department_id), department_id),
  job_title = COALESCE(sqlc.narg(job_title), job_title), manager_employee_id = COALESCE(sqlc.narg(manager_employee_id), manager_employee_id),
  updated_at = now(), row_version = row_version + 1
WHERE id = sqlc.arg(id);

-- name: CreateFormTemplateVersionNumbered :one
INSERT INTO form_template_versions (form_template_id, version, definition, offline_allowed, channels, retention_class_id, required_rule_keys, effective_from, created_by, updated_by)
VALUES (@form_template_id, @version, @definition, COALESCE(sqlc.narg(offline_allowed)::boolean, false), @channels, @retention_class_id, @required_rule_keys, @effective_from, @created_by, @created_by)
RETURNING *;

-- name: EffectiveFxRate :one
SELECT * FROM fx_rates WHERE from_currency = $1 AND to_currency = $2 AND rate_date <= $3 ORDER BY rate_date DESC, id DESC LIMIT 1;

-- name: ActiveRuleSnapshot :many
SELECT d."key", v.value FROM rule_versions v JOIN rule_definitions d ON d.id = v.rule_definition_id
WHERE v.status = 'active' AND v.scope_type = 'group';

-- name: OfflineTemplates :many
SELECT * FROM form_template_versions WHERE status = 'published' AND offline_allowed;

-- name: CreateSyncConflict :one
INSERT INTO sync_conflicts (batch_id, device_id, site_id, client_record_id, reason, client_payload)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: IdentityByPhone :one
SELECT * FROM whatsapp_identities WHERE phone_e164 = $1 AND status = 'verified';

-- name: InsertWhatsappMessage :exec
INSERT INTO whatsapp_messages (direction, identity_id, provider_message_id, message_type, sent_at, delivery_status)
VALUES (@direction, @identity_id, @provider_message_id, @message_type, @sent_at, @delivery_status);

-- name: IntegrationBySystem :one
SELECT * FROM integrations WHERE "system" = $1 LIMIT 1;

-- name: TouchIntegration :exec
UPDATE integrations SET last_success_at = now(), updated_at = now() WHERE id = $1;

-- name: RequeuePosting :one
UPDATE postings SET status = 'pending', attempts = COALESCE(attempts, 0) + 1, updated_at = now(), row_version = row_version + 1 WHERE id = $1 RETURNING *;

-- name: CreatePostingRow :exec
INSERT INTO postings (target_system, posting_type, subject_type, subject_id, site_id, payload, created_by, updated_by)
VALUES (@target_system, @posting_type, @subject_type, @subject_id, @site_id, @payload, @created_by, @created_by);
