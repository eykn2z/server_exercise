package router

import (
	"net/http"
	"server/handler"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			//Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{Method: "GET", Path: "/persons", HandlerFunc: handler.GetPersons},
}
