# Bug Reproduction

## Bug

Provenance graph validation and export do not consistently reject invalid records and chains. Error identity is lost, and exporting an invalid chain can panic.

## Reproduction

Run these commands independently from the repository root:

    go test ./internal/api/../domain/provenance -run '^TestGraphAddValidatesRecordCause$' -count=1
    go test ./internal/api/../domain/provenance -run '^TestRelationErrorPreservesLineage$' -count=1
    go test ./internal/api/../domain/provenance -run '^TestValidateChainPreservesSentinel$' -count=1
    go test ./internal/security/../application -run '^TestExportJSONRejectsInvalidChain$' -count=1

## Observed Errors

    model_test.go:11: invalid record cause missing: <nil>
    relations_test.go:8: self relation accepted
    prov_json_test.go:10: chain sentinel lost
    panic: runtime error: invalid memory address or nil pointer dereference
    internal/domain/provenance/model.go:90
    internal/application/provenance.go:29
