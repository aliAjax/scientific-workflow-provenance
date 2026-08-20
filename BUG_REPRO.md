# Bug Reproduction

## Bug

Zero-value nonce, sampler, metrics, and lock components are not usable without explicit initialization. Their first normal operation panics.

## Reproduction

Run these commands independently from the repository root:

    go test ./internal/security -run '^TestZeroNonceStoreRoundTrip$' -count=1
    go test ./internal/observability -run '^TestSamplerZeroValueDoesNotPanic$' -count=1
    go test ./internal/metrics -run '^TestZeroRegistryIsWritable$' -count=1
    go test ./internal/ports/../repository/. -run '^TestZeroLockTableIsUsable$' -count=1

## Observed Errors

    panic: assignment to entry in nil map
    internal/security/nonce.go:21
    panic: runtime error: invalid memory address or nil pointer dereference
    internal/observability/sampling.go:22
    panic: assignment to entry in nil map
    internal/metrics/registry.go:18
    panic: assignment to entry in nil map
    internal/repository/locks.go:27
