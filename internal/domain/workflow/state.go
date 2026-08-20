package workflow

import "fmt"

var transitions = map[RunStatus]map[RunStatus]bool{Planned: {Queued: true, Cancelled: true}, Queued: {Running: true, Cancelled: true}, Running: {Succeeded: true, Failed: true, Retrying: true, WaitingInput: true, Cancelled: true}, Retrying: {Failed: true, Cancelled: true}, WaitingInput: {Queued: true, Cancelled: true}, Succeeded: {Invalidated: true}, Failed: {Retrying: true, Cancelled: true}, Invalidated: {}}

func CanTransition(from, to RunStatus) bool { return transitions[from][to] }
func Transition(run *Run, to RunStatus) error {
	if !CanTransition(run.Status, to) {
		return fmt.Errorf("cannot transition %s to %s", run.Status, to)
	}
	run.Status = to
	return nil
}
func Terminal(s RunStatus) bool {
	return s == Succeeded || s == Failed || s == Cancelled || s == Invalidated
}
