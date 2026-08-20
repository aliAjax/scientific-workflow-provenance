package sample

import (
	"sort"
	"time"
)

type Link struct {
	Parent, Child string
	Relation      string
	At            time.Time
}

func Ancestors(root string, all []Sample) []Sample {
	byID := map[string]Sample{}
	for _, s := range all {
		byID[s.ID] = s
	}
	seen := map[string]bool{}
	out := []Sample{}
	var walk func(string)
	walk = func(id string) {
		for _, s := range byID[id].ParentIDs {
			if !seen[s] {
				seen[s] = true
				if p, ok := byID[s]; ok {
					out = append(out, cloneSample(p))
					walk(s)
				}
			}
		}
	}
	walk(root)
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}
func Descendants(root string, all []Sample) []Sample {
	children := map[string][]string{}
	for _, s := range all {
		for _, p := range s.ParentIDs {
			children[p] = append(children[p], s.ID)
		}
	}
	seen := map[string]bool{}
	out := []Sample{}
	var walk func(string)
	walk = func(id string) {
		for _, c := range children[id] {
			if !seen[c] {
				seen[c] = true
				for _, s := range all {
					if s.ID == c {
						out = append(out, s)
					}
				}
				walk(c)
			}
		}
	}
	walk(root)
	return out
}
