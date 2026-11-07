package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"foodies/api/internal/auth"
	"foodies/api/internal/db"
)

const (
	// maxFailedLogins and MaxFailedPins lock the credential. Spec 8.4 sets five
	// for PINs, and passwords follow it until a rule says otherwise.
	maxFailedLogins = 5
	// MaxFailedPins is the number of wrong PINs that locks signing.
	MaxFailedPins    = 5
	refreshTTL       = 30 * 24 * time.Hour
	mfaChallengeTTL  = 5 * time.Minute
	maxMfaAttempts   = 5
	minPasswordChars = 8
	channelWeb       = "web"
)

// Sign-in outcomes, matching TokenPair.status.
const (
	StatusAuthenticated          = "authenticated"
	StatusMfaRequired            = "mfa_required"
	StatusPasswordChangeRequired = "password_change_required"
)

// TokenPair is the result of a sign-in step.
type TokenPair struct {
	Status         string
	AccessToken    string
	RefreshToken   string
	ExpiresIn      int
	MfaChallengeID string
}

var errBadCredentials = Fail(http.StatusUnauthorized, "unauthenticated", "Sign-in failed", "Check your employee number and password.")

// dummyHash answers for unknown employee numbers, so timing reveals nothing.
const dummyHash = "$argon2id$v=19$m=65536,t=3,p=2$AAAAAAAAAAAAAAAAAAAAAA$AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"

// Login signs an employee in with their number and password.
func (s *Service) Login(ctx context.Context, number, password, ip string, device *uuid.UUID) (TokenPair, error) {
	q := s.q(ctx)
	u, err := q.GetUserForSignIn(ctx, strings.TrimSpace(number))
	if errors.Is(err, pgx.ErrNoRows) {
		_, _ = auth.Verify(password, dummyHash)
		return TokenPair{}, errBadCredentials
	}
	if err != nil {
		return TokenPair{}, fmt.Errorf("read user: %w", err)
	}
	cred, err := q.GetCredentials(ctx, u.ID)
	if errors.Is(err, pgx.ErrNoRows) || (err == nil && cred.PasswordHash == nil) {
		_, _ = auth.Verify(password, dummyHash)
		return TokenPair{}, errBadCredentials
	}
	if err != nil {
		return TokenPair{}, fmt.Errorf("read credentials: %w", err)
	}
	if u.Status == "locked" {
		return TokenPair{}, errLocked()
	}
	if u.Status == "disabled" {
		return TokenPair{}, errBadCredentials
	}

	ok, err := auth.Verify(password, *cred.PasswordHash)
	if err != nil {
		return TokenPair{}, fmt.Errorf("verify password: %w", err)
	}
	if !ok {
		return TokenPair{}, s.failedLogin(ctx, u.ID)
	}
	if u.Channels == "whatsapp" {
		return TokenPair{}, Fail(http.StatusForbidden, "channel_not_allowed", "Use WhatsApp", "This account works on WhatsApp only.")
	}
	if err := q.ResetFailedLogins(ctx, u.ID); err != nil {
		return TokenPair{}, fmt.Errorf("reset failed logins: %w", err)
	}

	if u.CredentialStatus != nil && *u.CredentialStatus == "must_change" {
		return s.startSession(ctx, u.ID, u.EmployeeID, device, ip, auth.LimitPasswordChange, StatusPasswordChangeRequired)
	}
	return s.afterPassword(ctx, u.ID, u.EmployeeID, device, ip, cred)
}

func errLocked() *Error {
	return Fail(http.StatusUnauthorized, "account_locked", "Account locked", "This account is locked. Ask ICT to unlock it.")
}

func (s *Service) failedLogin(ctx context.Context, user uuid.UUID) error {
	n, err := s.q(ctx).RecordFailedLogin(ctx, user)
	if err != nil {
		return fmt.Errorf("record failed login: %w", err)
	}
	if n >= maxFailedLogins {
		if err := s.q(ctx).SetUserAuthStatus(ctx, db.SetUserAuthStatusParams{ID: user, Status: "locked"}); err != nil {
			return fmt.Errorf("lock user: %w", err)
		}
		return errLocked()
	}
	return errBadCredentials
}

// afterPassword continues sign-in once the password has passed.
func (s *Service) afterPassword(ctx context.Context, user, employee uuid.UUID, device *uuid.UUID, ip string, cred db.UserCredential) (TokenPair, error) {
	needsMfa, err := s.q(ctx).UserRequiresMfa(ctx, user)
	if err != nil {
		return TokenPair{}, fmt.Errorf("check mfa: %w", err)
	}
	if !needsMfa {
		return s.startSession(ctx, user, employee, device, ip, auth.LimitNone, StatusAuthenticated)
	}
	if cred.TotpConfirmedAt == nil {
		// Enrolment needs a token, so it is one that can do nothing else.
		return s.startSession(ctx, user, employee, device, ip, auth.LimitMfaEnrolment, StatusAuthenticated)
	}
	id, err := s.q(ctx).CreateMfaChallenge(ctx, db.CreateMfaChallengeParams{
		UserID: user, DeviceID: device, Channel: channelWeb, ExpiresAt: s.now().Add(mfaChallengeTTL),
	})
	if err != nil {
		return TokenPair{}, fmt.Errorf("create mfa challenge: %w", err)
	}
	return TokenPair{Status: StatusMfaRequired, MfaChallengeID: id.String()}, nil
}

// VerifyMfa completes a sign-in that was waiting for a TOTP code.
func (s *Service) VerifyMfa(ctx context.Context, challengeID, code, ip string) (TokenPair, error) {
	bad := Fail(http.StatusUnauthorized, "unauthenticated", "Sign-in failed", "That code is not right. Sign in again if it keeps failing.")
	id, err := uuid.Parse(challengeID)
	if err != nil {
		return TokenPair{}, bad
	}
	q := s.q(ctx)
	ch, err := q.GetMfaChallenge(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return TokenPair{}, bad
	}
	if err != nil {
		return TokenPair{}, fmt.Errorf("read mfa challenge: %w", err)
	}
	if ch.ConsumedAt != nil || s.now().After(ch.ExpiresAt) || ch.FailedAttempts >= maxMfaAttempts {
		return TokenPair{}, bad
	}
	cred, err := q.GetCredentials(ctx, ch.UserID)
	if err != nil {
		return TokenPair{}, fmt.Errorf("read credentials: %w", err)
	}
	if cred.TotpSecret == nil {
		return TokenPair{}, bad
	}
	ok, err := auth.VerifyTOTP(*cred.TotpSecret, code, s.now())
	if err != nil {
		return TokenPair{}, fmt.Errorf("verify totp: %w", err)
	}
	if !ok {
		if err := q.RecordFailedMfa(ctx, id); err != nil {
			return TokenPair{}, fmt.Errorf("record failed mfa: %w", err)
		}
		return TokenPair{}, bad
	}
	n, err := q.ConsumeMfaChallenge(ctx, id)
	if err != nil {
		return TokenPair{}, fmt.Errorf("consume mfa challenge: %w", err)
	}
	if n == 0 {
		return TokenPair{}, bad
	}
	u, err := q.GetUserForRequest(ctx, ch.UserID)
	if err != nil {
		return TokenPair{}, fmt.Errorf("read user: %w", err)
	}
	return s.startSession(ctx, u.ID, u.EmployeeID, ch.DeviceID, ip, auth.LimitNone, StatusAuthenticated)
}

func (s *Service) startSession(ctx context.Context, user, employee uuid.UUID, device *uuid.UUID, ip string, limit auth.Limit, status string) (TokenPair, error) {
	q := s.q(ctx)
	session, err := q.CreateAuthSession(ctx, db.CreateAuthSessionParams{UserID: user, Channel: channelWeb, DeviceID: device, Ip: nonEmpty(ip)})
	if err != nil {
		return TokenPair{}, fmt.Errorf("create session: %w", err)
	}
	if err := q.TouchSignIn(ctx, user); err != nil {
		return TokenPair{}, fmt.Errorf("touch sign-in: %w", err)
	}
	pair, err := s.issue(ctx, session.ID, user, employee, channelWeb, limit)
	if err != nil {
		return TokenPair{}, err
	}
	pair.Status = status
	ictx := auth.With(ctx, &auth.Principal{UserID: user, EmployeeID: employee, Channel: channelWeb, IP: ip, DeviceID: device})
	return pair, s.audit(ictx, "auth.sign_in", "users", user, nil, nil, nil)
}

func (s *Service) issue(ctx context.Context, session, user, employee uuid.UUID, channel string, limit auth.Limit) (TokenPair, error) {
	access, err := s.signer.Issue(auth.Claims{UserID: user, EmployeeID: employee, SessionID: session, Channel: channel, Limit: limit}, s.now())
	if err != nil {
		return TokenPair{}, fmt.Errorf("sign token: %w", err)
	}
	refresh, hash, err := auth.NewOpaqueToken()
	if err != nil {
		return TokenPair{}, err
	}
	err = s.q(ctx).CreateRefreshToken(ctx, db.CreateRefreshTokenParams{SessionID: session, TokenHash: hash, ExpiresAt: s.now().Add(refreshTTL)})
	if err != nil {
		return TokenPair{}, fmt.Errorf("store refresh token: %w", err)
	}
	return TokenPair{Status: StatusAuthenticated, AccessToken: access, RefreshToken: refresh, ExpiresIn: int(auth.AccessTTL.Seconds())}, nil
}

// Refresh rotates a refresh token. Reusing a spent one revokes the session,
// because one of the two holders must be a thief.
func (s *Service) Refresh(ctx context.Context, token string) (TokenPair, error) {
	bad := Fail(http.StatusUnauthorized, "token_expired", "Session ended", "Sign in again.")
	q := s.q(ctx)
	row, err := q.GetRefreshToken(ctx, auth.HashOpaque(token))
	if errors.Is(err, pgx.ErrNoRows) {
		return TokenPair{}, bad
	}
	if err != nil {
		return TokenPair{}, fmt.Errorf("read refresh token: %w", err)
	}
	if row.UsedAt != nil {
		if err := q.RevokeAuthSession(ctx, row.SessionID); err != nil {
			return TokenPair{}, fmt.Errorf("revoke session: %w", err)
		}
		return TokenPair{}, bad
	}
	if row.RevokedAt != nil || s.now().After(row.ExpiresAt) {
		return TokenPair{}, bad
	}
	n, err := q.UseRefreshToken(ctx, row.ID)
	if err != nil {
		return TokenPair{}, fmt.Errorf("spend refresh token: %w", err)
	}
	if n == 0 {
		return TokenPair{}, bad
	}
	u, err := q.GetUserForRequest(ctx, row.UserID)
	if err != nil {
		return TokenPair{}, fmt.Errorf("read user: %w", err)
	}
	if u.Status != "active" {
		return TokenPair{}, bad
	}
	// A refreshed token never escapes a sign-in step still outstanding.
	limit, err := s.limitFor(ctx, u.ID)
	if err != nil {
		return TokenPair{}, err
	}
	return s.issue(ctx, row.SessionID, u.ID, u.EmployeeID, row.Channel, limit)
}

func (s *Service) limitFor(ctx context.Context, user uuid.UUID) (auth.Limit, error) {
	q := s.q(ctx)
	cred, err := q.GetCredentials(ctx, user)
	if err != nil {
		return auth.LimitNone, fmt.Errorf("read credentials: %w", err)
	}
	row, err := q.GetUser(ctx, user)
	if err != nil {
		return auth.LimitNone, fmt.Errorf("read user: %w", err)
	}
	if row.CredentialStatus != nil && *row.CredentialStatus == "must_change" {
		return auth.LimitPasswordChange, nil
	}
	needs, err := q.UserRequiresMfa(ctx, user)
	if err != nil {
		return auth.LimitNone, fmt.Errorf("check mfa: %w", err)
	}
	if needs && cred.TotpConfirmedAt == nil {
		return auth.LimitMfaEnrolment, nil
	}
	return auth.LimitNone, nil
}

// Logout ends the current session.
func (s *Service) Logout(ctx context.Context) error {
	p := auth.From(ctx)
	if err := s.q(ctx).RevokeAuthSession(ctx, p.SessionID); err != nil {
		return fmt.Errorf("revoke session: %w", err)
	}
	return s.audit(ctx, "auth.sign_out", "users", p.UserID, nil, nil, nil)
}

// TokenUser returns the user a token names, without touching the database.
func (s *Service) TokenUser(token string) (uuid.UUID, error) {
	c, err := s.signer.Parse(token, s.now())
	if err != nil {
		return uuid.Nil, fmt.Errorf("parse token: %w", err)
	}
	return c.UserID, nil
}

// Authenticate turns a bearer token into a principal. Permissions come from the
// database on every call, so a change applies at once.
func (s *Service) Authenticate(ctx context.Context, token, ip string) (*auth.Principal, error) {
	if token == "" {
		return nil, Fail(http.StatusUnauthorized, "unauthenticated", "Sign in", "Sign in to continue.")
	}
	claims, err := s.signer.Parse(token, s.now())
	if errors.Is(err, auth.ErrTokenExpired) {
		return nil, Fail(http.StatusUnauthorized, "token_expired", "Session expired", "Your session expired.")
	}
	if err != nil {
		return nil, Fail(http.StatusUnauthorized, "unauthenticated", "Sign in", "Sign in to continue.")
	}
	q := s.q(ctx)
	session, err := q.GetAuthSession(ctx, claims.SessionID)
	if err != nil || session.RevokedAt != nil {
		return nil, Fail(http.StatusUnauthorized, "unauthenticated", "Session ended", "Sign in again.")
	}
	u, err := q.GetUserForRequest(ctx, claims.UserID)
	if err != nil {
		return nil, fmt.Errorf("read user: %w", err)
	}
	if u.Status != "active" {
		return nil, Fail(http.StatusUnauthorized, "account_locked", "Account unavailable", "This account cannot be used.")
	}
	grants, err := s.grants(ctx, u.ID)
	if err != nil {
		return nil, err
	}
	if err := q.TouchAuthSession(ctx, session.ID); err != nil {
		return nil, fmt.Errorf("touch session: %w", err)
	}
	return &auth.Principal{
		UserID: u.ID, EmployeeID: u.EmployeeID, SiteID: u.SiteID, SessionID: session.ID, Channel: claims.Channel,
		DeviceID: session.DeviceID, IP: ip, Limit: claims.Limit, Grants: grants,
	}, nil
}

func (s *Service) grants(ctx context.Context, user uuid.UUID) ([]auth.Grant, error) {
	today := s.today()
	rows, err := s.q(ctx).ListUserGrants(ctx, db.ListUserGrantsParams{UserID: user, Today: today})
	if err != nil {
		return nil, fmt.Errorf("read grants: %w", err)
	}
	out := make([]auth.Grant, 0, len(rows))
	for _, r := range rows {
		g := auth.Grant{RoleKey: r.RoleKey, RoleName: r.RoleName, Permissions: r.PermissionKeys, ScopeType: r.ScopeType, ScopeID: r.ScopeID}
		g.Scope = auth.Scope{All: r.ScopeType == "all", Sites: r.SiteIds}
		out = append(out, g)
	}
	return out, nil
}

// today is the current business date in Harare.
func (s *Service) today() pgtype.Date {
	return pgtype.Date{Time: s.now().In(Harare), Valid: true}
}

// ChangePassword sets a new password and ends the forced-change step.
func (s *Service) ChangePassword(ctx context.Context, current, next string) error {
	p := auth.From(ctx)
	q := s.q(ctx)
	cred, err := q.GetCredentials(ctx, p.UserID)
	if err != nil {
		return fmt.Errorf("read credentials: %w", err)
	}
	if cred.PasswordHash == nil {
		return errInvalid("/current_password", "That is not your current password.")
	}
	ok, err := auth.Verify(current, *cred.PasswordHash)
	if err != nil {
		return fmt.Errorf("verify password: %w", err)
	}
	if !ok {
		return errInvalid("/current_password", "That is not your current password.")
	}
	if len(next) < minPasswordChars {
		return errInvalid("/new_password", fmt.Sprintf("Use at least %d characters.", minPasswordChars))
	}
	if next == current {
		return errInvalid("/new_password", "Choose a password you have not just used.")
	}
	hash, err := auth.Hash(next)
	if err != nil {
		return err
	}
	if err := q.SetPassword(ctx, db.SetPasswordParams{UserID: p.UserID, PasswordHash: &hash}); err != nil {
		return fmt.Errorf("store password: %w", err)
	}
	if err := q.SetCredentialStatus(ctx, db.SetCredentialStatusParams{ID: p.UserID, CredentialStatus: ptr("active")}); err != nil {
		return fmt.Errorf("mark credentials active: %w", err)
	}
	return s.audit(ctx, "auth.password_change", "users", p.UserID, nil, nil, nil)
}

func ptr[T any](v T) *T { return &v }

// SetPin sets the signing PIN, proved by the current PIN or password.
func (s *Service) SetPin(ctx context.Context, pin, current string) error {
	p := auth.From(ctx)
	q := s.q(ctx)
	if len(pin) != 6 || strings.Trim(pin, "0123456789") != "" {
		return errInvalid("/pin", "Enter 6 digits.")
	}
	cred, err := q.GetCredentials(ctx, p.UserID)
	if err != nil {
		return fmt.Errorf("read credentials: %w", err)
	}
	proved := false
	for _, hash := range []*string{cred.PasswordHash, cred.PinHash} {
		if hash == nil {
			continue
		}
		ok, err := auth.Verify(current, *hash)
		if err != nil {
			return fmt.Errorf("verify secret: %w", err)
		}
		proved = proved || ok
	}
	if !proved {
		return errInvalid("/current_secret", "That does not match your password or PIN.")
	}
	hash, err := auth.Hash(pin)
	if err != nil {
		return err
	}
	if err := q.SetPin(ctx, db.SetPinParams{UserID: p.UserID, PinHash: &hash}); err != nil {
		return fmt.Errorf("store pin: %w", err)
	}
	return s.audit(ctx, "auth.pin_set", "users", p.UserID, nil, nil, nil)
}

// CheckPin verifies a signer's PIN. It writes failures outside any transaction,
// so a wrong PIN still counts when the action it guards rolls back.
func (s *Service) CheckPin(ctx context.Context, user uuid.UUID, pin string) error {
	cred, err := s.queries.GetCredentials(ctx, user)
	if errors.Is(err, pgx.ErrNoRows) {
		return Fail(http.StatusUnprocessableEntity, "pin_invalid", "PIN not set", "Set a PIN in your account first.")
	}
	if err != nil {
		return fmt.Errorf("read credentials: %w", err)
	}
	if cred.PinHash == nil {
		return Fail(http.StatusUnprocessableEntity, "pin_invalid", "PIN not set", "Set a PIN in your account first.")
	}
	if cred.PinLockedAt != nil {
		return Fail(http.StatusUnprocessableEntity, "pin_locked", "PIN locked", "Too many wrong PINs. Ask ICT to reset it.")
	}
	ok, err := auth.Verify(pin, *cred.PinHash)
	if err != nil {
		return fmt.Errorf("verify pin: %w", err)
	}
	if ok {
		if err := s.queries.ResetFailedPins(ctx, user); err != nil {
			return fmt.Errorf("reset failed pins: %w", err)
		}
		return nil
	}
	n, err := s.queries.RecordFailedPin(ctx, db.RecordFailedPinParams{UserID: user, MaxAttempts: MaxFailedPins})
	if err != nil {
		return fmt.Errorf("record failed pin: %w", err)
	}
	if n >= MaxFailedPins {
		return Fail(http.StatusUnprocessableEntity, "pin_locked", "PIN locked", "Too many wrong PINs. Ask ICT to reset it.")
	}
	return Fail(http.StatusUnprocessableEntity, "pin_invalid", "Wrong PIN", fmt.Sprintf("Wrong PIN. %d tries left.", MaxFailedPins-int(n)))
}

// UnlockWithPin switches user on an enrolled device.
func (s *Service) UnlockWithPin(ctx context.Context, device uuid.UUID, number, pin, ip string) (TokenPair, error) {
	q := s.q(ctx)
	if _, err := q.GetActiveDevice(ctx, device); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return TokenPair{}, Fail(http.StatusForbidden, "forbidden", "Device not enrolled", "This device is not enrolled for PIN unlock.")
		}
		return TokenPair{}, fmt.Errorf("read device: %w", err)
	}
	u, err := q.GetUserForSignIn(ctx, strings.TrimSpace(number))
	if errors.Is(err, pgx.ErrNoRows) {
		return TokenPair{}, Fail(http.StatusUnprocessableEntity, "pin_invalid", "Wrong PIN", "Check your employee number and PIN.")
	}
	if err != nil {
		return TokenPair{}, fmt.Errorf("read user: %w", err)
	}
	if u.Status != "active" {
		return TokenPair{}, errLocked()
	}
	if err := s.CheckPin(ctx, u.ID, pin); err != nil {
		return TokenPair{}, err
	}
	return s.startSession(ctx, u.ID, u.EmployeeID, &device, ip, auth.LimitNone, StatusAuthenticated)
}

// EnrolMfa starts TOTP enrolment and returns the secret to show.
func (s *Service) EnrolMfa(ctx context.Context) (secret, uri string, err error) {
	p := auth.From(ctx)
	q := s.q(ctx)
	emp, err := q.GetEmployee(ctx, p.EmployeeID)
	if err != nil {
		return "", "", fmt.Errorf("read employee: %w", err)
	}
	secret, err = auth.NewTOTPSecret()
	if err != nil {
		return "", "", err
	}
	if err := q.SetTotpSecret(ctx, db.SetTotpSecretParams{UserID: p.UserID, TotpSecret: &secret}); err != nil {
		return "", "", fmt.Errorf("store totp secret: %w", err)
	}
	return secret, auth.TOTPURI(secret, emp.EmployeeNumber), nil
}

// ConfirmMfa proves the authenticator works and finishes enrolment.
func (s *Service) ConfirmMfa(ctx context.Context, code string) error {
	p := auth.From(ctx)
	q := s.q(ctx)
	cred, err := q.GetCredentials(ctx, p.UserID)
	if err != nil {
		return fmt.Errorf("read credentials: %w", err)
	}
	if cred.TotpSecret == nil {
		return errState("Start two-factor enrolment first.")
	}
	ok, err := auth.VerifyTOTP(*cred.TotpSecret, code, s.now())
	if err != nil {
		return fmt.Errorf("verify totp: %w", err)
	}
	if !ok {
		return errInvalid("/code", "That code is not right. Check the time on your phone.")
	}
	if err := q.ConfirmTotp(ctx, p.UserID); err != nil {
		return fmt.Errorf("confirm totp: %w", err)
	}
	return s.audit(ctx, "auth.mfa_enrol", "users", p.UserID, nil, nil, nil)
}

// Me is the signed-in user with their employee record.
type Me struct {
	User     db.User
	Employee db.Employee
	Grants   []auth.Grant
}

// GetMe returns the signed-in user.
func (s *Service) GetMe(ctx context.Context) (Me, error) {
	p := auth.From(ctx)
	q := s.q(ctx)
	user, err := q.GetUser(ctx, p.UserID)
	if err != nil {
		return Me{}, fmt.Errorf("read user: %w", err)
	}
	emp, err := q.GetEmployee(ctx, p.EmployeeID)
	if err != nil {
		return Me{}, fmt.Errorf("read employee: %w", err)
	}
	return Me{User: user, Employee: emp, Grants: p.Grants}, nil
}

// ListMySessions returns the signed-in user's live sessions.
func (s *Service) ListMySessions(ctx context.Context) ([]db.AuthSession, error) {
	return wrap2(s.q(ctx).ListAuthSessions(ctx, auth.From(ctx).UserID))("list sessions")
}

// RevokeMySession ends one of the signed-in user's sessions.
func (s *Service) RevokeMySession(ctx context.Context, id uuid.UUID) error {
	p := auth.From(ctx)
	q := s.q(ctx)
	session, err := q.GetAuthSession(ctx, id)
	if err != nil || session.UserID != p.UserID {
		return errNotFound("session")
	}
	return wrap(q.RevokeAuthSession(ctx, id), "revoke session")
}
