package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// 1 - Instalar mux para fazer o Router
func main() {
	config.Load()
	fmt.Printf("Rodando API na porta %d", config.Port)
	r := router.Generate()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
