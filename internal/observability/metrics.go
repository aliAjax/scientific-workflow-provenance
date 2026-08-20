package observability

import (
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	value uint64
}

func (c *Counter) Add(n uint64) { c.mu.Lock(); c.value += n; c.mu.Unlock() }
func (c *Counter) Get() uint64  { c.mu.Lock(); defer c.mu.Unlock(); return c.value }

type Histogram struct {
	mu     sync.Mutex
	values []float64
}

func (h *Histogram) Observe(v float64) { h.mu.Lock(); h.values = append(h.values, v); h.mu.Unlock() }
func (h *Histogram) Summary() (count int, avg float64) {
	h.mu.Lock()
	defer h.mu.Unlock()
	count = len(h.values)
	for _, v := range h.values {
		avg += v
	}
	if count > 0 {
		avg /= float64(count)
	}
	return
}

type Timer struct {
	start time.Time
	h     *Histogram
}

func Start(h *Histogram) Timer { return Timer{time.Now(), h} }
func (t Timer) Stop() {
	if t.h != nil {
		t.h.Observe(time.Since(t.start).Seconds())
	}
}
