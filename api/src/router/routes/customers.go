package routes

import (
	"cardapio-virtual-api/src/controllers"
	"net/http"
)

var customersRoutes = []Route{
	{
		URI:      "/customers",
		Method:   http.MethodPost,
		Function: controllers.CreateCustomer,
		Auth:     false,
	},
	{
		URI:      "/customers",
		Method:   http.MethodGet,
		Function: controllers.FindAllCustomer,
		Auth:     false,
	},
	{
		URI:      "/customers/{id}",
		Method:   http.MethodGet,
		Function: controllers.FindCustomer,
		Auth:     false,
	},
	{
		URI:      "/customers/{customerID}",
		Method:   http.MethodPut,
		Function: controllers.UpdateCustomer,
		Auth:     false,
	},
	{
		URI:      "/customer/{customerID}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteCustomer,
		Auth:     false,
	},
}
