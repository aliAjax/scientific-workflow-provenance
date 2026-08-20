package artifact

import (
	"sort"
	"time"
)

func SortByTime(values []Artifact) []Artifact {
	out := values
	sort.Slice(out, func(i, j int) bool { return out[i].CreatedAt.Before(out[j].CreatedAt) })
	return out
}
func Since(values []Artifact, t time.Time) []Artifact {
	out := []Artifact{}
	for _, v := range values {
		if v.CreatedAt.After(t) {
			out = append(out, v)
		}
	}
	return SortByTime(out)
}
