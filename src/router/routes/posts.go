package routes

import (
	"api/src/controllers"
	"net/http"
)

var postsRoute = []Routes{
	{
		Uri:                "/posts",
		Method:             http.MethodPost,
		Fn:                 controllers.CreatePost,
		NeedAuthentication: true,
	},
	{
		Uri:                "/posts",
		Method:             http.MethodGet,
		Fn:                 controllers.GetPosts,
		NeedAuthentication: true,
	},
	{
		Uri:                "/posts/{id}",
		Method:             http.MethodGet,
		Fn:                 controllers.GetPost,
		NeedAuthentication: true,
	},
	{
		Uri:                "/posts/{id}",
		Method:             http.MethodPut,
		Fn:                 controllers.UpdatePost,
		NeedAuthentication: true,
	},
	{
		Uri:                "/posts/{id}",
		Method:             http.MethodDelete,
		Fn:                 controllers.DeletePost,
		NeedAuthentication: true,
	},
	{
		Uri:                "/users/{userId}/posts",
		Method:             http.MethodGet,
		Fn:                 controllers.FindPostsByUser,
		NeedAuthentication: true,
	},
	{
		Uri:                "/posts/{postId}/like",
		Method:             http.MethodPost,
		Fn:                 controllers.LikePost,
		NeedAuthentication: true,
	},
	{
		Uri:                "/posts/{postId}/unlike",
		Method:             http.MethodPost,
		Fn:                 controllers.UnlikePost,
		NeedAuthentication: true,
	},
}
