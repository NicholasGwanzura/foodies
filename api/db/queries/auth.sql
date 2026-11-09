-- name: GetUserForSignIn :one
SELECT u.id, u.employee_id, u.status, u.mfa_required, u.credential_status, u.channels,
       e.employee_number, e.employment_status, e.site_id
FROM users u JOIN employees e ON e.id = u.employee_id
WHERE e.employee_number = $1;

-- name: GetUserForRequest :one
SELECT u.id, u.employee_id, u.status, e.site_id
FROM users u JOIN employees e ON e.id = u.employee_id
WHERE u.id = $1;

-- name: GetCredentials :one
SELECT * FROM user_credentials WHERE user_id = $1;

-- name: EnsureCredentials :exec
INSERT INTO user_credentials (user_id) VALUES ($1) ON CONFLICT (user_id) DO NOTHING;

-- name: SetPassword :exec
UPDATE user_credentials SET password_hash = $2, failed_logins = 0, updated_at = now() WHERE user_id = $1;

-- name: SetPin :exec
UPDATE user_credentials SET pin_hash = $2, failed_pins = 0, pin_locked_at = NULL, updated_at = now() WHERE user_id = $1;

-- name: ClearPin :exec
UPDATE user_credentials SET pin_hash = NULL, failed_pins = 0, pin_locked_at = NULL, updated_at = now() WHERE user_id = $1;

-- name: SetTotpSecret :exec
UPDATE user_credentials SET totp_secret = $2, totp_confirmed_at = NULL, updated_at = now() WHERE user_id = $1;

-- name: ClearMfa :exec
UPDATE user_credentials SET totp_secret = NULL, totp_confirmed_at = NULL, updated_at = now() WHERE user_id = $1;

-- name: ConfirmTotp :exec
UPDATE user_credentials SET totp_confirmed_at = now(), updated_at = now() WHERE user_id = $1;

-- name: RecordFailedLogin :one
UPDATE user_credentials SET failed_logins = failed_logins + 1, updated_at = now() WHERE user_id = $1 RETURNING failed_logins;

-- name: ResetFailedLogins :exec
UPDATE user_credentials SET failed_logins = 0, updated_at = now() WHERE user_id = $1;

-- name: RecordFailedPin :one
UPDATE user_credentials SET failed_pins = failed_pins + 1,
  pin_locked_at = CASE WHEN failed_pins + 1 >= sqlc.arg(max_attempts)::int THEN now() ELSE pin_locked_at END,
  updated_at = now()
WHERE user_id = sqlc.arg(user_id) RETURNING failed_pins;

-- name: ResetFailedPins :exec
UPDATE user_credentials SET failed_pins = 0, updated_at = now() WHERE user_id = $1;

-- name: SetUserAuthStatus :exec
UPDATE users SET status = $2, updated_at = now(), row_version = row_version + 1 WHERE id = $1;

-- name: SetCredentialStatus :exec
UPDATE users SET credential_status = $2, updated_at = now(), row_version = row_version + 1 WHERE id = $1;

-- name: TouchSignIn :exec
UPDATE users SET last_sign_in_at = now() WHERE id = $1;

-- name: SetUserMfaRequired :exec
UPDATE users SET mfa_required = $2 WHERE id = $1;

-- name: CreateAuthSession :one
INSERT INTO auth_sessions (user_id, channel, device_id, ip) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetAuthSession :one
SELECT * FROM auth_sessions WHERE id = $1;

-- name: TouchAuthSession :exec
UPDATE auth_sessions SET last_seen_at = now() WHERE id = $1 AND last_seen_at < now() - interval '1 minute';

-- name: ListAuthSessions :many
SELECT * FROM auth_sessions WHERE user_id = $1 AND revoked_at IS NULL ORDER BY last_seen_at DESC LIMIT 50;

-- name: RevokeAuthSession :exec
UPDATE auth_sessions SET revoked_at = now() WHERE id = $1 AND revoked_at IS NULL;

-- name: RevokeUserSessions :exec
UPDATE auth_sessions SET revoked_at = now() WHERE user_id = $1 AND revoked_at IS NULL;

-- name: CreateRefreshToken :exec
INSERT INTO refresh_tokens (session_id, token_hash, expires_at) VALUES ($1, $2, $3);

-- name: GetRefreshToken :one
SELECT r.*, s.user_id, s.channel, s.device_id, s.revoked_at
FROM refresh_tokens r JOIN auth_sessions s ON s.id = r.session_id
WHERE r.token_hash = $1;

-- name: UseRefreshToken :execrows
UPDATE refresh_tokens SET used_at = now() WHERE id = $1 AND used_at IS NULL;

-- name: CreateMfaChallenge :one
INSERT INTO mfa_challenges (user_id, device_id, channel, expires_at) VALUES ($1, $2, $3, $4) RETURNING id;

-- name: GetMfaChallenge :one
SELECT * FROM mfa_challenges WHERE id = $1;

-- name: ConsumeMfaChallenge :execrows
UPDATE mfa_challenges SET consumed_at = now() WHERE id = $1 AND consumed_at IS NULL;

-- name: RecordFailedMfa :exec
UPDATE mfa_challenges SET failed_attempts = failed_attempts + 1 WHERE id = $1;

-- name: ListUserGrants :many
SELECT r.key AS role_key, r.name AS role_name, r.permission_keys, a.scope_type, a.scope_id,
       CASE a.scope_type
         WHEN 'site' THEN ARRAY[a.scope_id]::uuid[]
         WHEN 'complex' THEN COALESCE((SELECT array_agg(s.id) FROM sites s WHERE s.complex_id = a.scope_id), '{}')::uuid[]
         ELSE '{}'::uuid[]
       END AS site_ids
FROM role_assignments a JOIN roles r ON r.id = a.role_id
WHERE a.user_id = sqlc.arg(user_id)
  AND a.effective_from <= sqlc.arg(today)::date
  AND (a.effective_to IS NULL OR a.effective_to >= sqlc.arg(today)::date);

-- name: UserRequiresMfa :one
SELECT COALESCE(bool_or(r.requires_mfa), false)::boolean
FROM role_assignments a JOIN roles r ON r.id = a.role_id
WHERE a.user_id = $1 AND (a.effective_to IS NULL OR a.effective_to >= current_date);

-- name: GetActiveDevice :one
SELECT * FROM devices WHERE id = $1 AND status = 'active';

-- name: ListUsersWithoutPassword :many
SELECT u.id FROM users u LEFT JOIN user_credentials c ON c.user_id = u.id
WHERE c.password_hash IS NULL;

-- name: DisableRoleMfa :execrows
UPDATE roles SET requires_mfa = false WHERE requires_mfa;
