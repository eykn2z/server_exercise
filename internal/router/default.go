package router

import (
	"net/http"
	"server/internal/handler"
)

func getRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/persons", handler.GetPersons)
	return router
}
