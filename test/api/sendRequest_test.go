package api_test

import (
	"io"
	"net/http"
	"net/http/httptest"
)

const (
	port    = "8080"
	baseURL = "http://localhost:" + port
)

func sendRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
