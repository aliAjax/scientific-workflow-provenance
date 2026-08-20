package workflow

type Impact struct {
	ChangedNodes  []string `json:"changedNodes"`
	AffectedNodes []string `json:"affectedNodes"`
	Breaking      bool     `json:"breaking"`
}

func Analyze(oldDef, newDef Definition) Impact {
	old := map[string]Node{}
	for _, n := range oldDef.Nodes {
		old[n.ID] = n
	}
	changed := []string{}
	for _, n := range newDef.Nodes {
		o, ok := old[n.ID]
		if !ok || o.Kind != n.Kind || o.Timeout != n.Timeout {
			changed = append(changed, n.ID)
		}
	}
	affected := map[string]bool{}
	for _, e := range newDef.Edges {
		for _, c := range changed {
			if e.From == c {
				affected[e.To] = true
			}
		}
	}
	a := []string{}
	for id := range affected {
		a = append(a, id)
	}
	return Impact{ChangedNodes: changed, AffectedNodes: a, Breaking: len(changed) > 0 && oldDef.Version == newDef.Version}
}
