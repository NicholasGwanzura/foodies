package service

import (
	"context"
	"fmt"
)

// A guard stops an operation before it happens, and an effect does what must
// follow it. Both are small functions keyed by the operation's Go name, so the
// scaffolded transition stays the one place that changes status, and each rule
// from the spec reads on its own. They run in the caller's transaction, so an
// effect that fails rolls the whole operation back.
type hook func(context.Context, *Service, any) error

var (
	guards  = map[string][]hook{}
	effects = map[string][]hook{}
)

func typed[T any](f func(context.Context, *Service, T) error) hook {
	return func(ctx context.Context, s *Service, v any) error {
		row, ok := v.(T)
		if !ok {
			return fmt.Errorf("hook expected %T, got %T", row, v)
		}
		return f(ctx, s, row)
	}
}

func onGuard[T any](op string, f func(context.Context, *Service, T) error) {
	guards[op] = append(guards[op], typed(f))
}

func onEffect[T any](op string, f func(context.Context, *Service, T) error) {
	effects[op] = append(effects[op], typed(f))
}

func (s *Service) guard(ctx context.Context, op string, v any) error {
	for _, h := range guards[op] {
		if err := h(ctx, s, v); err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) effect(ctx context.Context, op string, v any) error {
	for _, h := range effects[op] {
		if err := h(ctx, s, v); err != nil {
			return err
		}
	}
	return nil
}
