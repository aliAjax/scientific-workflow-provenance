package security

import (
	"sync"
	"time"
)

type Bucket struct {
	mu               sync.Mutex
	capacity, tokens int
	refill           time.Duration
	last             time.Time
}

func NewBucket(capacity int, refill time.Duration) *Bucket {
	if capacity < 1 {
		capacity = 1
	}
	return &Bucket{capacity: capacity, tokens: capacity, refill: refill, last: time.Now()}
}
func (b *Bucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	now := time.Now()
	if b.refill > 0 {
		n := int(now.Sub(b.last) / b.refill)
		if n > 0 {
			b.tokens += n
			if b.tokens > b.capacity {
				b.tokens = b.capacity
			}
			b.last = now
		}
	}
	if b.tokens < 1 {
		return false
	}
	b.tokens--
	return true
}
