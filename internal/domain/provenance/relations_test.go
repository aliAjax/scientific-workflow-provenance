package provenance

import "testing"

func TestRelationErrorPreservesLineage(t *testing.T) {
	_, err := New().LinkValidated(Relation{From: "same", To: "same", Type: "derived"})
	if err == nil {
		t.Fatal("self relation accepted")
	}
}
