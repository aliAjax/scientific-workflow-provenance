package workflow

import (
	"sort"
	"strings"
)

type Parameter struct {
	Name     string
	Type     string
	Required bool
	Default  string
}

func NormalizeParams(values map[string]string) map[string]string {
	out := map[string]string{}
	for k, v := range values {
		out[strings.TrimSpace(k)] = strings.TrimSpace(v)
	}
	return out
}
func ParameterNames(v []Parameter) []string {
	out := []string{}
	for _, p := range v {
		out = append(out, p.Name)
	}
	sort.Strings(out)
	return out
}
func RequiredMissing(def []Parameter, values map[string]string) []string {
	out := []string{}
	for _, p := range def {
		if p.Required && strings.TrimSpace(values[p.Name]) == "" {
			out = append(out, p.Name)
		}
	}
	sort.Strings(out)
	return out
}
