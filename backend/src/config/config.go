package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//String de conexão com o mysql
	StringConnectionDB = ""
	//Porta de conexão com a API
	Port = 0
	//SecretKey used to generate token
	SecretKey []byte
)

func Load() {
	var erro error
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	//por padrão as informações no env vem de string, então tem que converter o valor da porta
	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 8000
	}
	StringConnectionDB = fmt.Sprintf("postgresql://%s:%s@localhost:%s/%s?sslmode=disable", os.Getenv("DB_USUARIO"), os.Getenv("DB_SENHA"), os.Getenv("DB_PORT"), os.Getenv("DB_NOME"))

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
