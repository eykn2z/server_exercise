package handler

import (
	"log"
	"net/http"
)

func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Access: [%s] %s\n",r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}