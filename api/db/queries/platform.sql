-- name: LockAuditChain :exec
SELECT pg_advisory_xact_lock(4242);

-- name: LastAuditHash :one
SELECT hash FROM audit_events ORDER BY seq DESC LIMIT 1;

-- name: InsertAuditEvent :exec
INSERT INTO audit_events (occurred_at, actor_user_id, actor_role, action, subject_type, subject_id, site_id,
  channel, device_id, ip, before, after, prev_hash, hash)
VALUES (@occurred_at, @actor_user_id, @actor_role, @action, @subject_type, @subject_id, @site_id,
  @channel, @device_id, @ip, @before, @after, @prev_hash, @hash);

-- name: NextReference :one
INSERT INTO reference_sequences (prefix, site_code, year, last_value)
VALUES ($1, $2, $3, 1)
ON CONFLICT (prefix, site_code, year) DO UPDATE SET last_value = reference_sequences.last_value + 1
RETURNING last_value;

-- name: GetSiteCode :one
SELECT code FROM sites WHERE id = $1;

-- name: ClaimIdempotencyKey :one
WITH ins AS (
  INSERT INTO idempotency_keys (actor_id, key, request_hash)
  VALUES ($1, $2, $3)
  ON CONFLICT (actor_id, key) DO NOTHING
  RETURNING 1
)
SELECT EXISTS (SELECT 1 FROM ins)::boolean;

-- name: GetIdempotencyKey :one
SELECT * FROM idempotency_keys WHERE actor_id = $1 AND key = $2;

-- name: FinishIdempotencyKey :exec
UPDATE idempotency_keys SET status = $3, headers = $4, body = $5 WHERE actor_id = $1 AND key = $2;

-- name: ReleaseIdempotencyKey :exec
DELETE FROM idempotency_keys WHERE actor_id = $1 AND key = $2;

-- name: PurgeIdempotencyKeys :exec
DELETE FROM idempotency_keys WHERE created_at < now() - interval '72 hours';

-- name: SodConflicts :many
SELECT r."key", r.name, r.enforcement FROM sod_rules r
WHERE r.active AND r.scope = 'same_record'
  AND (r.action_a = sqlc.arg(action) OR r.action_b = sqlc.arg(action))
  AND EXISTS (
    SELECT 1 FROM audit_events e
    WHERE e.subject_type = sqlc.arg(subject_type) AND e.subject_id = sqlc.arg(subject_id)
      AND e.actor_user_id = sqlc.arg(actor)
      AND e.action = CASE WHEN r.action_a = sqlc.arg(action) THEN r.action_b ELSE r.action_a END);
