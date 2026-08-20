package retention

import "time"

type Policy struct {
	MaxAge            time.Duration
	MaxCount          int
	DeleteInvalidated bool
}

func Keep(created, now time.Time, p Policy, status string) bool {
	if p.MaxAge > 0 && now.Sub(created) > p.MaxAge {
		return false
	}
	if status == "invalidated" && p.DeleteInvalidated {
		return false
	}
	return true
}
