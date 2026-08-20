package provenance

import (
	"sort"
	"time"
)

func (g *Graph) LinkDerived(parent, child string) Relation {
	return g.Relate(Relation{ID: parent + "->" + child, From: parent, To: child, Type: "wasDerivedFrom", At: time.Now().UTC()})
}
func (g *Graph) ActivitiesFor(entity string) []Record {
	g.mu.RLock()
	defer g.mu.RUnlock()
	ids := map[string]bool{}
	for _, r := range g.Relations {
		if r.From == entity {
			ids[r.To] = true
		}
		if r.To == entity {
			ids[r.From] = true
		}
	}
	out := []Record{}
	for id := range ids {
		if r, ok := g.Records[id]; ok && r.Kind == Activity {
			out = append(out, r)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}
