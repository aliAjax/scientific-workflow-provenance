package planner

import (
	"scientific-workflow-provenance/internal/domain/workflow"
	"sort"
)

func Topological(d workflow.Definition) []string {
	in := map[string]int{}
	adj := map[string][]string{}
	for _, n := range d.Nodes {
		in[n.ID] = 0
	}
	for _, e := range d.Edges {
		in[e.To]++
		adj[e.From] = append(adj[e.From], e.To)
	}
	q := []string{}
	for id, n := range in {
		if n == 0 {
			q = append(q, id)
		}
	}
	sort.Strings(q)
	out := []string{}
	for len(q) > 0 {
		id := q[0]
		q = q[1:]
		out = append(out, id)
		for _, to := range adj[id] {
			in[to]--
			if in[to] == 0 {
				q = append(q, to)
				sort.Strings(q)
			}
		}
	}
	return out
}
func Sources(d workflow.Definition) []string {
	all := map[string]bool{}
	for _, n := range d.Nodes {
		all[n.ID] = true
	}
	for _, e := range d.Edges {
		delete(all, e.To)
	}
	out := []string{}
	for id := range all {
		out = append(out, id)
	}
	sort.Strings(out)
	return out
}
