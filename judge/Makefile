.DEFAULT_GOAL := help
SHELL         := /usr/bin/env bash -euo pipefail

.PHONY:	help
help:
	@bat Makefile || batcat Makefile || cat Makefile


.PHONY:	test
test:
	go test -v -race -shuffle=on ./...


.PHONY: fmt
fmt:
	go fmt ./...


.PHONY:	lint
lint:
	golangci-lint run ./...
