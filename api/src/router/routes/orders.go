package routes

var ordersRoutes = []Route {
	{
		URI:          "/orders",
		Method:       http.MethodPost,
		Function:     controllers.CreateOrder,
		RequiredAuth: false,
	},
	{
		URI:          "/orders",
		Method:       http.MethodGet,
		Function:     controllers.FetchOrder,
		RequiredAuth: false,
	},
	{
		URI:          "/orders/{id}",
		Method:       http.MethodGet,
		Function:     controllers.GetOrderByID,
		RequiredAuth: false,
	},
	{
		URI:          "/orders/{id}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateOrder,
		RequiredAuth: false,
	},
	{
		URI:          "/orders/{id}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteOrder,
		RequiredAuth: false,
	},