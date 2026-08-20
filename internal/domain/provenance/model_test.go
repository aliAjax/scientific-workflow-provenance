package provenance

import (
	"errors"
	"testing"
)

func TestGraphAddValidatesRecordCause(t *testing.T) {
	_, err := New().AddValidated(Record{Kind: Entity})
	if err == nil || !errors.Is(err, ErrInvalidChain) {
		t.Fatalf("invalid record cause missing: %v", err)
	}
}
