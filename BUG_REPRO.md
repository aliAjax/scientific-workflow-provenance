# Bug Reproduction

## Bug

Retry state transitions disagree across workflow, scheduler, query, and HTTP layers. Retrying work cannot follow the expected path, and terminal state can be reported as running.

## Reproduction

Run these commands independently from the repository root:

    go test ./internal/planner/../domain/workflow -run '^TestRetryTransitionReachesSuccess$' -count=1
    go test ./internal/worker/../scheduler -run '^TestShouldRetryHonorsRetryableClass$' -count=1
    go test ./internal/planner/../query -run '^TestRunsFilterRetrying$' -count=1
    go test ./internal/transport -run '^TestRunHTTPReportsTerminalState$' -count=1

## Observed Errors

    state_test.go:9: queued: cannot transition retrying to queued
    retry_test.go:11: non-retryable error was accepted
    query_test.go:11: in-progress states missing
    server_test.go:21: status was rewritten: running
