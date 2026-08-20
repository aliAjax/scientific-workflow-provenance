package worker

import (
	"strings"
	"sync"
	"time"
)

type LogLine struct {
	At      time.Time `json:"at"`
	Level   string    `json:"level"`
	Message string    `json:"message"`
}
type LogBuffer struct {
	mu    sync.RWMutex
	max   int
	lines []LogLine
}

func NewLogBuffer(max int) *LogBuffer {
	if max < 1 {
		max = 1000
	}
	return &LogBuffer{max: max}
}
func (b *LogBuffer) Append(level, msg string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.lines = append(b.lines, LogLine{At: time.Now().UTC(), Level: level, Message: Sanitize(msg)})
	if len(b.lines) > b.max {
		b.lines = b.lines[len(b.lines)-b.max:]
	}
}
func (b *LogBuffer) List() []LogLine {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return append([]LogLine(nil), b.lines...)
}
func Sanitize(s string) string {
	for _, secret := range []string{"TOKEN=", "PASSWORD=", "SECRET="} {
		for strings.Contains(s, secret) {
			i := strings.Index(s, secret)
			j := strings.IndexByte(s[i:], ' ')
			if j < 0 {
				j = len(s) - i
			}
			s = s[:i] + secret + "***" + s[i+j:]
		}
	}
	return s
}
