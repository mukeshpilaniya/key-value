# See `make help` for a list of all available commands.
# Configuration
PROJECT_NAME ?= kay-value
BUILD_TIMESTAMP := $(shell date +%Y-%m-%d-%H-%M-%S)
CI_COMMIT_SHORT_SHA := $(shell git rev-parse --short HEAD)
.ONESHELL:
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables

.PHONY: up
up: build start

.PHONY: down
down: stop

.PHONY: build
build:
	@echo "building backend..."
	@go build -o ./bin/key-value ./cmd/*
	@echo "backend built!"

.PHONY: start
start:
	@echo "starting backend..."
	@ ./bin/key-value &
	@echo "backend started!"

.PHONY: stop
stop:
	@echo "stopping backend ..."
	@pkill -SIGTERM  "key-value"
	@echo "stopped backend ..."