package execution

import (
	"context"
	"scientific-workflow-provenance/internal/planner"
	"scientific-workflow-provenance/internal/worker"
	"sync"
	"testing"
)

type mutatingAdapter struct{}

func (mutatingAdapter) Execute(_ context.Context, r worker.Request) (worker.Result, error) {
	r.Parameters["mode"] = "mutated"
	return worker.Result{}, nil
}

func TestEngineAttemptsSnapshot(t *testing.T) {
	e := New(mutatingAdapter{})
	params1 := map[string]string{"mode": "strict"}
	params2 := map[string]string{"mode": "strict"}
	start := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		<-start
		_, _, _ = e.Execute(context.Background(), "run-1", planner.Step{ID: "step"}, params1)
	}()
	go func() {
		defer wg.Done()
		_, _, _ = e.Execute(context.Background(), "run-1", planner.Step{ID: "step"}, params2)
	}()
	close(start)
	wg.Wait()
	params := params1
	if params["mode"] != "strict" {
		t.Fatal("adapter mutated caller parameters")
	}
}
