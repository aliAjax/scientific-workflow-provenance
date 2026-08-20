package query

import (
	"scientific-workflow-provenance/internal/domain/workflow"
	"testing"
)

func TestRunsFilterRetrying(t *testing.T) {
	items := InProgress([]workflow.Run{{ID: "q", Status: workflow.Queued}, {ID: "r", Status: workflow.Retrying}, {ID: "w", Status: workflow.WaitingInput}})
	if len(items) != 3 {
		t.Fatalf("in-progress states missing: %#v", items)
	}
}
