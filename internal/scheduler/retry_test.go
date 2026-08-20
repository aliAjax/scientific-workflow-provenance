package scheduler

import (
	"scientific-workflow-provenance/internal/domain/workflow"
	"testing"
)

func TestShouldRetryHonorsRetryableClass(t *testing.T) {
	p := workflow.RetryPolicy{MaxAttempts: 3, Retryable: []string{"timeout"}}
	if ShouldRetry(p, 1, "quota") {
		t.Fatal("non-retryable error was accepted")
	}
	if !ShouldRetry(p, 1, "timeout") {
		t.Fatal("retryable error was rejected")
	}
}
