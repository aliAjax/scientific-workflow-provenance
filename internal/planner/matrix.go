package planner

import (
	"fmt"
	"scientific-workflow-provenance/internal/domain/workflow"
	"sort"
)

func Expand(n workflow.Node) []Step {
	if len(n.Matrix) == 0 {
		return []Step{{ID: n.ID, NodeID: n.ID}}
	}
	out := make([]Step, 0, len(n.Matrix))
	for i, m := range n.Matrix {
		keys := make([]string, 0, len(m))
		for k := range m {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		canonical := ""
		for _, k := range keys {
			canonical += fmt.Sprintf("%s=%s;", k, m[k])
		}
		out = append(out, Step{ID: fmt.Sprintf("%s-%d", n.ID, i), NodeID: n.ID, Matrix: m, CacheKey: canonical})
	}
	return out
}
func MatrixValues(n workflow.Node, key string) []string {
	out := []string{}
	for _, m := range n.Matrix {
		if v, ok := m[key]; ok {
			out = append(out, v)
		}
	}
	sort.Strings(out)
	return out
}
