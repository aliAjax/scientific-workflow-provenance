package execution

import (
	"sort"
	"time"
)

type StateEvent struct {
	Sequence int
	RunID    string
	StepID   string
	State    string
	At       time.Time
	Message  string
}
type EventLog struct{ items []StateEvent }

func (l *EventLog) Append(e StateEvent) {
	if e.At.IsZero() {
		e.At = time.Now().UTC()
	}
	e.Sequence = len(l.items) + 1
	l.items = append(l.items, e)
}
func (l EventLog) List() []StateEvent { return append([]StateEvent(nil), l.items...) }
func (l EventLog) ForStep(id string) []StateEvent {
	o := []StateEvent{}
	for _, e := range l.items {
		if e.StepID == id {
			o = append(o, e)
		}
	}
	sort.Slice(o, func(i, j int) bool { return o[i].Sequence < o[j].Sequence })
	return o
}
