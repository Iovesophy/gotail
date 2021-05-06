NAME := gotail

.PHONY: all
all: docker-build docker-run

.PHONY: docker-build
docker-build:
	docker build -t $(NAME) .

.PHONY: docker-run
docker-run:
	docker run --rm $(NAME)

check: cover.out convert.html open

cover.out:
	go test main_test.go main.go -coverprofile=cover.out

convert.html:
	go tool cover -html=cover.out -o convert.html

.PHONY: open
open:
	open convert.html

.PHONY: clean
clean:
	rm -f cover.out convert.html
