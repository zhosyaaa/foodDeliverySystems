proto:
	protoc ./internal/protos/*.proto --go_out=. --go-grpc_out=.

server:
	export ENVIRONMENT=DEVELOPMENT;\
	go run cmd/main.go

postgres:
	docker run --rm -d --name orders -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1079 -p 8000:8000 postgres:15

createdb:
	docker exec -it orders createdb --username=postgres --owner=postgres orders

dropdb:
	docker exec -it orders dropdb -U postgres orders


build:
	CGO_ENABLED=0 go build -o bin/main cmd/main.go

.PHONY: proto server build