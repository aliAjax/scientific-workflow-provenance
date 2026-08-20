package sample

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type State string

const (
	Active    State = "active"
	Frozen    State = "frozen"
	Destroyed State = "destroyed"
)

type Quality struct {
	Metric    string  `json:"metric"`
	Value     float64 `json:"value"`
	Unit      string  `json:"unit"`
	Threshold float64 `json:"threshold"`
	Passed    bool    `json:"passed"`
}
type Sample struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Batch     string     `json:"batch"`
	Source    string     `json:"source"`
	ParentIDs []string   `json:"parentIds"`
	State     State      `json:"state"`
	Qualities []Quality  `json:"qualities"`
	CreatedAt time.Time  `json:"createdAt"`
	FrozenAt  *time.Time `json:"frozenAt,omitempty"`
}

var ErrNotFound = errors.New("sample not found")
var ErrTransition = errors.New("invalid sample transition")

func cloneSample(v Sample) Sample {
	if v.ParentIDs != nil {
		p := make([]string, len(v.ParentIDs))
		copy(p, v.ParentIDs)
		v.ParentIDs = p
	}
	if v.Qualities != nil {
		q := make([]Quality, len(v.Qualities))
		copy(q, v.Qualities)
		v.Qualities = q
	}
	return v
}

type Store struct {
	mu   sync.RWMutex
	data map[string]Sample
}

func NewStore() *Store { return &Store{data: map[string]Sample{}} }
func (s *Store) Put(v Sample) error {
	if v.ID == "" {
		return fmt.Errorf("sample id required")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if v.State == "" {
		v.State = Active
	}
	if v.CreatedAt.IsZero() {
		v.CreatedAt = time.Now().UTC()
	}
	s.data[v.ID] = cloneSample(v)
	return nil
}
func (s *Store) Get(id string) (Sample, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.data[id]
	if !ok {
		return Sample{}, ErrNotFound
	}
	return cloneSample(v), nil
}
func (s *Store) List() []Sample {
	s.mu.RLock()
	defer s.mu.RUnlock()
	out := make([]Sample, 0, len(s.data))
	for _, v := range s.data {
		out = append(out, v)
	}
	return out
}
func (s *Store) Freeze(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	v, ok := s.data[id]
	if !ok {
		return ErrNotFound
	}
	if v.State != Active {
		return ErrTransition
	}
	now := time.Now().UTC()
	v.State = Frozen
	v.FrozenAt = &now
	s.data[id] = v
	return nil
}
func (s *Store) Destroy(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	v, ok := s.data[id]
	if !ok {
		return ErrNotFound
	}
	if v.State == Destroyed {
		return ErrTransition
	}
	v.State = Destroyed
	s.data[id] = v
	return nil
}
