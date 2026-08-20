package security

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

var sensitive = []string{"token", "password", "secret", "authorization", "private_key"}

func RedactMap(in map[string]string) map[string]string {
	out := map[string]string{}
	for k, v := range in {
		low := strings.ToLower(k)
		hide := false
		for _, s := range sensitive {
			if strings.Contains(low, s) {
				hide = true
			}
		}
		if hide {
			out[k] = "[REDACTED]"
		} else {
			out[k] = v
		}
	}
	return out
}
func HashIdentity(value string) string {
	h := sha256.Sum256([]byte(value))
	return hex.EncodeToString(h[:])
}
func AllowedAlgorithm(value string) bool {
	switch strings.ToLower(value) {
	case "sha256", "sha384", "sha512":
		return true
	}
	return false
}
