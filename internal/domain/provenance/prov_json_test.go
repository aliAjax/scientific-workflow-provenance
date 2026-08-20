package provenance

import (
	"errors"
	"testing"
)

func TestValidateChainPreservesSentinel(t *testing.T) {
	if !errors.Is(ValidateChain(map[string]any{}), ErrInvalidChain) {
		t.Fatal("chain sentinel lost")
	}
}
