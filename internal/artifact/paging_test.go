package artifact

import (
	"testing"
	"time"
)

func TestArtifactSortSnapshot(t *testing.T) {
	in := []Artifact{{Name: "late", CreatedAt: time.Now().Add(time.Hour)}, {Name: "old", CreatedAt: time.Now()}}
	out := SortByTime(in)
	out[0].Name = "changed"
	if in[0].Name == "changed" {
		t.Fatal("artifact sort escaped input")
	}
}
