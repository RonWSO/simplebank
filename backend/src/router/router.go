package router

import (
	"backend/src/router/routes"

	"github.com/gorilla/mux"
)

// Gerar vai criar um router com as rotas configuradas
func Generate() *mux.Router {

	r := mux.NewRouter()
	return routes.Configurar(r)
}
