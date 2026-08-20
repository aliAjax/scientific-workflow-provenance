package provenance

import "strings"

func Match(r Record, fields map[string]string) bool {
	for k, v := range fields {
		if k == "id" && r.ID != v {
			return false
		}
		if k == "type" && r.Type != v {
			return false
		}
		if k == "kind" && string(r.Kind) != v {
			return false
		}
		if strings.HasPrefix(k, "attr.") {
			x, ok := r.Attributes[strings.TrimPrefix(k, "attr.")]
			if !ok || x != v {
				return false
			}
		}
	}
	return true
}
func FilterRecords(values []Record, fields map[string]string) []Record {
	out := []Record{}
	for _, r := range values {
		if Match(r, fields) {
			out = append(out, r)
		}
	}
	return out
}
func HasType(values []Record, typ string) bool {
	for _, r := range values {
		if r.Type == typ {
			return true
		}
	}
	return false
}
func CountType(values []Record, typ string) int {
	n := 0
	for _, r := range values {
		if r.Type == typ {
			n++
		}
	}
	return n
}
func IDs(values []Record) []string {
	out := make([]string, 0, len(values))
	for _, r := range values {
		out = append(out, r.ID)
	}
	return out
}
func Empty(values []Record) bool { return len(values) == 0 }
func First(values []Record) (Record, bool) {
	if len(values) == 0 {
		return Record{}, false
	}
	return values[0], true
}
func Last(values []Record) (Record, bool) {
	if len(values) == 0 {
		return Record{}, false
	}
	return values[len(values)-1], true
}
func TypesOf(values []Record) []Kind {
	out := make([]Kind, 0, len(values))
	for _, r := range values {
		out = append(out, r.Kind)
	}
	return out
}
func SameID(a, b Record) bool            { return a.ID == b.ID }
func IsEntity(r Record) bool             { return r.Kind == Entity }
func IsActivity(r Record) bool           { return r.Kind == Activity }
func IsAgent(r Record) bool              { return r.Kind == Agent }
func KindName(r Record) string           { return string(r.Kind) }
func RecordValid(r Record) bool          { return r.ID != "" && r.Type != "" }
func RelationCount(g *Graph) int         { return len(g.Relations) }
func HasRelations(g *Graph) bool         { return len(g.Relations) > 0 }
func HeadHash(g *Graph) string           { return g.head }
func GraphEmpty(g *Graph) bool           { return len(g.Records) == 0 }
func RecordCount(g *Graph) int           { return len(g.Records) }
func HasRecord(g *Graph, id string) bool { _, ok := g.Records[id]; return ok }
