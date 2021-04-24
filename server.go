package main

import (
	"log"
	"net/http"
	"server/handler"
	"server/router"
)

const serverType = "echo" // "net/http" "mux"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	port := ":8080"
	log.Printf("[START] server port: %s\n", port)

	if serverType != "echo" {
		if err := http.ListenAndServe(port, handler.Log(router.NewRouter())); err != nil {
			log.Fatalln("[FAILED] start sever error: %v", err)
		}
	} else {
		echoRouter := router.NewEchoRouter()
		echoRouter.Logger.Fatal(echoRouter.Start(port))
	}
}
