package storage

import (
	"sort"
	"time"
)

type Expiry struct {
	Key string
	At  time.Time
}

func Evict(items []Expiry, now time.Time, max int) []string {
	sort.Slice(items, func(i, j int) bool { return items[i].At.Before(items[j].At) })
	out := []string{}
	for _, v := range items {
		if max <= 0 || len(out) < max {
			if !v.At.After(now) {
				out = append(out, v.Key)
			}
		}
	}
	return out
}
