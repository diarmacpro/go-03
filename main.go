package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// Serve file statis
	http.HandleFunc("/static/", staticFileHandler)

	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Route paling spesifik ditaruh dulu
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/base", helloHandler)

	// Route paling umum di akhir (fallback)
	http.HandleFunc("/", indexHandler)

	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func staticFileHandler(w http.ResponseWriter, r *http.Request) {
	relPath := r.URL.Path[len("/static/"):]
	fullPath := "static/" + relPath

	// Coba buka file atau folder
	f, err := http.Dir("static").Open(relPath)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "static/404.html")
		return
	}

	// Jika direktori → blokir
	if fi.IsDir() {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Serve file jika valid
	http.ServeFile(w, r, fullPath)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "static/index.html")
	} else {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "static/404.html")
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Halo dari net/http!")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method tidak diizinkan", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Halo dari API"})
}
