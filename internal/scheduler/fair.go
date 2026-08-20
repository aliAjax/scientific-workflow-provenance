package scheduler

import (
	"sort"
	"sync"
	"time"
)

type Job struct {
	ID, Tenant string
	Priority   int
	EnqueuedAt time.Time
}
type FairQueue struct {
	mu      sync.Mutex
	items   []Job
	credits map[string]int
}

func NewFairQueue() *FairQueue { return &FairQueue{credits: map[string]int{}} }
func (q *FairQueue) Push(j Job) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if j.EnqueuedAt.IsZero() {
		j.EnqueuedAt = time.Now().UTC()
	}
	q.items = append(q.items, j)
	q.credits[j.Tenant]++
}
func (q *FairQueue) Pop() (Job, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.items) == 0 {
		return Job{}, false
	}
	sort.SliceStable(q.items, func(i, j int) bool {
		a, b := q.items[i], q.items[j]
		if a.Priority != b.Priority {
			return a.Priority > b.Priority
		}
		if q.credits[a.Tenant] != q.credits[b.Tenant] {
			return q.credits[a.Tenant] < q.credits[b.Tenant]
		}
		return a.EnqueuedAt.Before(b.EnqueuedAt)
	})
	j := q.items[0]
	q.items = q.items[1:]
	if q.credits[j.Tenant] > 0 {
		q.credits[j.Tenant]--
	}
	return j, true
}
