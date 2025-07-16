package handlers

import (
	"net/http"
)

func StaticFileHandler(w http.ResponseWriter, r *http.Request) {
	relPath := r.URL.Path[len("/static/"):]
	fullPath := "static/" + relPath

	f, err := http.Dir("static").Open(relPath)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "static/404.html")
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil || fi.IsDir() {
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "static/404.html")
		return
	}

	http.ServeFile(w, r, fullPath)
}
