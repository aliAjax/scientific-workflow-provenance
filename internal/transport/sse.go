package transport

import (
	"fmt"
	"net/http"
	"time"
)

func writeSSE(w http.ResponseWriter, event, data string) {
	fmt.Fprintf(w, "event: %s\\ndata: %s\\n\\n", event, data)
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}
func streamHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
}
func heartbeat(w http.ResponseWriter, done <-chan struct{}) {
	t := time.NewTicker(15 * time.Second)
	defer t.Stop()
	for {
		select {
		case <-done:
			return
		case <-t.C:
			writeSSE(w, "heartbeat", time.Now().UTC().Format(time.RFC3339))
		}
	}
}
