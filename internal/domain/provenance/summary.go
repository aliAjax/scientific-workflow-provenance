package provenance

import "sort"

type Summary struct{ Entities, Activities, Agents, Relations int }

func (g *Graph) Summary() Summary {
	g.mu.RLock()
	defer g.mu.RUnlock()
	s := Summary{Relations: len(g.Relations)}
	for _, r := range g.Records {
		switch r.Kind {
		case Entity:
			s.Entities++
		case Activity:
			s.Activities++
		case Agent:
			s.Agents++
		}
	}
	return s
}
func (g *Graph) RecordIDs() []string {
	g.mu.RLock()
	defer g.mu.RUnlock()
	out := make([]string, 0, len(g.Records))
	for id := range g.Records {
		out = append(out, id)
	}
	sort.Strings(out)
	return out
}
