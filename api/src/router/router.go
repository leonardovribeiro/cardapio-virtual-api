package router

import (
	"cardapio-virtual-api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate retorna um router com as rotas configuradas
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
