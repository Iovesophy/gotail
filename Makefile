NAME := gotail
NAME-C := gotailcheck

.PHONY: all
all: docker-build docker-run

.PHONY: docker-build
docker-build:
	docker build -t $(NAME) .

.PHONY: docker-run
docker-run:
	docker run -t $(NAME)

check: docker-build docker-run-covercheck open clean

docker-run-covercheck:
	docker run --rm --name $(NAME-C) -itd $(NAME) /bin/sh
	docker cp $(NAME-C):/root/convert.html $(shell pwd)
	docker stop $(NAME-C)
	cp convert.html $(shell pwd)/log/$(shell date +"%Y%m%d%I%M%S").html

.PHONY: open
open:
	open convert.html

.PHONY: clean
clean:
	sleep 10
	rm -f cover.out convert.html
