package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test APIHandler response
func TestAPIHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/api", nil)
	w := httptest.NewRecorder()

	APIHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code harus 200, dapat %d", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		t.Errorf("Content-Type harus application/json, dapat %s", contentType)
	}
}

// Test HelloHandler untuk /base
func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/base", nil)
	w := httptest.NewRecorder()

	HelloHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code harus 200, dapat %d", res.StatusCode)
	}
}

// Test index root
func TestIndexHandlerRoot(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	IndexHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code root '/' harus 200, dapat %d", res.StatusCode)
	}
}

// Test index not found
func TestIndexHandlerNotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/tidak-ada", nil)
	w := httptest.NewRecorder()

	IndexHandler(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Status code harus 404, dapat %d", res.StatusCode)
	}
}
