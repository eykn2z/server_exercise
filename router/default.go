package router

import (
	"net/http"
	"server/handler"
)

func getRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/persons", handler.GetPersons)
	return router
}
