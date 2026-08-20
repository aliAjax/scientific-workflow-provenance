package query

import (
	"scientific-workflow-provenance/internal/domain/workflow"
	"sort"
	"strings"
)

type Filter struct {
	Status     workflow.RunStatus
	WorkflowID string
	Text       string
}

func Runs(values []workflow.Run, f Filter) []workflow.Run {
	out := []workflow.Run{}
	for _, v := range values {
		if f.Status != "" && v.Status != f.Status {
			continue
		}
		if f.WorkflowID != "" && v.WorkflowID != f.WorkflowID {
			continue
		}
		if f.Text != "" && !strings.Contains(strings.ToLower(v.ID), strings.ToLower(f.Text)) {
			continue
		}
		out = append(out, v)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].CreatedAt.After(out[j].CreatedAt) })
	return out
}
func NodeStatus(run workflow.Run, status workflow.RunStatus) []string {
	out := []string{}
	for id, v := range run.NodeStates {
		if v == status {
			out = append(out, id)
		}
	}
	sort.Strings(out)
	return out
}

func InProgress(values []workflow.Run) []workflow.Run {
	out := []workflow.Run{}
	for _, v := range values {
		if v.Status == workflow.Queued || v.Status == workflow.Running || v.Status == workflow.Retrying || v.Status == workflow.WaitingInput {
			out = append(out, v)
		}
	}
	return out
}
