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
	defer t.mu.Unlock()
	if t.closed {
		return
	}
	for i := len(t.undos) - 1; i >= 0; i-- {
		t.undos[i]()
	}
	t.closed = true
}
