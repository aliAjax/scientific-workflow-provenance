package application

import (
	"fmt"
	"scientific-workflow-provenance/internal/domain/sample"
	"scientific-workflow-provenance/internal/domain/workflow"
	"scientific-workflow-provenance/internal/planner"
)

type ValidationService struct{}

func (ValidationService) Workflow(d workflow.Definition) []error {
	errs := []error{}
	if e := d.Validate(); e != nil {
		errs = append(errs, e)
	}
	for _, n := range d.Nodes {
		if e := planner.ValidateCondition(n); e != nil {
			errs = append(errs, e)
		}
	}
	return errs
}
func (ValidationService) Sample(v sample.Sample) error {
	if e := sample.Validate(v); e != nil {
		return fmt.Errorf("sample validation: %w", e)
	}
	return nil
}
