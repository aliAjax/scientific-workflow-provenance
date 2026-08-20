package adapters

import (
	"context"
	"scientific-workflow-provenance/internal/worker"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type cancellationProbe struct {
	started chan struct{}
	once    sync.Once
	calls   atomic.Int32
}

func (p *cancellationProbe) Execute(context.Context, worker.Request) (worker.Result, error) {
	p.calls.Add(1)
	p.once.Do(func() { close(p.started) })
	return worker.Result{}, nil
}

func TestBatchExecutorPropagatesCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	probe := &cancellationProbe{started: make(chan struct{})}
	requests := make([]worker.Request, 32)
	for i := range requests {
		requests[i] = worker.Request{StepID: "cancelled"}
	}
	results, errs := (BatchExecutor{Adapter: probe, Concurrency: 1}).Execute(ctx, requests)
	if len(results) != len(requests) || len(errs) != len(requests) {
		t.Fatalf("cancelled batch returned unexpected result: %#v %#v", results, errs)
	}
	for i, err := range errs {
		if err == nil {
			t.Fatalf("request %d was not cancelled", i)
		}
	}
	select {
	case <-probe.started:
		t.Fatalf("cancelled batch invoked adapter %d times", probe.calls.Load())
	case <-time.After(20 * time.Millisecond):
	}
}
