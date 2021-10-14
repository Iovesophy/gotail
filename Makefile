IMAGE_NAME := gotail
TEST_FUNC := ''
TEST_TAGS := 'unit_test'

.PHONY: all
all: docker-build docker-run

.PHONY: docker-build
docker-build:
	docker build -t $(IMAGE_NAME) .

.PHONY: docker-run
docker-run:
	docker run --rm $(IMAGE_NAME)

.PHONY: test
test:
	docker run -e GO111MODULE=auto --rm -v $(PWD):/app golang:1.17 bash -c \
		"cd /app && go test -tags=$(TEST_TAGS) -coverprofile=cover.out -run $(TEST_FUNC) -v && go tool cover -html=cover.out -o cover.html"

.PHONY: clean
clean:
	rm cover.*
