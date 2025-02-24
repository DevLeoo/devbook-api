package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	Uri                string
	Method             string
	Fn                 func(http.ResponseWriter, *http.Request)
	NeedAuthentication bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := usersRoute
	routes = append(routes, loginRoute)
	routes = append(routes, postsRoute...)

	for _, route := range routes {
		if route.NeedAuthentication {
			r.HandleFunc(route.Uri,
				middlewares.Logger(middlewares.Authenticate(route.Fn)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri, route.Fn).Methods(route.Method)

		}

	}

	return r
}
