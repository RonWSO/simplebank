package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://ronaldoliveira:mypostgrespassword@localhost:5432/simple_bank?sslmode=disable"
)

// var dbSource = fmt.Sprintf("postgresql://%s:%s@localhost:%s/%s?sslmode=disable", os.Getenv("DB_USUARIO"), os.Getenv("DB_SENHA"), os.Getenv("DB_PORT"), os.Getenv("DB_NOME"))
var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the database", err)
	}

	// Cria um contexto
	ctx := context.Background()
	fmt.Println(testDB)
	testQueries = New(testDB)

	// Você pode usar o contexto nas chamadas subsequentes
	err = testDB.PingContext(ctx) // Apenas para verificar a conexão com contexto
	if err != nil {
		log.Fatal("cannot ping the database", err)
	}

	os.Exit(m.Run())
}
