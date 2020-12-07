package routes

import (
	"cardapio-virtual-api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route representa uma rota
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	RequiredAuth bool
}

// Configure coloca todas as rotas dentro do router
func Configure(r *mux.Router) *mux.Router {
	routes := customersRoutes
	routes = append(routes, loginRoutes...)
	routes = append(routes, productsRoutes...)

	for _, route := range routes {
		if route.RequiredAuth {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
