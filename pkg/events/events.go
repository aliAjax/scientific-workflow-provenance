package events

import (
	"encoding/json"
	"sync"
	"time"
)

type Event struct {
	ID       string          `json:"id"`
	Type     string          `json:"type"`
	Subject  string          `json:"subject"`
	Data     json.RawMessage `json:"data"`
	At       time.Time       `json:"at"`
	Sequence uint64          `json:"sequence"`
}
type Bus struct {
	mu          sync.RWMutex
	seq         uint64
	subscribers map[string][]chan Event
}

func New() *Bus { return &Bus{subscribers: map[string][]chan Event{}} }
func (b *Bus) Subscribe(topic string, buffer int) <-chan Event {
	if buffer < 1 {
		buffer = 16
	}
	c := make(chan Event, buffer)
	b.mu.Lock()
	b.subscribers[topic] = append(b.subscribers[topic], c)
	b.mu.Unlock()
	return c
}
func (b *Bus) Publish(e Event) {
	b.mu.Lock()
	b.seq++
	e.Sequence = b.seq
	if e.At.IsZero() {
		e.At = time.Now().UTC()
	}
	for _, c := range b.subscribers[e.Type] {
		select {
		case c <- e:
		default:
		}
	}
	b.mu.Unlock()
}
