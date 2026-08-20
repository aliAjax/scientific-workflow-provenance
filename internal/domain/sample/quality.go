package sample

import (
	"fmt"
	"math"
	"sort"
)

type Report struct {
	SampleID string    `json:"sampleId"`
	Passed   bool      `json:"passed"`
	Metrics  []Quality `json:"metrics"`
	Warnings []string  `json:"warnings"`
}

func Evaluate(id string, metrics []Quality) Report {
	r := Report{SampleID: id, Passed: true, Metrics: metrics}
	for i := range r.Metrics {
		m := &r.Metrics[i]
		if math.IsNaN(m.Value) || math.IsInf(m.Value, 0) {
			m.Passed = false
			r.Passed = false
			r.Warnings = append(r.Warnings, fmt.Sprintf("%s is not finite", m.Metric))
			continue
		}
		if m.Threshold != 0 && m.Value < m.Threshold {
			m.Passed = false
			r.Passed = false
			r.Warnings = append(r.Warnings, fmt.Sprintf("%s below threshold", m.Metric))
		}
	}
	sort.Strings(r.Warnings)
	return r
}
