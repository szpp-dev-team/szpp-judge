.PHONY: prepare
prepare:
	@cd backend && go mod vendor

.PHONY: up
up: prepare .env
	@docker compose up --wait --build -d

.PHONY: clean
clean:
	@docker compose down --rmi local -v --remove-orphans \
	&& rm -rf ./.data

