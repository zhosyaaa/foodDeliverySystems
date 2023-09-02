proto:
	protoc ./internal/protos/*.proto --go_out=. --go-grpc_out=.

server:
	export ENVIRONMENT=DEVELOPMENT;\
	go run cmd/main.go

postgres:
	docker run --rm -d --name fooddeliverysystems-auth -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1079 -p 8080:8080 postgres:15

createdb:
	docker exec -it fooddeliverysystems-auth createdb --username=postgres --owner=postgres auth

dropdb:
	docker exec -it fooddeliverysystems-auth dropdb -U postgres auth
build:
	CGO_ENABLED=0 go build -o bin/main cmd/main.go

.PHONY: proto server build