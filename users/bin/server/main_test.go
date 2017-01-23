package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	w := GET("/status.json")
	assert.Equal(t, w.Code, 200, "status should be 200")
}

func TestRegister(t *testing.T) {
	user := []byte(`{"email":"lukemorton.dev@gmail.com"}`)
	w := POST("/register.json", user)
	assert.Contains(t, w.Body.String(), `"email":"lukemorton.dev@gmail.com"`, "includes email")
}

func TestBadRequest(t *testing.T) {
	var w *httptest.ResponseRecorder

	w = GET("/")
	assert.Equal(t, w.Code, 400, "status should be 400")

	w = POST("/", nil)
	assert.Equal(t, w.Code, 400, "status should be 400")

	w = GET("/nope")
	assert.Equal(t, w.Code, 400, "status should be 400")
}

func GET(path string) *httptest.ResponseRecorder {
	return request("GET", path, nil)
}

func POST(path string, body []byte) *httptest.ResponseRecorder {
	return request("POST", path, bytes.NewBuffer(body))
}

func request(method string, path string, body io.Reader) *httptest.ResponseRecorder {
	r, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	AppEngine().ServeHTTP(w, r)
	return w
}
