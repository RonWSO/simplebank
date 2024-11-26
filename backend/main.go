package main

import (
	"api/api"
	db "api/db/sqlc"
	"api/src/config"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver     = "postgres"
	dbSource     = "postgresql://ronaldoliveira:mypostgrespassword@localhost:5432/simple_bank?sslmode=disable"
	serverAddres = "0.0.0.0:8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to DB:", err)
	}
	config.Load()
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddres)
	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
