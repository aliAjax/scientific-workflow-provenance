package observability

import (
	"math/rand"
	"sync/atomic"
)

type Sampler struct {
	rate float64
	seen uint64
}

func NewSampler(rate float64) *Sampler {
	if rate < 0 {
		rate = 0
	}
	if rate > 1 {
		rate = 1
	}
	return &Sampler{rate: rate}
}
func (s *Sampler) Keep() bool   { atomic.AddUint64(&s.seen, 1); return rand.Float64() < s.rate }
func (s *Sampler) Seen() uint64 { return atomic.LoadUint64(&s.seen) }
