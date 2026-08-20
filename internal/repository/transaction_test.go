package repository

import (
	"testing"
	"time"
)

func TestTransactionFailureRollsBack(t *testing.T) {
	tx := Begin()
	done := make(chan struct{})
	if err := tx.Do(func() {}, func() { _ = tx.Do(func() {}, func() {}); close(done) }); err != nil {
		t.Fatal(err)
	}
	rollbackDone := make(chan struct{})
	go func() { tx.Rollback(); close(rollbackDone) }()
	select {
	case <-rollbackDone:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("rollback deadlocked on transaction lock")
	}
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("rollback callback deadlocked on transaction lock")
	}
}
