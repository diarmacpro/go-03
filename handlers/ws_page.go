package handlers

import (
	"net/http"
)

// WSPage menyajikan halaman WebSocket UI
func WSPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/ws/index.html")
}
