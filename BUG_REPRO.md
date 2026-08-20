# Bug Reproduction

## Bug

Object, artifact-quality, adapter, and API errors lose their original causes while crossing layers. Cause checks and API error classification consequently fail.

## Reproduction

Run these commands independently from the repository root:

    go test ./internal/infrastructure -run '^TestObjectStoreNotFoundChain$' -count=1
    go test ./internal/adapters -run '^TestObjectAdapterPreservesCause$' -count=1
    go test ./internal/quality/../application -run '^TestArtifactPublishPreservesQualityCause$' -count=1
    go test ./internal/api -run '^TestAPIErrorClassification$' -count=1

## Observed Errors

    objectstore_test.go:12: missing cause: object "missing": object not found
    object_test.go:13: adapter lost cause: object adapter: object "missing": object not found
    artifacts_test.go:14: quality cause lost: artifact quality rejected: value below minimum: artifact quality gate failed
    errors_test.go:11: api wrapper lost cause
