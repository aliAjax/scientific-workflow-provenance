package application

import (
	"scientific-workflow-provenance/internal/domain/provenance"
	"scientific-workflow-provenance/internal/domain/sample"
	"scientific-workflow-provenance/pkg/ids"
)

type SampleService struct {
	Repo  *sample.Store
	Graph *provenance.Graph
}

func (s *SampleService) Create(v sample.Sample) (sample.Sample, error) {
	if v.ID == "" {
		v.ID = ids.New("sample")
	}
	if e := s.Repo.Put(v); e != nil {
		return sample.Sample{}, e
	}
	s.Graph.Add(provenance.Record{ID: "sample:" + v.ID, Kind: provenance.Entity, Type: "sample", Attributes: map[string]any{"source": v.Source}})
	for _, p := range v.ParentIDs {
		s.Graph.Relate(provenance.Relation{ID: ids.New("rel"), From: "sample:" + p, To: "sample:" + v.ID, Type: "wasDerivedFrom"})
	}
	return v, nil
}
func (s *SampleService) Lineage(id string) []provenance.Record {
	return s.Graph.Lineage("sample:" + id)
}
