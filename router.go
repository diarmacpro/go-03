package main

import (
	"app/handlers"
	"net/http"
)

func setupRoutes() {
	http.HandleFunc("/static/", handlers.StaticFileHandler)
	http.HandleFunc("/api", handlers.APIHandler)
	http.HandleFunc("/base", handlers.HelloHandler)
	http.HandleFunc("/", handlers.IndexHandler)
}
