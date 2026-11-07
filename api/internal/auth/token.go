package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// AccessTTL is the life of an access token.
const AccessTTL = 15 * time.Minute

// Limit narrows what a token can do until the user finishes a sign-in step.
type Limit string

// Token limits.
const (
	LimitNone           Limit = ""
	LimitPasswordChange Limit = "password_change"
	LimitMfaEnrolment   Limit = "mfa_enrolment"
)

// Claims is the payload of an access token.
type Claims struct {
	UserID     uuid.UUID `json:"sub"`
	EmployeeID uuid.UUID `json:"emp"`
	SessionID  uuid.UUID `json:"sid"`
	Channel    string    `json:"chn"`
	Limit      Limit     `json:"lim,omitempty"`
	IssuedAt   int64     `json:"iat"`
	ExpiresAt  int64     `json:"exp"`
}

// Token errors.
var (
	ErrTokenInvalid = errors.New("token invalid")
	ErrTokenExpired = errors.New("token expired")
)

// Signer issues and checks HS256 JWTs.
type Signer struct{ secret []byte }

// NewSigner returns a Signer using secret.
func NewSigner(secret string) *Signer { return &Signer{secret: []byte(secret)} }

var jwtHeader = base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))

// Issue signs claims, stamping their issue and expiry times from now.
func (s *Signer) Issue(c Claims, now time.Time) (string, error) {
	c.IssuedAt = now.Unix()
	c.ExpiresAt = now.Add(AccessTTL).Unix()
	body, err := json.Marshal(c)
	if err != nil {
		return "", fmt.Errorf("encode claims: %w", err)
	}
	unsigned := jwtHeader + "." + base64.RawURLEncoding.EncodeToString(body)
	return unsigned + "." + s.sign(unsigned), nil
}

// Parse verifies a token and returns its claims.
func (s *Signer) Parse(token string, now time.Time) (Claims, error) {
	var c Claims
	head, rest, ok := strings.Cut(token, ".")
	if !ok || head != jwtHeader {
		return c, ErrTokenInvalid
	}
	payload, sig, ok := strings.Cut(rest, ".")
	if !ok || !hmac.Equal([]byte(sig), []byte(s.sign(head+"."+payload))) {
		return c, ErrTokenInvalid
	}
	body, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		return c, ErrTokenInvalid
	}
	if err := json.Unmarshal(body, &c); err != nil {
		return c, ErrTokenInvalid
	}
	if now.Unix() >= c.ExpiresAt {
		return c, ErrTokenExpired
	}
	return c, nil
}

func (s *Signer) sign(unsigned string) string {
	mac := hmac.New(sha256.New, s.secret)
	mac.Write([]byte(unsigned))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

// NewOpaqueToken returns a random token and the hash to store for it.
func NewOpaqueToken() (token, hash string, err error) {
	buf := make([]byte, 32)
	if _, err = rand.Read(buf); err != nil {
		return "", "", fmt.Errorf("read token: %w", err)
	}
	token = base64.RawURLEncoding.EncodeToString(buf)
	return token, HashOpaque(token), nil
}

// HashOpaque hashes a high-entropy token. It needs no salt or slow hash.
func HashOpaque(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}
