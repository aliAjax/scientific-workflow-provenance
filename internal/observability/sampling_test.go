package observability

import "testing"

func TestSamplerZeroValueDoesNotPanic(t *testing.T) {
	var s *Sampler
	if s.Keep() {
		t.Fatal("nil sampler should not keep")
	}
	if s.Seen() != 0 {
		t.Fatal("nil sampler count changed")
	}
}
