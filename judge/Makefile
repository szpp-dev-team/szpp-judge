.DEFAULT_GOAL      := help
SHELL              := /usr/bin/env bash -euo pipefail
DOCKERFILES_DIR    := _docker
IMAGE_NAME_PREFIX  := szpp-judge-images

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


.PHONY:	langs/build/images
langs/build/images:
	docker build -t $(IMAGE_NAME_PREFIX)-gcc -f $(DOCKERFILES_DIR)/gcc.Dockerfile $(DOCKERFILES_DIR)


.PHONY:	langs/ls/images
langs/ls/images:
	docker images -f label=szpp-judge-image=true
