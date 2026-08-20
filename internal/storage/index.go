package storage

import (
	"sort"
	"sync"
	"time"
)

type Record struct {
	Key   string
	Value []byte
	At    time.Time
}
type Index struct {
	mu       sync.RWMutex
	byKey    map[string]Record
	byPrefix map[string][]string
}

func NewIndex() *Index { return &Index{byKey: map[string]Record{}, byPrefix: map[string][]string{}} }
func (i *Index) Put(key string, value []byte) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.byKey[key] = Record{key, append([]byte(nil), value...), time.Now().UTC()}
	p := prefix(key)
	found := false
	for _, k := range i.byPrefix[p] {
		if k == key {
			found = true
		}
	}
	if !found {
		i.byPrefix[p] = append(i.byPrefix[p], key)
	}
}
func (i *Index) Get(key string) (Record, bool) {
	i.mu.RLock()
	defer i.mu.RUnlock()
	v, ok := i.byKey[key]
	v.Value = append([]byte(nil), v.Value...)
	return v, ok
}
func (i *Index) Keys(prefix string) []string {
	i.mu.RLock()
	defer i.mu.RUnlock()
	o := append([]string(nil), i.byPrefix[prefix]...)
	sort.Strings(o)
	return o
}
func prefix(k string) string {
	for n, c := range k {
		if c == '/' || c == ':' || c == '-' {
			return k[:n]
		}
	}
	return k
}
