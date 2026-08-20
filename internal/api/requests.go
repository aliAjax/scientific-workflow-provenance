package api

import (
	"net/http"
	"strings"
)

type CreateWorkflow struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Version int    `json:"version"`
}
type CreateRun struct {
	WorkflowID string            `json:"workflowId"`
	Version    int               `json:"version"`
	Parameters map[string]string `json:"parameters"`
}
type CreateArtifact struct {
	Name, Schema, Unit string
	Value              float64
}

func IdempotencyKey(r *http.Request) string {
	return strings.TrimSpace(r.Header.Get("Idempotency-Key"))
}
func Accepts(r *http.Request, content string) bool {
	v := r.Header.Get("Accept")
	return v == "" || strings.Contains(v, content)
}
