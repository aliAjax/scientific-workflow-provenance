package metrics

import "testing"

func TestZeroRegistryIsWritable(t *testing.T) {
	var r Registry
	r.Inc("jobs", 1)
	r.Set("load", 0.5)
	r.Observe("latency", 2)
	if got := r.Snapshot()["jobs"]; got != uint64(1) {
		t.Fatalf("counter missing: %#v", got)
	}
}
