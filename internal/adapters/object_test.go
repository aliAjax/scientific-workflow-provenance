package adapters

import (
	"context"
	"errors"
	"scientific-workflow-provenance/internal/infrastructure"
	"testing"
)

func TestObjectAdapterPreservesCause(t *testing.T) {
	_, err := NewObjectAdapter().Get(context.Background(), "missing")
	if !errors.Is(err, infrastructure.ErrObjectNotFound) {
		t.Fatalf("adapter lost cause: %v", err)
	}
}
