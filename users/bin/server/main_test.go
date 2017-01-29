package main

import (
	"bytes"
	"encoding/json"
	"github.com/lukemorton/api/users"
	"github.com/stretchr/testify/assert"
	"gopkg.in/gin-gonic/gin.v1"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	w := GET("/status.json")
	assert.Equal(t, w.Code, 200, "status should be 200")
}

func TestRegister(t *testing.T) {
	w := POST("/register.json", h{
		"email":    "lukemorton.dev@gmail.com",
		"password": "bob",
	})
	user := userFromResponse(w)
	assert.Equal(t, w.Code, 200, "status should be 200")
	assert.Equal(t, int64(1), user.Id, "includes ID")
	assert.Equal(t, "lukemorton.dev@gmail.com", user.Email, "includes email")
}

func TestVerify(t *testing.T) {
	app := testApp()

	app.POST("/register.json", h{
		"email":    "lukemorton.dev@gmail.com",
		"password": "bob",
	})

	w := app.POST("/verify.json", h{
		"email":    "lukemorton.dev@gmail.com",
		"password": "bob",
	})
	user := userFromResponse(w)
	assert.Equal(t, w.Code, 200, "status should be 200")
	assert.Equal(t, int64(1), user.Id, "includes ID")
}

func TestVerifyError(t *testing.T) {
	w := POST("/verify.json", h{
		"email":    "lukemorton.dev@gmail.com",
		"password": "bob",
	})
	assert.Equal(t, w.Code, 401, "status should be 401")
}

func TestResetPassword(t *testing.T) {
	app := testApp()

	app.POST("/register.json", h{
		"email":    "lukemorton.dev@gmail.com",
		"password": "bob",
	})

	w := app.POST("/password/reset.json", h{
		"email": "lukemorton.dev@gmail.com",
	})
	assert.Equal(t, w.Code, 200, "status should be 200")
}

func TestResetPasswordError(t *testing.T) {
	w := POST("/password/reset.json", h{
		"email": "lukemorton.dev@gmail.com",
	})
	assert.Equal(t, w.Code, 422, "status should be 422")
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
	return testApp().GET(path)
}

func POST(path string, body interface{}) *httptest.ResponseRecorder {
	return testApp().POST(path, body)
}

func testApp() testAppAgent {
	return testAppAgent{HTTP()}
}

type testAppAgent struct {
	engine *gin.Engine
}

func (app testAppAgent) GET(path string) *httptest.ResponseRecorder {
	return app.request("GET", path, nil)
}

func (app testAppAgent) POST(path string, body interface{}) *httptest.ResponseRecorder {
	jsonBody, _ := json.Marshal(body)
	return app.request("POST", path, bytes.NewBuffer(jsonBody))
}

type h map[string]interface{}

func (app testAppAgent) request(method string, path string, body io.Reader) *httptest.ResponseRecorder {
	r, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	app.engine.ServeHTTP(w, r)
	return w
}

func userFromResponse(w *httptest.ResponseRecorder) users.User {
	var user *users.User
	err := json.Unmarshal(w.Body.Bytes(), &user)

	if err != nil {
		log.Fatal(err)
	}

	return *user
}
