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
	{
		URI:          "/user/{id}/follow",
		Method:       http.MethodPost,
		Function:     controllers.FollowUser,
		RequiredAuth: true,
	},
	{
		URI:          "/user/{id}/unfollow",
		Method:       http.MethodPost,
		Function:     controllers.UnfollowUser,
		RequiredAuth: true,
	},
	{
		URI:          "/user/{id}/followers",
		Method:       http.MethodGet,
		Function:     controllers.GetFollowers,
		RequiredAuth: true,
	},
	{
		URI:          "/user/{id}/following",
		Method:       http.MethodGet,
		Function:     controllers.GetFollowing,
		RequiredAuth: true,
	},
	{
		URI:          "/user/update/password",
		Method:       http.MethodPut,
		Function:     controllers.UpdatePassword,
		RequiredAuth: true,
	},
}
