include .env

build:
	@go build

run: build
	@./create2-payment-gateway

migrate:
	@migrate -path migrations -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" up

migrate-down:
	@migrate -path migrations -database "mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" down
