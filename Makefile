COMPOSE=docker-compose

.PHONY: run
run:
	$(COMPOSE) up -d --build

.PHONY: up
up:
	$(COMPOSE) up -d

.PHONY: migrate-create
migrate-create:
	$(COMPOSE) exec api goose create ${FILE_NAME} sql

.PHONY: migrate-up
migrate-up:
	$(COMPOSE) exec api goose up
