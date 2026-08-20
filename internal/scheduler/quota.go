package scheduler

import (
	"errors"
	"sync"
)

type Quota struct {
	mu    sync.Mutex
	limit map[string]int
	used  map[string]int
}

func NewQuota() *Quota { return &Quota{limit: map[string]int{}, used: map[string]int{}} }
func (q *Quota) Set(tenant string, limit int) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.limit[tenant] = limit
}
func (q *Quota) Reserve(tenant string, n int) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	if n < 1 {
		return errors.New("reservation must be positive")
	}
	if q.limit[tenant] > 0 && q.used[tenant]+n > q.limit[tenant] {
		return errors.New("tenant quota exceeded")
	}
	q.used[tenant] += n
	return nil
}
func (q *Quota) Release(tenant string, n int) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.used[tenant] -= n
	if q.used[tenant] < 0 {
		q.used[tenant] = 0
	}
}
func (q *Quota) Usage(tenant string) int { q.mu.Lock(); defer q.mu.Unlock(); return q.used[tenant] }
