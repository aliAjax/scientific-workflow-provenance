package api

import (
	"errors"
	"testing"
)

func TestAPIErrorClassification(t *testing.T) {
	sentinel := errors.New("not found")
	if !errors.Is(WrapCause(sentinel), sentinel) {
		t.Fatal("api wrapper lost cause")
	}
}
