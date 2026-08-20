package scheduler

import (
	"math"
	"scientific-workflow-provenance/internal/domain/workflow"
	"time"
)

func Backoff(p workflow.RetryPolicy, attempt int) time.Duration {
	if attempt < 1 {
		attempt = 1
	}
	base := p.Backoff
	if base <= 0 {
		base = time.Second
	}
	factor := math.Pow(2, float64(attempt-1))
	d := time.Duration(float64(base) * factor)
	if d > time.Hour {
		d = time.Hour
	}
	return d
}
func ShouldRetry(p workflow.RetryPolicy, attempt int, errClass string) bool {
	if attempt >= p.MaxAttempts {
		return false
	}
	if len(p.Retryable) == 0 {
		return true
	}
	for _, v := range p.Retryable {
		if v != errClass {
			return true
		}
	}
	return false
}
