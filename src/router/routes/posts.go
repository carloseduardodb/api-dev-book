package routes

import "net/http"

var postRoutes = []Route{
	{
		URI:          "/post",
		Method:       http.MethodPost,
		Function:     controllers.CreatePost,
		RequiredAuth: false,
	},
	{
		URI:          "/posts",
		Method:       http.MethodGet,
		Function:     controllers.GetPosts,
		RequiredAuth: true,
	},
	{
		URI:          "/post/{id}",
		Method:       http.MethodGet,
		Function:     controllers.GetPost,
		RequiredAuth: true,
	},
	{
		URI:          "/post/{id}",
		Method:       http.MethodPut,
		Function:     controllers.UpdatePost,
		RequiredAuth: true,
	},
	{
		URI:          "/post/{id}",
		Method:       http.MethodDelete,
		Function:     controllers.DeletePost,
		RequiredAuth: true,
	},
	{
		URI:          "/post/{id}/like",
		Method:       http.MethodPost,
		Function:     controllers.LikePost,
		RequiredAuth: true,
	},
	{
		URI:          "/post/{id}/unlike",
		Method:       http.MethodPost,
		Function:     controllers.UnlikePost,
		RequiredAuth: true,
	},
}
