# Bug Reproduction

## Bug

Concurrent event lifecycle operations are unsafe: dispatcher and bus setup can panic, outbox delivery can miss due messages, and concurrent logging can dereference an uninitialized encoder.

## Reproduction

Run these commands independently from the repository root:

    go test -race ./internal/events -run '^TestDispatcherConcurrentLifecycle$' -count=1
    go test -race ./internal/events -run '^TestOutboxDueDoesNotHang$' -count=1
    go test -race ./pkg/events -run '^TestBusPublishUnderBackpressure$' -count=1
    go test -race ./internal/verify008 -run '^TestLoggerConcurrentWrites$' -count=1

## Observed Errors

    panic: assignment to entry in nil map
    internal/events/handlers.go:18
    outbox_test.go:13: unlimited due missed item: []events.Message{}
    panic: assignment to entry in nil map
    pkg/events/events.go:30
    panic: runtime error: invalid memory address or nil pointer dereference
    internal/observability/log.go:23
