export IMAGE_NAME=snp-api

.DEFAULT_GOAL := help
## Follow service logs
logs:
	@echo "Follow services' logs..."
	@docker-compose -f local-env/docker-compose.yml logs -f

## Run tests
test: up
	@echo "Running tests..."
	@go test -race -failfast ./...

## Generate swag docs
swag-docs:
	@echo "Updating swagger"
	@swag init --parseDependency --parseInternal --parseDepth 1 -g ./cmd/api/main.go

## Install swag client
swag-install:
	@go get -u github.com/swaggo/swag/cmd/swag

## Start all services
up:
	@docker-compose -f ./docker-compose.yml up -d

## Stop all services (if they are running)
down:
	@echo "Stopping services..."
	@docker-compose -f ./docker-compose.yml down

## Delete all docker resources
delete-containers:
	@docker ps -a -q
	@docker rm -f $(shell docker ps -a -q)
	@docker volume rm $(shell docker volume ls -q)

# -- help
# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)
TARGET_MAX_CHAR_NUM=20

## Show help
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
