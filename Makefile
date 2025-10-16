ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: run
run:
	@go mod tidy
	@go run .

.PHONY: docs
docs:
	@go tool swag fmt
	@go tool swag init
