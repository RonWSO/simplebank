start-project:
	go mod init api
	make install-dependencies

install-dependencies:
	go get github.com/joho/godotenv
	go get github.com/gorilla/mux
	go get github.com/dgrijalva/jwt-go
	go get github.com/golang-migrate/migrate/v4
	go get github.com/sqlc-dev/sqlc/cmd/sqlc@latest

createdb:
	docker exec -it postgres12 createdb -U ronaldoliveira --owner=ronaldoliveira simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=ronaldoliveira -e POSTGRES_PASSWORD=mypostgrespassword -d postgres:12-alpine

migrateup:
	cd API && migrate -path src/db/migration -database "postgresql://ronaldoliveira:mypostgrespassword@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	cd API && migrate -path src/db/migration -database "postgresql://ronaldoliveira:mypostgrespassword@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	cd API && sqlc generate

test:
	cd API && go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc-run