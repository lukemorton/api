all: start

test:
	cd ./users && make test
	cd ./authors && make test
	cd ./front && make test

build:
	cd ./users && make build
	cd ./authors && make build
	cd ./front && make build

compose:
	docker-compose build
	docker-compose up

start: test build compose
