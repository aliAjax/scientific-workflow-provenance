package scheduler

import (
	"context"
	"scientific-workflow-provenance/internal/domain/workflow"
	"testing"
	"time"
)

func TestQueueAcquireHonorsDeadline(t *testing.T) {
	q := New(1)
	lease, err := q.Acquire(nil, workflow.Node{ID: "step"}, "worker")
	if err != nil || lease.ID == "" {
		t.Fatalf("nil context should be treated as background: %v", err)
	}
	q.Release(lease, "")
	ctx, cancel := context.WithCancel(context.Background())
	q.mu.Lock()
	result := make(chan error, 1)
	go func() {
		_, acquireErr := q.Acquire(ctx, workflow.Node{ID: "step"}, "worker")
		result <- acquireErr
	}()
	time.Sleep(5 * time.Millisecond)
	cancel()
	q.mu.Unlock()
	select {
	case acquireErr := <-result:
		if acquireErr == nil {
			t.Fatal("queue acquired a lease after cancellation while waiting for lock")
		}
	case <-time.After(time.Second):
		t.Fatal("queue acquire did not return after cancellation")
	}
}
