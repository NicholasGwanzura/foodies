package auth

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMatches(t *testing.T) {
	t.Parallel()
	tests := []struct {
		pattern, key string
		want         bool
	}{
		{"*", "cash.cash_up.read", true},
		{"cash.cash_up.read", "cash.cash_up.read", true},
		{"cash.cash_up.read", "cash.cash_up.write", false},
		{"cash.*", "cash.cash_up.read", true},
		{"cash.*", "cashier.x", false},
		{"cash.*.read", "cash.cash_up.read", true},
		{"cash.*.read", "cash.cash_up.write", false},
		{"cash.cash_up.*", "cash.cash_up.approve", true},
		{"*.read", "org.site.read", true},
		{"*.read", "org.site.write", false},
		{"org.*.read", "org.get_effective_fx_rate", false},
	}
	for _, tt := range tests {
		if got := Matches(tt.pattern, tt.key); got != tt.want {
			t.Errorf("Matches(%q, %q) = %t, want %t", tt.pattern, tt.key, got, tt.want)
		}
	}
}

func TestPasswordHash(t *testing.T) {
	t.Parallel()
	hash, err := Hash("correct horse")
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		secret string
		want   bool
	}{{"correct horse", true}, {"wrong", false}}
	for _, tt := range tests {
		got, err := Verify(tt.secret, hash)
		if err != nil || got != tt.want {
			t.Errorf("Verify(%q) = %t, %v, want %t", tt.secret, got, err, tt.want)
		}
	}
	if _, err := Verify("x", "not-a-hash"); !errors.Is(err, ErrMalformedHash) {
		t.Errorf("malformed hash error = %v", err)
	}
}

func TestTokens(t *testing.T) {
	t.Parallel()
	signer := NewSigner("secret")
	now := time.Unix(1_800_000_000, 0)
	token, err := signer.Issue(Claims{UserID: uuid.New()}, now)
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name    string
		token   string
		signer  *Signer
		at      time.Time
		wantErr error
	}{
		{"valid", token, signer, now, nil},
		{"expired", token, signer, now.Add(AccessTTL), ErrTokenExpired},
		{"wrong key", token, NewSigner("other"), now, ErrTokenInvalid},
		{"tampered", token + "x", signer, now, ErrTokenInvalid},
		{"garbage", "abc", signer, now, ErrTokenInvalid},
	}
	for _, tt := range tests {
		if _, err := tt.signer.Parse(tt.token, tt.at); !errors.Is(err, tt.wantErr) {
			t.Errorf("%s: err = %v, want %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestTOTP(t *testing.T) {
	t.Parallel()
	// RFC 6238 appendix B, SHA-1, reduced to six digits.
	const secret = "GEZDGNBVGY3TQOJQGEZDGNBVGY3TQOJQ"
	code, err := TOTPCode(secret, time.Unix(59, 0))
	if err != nil {
		t.Fatal(err)
	}
	if code != "287082" {
		t.Errorf("code = %s, want 287082", code)
	}
	now := time.Unix(1_800_000_000, 0)
	for _, tt := range []struct {
		name string
		at   time.Time
		want bool
	}{{"current", now, true}, {"one step old", now.Add(-totpStep), true}, {"two steps old", now.Add(-2 * totpStep), false}} {
		c, err := TOTPCode(secret, tt.at)
		if err != nil {
			t.Fatal(err)
		}
		if got, _ := VerifyTOTP(secret, c, now); got != tt.want {
			t.Errorf("%s: %t, want %t", tt.name, got, tt.want)
		}
	}
}
