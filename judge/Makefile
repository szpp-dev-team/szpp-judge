.DEFAULT_GOAL      := help
SHELL              := /usr/bin/env bash -euo pipefail
DOCKER_DIR         := _docker
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
	docker build -t $(IMAGE_NAME_PREFIX)-gcc $(DOCKER_DIR)/gcc
	docker build -t $(IMAGE_NAME_PREFIX)-python $(DOCKER_DIR)/python
	docker build -t $(IMAGE_NAME_PREFIX)-java $(DOCKER_DIR)/java


.PHONY:	langs/ls/images
langs/ls/images:
	docker images -f label=szpp-judge-image=true
