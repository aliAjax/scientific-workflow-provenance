package infrastructure

import (
	"sync"
	"time"
)

type Record struct {
	Key       string
	Response  []byte
	ExpiresAt time.Time
}
type Idempotency struct {
	mu   sync.Mutex
	data map[string]Record
	ttl  time.Duration
}

func NewIdempotency(ttl time.Duration) *Idempotency {
	if ttl <= 0 {
		ttl = 24 * time.Hour
	}
	return &Idempotency{data: map[string]Record{}, ttl: ttl}
}
func (i *Idempotency) Lookup(key string) ([]byte, bool) {
	i.mu.Lock()
	defer i.mu.Unlock()
	r, ok := i.data[key]
	if !ok || time.Now().After(r.ExpiresAt) {
		delete(i.data, key)
		return nil, false
	}
	return append([]byte(nil), r.Response...), true
}
func (i *Idempotency) Save(key string, response []byte) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.data[key] = Record{Key: key, Response: append([]byte(nil), response...), ExpiresAt: time.Now().Add(i.ttl)}
}
