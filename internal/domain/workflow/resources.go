package workflow

import (
	"fmt"
	"sort"
)

type ResourcePool struct {
	ID       string
	Capacity ResourceRequest
	Used     ResourceRequest
}

func (p *ResourcePool) CanReserve(r ResourceRequest) bool {
	return p.Used.CPU+r.CPU <= p.Capacity.CPU && p.Used.MemoryMB+r.MemoryMB <= p.Capacity.MemoryMB
}
func (p *ResourcePool) Reserve(r ResourceRequest) error {
	if !p.CanReserve(r) {
		return fmt.Errorf("resource pool %s exhausted", p.ID)
	}
	p.Used.CPU += r.CPU
	p.Used.MemoryMB += r.MemoryMB
	return nil
}
func (p *ResourcePool) Release(r ResourceRequest) {
	p.Used.CPU -= r.CPU
	p.Used.MemoryMB -= r.MemoryMB
	if p.Used.CPU < 0 {
		p.Used.CPU = 0
	}
	if p.Used.MemoryMB < 0 {
		p.Used.MemoryMB = 0
	}
}
func SortPools(p []ResourcePool) { sort.Slice(p, func(i, j int) bool { return p[i].ID < p[j].ID }) }
