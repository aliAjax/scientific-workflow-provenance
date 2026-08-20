package repository

import (
	"errors"
	"sync"
)

type Tx struct {
	mu     sync.Mutex
	closed bool
	undos  []func()
}

func Begin() *Tx { return &Tx{} }
func (t *Tx) Do(apply, undo func()) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.closed {
		return errors.New("transaction closed")
	}
	apply()
	t.undos = append(t.undos, undo)
	return nil
}
func (t *Tx) Commit() { t.mu.Lock(); defer t.mu.Unlock(); t.closed = true; t.undos = nil }
func (t *Tx) Rollback() {
	t.mu.Lock()
	if t.closed {
		t.mu.Unlock()
		return
	}
	// Mark closed and detach the undo list before running callbacks. An undo
	// callback may itself call Do (e.g. nested cleanup); holding the lock across
	// the callbacks would deadlock that reentrant call, and leaving the list
	// mutable would let a concurrent Do append to a slice we are unwinding.
	undos := t.undos
	t.undos = nil
	t.closed = true
	t.mu.Unlock()
	for i := len(undos) - 1; i >= 0; i-- {
		undos[i]()
	}
}
