package security

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"
)

type NonceStore struct {
	mu    sync.Mutex
	items map[string]time.Time
}

func NewNonceStore() *NonceStore { return &NonceStore{items: map[string]time.Time{}} }
func (n *NonceStore) Issue(ttl time.Duration) string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	v := hex.EncodeToString(b)
	n.mu.Lock()
	n.items[v] = time.Now().Add(ttl)
	n.mu.Unlock()
	return v
}
func (n *NonceStore) Consume(v string) bool {
	n.mu.Lock()
	defer n.mu.Unlock()
	t, ok := n.items[v]
	if !ok || time.Now().After(t) {
		delete(n.items, v)
		return false
	}
	delete(n.items, v)
	return true
}
