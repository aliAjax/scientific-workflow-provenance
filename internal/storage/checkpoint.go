package storage

import (
	"encoding/json"
	"errors"
	"sync"
	"time"
)

type Checkpoint struct {
	RunID    string            `json:"runId"`
	Sequence int               `json:"sequence"`
	State    map[string]string `json:"state"`
	At       time.Time         `json:"at"`
}
type Checkpoints struct {
	mu    sync.RWMutex
	items map[string][]Checkpoint
}

func NewCheckpoints() *Checkpoints { return &Checkpoints{items: map[string][]Checkpoint{}} }
func (c *Checkpoints) Save(v Checkpoint) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v.At.IsZero() {
		v.At = time.Now().UTC()
	}
	c.items[v.RunID] = append(c.items[v.RunID], v)
}
func (c *Checkpoints) Latest(run string) (Checkpoint, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v := c.items[run]
	if len(v) == 0 {
		return Checkpoint{}, errors.New("checkpoint not found")
	}
	return v[len(v)-1], nil
}
func (c *Checkpoints) Encode(run string) ([]byte, error) {
	v, e := c.Latest(run)
	if e != nil {
		return nil, e
	}
	return json.Marshal(v)
}
