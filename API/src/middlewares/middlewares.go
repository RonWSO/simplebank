package middlewares

import (
	"api/src/answer"
	"api/src/authentication"
	"log"
	"net/http"
)

// Function to write request on logger, configurate later
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Check if the route needs authentication and validate token
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	//Retorna a função recebida após checar se o token na requisição é válido ele executa a função.
	return func(w http.ResponseWriter, r *http.Request) {
		if _, erro := authentication.ValidarToken(r); erro != nil {
			answer.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r)
	}
}
