package repository

import (
	"scientific-workflow-provenance/internal/domain/workflow"
	"sync"
	"testing"
)

func sampleRun() workflow.Run {
	return workflow.Run{ID: "run-1", NodeStates: map[string]workflow.RunStatus{"step": workflow.Queued}, Attempts: map[string]int{"step": 1}, Parameters: map[string]string{"mode": "strict"}}
}

func TestRunRepoGetSnapshot(t *testing.T) {
	r := NewRunRepo()
	r.Put(sampleRun())
	start := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		<-start
		v, _ := r.Get("run-1")
		v.NodeStates["step"] = workflow.Failed
		v.Parameters["mode"] = "fast"
	}()
	go func() { defer wg.Done(); <-start; v, _ := r.Get("run-1"); v.NodeStates["step"] = workflow.Running }()
	close(start)
	wg.Wait()
	again, _ := r.Get("run-1")
	if again.NodeStates["step"] != workflow.Queued || again.Parameters["mode"] != "strict" {
		t.Fatal("repository state was aliased")
	}
}

func TestRunRepoListSnapshot(t *testing.T) {
	r := NewRunRepo()
	r.Put(sampleRun())
	items := r.List()
	if len(items) != 1 {
		t.Fatal("list did not return the stored run")
	}
	start := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { defer wg.Done(); <-start; items[0].Attempts["step"] = 99 }()
	go func() { defer wg.Done(); <-start; v, _ := r.Get("run-1"); v.Attempts["step"] = 7 }()
	close(start)
	wg.Wait()
	again, _ := r.Get("run-1")
	if again.Attempts["step"] != 1 {
		t.Fatal("list leaked mutable attempt map")
	}
}
