package execution

import "testing"

func TestCheckpointLatestSnapshot(t *testing.T) {
	s := NewCheckpointStore()
	s.Save(Snapshot{RunID: "r", States: map[string]string{"step": "queued"}})
	v, _ := s.Latest("r")
	v.States["step"] = "failed"
	again, _ := s.Latest("r")
	if again.States["step"] != "queued" {
		t.Fatal("checkpoint state escaped")
	}
}
