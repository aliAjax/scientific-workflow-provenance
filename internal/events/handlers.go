package events

import (
	"encoding/json"
	"scientific-workflow-provenance/pkg/events"
	"sync"
)

type Handler func(events.Event)
type Dispatcher struct {
	mu       sync.RWMutex
	handlers map[string][]Handler
}

func New() *Dispatcher { return &Dispatcher{handlers: map[string][]Handler{}} }
func (d *Dispatcher) On(typ string, h Handler) {
	d.mu.Lock()
	d.handlers[typ] = append(d.handlers[typ], h)
	d.mu.Unlock()
}
func (d *Dispatcher) Dispatch(e events.Event) {
	d.mu.RLock()
	hs := append([]Handler(nil), d.handlers[e.Type]...)
	d.mu.RUnlock()
	for _, h := range hs {
		h(e)
	}
}
func JSON(typ, subject string, v any) events.Event {
	b, _ := json.Marshal(v)
	return events.Event{Type: typ, Subject: subject, Data: b}
}
