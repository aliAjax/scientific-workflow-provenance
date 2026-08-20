package application

import (
	"fmt"
	"scientific-workflow-provenance/internal/domain/workflow"
	"scientific-workflow-provenance/internal/planner"
	"scientific-workflow-provenance/internal/repository"
	"scientific-workflow-provenance/pkg/ids"
)

type WorkflowService struct {
	Repo    *repository.WorkflowRepo
	Planner planner.Planner
}

func NewWorkflowService(r *repository.WorkflowRepo) *WorkflowService {
	return &WorkflowService{Repo: r}
}
func (s *WorkflowService) Create(d workflow.Definition) (workflow.Definition, error) {
	if d.ID == "" {
		d.ID = ids.New("wf")
	}
	if d.Version == 0 {
		d.Version = 1
	}
	if d.Digest == "" {
		d.Digest = d.ComputeDigest()
	}
	if err := s.Repo.Put(d); err != nil {
		return workflow.Definition{}, err
	}
	return s.Repo.Get(d.ID, d.Version)
}
func (s *WorkflowService) Plan(id string, version int, params map[string]string) (planner.Plan, error) {
	d, e := s.Repo.Get(id, version)
	if e != nil {
		return planner.Plan{}, e
	}
	p, e := s.Planner.Build(d, params)
	if e != nil {
		return planner.Plan{}, fmt.Errorf("plan workflow: %w", e)
	}
	return p, nil
}
