package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1" //nolint:gosec // RFC 6238 authenticator apps use HMAC-SHA1
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"net/url"
	"strings"
	"time"
)

const (
	totpStep   = 30 * time.Second
	totpDigits = 1_000_000
	totpIssuer = "Foodies"
)

var base32NoPad = base32.StdEncoding.WithPadding(base32.NoPadding)

// NewTOTPSecret returns a fresh base32 secret.
func NewTOTPSecret() (string, error) {
	buf := make([]byte, 20)
	if _, err := rand.Read(buf); err != nil {
		return "", fmt.Errorf("read totp secret: %w", err)
	}
	return base32NoPad.EncodeToString(buf), nil
}

// TOTPURI returns the otpauth URI authenticator apps scan.
func TOTPURI(secret, account string) string {
	label := url.PathEscape(totpIssuer + ":" + account)
	return fmt.Sprintf("otpauth://totp/%s?secret=%s&issuer=%s", label, secret, totpIssuer)
}

// TOTPCode returns the code for secret at t.
func TOTPCode(secret string, t time.Time) (string, error) {
	key, err := base32NoPad.DecodeString(strings.ToUpper(secret))
	if err != nil {
		return "", fmt.Errorf("decode totp secret: %w", err)
	}
	var counter [8]byte
	binary.BigEndian.PutUint64(counter[:], uint64(t.Unix()/int64(totpStep.Seconds()))) //nolint:gosec // time is positive
	mac := hmac.New(sha1.New, key)
	mac.Write(counter[:])
	sum := mac.Sum(nil)
	offset := sum[len(sum)-1] & 0x0f
	value := binary.BigEndian.Uint32(sum[offset:offset+4]) & 0x7fffffff
	return fmt.Sprintf("%06d", value%totpDigits), nil
}

// VerifyTOTP accepts the current code, and one step either side for clock skew.
func VerifyTOTP(secret, code string, now time.Time) (bool, error) {
	for _, skew := range []time.Duration{0, -totpStep, totpStep} {
		want, err := TOTPCode(secret, now.Add(skew))
		if err != nil {
			return false, err
		}
		if hmac.Equal([]byte(want), []byte(code)) {
			return true, nil
		}
	}
	return false, nil
}
