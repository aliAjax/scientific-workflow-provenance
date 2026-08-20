package ports

import (
	"context"
	"scientific-workflow-provenance/internal/domain/workflow"
)

type WorkflowRepository interface {
	Put(workflow.Definition) error
	Get(string, int) (workflow.Definition, error)
	List() []workflow.Definition
}
type RunRepository interface {
	Put(workflow.Run)
	Get(string) (workflow.Run, bool)
	List() []workflow.Run
}
type Planner interface {
	Build(workflow.Definition, map[string]string) (any, error)
}
type Worker interface {
	Execute(context.Context, string, string, map[string]string) (any, error)
}
