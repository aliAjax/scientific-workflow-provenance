package storage

import "testing"

func TestCompactDoesNotMutateInput(t *testing.T) {
	in := []Segment{{Sequence: 2}, {Sequence: 1}}
	_ = Compact(in)
	if in[0].Sequence != 2 {
		t.Fatal("compact reordered caller slice")
	}
}
