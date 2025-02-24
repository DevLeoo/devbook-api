package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Create will return a new configured router
func Create() *mux.Router {
	router := mux.NewRouter()
	return routes.Configure(router)
}
