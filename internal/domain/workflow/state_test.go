package workflow

import "testing"

func TestRetryTransitionReachesSuccess(t *testing.T) {
	r := Run{Status: Failed}
	for _, next := range []RunStatus{Retrying, Queued, Running, Succeeded} {
		if err := Transition(&r, next); err != nil {
			t.Fatalf("%s: %v", next, err)
		}
	}
}
