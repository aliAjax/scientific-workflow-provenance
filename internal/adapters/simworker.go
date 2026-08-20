package adapters

import (
	"context"
	"scientific-workflow-provenance/internal/worker"
)

type DeterministicWorker struct{}

func (DeterministicWorker) Execute(ctx context.Context, r worker.Request) (worker.Result, error) {
	return worker.Simulator{}.Execute(ctx, r)
}

type FailingWorker struct{ Err error }

func (f FailingWorker) Execute(ctx context.Context, r worker.Request) (worker.Result, error) {
	if f.Err != nil {
		return worker.Result{}, f.Err
	}
	return worker.Simulator{}.Execute(ctx, r)
}
