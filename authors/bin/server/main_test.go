package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	w := GET("/status.json")
	assert.Equal(t, w.Code, 200, "status should be 200")
}

func TestAuthors(t *testing.T) {
	w := GET("/authors.json")
	assert.Equal(t, w.Code, 200, "status should be 200")
}

func TestBadRequest(t *testing.T) {
	var w *httptest.ResponseRecorder

	w = GET("/")
	assert.Equal(t, w.Code, 400, "status should be 400")

	w = POST("/")
	assert.Equal(t, w.Code, 400, "status should be 400")

	w = GET("/nope")
	assert.Equal(t, w.Code, 400, "status should be 400")
}

func GET(path string) *httptest.ResponseRecorder {
	return request("GET", path)
}

func POST(path string) *httptest.ResponseRecorder {
	return request("POST", path)
}

func request(method string, path string) *httptest.ResponseRecorder {
	r, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	AppEngine().ServeHTTP(w, r)
	return w
}
