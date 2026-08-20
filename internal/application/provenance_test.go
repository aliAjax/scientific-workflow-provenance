package application

import "testing"

func TestExportJSONRejectsInvalidChain(t *testing.T) {
	s := NewProvenanceService(nil)
	if _, err := s.ExportJSON(); err == nil {
		t.Fatal("invalid empty graph exported")
	}
}
