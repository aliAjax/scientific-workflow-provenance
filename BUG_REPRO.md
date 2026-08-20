# Bug Reproduction

## Bug

A failed transaction can deadlock during rollback, and segment, checkpoint, and audit snapshot results expose mutable state to callers.

## Reproduction

Run these commands independently from the repository root:

    go test ./internal/storage/../repository/. -run '^TestTransactionFailureRollsBack$' -count=1
    go test ./internal/lifecycle/../storage -run '^TestSegmentPendingSnapshot$' -count=1
    go test ./internal/domain/../execution -run '^TestCheckpointLatestSnapshot$' -count=1
    go test ./pkg/audit -run '^TestAuditListMetadataSnapshot$' -count=1

## Observed Errors

    transaction_test.go:19: rollback deadlocked on transaction lock
    segments_test.go:11: pending payload escaped
    checkpoint_test.go:12: checkpoint state escaped
    audit_test.go:11: audit metadata escaped
