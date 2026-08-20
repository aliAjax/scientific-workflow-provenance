package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"sync"
	"time"
)

type Segment struct {
	ID        string
	Sequence  uint64
	Payload   []byte
	Digest    string
	CreatedAt time.Time
	Uploaded  bool
}
type SegmentLog struct {
	mu    sync.Mutex
	items []Segment
	max   int
}

func NewSegmentLog(max int) *SegmentLog {
	if max < 1 {
		max = 100
	}
	return &SegmentLog{max: max}
}
func (l *SegmentLog) Append(id string, payload []byte) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if len(l.items) >= l.max {
		return errors.New("segment capacity exceeded")
	}
	h := sha256.Sum256(payload)
	l.items = append(l.items, Segment{ID: id, Sequence: uint64(len(l.items) + 1), Payload: append([]byte(nil), payload...), Digest: hex.EncodeToString(h[:]), CreatedAt: time.Now().UTC()})
	return nil
}
func (l *SegmentLog) Pending() []Segment {
	l.mu.Lock()
	defer l.mu.Unlock()
	o := []Segment{}
	for _, s := range l.items {
		if !s.Uploaded {
			o = append(o, s)
		}
	}
	return o
}
func (l *SegmentLog) MarkUploaded(seq uint64) {
	l.mu.Lock()
	defer l.mu.Unlock()
	for i := range l.items {
		if l.items[i].Sequence == seq {
			l.items[i].Uploaded = true
		}
	}
}
