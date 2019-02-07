SERVICE=backend	

run: up

build-deps:
	$(MAKE) -C ./$(SERVICE) MAKEFLAGS=build-deps

push:
	docker-compose push

build:
	docker-compose build

build-with-deps: build-deps
	docker-compose build --no-cache

start: up

up:
	docker-compose up -d

stop: down

down:
	docker-compose down

restart:
	docker-compose restart

rm:
	docker-compose rm -f

log: logs

logf:
	docker-compose logs --follow

logs:
	docker-compose logs

env:
	cp .env.example .env

envs:
	docker-compose run $(SERVICE) env

enter:
	docker-compose run $(SERVICE) /bin/sh

test-docker:
	docker-compose run $(SERVICE) /bin/sh -c "go test ./test/..."

test-local:
	$(MAKE) -C ./$(SERVICE) MAKEFLAGS=test-local
