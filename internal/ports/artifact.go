package ports

import (
	"context"
	"scientific-workflow-provenance/internal/artifact"
)

type ArtifactRepository interface {
	Put(artifact.Artifact) error
	Get(string) (artifact.Artifact, bool)
	Invalidate(string, string) error
	List() []artifact.Artifact
}
type ObjectStore interface {
	Put(context.Context, string, []byte) error
	Get(context.Context, string) ([]byte, error)
	Delete(context.Context, string) error
}
