proto:
	protoc ./internal/protos/*.proto --go_out=. --go-grpc_out=.

server:
	export ENVIRONMENT=DEVELOPMENT;\
	go run cmd/main.go

postgres:
	docker run --rm -d --name payment -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1079 -p 6543:6543 postgres:15

createdb:
	docker exec -it payment createdb --username=postgres --owner=postgres payment

dropdb:
	docker exec -it payment dropdb -U postgres payment

build:
	CGO_ENABLED=0 go build -o bin/main cmd/main.go

.PHONY: proto server build