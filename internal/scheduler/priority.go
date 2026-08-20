package scheduler

import (
	"container/heap"
	"time"
)

type PriorityItem struct {
	ID       string
	Priority int
	At       time.Time
	index    int
}
type PriorityQueue []*PriorityItem

func (p PriorityQueue) Len() int { return len(p) }
func (p PriorityQueue) Less(i, j int) bool {
	if p[i].Priority == p[j].Priority {
		return p[i].At.Before(p[j].At)
	}
	return p[i].Priority > p[j].Priority
}
func (p PriorityQueue) Swap(i, j int) { p[i], p[j] = p[j], p[i]; p[i].index = i; p[j].index = j }
func (p *PriorityQueue) Push(x any)   { v := x.(*PriorityItem); v.index = len(*p); *p = append(*p, v) }
func (p *PriorityQueue) Pop() any {
	old := *p
	n := len(old)
	v := old[n-1]
	*p = old[:n-1]
	v.index = -1
	return v
}
func NewPriority() *PriorityQueue { p := PriorityQueue{}; heap.Init(&p); return &p }
