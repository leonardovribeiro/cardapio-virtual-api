package routes

import (
	"cardapio-virtual-api/src/controllers"
	"net/http"
)

var loginRoutes = []Route{
	{
		URI:          "/login",
		Method:       http.MethodPost,
		Function:     controllers.Login,
		RequiredAuth: false,
	},
}
