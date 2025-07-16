package main

import (
	"app/handlers"
	"app/ws"
	"net/http"
)

func setupRoutes() {
	http.HandleFunc("/static/", handlers.StaticFileHandler)
	http.HandleFunc("/api", handlers.APIHandler)
	http.HandleFunc("/base", handlers.HelloHandler)
	http.HandleFunc("/ws/room", ws.RoomHandler)
	http.HandleFunc("/", handlers.IndexHandler)
}
