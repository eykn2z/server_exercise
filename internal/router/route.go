package router

import (
	"net/http"
	"server/internal/handler"

	"github.com/labstack/echo"
)

type Route struct {
	Name            string
	Method          string
	Path            string
	HandlerFunc     http.HandlerFunc
	EchoHandlerFunc echo.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{Method: "GET", Path: "/persons", HandlerFunc: handler.GetPersons, EchoHandlerFunc: handler.GetEchoPersons},
}
