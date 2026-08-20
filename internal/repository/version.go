package repository

import (
	"errors"
	"scientific-workflow-provenance/internal/domain/workflow"
	"sync"
)

type Versioned[T any] struct {
	mu     sync.RWMutex
	values map[string][]T
}

func NewVersioned[T any]() *Versioned[T] { return &Versioned[T]{values: map[string][]T{}} }
func (v *Versioned[T]) Put(id string, value T) {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.values[id] = append(v.values[id], value)
}
func (v *Versioned[T]) Latest(id string) (T, error) {
	v.mu.RLock()
	defer v.mu.RUnlock()
	x := v.values[id]
	if len(x) == 0 {
		var z T
		return z, errors.New("version not found")
	}
	return x[len(x)-1], nil
}
func (v *Versioned[T]) All(id string) []T {
	v.mu.RLock()
	defer v.mu.RUnlock()
	return append([]T(nil), v.values[id]...)
}
func DefinitionCompatible(a, b workflow.Definition) bool {
	return a.ID == b.ID && b.Version >= a.Version
}
