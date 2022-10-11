include ./config/config_example.env
CALENDAR := "calendar"
SCHEDULER := "scheduler"
SENDER := "sender"

build:
	go build -v -o ./bin/$(CALENDAR) -ldflags "$(LDFLAGS)" ./cmd/$(CALENDAR)
	go build -v -o ./bin/$(SCHEDULER) -ldflags "$(LDFLAGS)" ./cmd/$(SCHEDULER)
	go build -v -o ./bin/$(SENDER) -ldflags "$(LDFLAGS)" ./cmd/$(SENDER)

up:  config-copy build docker-up migrate-up
down: docker-down

test:
	go test -race ./internal/...

integration-tests: testi

fmt:
	gofmt -s -w .
	gofumpt -l -w .

lint: install-lint-deps
	golangci-lint run ./...

generate:
	protoc ./api/eventService.proto --go_out=./internal/grpc/protobuf --go-grpc_out=./internal/grpc/protobuf

docker-up:
	docker-compose -f ./docker/docker-compose.yml  --env-file ./config/config.env up -d
docker-down:
	docker-compose -f ./docker/docker-compose.yml --env-file ./config/config.env down --remove-orphans
docker-build:
	docker-compose -f ./docker/docker-compose.yml  --env-file ./config/config.env up -d --build

# for example `make migrate-create table=event`
migrate-create:
	docker-compose -f ./docker/docker-compose.yml  --env-file ./config/config.env run migration /go/bin/goose postgres "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"  create $(table) sql

migrate-up:
	docker-compose -f ./docker/docker-compose.yml  --env-file ./config/config.env run migration /go/bin/goose postgres "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"  up

migrate-down:
	docker-compose -f ./docker/docker-compose.yml  --env-file ./config/config.env run migration /go/bin/goose postgres "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"  down


testi: build
	docker-compose -f ./docker/docker-compose.integr.yml --env-file ./config/config_test.env up -d --build;\
	test_status_code=0 ;\
	docker-compose -f ./docker/docker-compose.integr.yml --env-file ./config/config_test.env run integr go test -race ./test... || test_status_code=$$? ;\
	docker-compose -f ./docker/docker-compose.integr.yml --env-file ./config/config_test.env down --remove-orphans;\
	exit $$test_status_code ;

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

.PHONY: build run test lint

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

config-copy:
	cp -n ./config/config_example.env ./config/config.env
