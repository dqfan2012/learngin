
.PHONY: all server migrate-up migrate-down make-migration model clean dep-tidy gofumpt lint lint-go build-docker seed setup-dev update-dep vendor

BIN             := learngin
PKG             := github.com/dqfan2012/$(BIN)
LOCAL_VERSION   := $(shell git describe --tags --always --dirty)
TEST_IMAGE_NAME ?= $(BIN)-test:$(LOCAL_VERSION)
GO_VERSION      := 1.22
GO_CONTAINER    := arm64v8/golang:$(GO_VERSION)

# Default target
all: server

# Build Docker image for linting and testing
build-docker:
	docker build -t $(TEST_IMAGE_NAME) -f build/Dockerfile .

# Clean any intermediate temp files
clean:
	rm -rf vendor/
	rm -rf artifacts/

# Remove unused dependencies and add missing dependencies
dep-tidy: build-docker
	go mod tidy -go=$(GO_VERSION)

gofumpt: build-docker
	docker run --rm -u `id -u` -v $(shell pwd):/app -w /app $(TEST_IMAGE_NAME) sh -c "gofumpt -l -w ."

lint: gofumpt lint-go

lint-go: build-docker
	docker run --rm -v $(shell pwd):/app -w /app $(TEST_IMAGE_NAME) sh -c 'golangci-lint run --fix'

# Apply all up migrations
migrate-up:
	go run ./cmd/migrate/main.go -up

# Rollback all migrations
migrate-down:
	go run ./cmd/migrate/main.go -down

# Create new migration files
migration:
	go run ./cmd/migrate/main.go -new $(name)

# Create a new model
# I.E., make model name=Blog/Post
model:
	go run ./cmd/model/main.go -name=$(name)

# Seed the database
seed:
	go run ./cmd/seeder/main.go

# Run the web server
server:
	go run ./cmd/server/main.go

# Automatically rollback migrations, rerun migrations, seed the db, and run the server
setup-dev: migrate-down migrate-up seed server

## Update specific dependencies. Example usage: make update-dep dependency=golang.org/x/crypto/ ...
update-dep: build-docker
	docker run --rm -u `id -u` -v $(shell go env GOPATH)/pkg/mod:/go/pkg/mod -v $(shell pwd):/app -w /app $(TEST_IMAGE_NAME) sh -c "go get -u $(dependency)"

vendor:  ## Update vendor packages and lock file
	go mod vendor
