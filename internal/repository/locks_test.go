package repository

import "testing"

func TestZeroLockTableIsUsable(t *testing.T) {
	var l LockTable
	if err := l.Acquire("sample", "worker", 0); err != nil {
		t.Fatal(err)
	}
	l.Release("sample", "worker")
}
