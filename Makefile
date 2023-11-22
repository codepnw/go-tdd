include .env

run:
	@go run ./cmd/main.go
test:
	@GOFLAGS="-count=1" go test -v -cover -race ./...


postgresinit:
	docker run -p 5433:5432 --name ${DB_NAME} -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} -d postgres:15.4
postgres:
	docker exec -it ${DB_NAME} psql

createdb:
	docker exec -it ${DB_NAME} createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}
dropdb:
	docker exec -it ${DB_NAME} dropdb ${DB_NAME}