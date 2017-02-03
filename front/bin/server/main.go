package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	hs := hostSwitch{}

	hs["default"] = defaultRouter()
	hs["users.localhost:3000"] = wrapProxy(users())
	hs["authors.localhost:3000"] = wrapProxy(authors())

	log.Fatal(http.ListenAndServe("0.0.0.0:3000", hs))
}

type hostSwitch map[string]http.Handler

func (hs hostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host)
	if handler := hs[r.Host]; handler != nil {
		handler.ServeHTTP(w, r)
	} else {
		hs["default"].ServeHTTP(w, r)
	}
}

func wrapProxy(proxy http.Handler) *gin.Engine {
	http := gin.Default()

	http.Use(func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
		c.Abort()
	})

	return http
}

func defaultRouter() *gin.Engine {
	http := gin.Default()

	http.NoRoute(func(c *gin.Context) {
		c.JSON(400, gin.H{"error": "Bad request, check the docs ;)"})
	})

	return http
}

func users() *httputil.ReverseProxy {
	url, _ := url.Parse("http://localhost:3001")
	return httputil.NewSingleHostReverseProxy(url)
}

func authors() *httputil.ReverseProxy {
	url, _ := url.Parse("http://localhost:3002")
	return httputil.NewSingleHostReverseProxy(url)
}
