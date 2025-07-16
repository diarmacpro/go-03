package main

import (
	"app/handlers"
	"app/ws"
	"net/http"
)

func setupRoutes() {
	// Redirect "/static" → "/static/"
	http.HandleFunc("/static", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/", http.StatusMovedPermanently)
	})

	// Serve static files
	http.HandleFunc("/static/", handlers.StaticFileHandler)

	// WebSocket endpoints
	http.HandleFunc("/ws", handlers.WSPage)
	http.HandleFunc("/ws/room", ws.RoomHandler)

	// API & base
	http.HandleFunc("/api", handlers.APIHandler)
	http.HandleFunc("/base", handlers.HelloHandler)

	// Fallback untuk "/"
	http.HandleFunc("/", handlers.IndexHandler)
}
