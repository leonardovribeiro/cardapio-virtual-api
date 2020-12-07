package routes

import (
	"cardapio-virtual-api/src/controllers"
	"net/http"
)

var productsRoutes = []Route{
	{
		URI:          "/products",
		Method:       http.MethodPost,
		Function:     controllers.CreateProduct,
		RequiredAuth: false,
	},
	{
		URI:          "/products",
		Method:       http.MethodGet,
		Function:     controllers.FetchProduct,
		RequiredAuth: false,
	},
	{
		URI:          "/products/{id}",
		Method:       http.MethodGet,
		Function:     controllers.GetProductByID,
		RequiredAuth: false,
	},
	{
		URI:          "/products/{id}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateProduct,
		RequiredAuth: false,
	},
	{
		URI:          "/products/{id}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteProduct,
		RequiredAuth: false,
	},
}
