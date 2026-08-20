package provenance

import (
	"sort"
	"strings"
)

func (g *Graph) Query(kind Kind, typeName string) []Record {
	g.mu.RLock()
	defer g.mu.RUnlock()
	out := []Record{}
	for _, r := range g.Records {
		if (kind == "" || r.Kind == kind) && (typeName == "" || strings.EqualFold(r.Type, typeName)) {
			out = append(out, r)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}
func (g *Graph) Neighbors(id string) []Relation {
	g.mu.RLock()
	defer g.mu.RUnlock()
	out := []Relation{}
	for _, r := range g.Relations {
		if r.From == id || r.To == id {
			out = append(out, r)
		}
	}
	return out
}
