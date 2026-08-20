package transport

import (
	"encoding/json"
	"net/http/httptest"
	"scientific-workflow-provenance/internal/domain/workflow"
	"testing"
)

func TestRunHTTPReportsTerminalState(t *testing.T) {
	s := NewServer()
	s.runs.Put(workflow.Run{ID: "r", Status: workflow.Retrying})
	req := httptest.NewRequest("GET", "/api/v1/runs/r", nil)
	rr := httptest.NewRecorder()
	s.Router().ServeHTTP(rr, req)
	var got workflow.Run
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatal(err)
	}
	if got.Status != workflow.Retrying {
		t.Fatalf("status was rewritten: %s", got.Status)
	}
}
