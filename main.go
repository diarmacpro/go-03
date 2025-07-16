package main

import (
	"fmt"
	"log"
	"net/http"
)

// func main() {
// 	// Inisialisasi semua routing
// 	setupRoutes()

// 	fmt.Println("Server berjalan di http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// }

// func main() {
// 	setupRoutes()

// 	fmt.Println("🔒 Server HTTPS berjalan di https://localhost:8443")
// 	err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
// 	if err != nil {
// 		log.Fatal("Gagal menjalankan server HTTPS:", err)
// 	}
// }

func main() {
	setupRoutes()

	fmt.Println("🌐 Server berjalan di http://0.0.0.0:8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil) // Tanpa TLS
	if err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}
