.PHONY: all

all: test docker

test:
	go test ./...

run:
	docker-compose up -d

down:
	docker-compose down

build-app:
	cd app && yarn && yarn web:build && mv dist ../build

docker: docker-init docker-mysql docker-review360

docker-init:
	docker build --pull \
	--file Dockerfile.init \
	-t review360-init:latest .

docker-review360:
	docker build --pull \
	--file Dockerfile \
	-t review360:latest .

docker-mysql:
	docker build --pull \
	--file Dockerfile.mysql \
	-t review360-mysql:latest .