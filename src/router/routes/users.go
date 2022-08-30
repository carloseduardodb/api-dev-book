package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:          "/user",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		RequiredAuth: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Function:     controllers.GetUsers,
		RequiredAuth: true,
	},
	{
		URI:          "/user/{id}",
		Method:       http.MethodGet,
		Function:     controllers.GetUser,
		RequiredAuth: true,
	},
	{
		URI:          "/user/{id}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		RequiredAuth: true,
	},
	{
		URI:          "/user/{id}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		RequiredAuth: true,
	},
}
