package sample

import (
	"sort"
	"time"
)

type Batch struct {
	ID        string     `json:"id"`
	Samples   []string   `json:"samples"`
	StartedAt time.Time  `json:"startedAt"`
	ClosedAt  *time.Time `json:"closedAt,omitempty"`
}

func GroupByBatch(values []Sample) map[string][]Sample {
	out := map[string][]Sample{}
	for _, v := range values {
		out[v.Batch] = append(out[v.Batch], v)
	}
	for k := range out {
		sort.Slice(out[k], func(i, j int) bool { return out[k][i].ID < out[k][j].ID })
	}
	return out
}
func CloseBatch(b *Batch) bool {
	if b.ClosedAt != nil {
		return false
	}
	now := time.Now().UTC()
	b.ClosedAt = &now
	return true
}
