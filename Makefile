GO ?= go
.PHONY: fmt test race vet build run smoke
fmt:
	$(GO)fmt -w $$(find . -name '*.go')
test:
	$(GO) test ./...
race:
	$(GO) test -race ./...
vet:
	$(GO) vet ./...
build:
	$(GO) build ./...
run:
	$(GO) run ./cmd/workflowd
smoke:
	sh scripts/smoke.sh
