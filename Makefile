.PHONY: re down run all test
re:
	docker-compose -f ./.docker/docker-compose.yml up --build -d
down:
	docker-compose -f ./.docker/docker-compose.yml down -v
run:
	docker-compose -f ./.docker/docker-compose.yml up -d
test:
	go run ./.docker/pub_script.go
all:	re

