include .env

ifeq ($(OS),Windows_NT)
    PROJECT_ROOT := $(CURDIR)
else
    PROJECT_ROOT := $(shell pwd)
endif
export PROJECT_ROOT

env-up:
	@docker compose up -d pollify-postgres

env-down:
	@docker compose down pollify-postgres

env-cleanup:
	@read -p "Clean all environment volume files? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
	  docker compose down pollify-postgres port-forwarder && \
	  rm -rf ${PROJECT_ROOT}/out/pgdata && \
	  echo "Files cleanup"; \
	else \
	  echo "Cleanup is closed"; \
	fi

env-port-forward:
	@docker composeup -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

migrate-create:
	@if [ -z "$(seq)" ]; then \
  		echo "Param seq is not exist: " \
  		exit 1; \
  	fi
	@docker compose run --rm pollify-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
    		echo "Param action is not exist: "; \
    		exit 1; \
    	fi
	@docker compose run --rm --use-aliases pollify-postgres-migrate \
		-path=//migrations \
		-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@todoapp-postgres:5432/${POSTGRES_DB}?sslmode=disable" \
		"$(action)"


pollify-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
    	export POSTGRES_HOST=localhost && \
    	go mod tidy && \
    	go run ${PROJECT_ROOT}/cmd/pollify/main.go