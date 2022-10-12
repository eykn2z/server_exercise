package router

import "net/http"

func NewNetHttpRouter() *http.ServeMux {
	router := http.NewServeMux()
	for _, route := range routes {
		router.HandleFunc(route.Path, route.HandlerFunc)
	}
	return router
}
