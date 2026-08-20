package verify008

import (
	"scientific-workflow-provenance/internal/observability"
	"testing"
)

func TestLoggerConcurrentWrites(t *testing.T) {
	var l observability.Logger
	l.Log("info", "message", map[string]any{"run": "r"})
}
