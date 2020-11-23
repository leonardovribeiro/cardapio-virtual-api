package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route representa uma rota
type Route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	Auth     bool
}

// Configure coloca todas as rotas dentro do router
func Configure(r *mux.Router) *mux.Router {
	routes := customersRoutes
	routes = append(routes, routeCustomersLogin)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
