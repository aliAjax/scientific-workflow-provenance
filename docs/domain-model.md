# Domain Model

WorkflowDefinition owns versioned DAG nodes and edges. A Run references one immutable definition version and has node state/attempt maps. Sample entities form a parent-child derivation graph and can transition active -> frozen -> destroyed. Artifacts become publishable only after schema/unit/range quality gates. Every entity/activity is appended to a provenance graph with a chained SHA-256 hash.
