package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
)

func Sign(secret, data []byte) string {
	h := hmac.New(sha256.New, secret)
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
func Verify(secret, data []byte, signature string) bool {
	expected := Sign(secret, data)
	return subtle.ConstantTimeCompare([]byte(expected), []byte(signature)) == 1
}
