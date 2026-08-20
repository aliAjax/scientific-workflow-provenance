# Bug Reproduction

## Bug

Cancellation and deadlines are not consistently honored by retry, batch-worker, and scheduler paths. Work continues after cancellation, while a deadline path can panic.

## Reproduction

Run these commands independently from the repository root:

    go test ./internal/config/../infrastructure/./. -run '^TestRetryerStopsBeforeFirstAttempt$' -count=1
    go test ./internal/adapters -run '^TestBatchExecutorPropagatesCancellation$' -count=1
    go test ./internal/adapters -run '^TestFailingWorkerHonorsContext$' -count=1
    go test ./internal/scheduler -run '^TestQueueAcquireHonorsDeadline$' -count=1

## Observed Errors

    retry ignored cancellation: err=<nil> calls=1
    request 1 was not cancelled
    worker ignored cancellation: worker failed
    panic: runtime error: invalid memory address or nil pointer dereference
    internal/scheduler/scheduler.go:33
