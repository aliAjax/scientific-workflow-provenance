package lifecycle

import (
	"fmt"
	"sync"
	"time"
)

type State string

const (
	Queued    State = "queued"
	Running   State = "running"
	Succeeded State = "succeeded"
	Failed    State = "failed"
	Cancelled State = "cancelled"
)

var allowed = map[State]map[State]bool{Queued: {Running: true, Cancelled: true}, Running: {Succeeded: true, Failed: true, Cancelled: true}, Failed: {Queued: true}, Succeeded: {}, Cancelled: {}}

type Machine struct {
	mu      sync.Mutex
	state   State
	history []State
}

func New(initial State) *Machine {
	if initial == "" {
		initial = Queued
	}
	return &Machine{state: initial, history: []State{initial}}
}
func (m *Machine) Transition(to State) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if !allowed[m.state][to] {
		return fmt.Errorf("invalid transition %s -> %s", m.state, to)
	}
	m.state = to
	m.history = append(m.history, to)
	return nil
}
func (m *Machine) State() State { m.mu.Lock(); defer m.mu.Unlock(); return m.state }
func (m *Machine) History() []State {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]State(nil), m.history...)
}

type Marker struct {
	State State
	At    time.Time
}
