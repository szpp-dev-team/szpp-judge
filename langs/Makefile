.DEFAULT_GOAL      := help
SHELL              := /usr/bin/env bash -euo pipefail
DOCKER_DIR         := _docker
SZPP_IMAGE_LABEL   := szpp-judge-image=true

.PHONY:	help
help:
	@bat Makefile || batcat Makefile || cat Makefile


.PHONY:	test
test:
	go test -v -shuffle=on ./...


.PHONY: fmt
fmt:
	go fmt ./...


.PHONY:	lint
lint:
	go vet ./...


.PHONY:	gen
gen: lint test __gen


.PHONY:	__gen
__gen:
	go run ./cmd/gen-ts/main.go -out ../frontend/src/gen/langs.ts


.PHONY:	docker/build/szpprun
docker/build/szpprun:
	./build-image.sh _docker/_szpprun


.PHONY:	docker/build/langs
docker/build/langs:	docker/build/szpprun
	find _docker -maxdepth 1 -type d -name '[a-zA-Z]*' | xargs -I@ ./build-image.sh @


.PHONY:	docker/image/gc
docker/image/gc:
	docker images -f label=$(SZPP_IMAGE_LABEL) -f dangling=true -q | xargs docker image rm


.PHONY:	docker/image/ls
docker/image/ls:
	docker images -f label=$(SZPP_IMAGE_LABEL)
