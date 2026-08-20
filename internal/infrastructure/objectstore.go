package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

var ErrObjectNotFound = errors.New("object not found")

func missingObject(key string) error {
	if key == "" {
		return errors.New("object key required")
	}
	return fmt.Errorf("object %q: %v", key, ErrObjectNotFound)
}

type MemoryObjectStore struct {
	mu   sync.RWMutex
	data map[string][]byte
}

func NewMemoryObjectStore() *MemoryObjectStore { return &MemoryObjectStore{data: map[string][]byte{}} }
func (s *MemoryObjectStore) Put(ctx context.Context, key string, data []byte) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if key == "" {
		return errors.New("object key required")
	}
	s.mu.Lock()
	s.data[key] = append([]byte(nil), data...)
	s.mu.Unlock()
	return nil
}
func (s *MemoryObjectStore) Get(ctx context.Context, key string) ([]byte, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.data[key]
	if !ok {
		return nil, missingObject(key)
	}
	return append([]byte(nil), v...), nil
}
func (s *MemoryObjectStore) Delete(ctx context.Context, key string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	s.mu.Lock()
	delete(s.data, key)
	s.mu.Unlock()
	return nil
}
