package application

import (
	"fmt"
	"scientific-workflow-provenance/internal/artifact"
	"scientific-workflow-provenance/internal/domain/provenance"
	"scientific-workflow-provenance/pkg/ids"
)

type ArtifactService struct {
	Store *artifact.Store
	Graph *provenance.Graph
}

func NewArtifactService(s *artifact.Store, g *provenance.Graph) *ArtifactService {
	return &ArtifactService{Store: s, Graph: g}
}
func (s *ArtifactService) Publish(a artifact.Artifact, g artifact.QualityGate) error {
	if err := artifact.Validate(a, g); err != nil {
		return err
	}
	if a.Digest == "" {
		a.Digest = artifact.Digest([]byte(fmt.Sprintf("%s:%f:%s", a.Name, a.Value, a.Unit)))
	}
	if err := s.Store.Put(a); err != nil {
		return err
	}
	s.Graph.Add(provenance.Record{ID: "artifact:" + a.Digest, Kind: provenance.Entity, Type: "artifact", Attributes: map[string]any{"name": a.Name, "schema": a.Schema}})
	return nil
}
func (s *ArtifactService) Invalidate(d, reason string) error {
	if err := s.Store.Invalidate(d, reason); err != nil {
		return err
	}
	s.Graph.Add(provenance.Record{ID: ids.New("invalidation"), Kind: provenance.Activity, Type: "invalidation", Attributes: map[string]any{"digest": d, "reason": reason}})
	return nil
}
