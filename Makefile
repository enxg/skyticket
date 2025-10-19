ifneq (,$(wildcard ./.env))
    include .env
    export
endif

.PHONY: run
run:
	@go mod tidy
	@go run .

.PHONY: build-prod
build-prod:
	@CGO_ENABLED=0 GOOS=linux go build -o skyticket ./main.go

.PHONY: docs
docs:
	@go tool swag fmt
	@go tool swag init
