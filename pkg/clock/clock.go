package clock

import "time"

type Clock interface {
	Now() time.Time
	Since(time.Time) time.Duration
}
type Real struct{}

func (Real) Now() time.Time                  { return time.Now().UTC() }
func (Real) Since(t time.Time) time.Duration { return time.Since(t) }

type Fixed struct{ Value time.Time }

func (f Fixed) Now() time.Time                  { return f.Value }
func (f Fixed) Since(t time.Time) time.Duration { return f.Value.Sub(t) }
