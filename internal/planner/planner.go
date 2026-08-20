package planner

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"scientific-workflow-provenance/internal/domain/workflow"
	"sort"
	"strings"
)

type Step struct {
	ID        string            `json:"id"`
	NodeID    string            `json:"nodeId"`
	Matrix    map[string]string `json:"matrix,omitempty"`
	DependsOn []string          `json:"dependsOn"`
	CacheKey  string            `json:"cacheKey"`
}
type Plan struct {
	WorkflowID string `json:"workflowId"`
	Version    int    `json:"version"`
	Steps      []Step `json:"steps"`
	Digest     string `json:"digest"`
}
type Planner struct{}

func (Planner) Build(d workflow.Definition, params map[string]string) (Plan, error) {
	if err := d.Validate(); err != nil {
		return Plan{}, err
	}
	indeg := map[string]int{}
	next := map[string][]string{}
	for _, n := range d.Nodes {
		indeg[n.ID] = 0
	}
	for _, e := range d.Edges {
		indeg[e.To]++
		next[e.From] = append(next[e.From], e.To)
	}
	ready := []string{}
	for id, v := range indeg {
		if v == 0 {
			ready = append(ready, id)
		}
	}
	sort.Strings(ready)
	order := []string{}
	for len(ready) > 0 {
		id := ready[0]
		ready = ready[1:]
		order = append(order, id)
		for _, to := range next[id] {
			indeg[to]--
			if indeg[to] == 0 {
				ready = append(ready, to)
				sort.Strings(ready)
			}
		}
	}
	steps := []Step{}
	for _, id := range order {
		n, _ := d.Node(id)
		matrices := n.Matrix
		if len(matrices) == 0 {
			matrices = []map[string]string{nil}
		}
		for i, m := range matrices {
			sid := fmt.Sprintf("%s-%d", id, i)
			deps := []string{}
			for _, e := range d.Edges {
				if e.To == id {
					deps = append(deps, e.From)
				}
			}
			sort.Strings(deps)
			steps = append(steps, Step{ID: sid, NodeID: id, Matrix: m, DependsOn: deps, CacheKey: cacheKey(d, id, m, params)})
		}
	}
	b := []byte{}
	for _, s := range steps {
		b = append(b, []byte(s.ID+s.CacheKey)...)
	}
	h := sha256.Sum256(b)
	return Plan{WorkflowID: d.ID, Version: d.Version, Steps: steps, Digest: hex.EncodeToString(h[:])}, nil
}
func cacheKey(d workflow.Definition, id string, m map[string]string, p map[string]string) string {
	keys := []string{}
	for k, v := range m {
		keys = append(keys, k+"="+v)
	}
	for k, v := range p {
		keys = append(keys, k+"="+v)
	}
	sort.Strings(keys)
	h := sha256.Sum256([]byte(d.Digest + "|" + id + "|" + strings.Join(keys, ";")))
	return hex.EncodeToString(h[:])
}
func (p Plan) Ready(done map[string]bool) []Step {
	out := []Step{}
	for _, s := range p.Steps {
		if done[s.ID] {
			continue
		}
		ok := true
		for _, d := range s.DependsOn {
			found := false
			for _, x := range p.Steps {
				if x.NodeID == d && done[x.ID] {
					found = true
				}
			}
			if !found {
				ok = false
			}
		}
		if ok {
			out = append(out, s)
		}
	}
	return out
}
