.PHONY: ent/generate
ent/generate:
	@GOFLAGS="-mod=mod" go generate ./domain/repository/ent

.PHONY: tbls/generate
tbls/generate:
	@rm -rf ./docs \
	&& tbls doc

MYSQL_DATABASE=szpp-judge-test-db
MYSQL_ROOT_PASSWORD=root

.PHONY: mysql/start
mysql/start:
	$(eval CONTAINER_ID := $(shell docker run -d -P \
		-e MYSQL_DATABASE=${MYSQL_DATABASE} \
		-e MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} \
		mysql:8.0))
	@echo $(CONTAINER_ID)


.PHONY: test
test: mysql/start
	@TEST_MYSQL_PORT=$(shell docker port $(CONTAINER_ID) | rg 3306/tcp | awk -F ' -> ' '{print $$2}' | awk -F ':' '{print $$2}') \
	MYSQL_DATABASE=${MYSQL_DATABASE} \
	MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} \
	go test ./usecases/... -v \
	; docker rm -f $(CONTAINER_ID)
