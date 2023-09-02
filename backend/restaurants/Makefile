proto:
	protoc ./internal/protos/*.proto --go_out=. --go-grpc_out=.

server:
	export ENVIRONMENT=DEVELOPMENT;\
	go run cmd/main.go

postgres:
	docker run --rm -d --name restaurants -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1079 -p 5432:5432 postgres:15

createdb:
	docker exec -it postgres15 createdb --username=postgres --owner=postgres restaurants

dropdb:
	docker exec -it postgres15 dropdb -U postgres restaurants
build:
	CGO_ENABLED=0 go build -o bin/main cmd/main.go

.PHONY: proto server build