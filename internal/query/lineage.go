package query

import (
	"scientific-workflow-provenance/internal/domain/provenance"
	"sort"
)

func Related(g *provenance.Graph, id string) []string {
	rels := g.Neighbors(id)
	out := []string{}
	for _, r := range rels {
		if r.From == id {
			out = append(out, r.To)
		} else {
			out = append(out, r.From)
		}
	}
	sort.Strings(out)
	return out
}
func Types(g *provenance.Graph) map[string]int {
	out := map[string]int{}
	for _, r := range g.Query("", "") {
		out[r.Type]++
	}
	return out
}
