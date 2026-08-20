package worker

import (
	"errors"
	"sync"
)

type Limits struct {
	MaxCPUSeconds  int64
	MaxMemoryMB    int64
	MaxOutputBytes int64
}
type Usage struct {
	CPUSeconds  int64
	MemoryMB    int64
	OutputBytes int64
}

func (l Limits) Check(u Usage) error {
	if l.MaxCPUSeconds > 0 && u.CPUSeconds > l.MaxCPUSeconds {
		return errors.New("cpu limit exceeded")
	}
	if l.MaxMemoryMB > 0 && u.MemoryMB > l.MaxMemoryMB {
		return errors.New("memory limit exceeded")
	}
	if l.MaxOutputBytes > 0 && u.OutputBytes > l.MaxOutputBytes {
		return errors.New("output limit exceeded")
	}
	return nil
}

type Meter struct {
	mu    sync.Mutex
	usage Usage
}

func (m *Meter) Add(u Usage) {
	m.mu.Lock()
	m.usage.CPUSeconds += u.CPUSeconds
	m.usage.MemoryMB = max(m.usage.MemoryMB, u.MemoryMB)
	m.usage.OutputBytes += u.OutputBytes
	m.mu.Unlock()
}
func (m *Meter) Read() Usage { m.mu.Lock(); defer m.mu.Unlock(); return m.usage }
func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
