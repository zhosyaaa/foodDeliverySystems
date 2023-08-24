proto:
	protoc ./pkg/proto/auth.proto --go_out=. --go-grpc_out=.

server:
	export ENVIRONMENT=DEVELOPMENT;\
	go run cmd/main.go

build:
	CGO_ENABLED=0 go build -o bin/main cmd/main.go

postgres:
	docker run --rm -d --name postgres15 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1079 -p 5432:5432 postgres:15

createdb:
	docker exec -it postgres15 createdb --username=postgres --owner=postgres food_delivery_auth

dropdb:
	docker exec -it postgres15 dropdb -U postgres food_delivery_auth

.PHONY: proto server build