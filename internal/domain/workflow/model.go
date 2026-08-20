package workflow

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

type RunStatus string

const (
	Planned      RunStatus = "planned"
	Queued       RunStatus = "queued"
	Running      RunStatus = "running"
	WaitingInput RunStatus = "waiting_input"
	Retrying     RunStatus = "retrying"
	Succeeded    RunStatus = "succeeded"
	Failed       RunStatus = "failed"
	Cancelled    RunStatus = "cancelled"
	Invalidated  RunStatus = "invalidated"
)

type NodeKind string

const (
	Task      NodeKind = "task"
	Condition NodeKind = "condition"
	Matrix    NodeKind = "matrix"
	Aggregate NodeKind = "aggregate"
	Approval  NodeKind = "approval"
)

type RetryPolicy struct {
	MaxAttempts int           `json:"maxAttempts"`
	Backoff     time.Duration `json:"backoff"`
	Retryable   []string      `json:"retryable"`
}
type ResourceRequest struct {
	CPU      int    `json:"cpu"`
	MemoryMB int    `json:"memoryMb"`
	Pool     string `json:"pool"`
}
type Node struct {
	ID        string              `json:"id"`
	Name      string              `json:"name"`
	Kind      NodeKind            `json:"kind"`
	Inputs    []string            `json:"inputs"`
	Outputs   []string            `json:"outputs"`
	Condition string              `json:"condition,omitempty"`
	Matrix    []map[string]string `json:"matrix,omitempty"`
	Retry     RetryPolicy         `json:"retry"`
	Timeout   time.Duration       `json:"timeout"`
	Resources ResourceRequest     `json:"resources"`
	Cacheable bool                `json:"cacheable"`
}
type Edge struct {
	From string `json:"from"`
	To   string `json:"to"`
	When string `json:"when,omitempty"`
}
type Definition struct {
	ID         string            `json:"id"`
	Name       string            `json:"name"`
	Version    int               `json:"version"`
	Nodes      []Node            `json:"nodes"`
	Edges      []Edge            `json:"edges"`
	Parameters map[string]string `json:"parameters"`
	CreatedAt  time.Time         `json:"createdAt"`
	Digest     string            `json:"digest"`
}
type Run struct {
	ID         string               `json:"id"`
	WorkflowID string               `json:"workflowId"`
	Version    int                  `json:"version"`
	Status     RunStatus            `json:"status"`
	NodeStates map[string]RunStatus `json:"nodeStates"`
	Attempts   map[string]int       `json:"attempts"`
	Parameters map[string]string    `json:"parameters"`
	CreatedAt  time.Time            `json:"createdAt"`
	UpdatedAt  time.Time            `json:"updatedAt"`
	Error      string               `json:"error,omitempty"`
}

var ErrCycle = errors.New("workflow contains a cycle")
var ErrInvalid = errors.New("invalid workflow definition")

func (d *Definition) Validate() error {
	if d.ID == "" || d.Name == "" || d.Version < 1 || len(d.Nodes) == 0 {
		return ErrInvalid
	}
	ids := make(map[string]bool)
	for _, n := range d.Nodes {
		if n.ID == "" || ids[n.ID] {
			return fmt.Errorf("duplicate node: %s", n.ID)
		}
		ids[n.ID] = true
		if n.Retry.MaxAttempts < 0 {
			return ErrInvalid
		}
	}
	indegree := make(map[string]int)
	next := make(map[string][]string)
	for _, e := range d.Edges {
		if !ids[e.From] || !ids[e.To] || e.From == e.To {
			return ErrInvalid
		}
		indegree[e.To]++
		next[e.From] = append(next[e.From], e.To)
	}
	q := make([]string, 0)
	for id := range ids {
		if indegree[id] == 0 {
			q = append(q, id)
		}
	}
	seen := 0
	for len(q) > 0 {
		id := q[0]
		q = q[1:]
		seen++
		for _, to := range next[id] {
			indegree[to]--
			if indegree[to] == 0 {
				q = append(q, to)
			}
		}
	}
	if seen != len(ids) {
		return ErrCycle
	}
	return nil
}
func (d *Definition) ComputeDigest() string {
	h := sha256.New()
	ids := make([]string, len(d.Nodes))
	for i, n := range d.Nodes {
		ids[i] = n.ID + ":" + string(n.Kind)
	}
	sort.Strings(ids)
	h.Write([]byte(strings.Join(ids, "|")))
	es := make([]string, len(d.Edges))
	for i, e := range d.Edges {
		es[i] = e.From + ">" + e.To + ":" + e.When
	}
	sort.Strings(es)
	h.Write([]byte(strings.Join(es, "|")))
	return hex.EncodeToString(h.Sum(nil))
}
func (d *Definition) Node(id string) (Node, bool) {
	for _, n := range d.Nodes {
		if n.ID == id {
			return n, true
		}
	}
	return Node{}, false
}
