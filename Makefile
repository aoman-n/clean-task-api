COMPOSE=docker-compose

.PHONY: run
run:
	$(COMPOSE) up -d --build

.PHONY: up
up:
	$(COMPOSE) up -d

.PHONY: down
down:
	$(COMPOSE) down

.PHONY: migrate-create
migrate-create:
	$(COMPOSE) run --rm api goose create ${FILE_NAME} sql

.PHONY: migrate-up
migrate-up:
	$(COMPOSE) run --rm api goose up

.PHONY: in-api
in-api:
	$(COMPOSE) exec api /bin/bash

.PHONY: test-api-integration
test-api-integration:
	docker-compose exec \
	-e DB_NAME=test_db \
	api richgo test -v -tags=integration ./...
