NAME := gotail

.PHONY: main
main: docker-build docker-run

.PHONY: docker-build
docker-build:
	docker build -t $(NAME) .

.PHONY: docker-run
docker-run:
	docker run --rm $(NAME)

