package planner

import (
	"fmt"
	"scientific-workflow-provenance/internal/domain/workflow"
	"scientific-workflow-provenance/pkg/expr"
)

func EvaluateCondition(n workflow.Node, values map[string]float64) bool {
	if n.Condition == "" {
		return true
	}
	p, e := expr.Parse(n.Condition)
	return e == nil && p.Eval(values)
}
func ValidateCondition(n workflow.Node) error {
	if n.Condition == "" {
		return nil
	}
	if _, e := expr.Parse(n.Condition); e != nil {
		return fmt.Errorf("node %s: %w", n.ID, e)
	}
	return nil
}
