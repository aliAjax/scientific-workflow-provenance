package worker

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
)

func ResultDigest(outputs map[string][]byte) string {
	keys := make([]string, 0, len(outputs))
	for k := range outputs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	h := sha256.New()
	for _, k := range keys {
		h.Write([]byte(k))
		h.Write(outputs[k])
	}
	return hex.EncodeToString(h.Sum(nil))
}
func Merge(a, b map[string][]byte) map[string][]byte {
	out := map[string][]byte{}
	for k, v := range a {
		out[k] = append([]byte(nil), v...)
	}
	for k, v := range b {
		out[k] = append([]byte(nil), v...)
	}
	return out
}
