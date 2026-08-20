# Bug Reproduction

## Bug

Run snapshots expose mutable maps and parameter state. Mutating returned or submitted values can change stored state, and concurrent mutations are reported as data races.

## Reproduction

Run these commands independently from the repository root:

    go test -race ./internal/repository -run '^TestRunRepoListSnapshot$' -count=1
    go test -race ./internal/execution -run '^TestEngineAttemptsSnapshot$' -count=1
    go test -race ./internal/application -run '^TestRunServiceRunIsolation$' -count=1
    go test -race ./internal/repository -run '^TestRunRepoGetSnapshot$' -count=1

## Observed Errors

    WARNING: DATA RACE
    memory_test.go:51: list leaked mutable attempt map
    engine_test.go:38: adapter mutated caller parameters
    runs_test.go:35: service leaked run parameters
    memory_test.go:31: repository state was aliased
    testing.go:1399: race detected during execution of test
