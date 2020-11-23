package routes

import (
	"cardapio-virtual-api/src/controllers"
	"net/http"
)

var routeCustomersLogin = Route{
	URI:      "/login",
	Method:   http.MethodPost,
	Function: controllers.Login,
	Auth:     false,
}
