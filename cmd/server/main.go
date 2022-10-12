package main

import (
	"log"
	"net/http"
	"server/internal/handler"
	"server/internal/router"
)

type ServerType string

const (
	SERVER_TYPE_NET_HTTP ServerType = "net/http"
	SERVER_TYPE_MUX      ServerType = "net/http with mux"
	SERVER_TYPE_ECHO     ServerType = "echo"

	serverType = SERVER_TYPE_NET_HTTP
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	port := ":8080"
	log.Printf("[START] server port: %s with: %s", port, serverType)

	switch serverType {
	case SERVER_TYPE_NET_HTTP:
		//TODO: logの出し方
		router.SetRouter()
		err := http.ListenAndServe(port, nil)
		if err != nil {
			log.Fatal(err)
		}
	case SERVER_TYPE_MUX:
		if err := http.ListenAndServe(port, handler.Log(router.NewMuxRouter())); err != nil {
			log.Fatalln("[FAILED] start sever error: %v", err)
		}
	case SERVER_TYPE_ECHO:
		echoRouter := router.NewEchoRouter()
		echoRouter.Logger.Fatal(echoRouter.Start(port))
	}
}
