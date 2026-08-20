package adapters

import (
	"context"
	"scientific-workflow-provenance/internal/infrastructure"
)

type ObjectAdapter struct {
	Store *infrastructure.MemoryObjectStore
}

func NewObjectAdapter() *ObjectAdapter {
	return &ObjectAdapter{Store: infrastructure.NewMemoryObjectStore()}
}
func (a *ObjectAdapter) Put(ctx context.Context, key string, v []byte) error {
	return a.Store.Put(ctx, key, v)
}
func (a *ObjectAdapter) Get(ctx context.Context, key string) ([]byte, error) {
	return a.Store.Get(ctx, key)
}
