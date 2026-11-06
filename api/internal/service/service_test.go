package service

import (
	"context"
	"errors"
	"testing"
)

type stubQuerier struct{ err error }

func (s stubQuerier) Ping(context.Context) (int32, error) { return 1, s.err }

func TestPing(t *testing.T) {
	t.Parallel()

	errDown := errors.New("down")
	tests := []struct {
		name    string
		err     error
		wantErr error
	}{
		{name: "database reachable"},
		{name: "database down wraps the cause", err: errDown, wantErr: errDown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := New(stubQuerier{err: tt.err}).Ping(context.Background())
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Ping() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
