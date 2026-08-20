package storage

import "sort"

func Compact(values []Segment) []Segment {
	sort.Slice(values, func(i, j int) bool { return values[i].Sequence < values[j].Sequence })
	seen := map[uint64]bool{}
	out := []Segment{}
	for _, v := range values {
		if !seen[v.Sequence] {
			seen[v.Sequence] = true
			out = append(out, v)
		}
	}
	return out
}
func Uploaded(values []Segment) int {
	n := 0
	for _, v := range values {
		if v.Uploaded {
			n++
		}
	}
	return n
}
