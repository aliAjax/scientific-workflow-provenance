package repository

import (
	"errors"
	"sync"
	"time"
)

type Lock struct {
	ID        string
	Owner     string
	ExpiresAt time.Time
}
type LockTable struct {
	mu    sync.Mutex
	items map[string]Lock
}

func NewLocks() *LockTable { return &LockTable{items: map[string]Lock{}} }
func (l *LockTable) Acquire(name, owner string, ttl time.Duration) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	now := time.Now()
	if x, ok := l.items[name]; ok && x.ExpiresAt.After(now) && x.Owner != owner {
		return errors.New("lock held")
	}
	l.items[name] = Lock{ID: name, Owner: owner, ExpiresAt: now.Add(ttl)}
	return nil
}
func (l *LockTable) Release(name, owner string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if x, ok := l.items[name]; ok && x.Owner == owner {
		delete(l.items, name)
	}
}
