package query

import (
	"scientific-workflow-provenance/internal/domain/provenance"
	"testing"
)

func TestProvenancePageCursorStable(t *testing.T) {
	g := provenance.New()
	g.Add(provenance.Record{ID: "a", Kind: provenance.Entity, Type: "sample"})
	g.Add(provenance.Record{ID: "b", Kind: provenance.Entity, Type: "sample"})
	p := ProvenancePage(g, provenance.Entity, 0, 1)
	if p.Next != "b" {
		t.Fatalf("cursor points to wrong item: %s", p.Next)
	}
}
