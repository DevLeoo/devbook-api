package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoute = []Routes{
	{
		Uri:                "/users",
		Method:             http.MethodPost,
		Fn:                 controllers.CreateUser,
		NeedAuthentication: false,
	},
	{
		Uri:                "/users",
		Method:             http.MethodGet,
		Fn:                 controllers.GetUsers,
		NeedAuthentication: false,
	},
	{
		Uri:                "/users/{id}",
		Method:             http.MethodGet,
		Fn:                 controllers.GetUser,
		NeedAuthentication: false,
	},
	{
		Uri:                "/users/{id}",
		Method:             http.MethodPut,
		Fn:                 controllers.UpdateUser,
		NeedAuthentication: false,
	},
	{
		Uri:                "/users/{id}",
		Method:             http.MethodDelete,
		Fn:                 controllers.DeleteUser,
		NeedAuthentication: false,
	},
	{
		Uri:                "/users/{followedId}/follow",
		Method:             http.MethodPost,
		Fn:                 controllers.FollowUser,
		NeedAuthentication: true,
	},
	{
		Uri:                "/users/{followedId}/stop-follow",
		Method:             http.MethodPost,
		Fn:                 controllers.UnfollowUser,
		NeedAuthentication: true,
	},
	{
		Uri:                "/users/{followedId}/followers",
		Method:             http.MethodGet,
		Fn:                 controllers.SearchFollowers,
		NeedAuthentication: true,
	},
	{
		Uri:                "/users/{followedId}/following",
		Method:             http.MethodGet,
		Fn:                 controllers.SearchFollowing,
		NeedAuthentication: true,
	},
	{
		Uri:                "/users/{userId}/update-password",
		Method:             http.MethodPost,
		Fn:                 controllers.UpdatePassword,
		NeedAuthentication: true,
	},
}
