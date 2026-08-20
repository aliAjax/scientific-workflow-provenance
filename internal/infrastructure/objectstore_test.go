package infrastructure

import (
	"context"
	"errors"
	"testing"
)

func TestObjectStoreNotFoundChain(t *testing.T) {
	_, err := NewMemoryObjectStore().Get(context.Background(), "missing")
	if !errors.Is(err, ErrObjectNotFound) {
		t.Fatalf("missing cause: %v", err)
	}
}
