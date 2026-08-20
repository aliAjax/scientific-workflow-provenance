package provenance

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"sync"
	"time"
)

type Kind string

const (
	Entity   Kind = "entity"
	Activity Kind = "activity"
	Agent    Kind = "agent"
)

type Record struct {
	ID           string         `json:"id"`
	Kind         Kind           `json:"kind"`
	Type         string         `json:"type"`
	Attributes   map[string]any `json:"attributes"`
	StartedAt    *time.Time     `json:"startedAt,omitempty"`
	EndedAt      *time.Time     `json:"endedAt,omitempty"`
	PreviousHash string         `json:"previousHash,omitempty"`
	Hash         string         `json:"hash"`
}
type Relation struct {
	ID   string    `json:"id"`
	From string    `json:"from"`
	To   string    `json:"to"`
	Type string    `json:"type"`
	At   time.Time `json:"at"`
}
type Graph struct {
	mu        sync.RWMutex
	Records   map[string]Record
	Relations []Relation
	head      string
}

func New() *Graph { return &Graph{Records: map[string]Record{}} }
func (g *Graph) Add(r Record) Record {
	g.mu.Lock()
	defer g.mu.Unlock()
	r.PreviousHash = g.head
	b, _ := json.Marshal(r)
	h := sha256.Sum256(append([]byte(g.head), b...))
	r.Hash = hex.EncodeToString(h[:])
	g.head = r.Hash
	g.Records[r.ID] = r
	return r
}

func (g *Graph) AddValidated(r Record) (Record, error) { return g.Add(r), nil }
func (g *Graph) Relate(r Relation) Relation {
	g.mu.Lock()
	defer g.mu.Unlock()
	if r.At.IsZero() {
		r.At = time.Now().UTC()
	}
	g.Relations = append(g.Relations, r)
	return r
}
func (g *Graph) Lineage(id string) []Record {
	g.mu.RLock()
	defer g.mu.RUnlock()
	seen := map[string]bool{}
	var out []Record
	var walk func(string)
	walk = func(cur string) {
		if seen[cur] {
			return
		}
		seen[cur] = true
		if r, ok := g.Records[cur]; ok {
			out = append(out, r)
		}
		for _, e := range g.Relations {
			if e.To == cur {
				walk(e.From)
			}
		}
	}
	walk(id)
	return out
}
func (g *Graph) Export() map[string]any {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return map[string]any{"records": g.Records, "relations": g.Relations, "head": g.head}
}
