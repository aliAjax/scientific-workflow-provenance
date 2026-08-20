package workflow

import (
	"sort"
	"strings"
)

type Template struct {
	ID          string
	Name        string
	Description string
	Defaults    map[string]string
}

func (t Template) Apply(params map[string]string) map[string]string {
	out := map[string]string{}
	for k, v := range t.Defaults {
		out[k] = v
	}
	for k, v := range params {
		out[k] = v
	}
	return out
}
func (t Template) Keywords() []string {
	set := map[string]bool{}
	for _, v := range strings.Fields(strings.ToLower(t.Name + " " + t.Description)) {
		if len(v) > 2 {
			set[v] = true
		}
	}
	out := []string{}
	for k := range set {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
