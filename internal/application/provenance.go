package application

import (
	"encoding/json"
	"scientific-workflow-provenance/internal/domain/provenance"
	"scientific-workflow-provenance/pkg/audit"
	"scientific-workflow-provenance/pkg/ids"
)

type ProvenanceService struct {
	Graph *provenance.Graph
	Audit *audit.Chain
}

func NewProvenanceService(g *provenance.Graph) *ProvenanceService {
	return &ProvenanceService{Graph: g, Audit: audit.New()}
}
func (s *ProvenanceService) Entity(id, typ string, attrs map[string]any) provenance.Record {
	r := s.Graph.Add(provenance.Record{ID: id, Kind: provenance.Entity, Type: typ, Attributes: attrs})
	s.Audit.Append("system", "entity.created", id, nil)
	return r
}
func (s *ProvenanceService) Activity(typ string, attrs map[string]any) provenance.Record {
	id := ids.New("activity")
	r := s.Graph.Add(provenance.Record{ID: id, Kind: provenance.Activity, Type: typ, Attributes: attrs})
	s.Audit.Append("system", "activity.created", id, nil)
	return r
}
func (s *ProvenanceService) ExportJSON() ([]byte, error) {
	var data map[string]any
	if s.Graph != nil {
		data = s.Graph.Export()
	}
	if err := provenance.ValidateChain(data); err != nil {
		return nil, err
	}
	return json.Marshal(data)
}
func (s *ProvenanceService) VerifyAudit() bool           { return s.Audit.Verify() }
