package workflow

import (
	"sort"
	"strings"
)

type DependencySet struct {
	Node    string
	Parents []string
}

func Dependencies(d Definition) map[string]DependencySet {
	m := map[string]DependencySet{}
	for _, n := range d.Nodes {
		m[n.ID] = DependencySet{Node: n.ID}
	}
	for _, e := range d.Edges {
		x := m[e.To]
		x.Parents = append(x.Parents, e.From)
		m[e.To] = x
	}
	for k := range m {
		sort.Strings(m[k].Parents)
	}
	return m
}
func MissingInputs(n Node, available map[string]bool) []string {
	out := []string{}
	for _, v := range n.Inputs {
		if !available[v] {
			out = append(out, v)
		}
	}
	sort.Strings(out)
	return out
}
func OutputNames(d Definition) []string {
	set := map[string]bool{}
	for _, n := range d.Nodes {
		for _, v := range n.Outputs {
			set[strings.TrimSpace(v)] = true
		}
	}
	out := []string{}
	for v := range set {
		if v != "" {
			out = append(out, v)
		}
	}
	sort.Strings(out)
	return out
}
