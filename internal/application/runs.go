package application

import (
	"context"
	"errors"
	"scientific-workflow-provenance/internal/domain/workflow"
	"scientific-workflow-provenance/internal/execution"
	"scientific-workflow-provenance/internal/planner"
	"scientific-workflow-provenance/internal/repository"
	"scientific-workflow-provenance/internal/worker"
	"scientific-workflow-provenance/pkg/ids"
	"time"
)

type RunService struct {
	Workflows *repository.WorkflowRepo
	Runs      *repository.RunRepo
	Engine    *execution.Engine
}

func NewRunService(w *repository.WorkflowRepo, r *repository.RunRepo, a worker.Adapter) *RunService {
	return &RunService{Workflows: w, Runs: r, Engine: execution.New(a)}
}
func (s *RunService) Create(ctx context.Context, wid string, version int, params map[string]string) (workflow.Run, error) {
	d, e := s.Workflows.Get(wid, version)
	if e != nil {
		return workflow.Run{}, e
	}
	p, e := planner.Planner{}.Build(d, params)
	if e != nil {
		return workflow.Run{}, e
	}
	r := workflow.Run{ID: ids.New("run"), WorkflowID: wid, Version: d.Version, Status: workflow.Planned, NodeStates: map[string]workflow.RunStatus{}, Attempts: map[string]int{}, Parameters: params, CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC()}
	for _, st := range p.Steps {
		r.NodeStates[st.ID] = workflow.Queued
	}
	s.Runs.Put(r)
	return r, nil
}
func (s *RunService) Cancel(id string) (workflow.Run, error) {
	r, ok := s.Runs.Get(id)
	if !ok {
		return workflow.Run{}, errors.New("run not found")
	}
	if workflow.Terminal(r.Status) {
		return r, nil
	}
	if e := workflow.Transition(&r, workflow.Cancelled); e != nil {
		return workflow.Run{}, e
	}
	r.UpdatedAt = time.Now().UTC()
	s.Runs.Put(r)
	return r, nil
}
func (s *RunService) ExecuteStep(ctx context.Context, id, step string) (execution.Attempt, worker.Result, error) {
	return s.Engine.Execute(ctx, id, planner.Step{ID: step}, nil)
}
