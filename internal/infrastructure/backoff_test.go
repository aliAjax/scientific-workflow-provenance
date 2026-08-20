package infrastructure

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRetryerStopsBeforeFirstAttempt(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	calls := 0
	if err := (Retryer{Attempts: 3}).Do(ctx, func() error { calls++; return nil }); err == nil || calls != 0 {
		t.Fatalf("retry ignored cancellation: err=%v calls=%d", err, calls)
	}
	var attempts int
	ctx, cancel = context.WithCancel(context.Background())
	err := (Retryer{Attempts: 3, Base: time.Millisecond}).Do(ctx, func() error {
		attempts++
		cancel()
		return errors.New("retry")
	})
	if !errors.Is(err, context.Canceled) || attempts != 1 {
		t.Fatalf("retry continued after cancellation: err=%v attempts=%d", err, attempts)
	}
}
