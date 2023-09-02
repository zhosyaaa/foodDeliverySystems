proto:
	protoc ./internal/protos/*.proto --go_out=. --go-grpc_out=.

server:
	export ENVIRONMENT=DEVELOPMENT;\
	go run cmd/main.go

postgres:
	docker run --rm -d --name restSelection -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1079 -p 8008:8008 postgres:15

createdb:
	docker exec -it restSelection createdb --username=postgres --owner=postgres restSelection

dropdb:
	docker exec -it restSelection dropdb -U postgres restSelection

build:
	CGO_ENABLED=0 go build -o bin/main cmd/main.go

.PHONY: proto server build