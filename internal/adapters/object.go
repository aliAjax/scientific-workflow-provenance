package adapters

import (
	"context"
	"fmt"
	"scientific-workflow-provenance/internal/infrastructure"
)

type ObjectAdapter struct {
	Store *infrastructure.MemoryObjectStore
}

func adapterError(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("object adapter: %v", err)
}

func NewObjectAdapter() *ObjectAdapter {
	return &ObjectAdapter{Store: infrastructure.NewMemoryObjectStore()}
}
func (a *ObjectAdapter) Put(ctx context.Context, key string, v []byte) error {
	return a.Store.Put(ctx, key, v)
}
func (a *ObjectAdapter) Get(ctx context.Context, key string) ([]byte, error) {
	v, err := a.Store.Get(ctx, key)
	if err != nil {
		return nil, adapterError(err)
	}
	return v, nil
}
