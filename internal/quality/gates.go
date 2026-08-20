package quality

import (
	"fmt"
	"math"
)

type Gate struct {
	Metric    string
	Min, Max  *float64
	Tolerance float64
}
type Result struct {
	Metric string
	Value  float64
	Passed bool
	Reason string
}

func Check(g Gate, v float64) Result {
	r := Result{Metric: g.Metric, Value: v, Passed: true}
	if math.IsNaN(v) || math.IsInf(v, 0) {
		r.Passed = false
		r.Reason = "not finite"
		return r
	}
	if g.Min != nil && v < *g.Min-g.Tolerance {
		r.Passed = false
		r.Reason = fmt.Sprintf("below %.4f", *g.Min)
	}
	if g.Max != nil && v > *g.Max+g.Tolerance {
		r.Passed = false
		r.Reason = fmt.Sprintf("above %.4f", *g.Max)
	}
	return r
}
func All(results []Result) bool {
	for _, r := range results {
		if !r.Passed {
			return false
		}
	}
	return true
}
