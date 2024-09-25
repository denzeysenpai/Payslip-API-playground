package routes

import (
	"gomysql/models"

	"github.com/gorilla/mux"
)

// These routes are just for views, not requests etc...
var Routes []models.Route = []models.Route{ // so this doesn't get exported if you don't capitalize the first letter of the field
	{
		Path:   "/", // home
		View:   "views/home.html",
		Method: "GET",
	},
	{
		Path:   "/about",
		View:   "views/about.html",
		Method: "GET",
	},
}

// Attaches the routes for serving the static html pages
func AttachRoutes(router *mux.Router) {
	for _, route := range Routes {
		router.HandleFunc(route.Path, CreateView(route.View).Handler).Methods(route.Method)
	}
}
