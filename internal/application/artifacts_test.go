package application

import (
	"errors"
	"scientific-workflow-provenance/internal/artifact"
	"testing"
)

func TestArtifactPublishPreservesQualityCause(t *testing.T) {
	svc := NewArtifactService(artifact.NewStore(), nil)
	min := 10.0
	err := svc.Publish(artifact.Artifact{Digest: "d", Value: 1}, artifact.QualityGate{Min: &min})
	if !errors.Is(err, artifact.ErrInvalid) {
		t.Fatalf("quality cause lost: %v", err)
	}
}
