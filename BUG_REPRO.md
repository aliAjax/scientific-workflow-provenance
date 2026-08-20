# Bug Reproduction

## Bug

Sample lineage, batch, metric, and store results share nested slice storage with their inputs. Later mutations alter previously returned or stored sample data.

## Reproduction

Run these commands independently from the repository root:

    go test ./internal/domain/sample -run '^TestLineageDoesNotAliasParents$' -count=1
    go test ./internal/domain/sample -run '^TestBatchGroupDoesNotAliasSample$' -count=1
    go test ./internal/domain/sample -run '^TestMetricValuesAreIsolated$' -count=1
    go test ./internal/domain/sample -run '^TestSampleStoreDeepSnapshot$' -count=1

## Observed Errors

    model_test.go:29: lineage result aliased input
    model_test.go:38: batch group aliased input
    model_test.go:48: metric values escaped
    model_test.go:17: sample store leaked nested slices
