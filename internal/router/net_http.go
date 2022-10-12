package router

import "net/http"

func SetRouter() {
	for _, route := range routes {
		http.HandleFunc(route.Path, route.HandlerFunc)
	}
}
