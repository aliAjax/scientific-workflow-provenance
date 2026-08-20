package query

import (
	"scientific-workflow-provenance/internal/domain/sample"
	"scientific-workflow-provenance/internal/domain/workflow"
	"sort"
)

func SamplesByID(v []sample.Sample) []sample.Sample {
	o := append([]sample.Sample(nil), v...)
	sort.Slice(o, func(i, j int) bool { return o[i].ID < o[j].ID })
	return o
}
func SamplesByCreated(v []sample.Sample) []sample.Sample {
	o := append([]sample.Sample(nil), v...)
	sort.Slice(o, func(i, j int) bool { return o[i].CreatedAt.Before(o[j].CreatedAt) })
	return o
}
func RunsByStatus(v []workflow.Run) []workflow.Run {
	o := append([]workflow.Run(nil), v...)
	sort.Slice(o, func(i, j int) bool { return o[i].Status < o[j].Status })
	return o
}
func RunsByUpdated(v []workflow.Run) []workflow.Run {
	o := append([]workflow.Run(nil), v...)
	sort.Slice(o, func(i, j int) bool { return o[i].UpdatedAt.After(o[j].UpdatedAt) })
	return o
}
