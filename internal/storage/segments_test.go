package storage

import "testing"

func TestSegmentPendingSnapshot(t *testing.T) {
	l := NewSegmentLog(2)
	_ = l.Append("s", []byte("payload"))
	out := l.Pending()
	out[0].Payload[0] = 'X'
	if string(l.Pending()[0].Payload) != "payload" {
		t.Fatal("pending payload escaped")
	}
}
