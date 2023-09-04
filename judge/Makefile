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


.PHONY:	docker/build/monitoring_exec
docker/build/monitoring_exec:
	docker build -t $(IMAGE_NAME_PREFIX)-monitoring_exec:latest $(DOCKER_DIR)/monitoring_exec


.PHONY:	docker/build/langs
docker/build/langs:	docker/build/monitoring_exec
	docker build -t $(IMAGE_NAME_PREFIX)-gcc:latest $(DOCKER_DIR)/gcc
	docker build -t $(IMAGE_NAME_PREFIX)-python:latest $(DOCKER_DIR)/python
	docker build -t $(IMAGE_NAME_PREFIX)-java:latest $(DOCKER_DIR)/java


.PHONY:	docker/image/gc
docker/image/gc:
	docker images -f label=szpp-judge-image=true -f dangling=true -q | xargs docker image rm


.PHONY:	docker/image/ls
docker/image/ls:
	docker images -f label=szpp-judge-image=true
