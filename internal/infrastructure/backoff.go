package infrastructure

import (
	"context"
	"math/rand"
	"time"
)

type Retryer struct {
	Attempts int
	Base     time.Duration
	Max      time.Duration
}

func (r Retryer) Do(ctx context.Context, fn func() error) error {
	if r.Attempts < 1 {
		r.Attempts = 1
	}
	if r.Base <= 0 {
		r.Base = 10 * time.Millisecond
	}
	if r.Max <= 0 {
		r.Max = time.Second
	}
	var err error
	for n := 0; n < r.Attempts; n++ {
		if err = fn(); err == nil {
			return nil
		}
		d := r.Base * time.Duration(1<<min(n, 10))
		if d > r.Max {
			d = r.Max
		}
		d += time.Duration(rand.Int63n(int64(d/4 + 1)))
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(d):
		}
	}
	return err
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
