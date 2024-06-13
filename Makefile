.PHONY: re down run all test logs
re:
	docker-compose -f ./.docker/docker-compose.yml up --build -d
down:
	docker-compose -f ./.docker/docker-compose.yml down -v
run:
	docker-compose -f ./.docker/docker-compose.yml up -d
pub:
	go run ./.docker/pub_script.go
logs:
	docker-compose -f ./.docker/docker-compose.yml logs -f
all:	re

