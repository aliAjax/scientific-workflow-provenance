package scheduler

import (
	"context"
	"errors"
	"fmt"
	"scientific-workflow-provenance/internal/domain/workflow"
	"sync"
	"time"
)

type Lease struct {
	ID        string    `json:"id"`
	StepID    string    `json:"stepId"`
	Owner     string    `json:"owner"`
	ExpiresAt time.Time `json:"expiresAt"`
}
type Queue struct {
	mu      sync.Mutex
	leases  map[string]Lease
	running map[string]int
	limit   int
}

func New(limit int) *Queue {
	if limit < 1 {
		limit = 1
	}
	return &Queue{leases: map[string]Lease{}, running: map[string]int{}, limit: limit}
}
func (q *Queue) Acquire(ctx context.Context, step workflow.Node, owner string) (Lease, error) {
	select {
	case <-ctx.Done():
		return Lease{}, ctx.Err()
	default:
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	pool := step.Resources.Pool
	if q.running[pool] >= q.limit {
		return Lease{}, errors.New("resource pool exhausted")
	}
	id := fmt.Sprintf("lease-%d", time.Now().UnixNano())
	l := Lease{ID: id, StepID: step.ID, Owner: owner, ExpiresAt: time.Now().Add(2 * time.Minute)}
	q.leases[id] = l
	q.running[pool]++
	return l, nil
}
func (q *Queue) Release(l Lease, pool string) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if _, ok := q.leases[l.ID]; ok {
		delete(q.leases, l.ID)
		if q.running[pool] > 0 {
			q.running[pool]--
		}
	}
}
func (q *Queue) Recover(now time.Time) []Lease {
	q.mu.Lock()
	defer q.mu.Unlock()
	out := []Lease{}
	for id, l := range q.leases {
		if now.After(l.ExpiresAt) {
			delete(q.leases, id)
			out = append(out, l)
		}
	}
	return out
}
func (q *Queue) Active() []Lease {
	q.mu.Lock()
	defer q.mu.Unlock()
	o := make([]Lease, 0, len(q.leases))
	for _, l := range q.leases {
		o = append(o, l)
	}
	return o
}
