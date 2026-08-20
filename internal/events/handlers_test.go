package events

import (
	"scientific-workflow-provenance/pkg/events"
	"sync"
	"sync/atomic"
	"testing"
)

func TestDispatcherConcurrentLifecycle(t *testing.T) {
	var d Dispatcher
	start := make(chan struct{})
	var wg sync.WaitGroup
	var called atomic.Int32
	h := func(events.Event) { called.Add(1) }
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			d.On("run", h)
		}()
	}
	close(start)
	wg.Wait()
	d.Dispatch(events.Event{Type: "run"})
	if called.Load() != 2 {
		t.Fatalf("concurrent dispatcher delivered %d handlers", called.Load())
	}
}
