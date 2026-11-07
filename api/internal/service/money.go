package service

import (
	"encoding/json"
	"fmt"
	"math/big"
)

// Decimals stay strings everywhere, and arithmetic here uses exact rationals,
// because money must never pick up floating-point error.

func rat(s string) *big.Rat {
	r, ok := new(big.Rat).SetString(s)
	if !ok {
		return new(big.Rat)
	}
	return r
}

func ratp(s *string) *big.Rat {
	if s == nil {
		return new(big.Rat)
	}
	return rat(*s)
}

// dec formats a rational as decimal(18,4).
func dec(r *big.Rat) string { return r.FloatString(4) }

func sub(a, b *big.Rat) *big.Rat { return new(big.Rat).Sub(a, b) }
func add(a, b *big.Rat) *big.Rat { return new(big.Rat).Add(a, b) }
func mul(a, b *big.Rat) *big.Rat { return new(big.Rat).Mul(a, b) }

// moneyRule is a rule value holding an amount in one currency.
type moneyRule struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

func decodeMoney(raw json.RawMessage) (moneyRule, error) {
	var m moneyRule
	if err := json.Unmarshal(raw, &m); err != nil || m.Amount == "" {
		return m, fmt.Errorf("rule value is not an amount: %s", raw)
	}
	return m, nil
}
