package execution

import (
	"encoding/json"
	"errors"
	"sync"
	"time"
)

type Snapshot struct {
	RunID    string            `json:"runId"`
	States   map[string]string `json:"states"`
	At       time.Time         `json:"at"`
	Sequence int               `json:"sequence"`
}
type CheckpointStore struct {
	mu   sync.RWMutex
	data map[string][]Snapshot
}

func NewCheckpointStore() *CheckpointStore { return &CheckpointStore{data: map[string][]Snapshot{}} }
func (s *CheckpointStore) Save(v Snapshot) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if v.At.IsZero() {
		v.At = time.Now().UTC()
	}
	v.Sequence = len(s.data[v.RunID]) + 1
	s.data[v.RunID] = append(s.data[v.RunID], v)
}
func (s *CheckpointStore) Latest(run string) (Snapshot, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v := s.data[run]
	if len(v) == 0 {
		return Snapshot{}, errors.New("no checkpoint")
	}
	out := v[len(v)-1]
	if out.States != nil {
		cp := make(map[string]string, len(out.States))
		for k, val := range out.States {
			cp[k] = val
		}
		out.States = cp
	}
	return out, nil
}
func (s *CheckpointStore) JSON(run string) ([]byte, error) {
	v, e := s.Latest(run)
	if e != nil {
		return nil, e
	}
	return json.Marshal(v)
}
