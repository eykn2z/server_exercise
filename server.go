package main

import (
	"log"
	"net/http"
	"server/handler"
)


func getRouter() *http.ServeMux{
	router := http.NewServeMux()
	router.HandleFunc("/persons", handler.GetPersonsHandler)
	return router
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	port := ":8080"
	log.Printf("[START] server port: %s\n", port)
	if err := http.ListenAndServe(port, handler.Log(getRouter())); err != nil {
		log.Fatalln("[FAILED] start sever error: %v", err)
	}
}