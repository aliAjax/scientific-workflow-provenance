package audit

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"sync"
	"time"
)

type Entry struct {
	Actor, Action, Subject string
	At                     time.Time
	Previous, Hash         string
	Metadata               map[string]string
}
type Chain struct {
	mu      sync.Mutex
	entries []Entry
	head    string
}

func New() *Chain { return &Chain{} }
func (c *Chain) Append(actor, action, subject string, m map[string]string) Entry {
	c.mu.Lock()
	defer c.mu.Unlock()
	e := Entry{Actor: actor, Action: action, Subject: subject, At: time.Now().UTC(), Previous: c.head, Metadata: m}
	b, _ := json.Marshal(e)
	h := sha256.Sum256(append([]byte(c.head), b...))
	e.Hash = hex.EncodeToString(h[:])
	c.head = e.Hash
	c.entries = append(c.entries, e)
	return e
}
func (c *Chain) Verify() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	prev := ""
	for _, e := range c.entries {
		if e.Previous != prev {
			return false
		}
		b, _ := json.Marshal(Entry{Actor: e.Actor, Action: e.Action, Subject: e.Subject, At: e.At, Previous: e.Previous, Metadata: e.Metadata})
		h := sha256.Sum256(append([]byte(prev), b...))
		if hex.EncodeToString(h[:]) != e.Hash {
			return false
		}
		prev = e.Hash
	}
	return true
}
func (c *Chain) List() []Entry {
	c.mu.Lock()
	defer c.mu.Unlock()
	out := make([]Entry, len(c.entries))
	for i, e := range c.entries {
		out[i] = e
		if e.Metadata != nil {
			cp := make(map[string]string, len(e.Metadata))
			for k, v := range e.Metadata {
				cp[k] = v
			}
			out[i].Metadata = cp
		}
	}
	return out
}
