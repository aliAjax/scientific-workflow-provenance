package query

import (
	"scientific-workflow-provenance/internal/domain/workflow"
	"sort"
)

func Facets(values []workflow.Run) map[string]int {
	out := map[string]int{}
	for _, v := range values {
		out[string(v.Status)]++
	}
	return out
}
func WorkflowIDs(values []workflow.Run) []string {
	set := map[string]bool{}
	for _, v := range values {
		set[v.WorkflowID] = true
	}
	out := []string{}
	for k := range set {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
