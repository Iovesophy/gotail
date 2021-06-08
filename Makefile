IMAGE_NAME := gotail
VERSION := $(godump show -r)
TEST_FUNC := ''

## Integration build and run by docker
.PHONY: all
all: docker-build docker-run

## Setup
.PHONY: deps
deps:
	GO111MODULE=off go get github.com/Songmu/make2help/cmd/make2help

## Docker-build phase
.PHONY: docker-build
docker-build:
	docker build -t $(IMAGE_NAME) .

## Docker-run phase
.PHONY: docker-run
docker-run:
	docker run --rm $(IMAGE_NAME)

## Integration Test
.PHONY: test
test:
	docker run -e GO111MODULE=auto --rm -v $(PWD):/go golang:latest bash -c \
		"go test -tags=integration -coverprofile=cover.out -run $(TEST_FUNC) -v && \
		 go tool cover -html=cover.out -o cover.html"

## Clean cover file
.PHONY: clean
clean:
	rm cover.*

## Show help
.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)

