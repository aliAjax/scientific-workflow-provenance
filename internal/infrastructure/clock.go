package infrastructure

import "time"

type Clock struct{}

func (Clock) Now() time.Time                     { return time.Now().UTC() }
func (Clock) Deadline(d time.Duration) time.Time { return time.Now().UTC().Add(d) }
