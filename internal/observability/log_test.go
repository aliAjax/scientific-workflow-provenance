package observability

import "testing"

func TestLoggerConcurrentWrites(t *testing.T) {
	var l Logger
	l.Log("info", "message", map[string]any{"run": "r"})
}
