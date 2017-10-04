package main

import (
	"net/http"

	"github.com/dougneal/playground/api"
)

var server *http.Server

func main() {
	server = &http.Server{
		Handler:      api.GetRouter(),
		Addr:         "127.0.0.1:8888",
		WriteTimeout: 0,
		ReadTimeout:  0,
	}

	panic(server.ListenAndServe())
}
