package router

import (
	"github.com/gorilla/mux"
)

func NewMuxRouter() *mux.Router {
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
