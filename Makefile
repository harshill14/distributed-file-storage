.PHONY: all proto build up down

all: proto build

proto:
	cd proto && ./generate.sh

build:
	docker-compose build

up:
	docker-compose up

down:
	docker-compose down
