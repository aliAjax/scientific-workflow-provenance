package sample

import "testing"

func sampleForIsolation() Sample {
	return Sample{ID: "s1", ParentIDs: []string{"p1"}, Qualities: []Quality{{Metric: "purity", Value: 0.9}}}
}

func TestSampleStoreDeepSnapshot(t *testing.T) {
	s := NewStore()
	_ = s.Put(sampleForIsolation())
	v, _ := s.Get("s1")
	v.ParentIDs[0] = "changed"
	v.Qualities[0].Value = 0
	again, _ := s.Get("s1")
	if again.ParentIDs[0] != "p1" || again.Qualities[0].Value != 0.9 {
		t.Fatal("sample store leaked nested slices")
	}
}

func TestLineageDoesNotAliasParents(t *testing.T) {
	values := []Sample{{ID: "p1", ParentIDs: []string{"grandparent"}}, {ID: "c1", ParentIDs: []string{"p1"}}}
	out := Ancestors("c1", values)
	if len(out) != 1 {
		t.Fatal("missing ancestor")
	}
	out[0].ParentIDs[0] = "changed"
	if values[0].ParentIDs[0] != "grandparent" {
		t.Fatal("lineage result aliased input")
	}
}

func TestBatchGroupDoesNotAliasSample(t *testing.T) {
	values := []Sample{sampleForIsolation()}
	groups := GroupByBatch(values)
	groups[""][0].ParentIDs[0] = "changed"
	if values[0].ParentIDs[0] != "p1" {
		t.Fatal("batch group aliased input")
	}
}

func TestMetricValuesAreIsolated(t *testing.T) {
	m := NewMetricSet()
	m.Add("signal", 1)
	values := m.ValuesCopy("signal")
	values[0] = 99
	if m.Mean("signal") != 1 {
		t.Fatal("metric values escaped")
	}
}
