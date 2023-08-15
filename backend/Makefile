.PHONY: ent/generate
ent/generate:
	@GOFLAGS="-mod=mod" go generate ./domain/repository/ent

.PHONY: tbls/generate
tbls/generate:
	@rm -rf ./docs \
	&& tbls doc

.PHONY: mysql/start
mysql/start:
	$(eval CONTAINER_ID := $(shell docker run -d -P \
		-e MYSQL_DATABASE=szpp-judge-db \
		-e MYSQL_ROOT_PASSWORD=szpp \
		mysql:8.0))
	@echo $(CONTAINER_ID)

