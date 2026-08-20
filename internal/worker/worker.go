package worker

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type Request struct {
	RunID, StepID string
	Inputs        map[string]string
	Parameters    map[string]string
}
type Result struct {
	Outputs               map[string][]byte
	Logs                  []string
	StartedAt, FinishedAt time.Time
	Digest                string
}
type Adapter interface {
	Execute(context.Context, Request) (Result, error)
}
type Simulator struct{}

func (Simulator) Execute(ctx context.Context, r Request) (Result, error) {
	start := time.Now().UTC()
	select {
	case <-ctx.Done():
		return Result{}, ctx.Err()
	case <-time.After(5 * time.Millisecond):
	}
	if r.StepID == "" {
		return Result{}, errors.New("step id required")
	}
	out := map[string][]byte{"result": []byte(fmt.Sprintf("run=%s step=%s", r.RunID, r.StepID))}
	return Result{Outputs: out, Logs: []string{"simulator completed"}, StartedAt: start, FinishedAt: time.Now().UTC(), Digest: fmt.Sprintf("%x", []byte(r.RunID+":"+r.StepID))}, nil
}
