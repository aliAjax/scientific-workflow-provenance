# Bug Reproduction

## Bug

Storage compaction, eviction, provenance pagination, and artifact sorting mutate caller-owned slice data. Earlier page and cursor state changes after later operations.

## Reproduction

Run these commands independently from the repository root:

    go test ./internal/quality/../storage -run '^TestCompactDoesNotMutateInput$' -count=1
    go test ./internal/quality/../storage -run '^TestEvictKeepsCursorInputStable$' -count=1
    go test scientific-workflow-provenance/internal/query/. -run '^TestProvenancePageCursorStable$' -count=1
    go test scientific-workflow-provenance/internal/artifact/. -run '^TestArtifactSortSnapshot$' -count=1

## Observed Errors

    compaction_test.go:9: compact reordered caller slice
    eviction_test.go:12: eviction reordered cursor input
    pagination_test.go:14: cursor points to wrong item: a
    paging_test.go:13: artifact sort escaped input
