package adapters

import (
	"context"
	"errors"
	"scientific-workflow-provenance/internal/worker"
	"testing"
)

func TestFailingWorkerHonorsContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	sentinel := errors.New("worker failed")
	_, err := (FailingWorker{Err: sentinel}).Execute(ctx, worker.Request{StepID: "s"})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("worker ignored cancellation: %v", err)
	}
}
