package ports

import (
	"scientific-workflow-provenance/internal/domain/provenance"
	"scientific-workflow-provenance/internal/domain/sample"
)

type SampleRepository interface {
	Put(sample.Sample) error
	Get(string) (sample.Sample, error)
	List() []sample.Sample
	Freeze(string) error
	Destroy(string) error
}
type ProvenanceStore interface {
	Add(provenance.Record) provenance.Record
	Lineage(string) []provenance.Record
	Export() map[string]any
}
