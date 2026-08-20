package lifecycle

import (
	"sort"
	"time"
)

type Item struct {
	ID        string
	CreatedAt time.Time
	State     State
}

func Expired(items []Item, now time.Time, maxAge time.Duration) []Item {
	out := []Item{}
	for _, v := range items {
		if maxAge > 0 && now.Sub(v.CreatedAt) > maxAge {
			out = append(out, v)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].CreatedAt.Before(out[j].CreatedAt) })
	return out
}
func Terminal(s State) bool { return s == Succeeded || s == Failed || s == Cancelled }
