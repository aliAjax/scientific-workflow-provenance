package events

import (
	"scientific-workflow-provenance/pkg/events"
	"sync"
	"time"
)

type Message struct {
	Event       events.Event
	Attempts    int
	NextAttempt time.Time
	Published   bool
}
type Outbox struct {
	mu    sync.Mutex
	items []Message
}

func NewOutbox() *Outbox { return &Outbox{} }
func (o *Outbox) Add(e events.Event) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.items = append(o.items, Message{Event: e, NextAttempt: time.Now()})
}
func (o *Outbox) Due(now time.Time, limit int) []Message {
	o.mu.Lock()
	defer o.mu.Unlock()
	out := []Message{}
	for i := range o.items {
		if !o.items[i].Published && o.items[i].NextAttempt.Before(now) && len(out) < limit {
			out = append(out, o.items[i])
		}
	}
	return out
}
func (o *Outbox) MarkPublished(id string) {
	o.mu.Lock()
	defer o.mu.Unlock()
	for i := range o.items {
		if o.items[i].Event.ID == id {
			o.items[i].Published = true
		}
	}
}
