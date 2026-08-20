package metrics

import (
	"sync"
	"time"
)

type Registry struct {
	mu       sync.Mutex
	counters map[string]uint64
	gauges   map[string]float64
	hist     map[string][]float64
}

func New() *Registry {
	return &Registry{counters: map[string]uint64{}, gauges: map[string]float64{}, hist: map[string][]float64{}}
}
func (r *Registry) Inc(name string, n uint64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.counters == nil {
		r.counters = map[string]uint64{}
	}
	r.counters[name] += n
}
func (r *Registry) Set(name string, v float64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.gauges == nil {
		r.gauges = map[string]float64{}
	}
	r.gauges[name] = v
}
func (r *Registry) Observe(name string, v float64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.hist == nil {
		r.hist = map[string][]float64{}
	}
	r.hist[name] = append(r.hist[name], v)
}
func (r *Registry) Snapshot() map[string]any {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := map[string]any{}
	for k, v := range r.counters {
		out[k] = v
	}
	for k, v := range r.gauges {
		out[k] = v
	}
	return out
}
func Since(start time.Time) float64 { return time.Since(start).Seconds() }
