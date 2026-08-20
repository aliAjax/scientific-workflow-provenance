package events

import "testing"

func TestBusPublishUnderBackpressure(t *testing.T) {
	var b Bus
	ch := b.Subscribe("run", 1)
	b.Publish(Event{Type: "run"})
	select {
	case <-ch:
	default:
		t.Fatal("zero bus dropped event")
	}
}
