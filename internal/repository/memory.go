package repository

import (
	"errors"
	"scientific-workflow-provenance/internal/domain/workflow"
	"sync"
)

type WorkflowRepo struct {
	mu   sync.RWMutex
	data map[string][]workflow.Definition
}

func NewWorkflowRepo() *WorkflowRepo { return &WorkflowRepo{data: map[string][]workflow.Definition{}} }
func (r *WorkflowRepo) Put(d workflow.Definition) error {
	if err := d.Validate(); err != nil {
		return err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	d.Digest = d.ComputeDigest()
	r.data[d.ID] = append(r.data[d.ID], d)
	return nil
}
func (r *WorkflowRepo) Get(id string, version int) (workflow.Definition, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, d := range r.data[id] {
		if version == 0 || d.Version == version {
			return d, nil
		}
	}
	return workflow.Definition{}, errors.New("workflow not found")
}
func (r *WorkflowRepo) List() []workflow.Definition {
	r.mu.RLock()
	defer r.mu.RUnlock()
	o := []workflow.Definition{}
	for _, v := range r.data {
		if len(v) > 0 {
			o = append(o, v[len(v)-1])
		}
	}
	return o
}

type RunRepo struct {
	mu   sync.RWMutex
	data map[string]workflow.Run
}

func NewRunRepo() *RunRepo            { return &RunRepo{data: map[string]workflow.Run{}} }
func (r *RunRepo) Put(v workflow.Run) { r.mu.Lock(); defer r.mu.Unlock(); r.data[v.ID] = v }
func (r *RunRepo) Get(id string) (workflow.Run, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.data[id]
	return v, ok
}
func (r *RunRepo) List() []workflow.Run {
	r.mu.RLock()
	defer r.mu.RUnlock()
	o := []workflow.Run{}
	for _, v := range r.data {
		o = append(o, v)
	}
	return o
}
