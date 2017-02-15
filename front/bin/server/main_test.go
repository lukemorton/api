package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	w := GET("/status.json")
	assert.Equal(t, 200, w.Code, "status should be 200")
}

func TestVerify(t *testing.T) {
	w := POST("http://users/verify.json", h{
		"email":    "lukemorton.dev@gmail.com",
		"password": "bob",
	})
	assert.Equal(t, 200, w.Code, "status should be 200")
}

func TestBadRequest(t *testing.T) {
	var w *httptest.ResponseRecorder

	w = GET("/")
	assert.Equal(t, 400, w.Code, "status should be 400")

	w = POST("/", nil)
	assert.Equal(t, 400, w.Code, "status should be 400")

	w = GET("/nope")
	assert.Equal(t, 400, w.Code, "status should be 400")
}

func GET(path string) *httptest.ResponseRecorder {
	return request(testApp(), "GET", path, nil)
}

func POST(path string, body interface{}) *httptest.ResponseRecorder {
	jsonBody, _ := json.Marshal(body)
	return request(testApp(), "POST", path, bytes.NewBuffer(jsonBody))
}

func request(app hostSwitch, method string, path string, body io.Reader) *httptest.ResponseRecorder {
	r, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	app.ServeHTTP(w, r)
	return w
}

func testApp() hostSwitch {
	app := hostSwitch{}

	app["default"] = defaultRouter()

	users := http.NewServeMux()
	users.HandleFunc("/verify.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	app["users"] = wrapProxy(users)

	return app
}

type h map[string]interface{}
