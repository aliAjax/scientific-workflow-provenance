package security

import (
	"testing"
	"time"
)

func TestZeroNonceStoreRoundTrip(t *testing.T) {
	var n NonceStore
	v := n.Issue(time.Second)
	if v == "" || !n.Consume(v) || n.Consume(v) {
		t.Fatal("nonce lifecycle failed")
	}
}
