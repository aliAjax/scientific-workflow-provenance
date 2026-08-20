package ids

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

func New(prefix string) string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return fmt.Sprintf("%s-%d", prefix, time.Now().UnixNano())
	}
	return prefix + "-" + hex.EncodeToString(b)
}
func Digest(parts ...string) string {
	h := uint64(1469598103934665603)
	for _, p := range parts {
		for i := 0; i < len(p); i++ {
			h ^= uint64(p[i])
			h *= 1099511628211
		}
	}
	return fmt.Sprintf("%016x", h)
}
