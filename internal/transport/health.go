package transport

import (
	"net/http"
	"time"
)

type Health struct {
	Started    time.Time
	Components map[string]func() error
}

func (h Health) Check() map[string]any {
	out := map[string]any{"status": "ok", "uptime": time.Since(h.Started).String()}
	for k, f := range h.Components {
		if f != nil {
			if e := f(); e != nil {
				out["status"] = "degraded"
				out[k] = e.Error()
			} else {
				out[k] = "ok"
			}
		}
	}
	return out
}
func (h Health) Handler(w http.ResponseWriter, r *http.Request) { write(w, http.StatusOK, h.Check()) }
