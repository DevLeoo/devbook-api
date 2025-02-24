package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginRoute = Routes{
	Uri:                "/login",
	Method:             http.MethodPost,
	Fn:                 controllers.Login,
	NeedAuthentication: false,
}
