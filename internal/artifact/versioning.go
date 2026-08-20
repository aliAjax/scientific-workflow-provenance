package artifact

import (
	"sort"
	"strings"
)

type Catalog struct{ Items []Artifact }

func (c Catalog) BySchema(schema string) []Artifact {
	o := []Artifact{}
	for _, a := range c.Items {
		if strings.EqualFold(a.Schema, schema) {
			o = append(o, a)
		}
	}
	sort.Slice(o, func(i, j int) bool { return o[i].CreatedAt.Before(o[j].CreatedAt) })
	return o
}
func (c Catalog) Latest(name string) (Artifact, bool) {
	var out Artifact
	found := false
	for _, a := range c.Items {
		if a.Name == name && (!found || a.CreatedAt.After(out.CreatedAt)) {
			out = a
			found = true
		}
	}
	return out, found
}
