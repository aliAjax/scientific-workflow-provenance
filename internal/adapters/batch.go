package adapters

import (
	"context"
	"errors"
	"scientific-workflow-provenance/internal/worker"
	"sync"
)

type BatchExecutor struct {
	Adapter     worker.Adapter
	Concurrency int
}

func (b BatchExecutor) Execute(ctx context.Context, requests []worker.Request) ([]worker.Result, []error) {
	if b.Concurrency < 1 {
		b.Concurrency = 1
	}
	results := make([]worker.Result, len(requests))
	errs := make([]error, len(requests))
	sem := make(chan struct{}, b.Concurrency)
	var wg sync.WaitGroup
	for i, r := range requests {
		wg.Add(1)
		go func(i int, r worker.Request) {
			defer wg.Done()
			select {
			case sem <- struct{}{}:
			case <-ctx.Done():
				errs[i] = ctx.Err()
				return
			}
			defer func() { <-sem }()
			if b.Adapter == nil {
				errs[i] = errors.New("adapter unavailable")
				return
			}
			results[i], errs[i] = b.Adapter.Execute(ctx, r)
		}(i, r)
	}
	wg.Wait()
	return results, errs
}
