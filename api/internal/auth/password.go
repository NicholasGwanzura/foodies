// Package auth holds credential hashing, tokens, TOTP and the signed-in
// principal. It knows nothing about HTTP or the database.
package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Argon2id parameters, per the OWASP baseline.
const (
	argonTime    = 3
	argonMemory  = 64 * 1024
	argonThreads = 2
	argonKeyLen  = 32
	saltLen      = 16
)

// ErrMalformedHash means a stored hash could not be parsed.
var ErrMalformedHash = errors.New("malformed password hash")

// Hash returns an encoded Argon2id hash of secret, for passwords and PINs.
func Hash(secret string) (string, error) {
	salt := make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("read salt: %w", err)
	}
	key := argon2.IDKey([]byte(secret), salt, argonTime, argonMemory, argonThreads, argonKeyLen)
	enc := base64.RawStdEncoding
	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, argonMemory, argonTime, argonThreads, enc.EncodeToString(salt), enc.EncodeToString(key)), nil
}

// Verify reports whether secret matches an encoded hash from Hash.
func Verify(secret, encoded string) (bool, error) {
	parts := strings.Split(encoded, "$")
	if len(parts) != 6 || parts[1] != "argon2id" {
		return false, ErrMalformedHash
	}
	var memory, timeCost uint32
	var threads uint8
	if _, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &timeCost, &threads); err != nil {
		return false, fmt.Errorf("%w: %w", ErrMalformedHash, err)
	}
	enc := base64.RawStdEncoding
	salt, err := enc.DecodeString(parts[4])
	if err != nil {
		return false, fmt.Errorf("%w: %w", ErrMalformedHash, err)
	}
	want, err := enc.DecodeString(parts[5])
	if err != nil {
		return false, fmt.Errorf("%w: %w", ErrMalformedHash, err)
	}
	got := argon2.IDKey([]byte(secret), salt, timeCost, memory, threads, uint32(len(want))) //nolint:gosec // length of a 32-byte key
	return subtle.ConstantTimeCompare(got, want) == 1, nil
}
