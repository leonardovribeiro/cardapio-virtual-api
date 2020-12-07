package routes

import (
	"cardapio-virtual-api/src/controllers"
	"net/http"
)

var customersRoutes = []Route{
	{
		URI:          "/customers",
		Method:       http.MethodPost,
		Function:     controllers.CreateCustomer,
		RequiredAuth: false,
	},
	{
		URI:          "/customers",
		Method:       http.MethodGet,
		Function:     controllers.FetchCustomer,
		RequiredAuth: false,
	},
	{
		URI:          "/customers/{id}",
		Method:       http.MethodGet,
		Function:     controllers.GetCustomerByID,
		RequiredAuth: false,
	},
	{
		URI:          "/customers/{id}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateCustomer,
		RequiredAuth: false,
	},
	{
		URI:          "/customers/{id}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteCustomer,
		RequiredAuth: false,
	},
}
