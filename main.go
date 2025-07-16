package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Inisialisasi semua routing
	setupRoutes()

	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
