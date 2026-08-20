package observability

import (
	"encoding/json"
	"io"
	"sync"
	"time"
)

type Logger struct {
	mu  sync.Mutex
	out io.Writer
}

func NewLogger(out io.Writer) *Logger { return &Logger{out: out} }
func (l *Logger) Log(level, msg string, fields map[string]any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	m := map[string]any{"time": time.Now().UTC(), "level": level, "message": msg}
	for k, v := range fields {
		m[k] = v
	}
	_ = json.NewEncoder(l.out).Encode(m)
}
func (l *Logger) Info(msg string, fields map[string]any)  { l.Log("info", msg, fields) }
func (l *Logger) Error(msg string, fields map[string]any) { l.Log("error", msg, fields) }
