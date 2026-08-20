package storage

import (
	"testing"
	"time"
)

func TestEvictKeepsCursorInputStable(t *testing.T) {
	in := []Expiry{{Key: "late", At: time.Now().Add(time.Hour)}, {Key: "old", At: time.Now().Add(-time.Hour)}}
	_ = Evict(in, time.Now(), 1)
	if in[0].Key != "late" {
		t.Fatal("eviction reordered cursor input")
	}
}
