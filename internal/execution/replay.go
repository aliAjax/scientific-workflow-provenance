package execution

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"scientific-workflow-provenance/internal/domain/workflow"
	"scientific-workflow-provenance/internal/planner"
	"sort"
)

type Event struct {
	Sequence     int                `json:"sequence"`
	StepID       string             `json:"stepId"`
	Status       workflow.RunStatus `json:"status"`
	OutputDigest string             `json:"outputDigest"`
}

func Replay(events []Event) map[string]workflow.RunStatus {
	sort.Slice(events, func(i, j int) bool { return events[i].Sequence < events[j].Sequence })
	out := map[string]workflow.RunStatus{}
	for _, e := range events {
		out[e.StepID] = e.Status
	}
	return out
}
func ReplayDigest(events []Event) string {
	h := sha256.New()
	for _, e := range events {
		fmt.Fprintf(h, "%d:%s:%s:%s|", e.Sequence, e.StepID, e.Status, e.OutputDigest)
	}
	return hex.EncodeToString(h.Sum(nil))
}
func CacheCompatible(p planner.Plan, events []Event) bool {
	return p.Digest == ReplayDigest(events) || len(events) == 0
}
