package scheduler

import (
	"errors"
	"sync"
	"time"
)

type Limit struct {
	MaxConcurrent int
	MaxDuration   time.Duration
}
type Limiter struct {
	mu     sync.Mutex
	limits map[string]Limit
	active map[string]int
}

func NewLimiter() *Limiter                        { return &Limiter{limits: map[string]Limit{}, active: map[string]int{}} }
func (l *Limiter) Configure(kind string, v Limit) { l.mu.Lock(); l.limits[kind] = v; l.mu.Unlock() }
func (l *Limiter) Enter(kind string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	cfg := l.limits[kind]
	if cfg.MaxConcurrent > 0 && l.active[kind] >= cfg.MaxConcurrent {
		return errors.New("concurrency limit reached")
	}
	l.active[kind]++
	return nil
}
func (l *Limiter) Leave(kind string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.active[kind] > 0 {
		l.active[kind]--
	}
}
