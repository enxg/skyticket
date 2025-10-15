ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: run
run:
	@go mod tidy
	@go run .
