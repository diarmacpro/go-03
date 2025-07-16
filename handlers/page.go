package handlers

import (
	"fmt"
	"net/http"
)

// /base
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Halo dari net/http!")
}

// /
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "static/index.html")
	} else {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "static/404.html")
	}
}
