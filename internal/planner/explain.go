package planner

import (
	"fmt"
	"strings"
)

type Explanation struct {
	StepID  string   `json:"stepId"`
	Reasons []string `json:"reasons"`
}

func Explain(p Plan) []Explanation {
	out := make([]Explanation, 0, len(p.Steps))
	for _, s := range p.Steps {
		r := []string{}
		if len(s.DependsOn) == 0 {
			r = append(r, "source node")
		}
		if len(s.Matrix) > 0 {
			r = append(r, fmt.Sprintf("matrix expansion with %d dimensions", len(s.Matrix)))
		}
		if s.CacheKey != "" {
			r = append(r, "deterministic cache key computed")
		}
		out = append(out, Explanation{StepID: s.ID, Reasons: r})
	}
	return out
}
func Format(e Explanation) string { return e.StepID + ": " + strings.Join(e.Reasons, ", ") }
