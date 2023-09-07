.PHONY: ent/generate
ent/generate:
	@GOFLAGS="-mod=mod" go generate ./domain/repository/ent

.PHONY: tbls/generate
tbls/generate:
	@rm -rf ./docs \
	&& tbls doc

.PHONY: mysql/start
mysql/start:
ifeq ($(CI),true)
	$(eval CONTAINER_ID := docker ps | grep mariadb | awk '{print $1}')
else
	$(eval CONTAINER_ID := $(shell docker run -d -P \
		-e MYSQL_DATABASE=szpp-judge-db \
		-e MYSQL_ROOT_PASSWORD=szpp \
		mariadb:11.1))
endif
	@echo $(CONTAINER_ID)

.PHONY: test
test: mysql/start
	@TEST_MYSQL_PORT=$(shell docker port $(CONTAINER_ID) | rg 3306/tcp | awk -F ' -> ' '{print $$2}' | awk -F ':' '{print $$2}') \
	MYSQL_DATABASE=szpp-judge-db \
	MYSQL_ROOT_PASSWORD=szpp \
	go test ./... -v \
	; docker rm -f $(CONTAINER_ID)

.PHONY: test-ci
test-ci: mysql/start
	@TEST_MYSQL_PORT=$(shell docker port $(CONTAINER_ID) | rg 3306/tcp | awk -F ' -> ' '{print $$2}' | awk -F ':' '{print $$2}') \
	MYSQL_DATABASE=szpp-judge-db \
	MYSQL_ROOT_PASSWORD=szpp \
	go test ./... -v -coverprofile=coverage.out
