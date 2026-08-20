package artifact

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Status string

const (
	Published   Status = "published"
	Invalidated Status = "invalidated"
	Pending     Status = "pending"
)

type Artifact struct {
	Digest        string    `json:"digest"`
	Name          string    `json:"name"`
	Schema        string    `json:"schema"`
	Unit          string    `json:"unit"`
	Value         float64   `json:"value"`
	Status        Status    `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
	InvalidReason string    `json:"invalidReason,omitempty"`
}
type QualityGate struct {
	Min, Max      *float64
	RequiredUnit  string
	RequireSchema bool
}

var ErrInvalid = errors.New("artifact quality gate failed")

func Validate(a Artifact, g QualityGate) error {
	if a.Digest == "" {
		return ErrInvalid
	}
	if g.Min != nil && a.Value < *g.Min {
		return fmt.Errorf("value below minimum: %w", ErrInvalid)
	}
	if g.Max != nil && a.Value > *g.Max {
		return fmt.Errorf("value above maximum: %w", ErrInvalid)
	}
	if g.RequiredUnit != "" && a.Unit != g.RequiredUnit {
		return fmt.Errorf("unit mismatch: %w", ErrInvalid)
	}
	if g.RequireSchema && a.Schema == "" {
		return fmt.Errorf("schema required: %w", ErrInvalid)
	}
	return nil
}
func Digest(data []byte) string { h := sha256.Sum256(data); return hex.EncodeToString(h[:]) }

type Store struct {
	mu   sync.RWMutex
	data map[string]Artifact
}

func NewStore() *Store { return &Store{data: map[string]Artifact{}} }
func (s *Store) Put(a Artifact) error {
	if a.Digest == "" {
		return ErrInvalid
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if a.CreatedAt.IsZero() {
		a.CreatedAt = time.Now().UTC()
	}
	if a.Status == "" {
		a.Status = Published
	}
	s.data[a.Digest] = a
	return nil
}
func (s *Store) Get(d string) (Artifact, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	a, ok := s.data[d]
	return a, ok
}
func (s *Store) Invalidate(d, reason string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	a, ok := s.data[d]
	if !ok {
		return errors.New("artifact not found")
	}
	a.Status = Invalidated
	a.InvalidReason = reason
	s.data[d] = a
	return nil
}
func (s *Store) List() []Artifact {
	s.mu.RLock()
	defer s.mu.RUnlock()
	o := make([]Artifact, 0, len(s.data))
	for _, a := range s.data {
		o = append(o, a)
	}
	return o
}
