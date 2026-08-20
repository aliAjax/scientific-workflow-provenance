package worker

import (
	"sync"
	"time"
)

type Heartbeat struct {
	WorkerID string    `json:"workerId"`
	At       time.Time `json:"at"`
	Load     float64   `json:"load"`
	Capacity int       `json:"capacity"`
}
type Registry struct {
	mu    sync.RWMutex
	nodes map[string]Heartbeat
}

func NewRegistry() *Registry { return &Registry{nodes: map[string]Heartbeat{}} }
func (r *Registry) Report(h Heartbeat) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if h.At.IsZero() {
		h.At = time.Now().UTC()
	}
	r.nodes[h.WorkerID] = h
}
func (r *Registry) Healthy(now time.Time, ttl time.Duration) []Heartbeat {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := []Heartbeat{}
	for _, h := range r.nodes {
		if now.Sub(h.At) <= ttl {
			out = append(out, h)
		}
	}
	return out
}
func (r *Registry) Remove(id string) { r.mu.Lock(); defer r.mu.Unlock(); delete(r.nodes, id) }
