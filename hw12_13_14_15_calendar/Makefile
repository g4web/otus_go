CALENDAR := "calendar"
SCHEDULER := "scheduler"
SENDER := "sender"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o ./bin/$(CALENDAR) -ldflags "$(LDFLAGS)" ./cmd/$(CALENDAR)
	go build -v -o ./bin/$(SCHEDULER) -ldflags "$(LDFLAGS)" ./cmd/$(SCHEDULER)
	go build -v -o ./bin/$(SENDER) -ldflags "$(LDFLAGS)" ./cmd/$(SENDER)

run: build docker-up migrate-up

test:
	go test -race ./...

fmt:
	gofmt -s -w .
	gofumpt -l -w .

install-lint-deps:
	(which golangci-lint > /dev/null) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.41.1

lint: install-lint-deps
	golangci-lint run ./...

.PHONY: build run test lint

generate:
	protoc ./api/eventService.proto --go_out=./app/calendar/server/grpc/protobuf --go-grpc_out=./app/calendar/server/grpc/protobuf


docker-up:
	docker-compose -f ./docker/docker-compose.yml up -d
docker-down:
	docker-compose -f ./docker/docker-compose.yml down --remove-orphans
docker-build:
	docker-compose -f ./docker/docker-compose.yml up -d --build

# for example `make migrate-create table=event`
migrate-create:
	migrate create -ext sql -dir ./internal/storage/sql/migrations $(table)
migrate-up:
	migrate -source file://internal/storage/sql/migrations -database postgres://otus_calendar_user:otus_calendar_password@localhost:5432/otus_calendar?sslmode=disable up
migrate-down:
	migrate -source file://internal/storage/sql/migrations -database postgres://otus_calendar_user:otus_calendar_password@localhost:5432/otus_calendar?sslmode=disable down 1
