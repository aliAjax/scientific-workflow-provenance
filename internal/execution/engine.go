package execution

import (
	"context"
	"errors"
	"fmt"
	"scientific-workflow-provenance/internal/domain/workflow"
	"scientific-workflow-provenance/internal/planner"
	"scientific-workflow-provenance/internal/worker"
	"sync"
	"time"
)

type Attempt struct {
	ID         string             `json:"id"`
	StepID     string             `json:"stepId"`
	Number     int                `json:"number"`
	Status     workflow.RunStatus `json:"status"`
	StartedAt  time.Time          `json:"startedAt"`
	FinishedAt *time.Time         `json:"finishedAt,omitempty"`
	Error      string             `json:"error,omitempty"`
}
type Engine struct {
	mu       sync.Mutex
	adapter  worker.Adapter
	attempts map[string][]Attempt
}

func New(a worker.Adapter) *Engine {
	if a == nil {
		a = worker.Simulator{}
	}
	return &Engine{adapter: a, attempts: map[string][]Attempt{}}
}
func (e *Engine) Execute(ctx context.Context, runID string, step planner.Step, params map[string]string) (Attempt, worker.Result, error) {
	e.mu.Lock()
	key := runID + "/" + step.ID
	num := len(e.attempts[key]) + 1
	a := Attempt{ID: fmt.Sprintf("%s-%d", key, num), StepID: step.ID, Number: num, Status: workflow.Running, StartedAt: time.Now().UTC()}
	e.attempts[key] = append(e.attempts[key], a)
	e.mu.Unlock()
	res, err := e.adapter.Execute(ctx, worker.Request{RunID: runID, StepID: step.ID, Parameters: params})
	e.mu.Lock()
	defer e.mu.Unlock()
	v := e.attempts[key][num-1]
	now := time.Now().UTC()
	v.FinishedAt = &now
	if err != nil {
		v.Status = workflow.Failed
		v.Error = err.Error()
	} else {
		v.Status = workflow.Succeeded
	}
	e.attempts[key][num-1] = v
	return v, res, err
}
func (e *Engine) Attempts(runID, stepID string) []Attempt {
	e.mu.Lock()
	defer e.mu.Unlock()
	return append([]Attempt(nil), e.attempts[runID+"/"+stepID]...)
}
func (e *Engine) Cancel(ctx context.Context) error {
	if ctx == nil {
		return errors.New("context required")
	}
	return nil
}
