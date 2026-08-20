package application

import (
	"context"
	"scientific-workflow-provenance/internal/domain/workflow"
	"scientific-workflow-provenance/internal/repository"
	"sync"
	"testing"
)

func TestRunServiceRunIsolation(t *testing.T) {
	w := repository.NewWorkflowRepo()
	_ = w.Put(workflow.Definition{ID: "wf", Name: "demo", Version: 1, Nodes: []workflow.Node{{ID: "step"}}})
	r := repository.NewRunRepo()
	svc := NewRunService(w, r, nil)
	params1 := map[string]string{"mode": "strict"}
	params2 := map[string]string{"mode": "strict"}
	start := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	var run workflow.Run
	var err error
	go func() { defer wg.Done(); <-start; run, err = svc.Create(context.Background(), "wf", 1, params1) }()
	go func() { defer wg.Done(); _, _ = svc.Create(context.Background(), "wf", 1, params2) }()
	close(start)
	wg.Wait()
	if err != nil {
		t.Fatal(err)
	}
	params := params1
	params["mode"] = "changed"
	run.Parameters["mode"] = "changed"
	saved, _ := r.Get(run.ID)
	if saved.Parameters["mode"] != "strict" {
		t.Fatal("service leaked run parameters")
	}
}
