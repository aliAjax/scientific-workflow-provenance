package events

import (
	"scientific-workflow-provenance/pkg/events"
	"testing"
	"time"
)

func TestOutboxDueDoesNotHang(t *testing.T) {
	o := NewOutbox()
	o.Add(events.Event{ID: "e"})
	if got := o.Due(time.Now().Add(time.Second), 0); len(got) != 1 {
		t.Fatalf("unlimited due missed item: %#v", got)
	}
}
