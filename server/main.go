package main

import (
	"github.com/lukemorton/api/app"
	"github.com/lukemorton/api/authors"
	"net/http"
)

func main() {
	s := app.NewServer(":3000", http.NewServeMux())

	s.Handle("/", func(r *http.Request) app.Response {
		return app.Error(400, "Bad request, check the docs.")
	})

	s.Handle("/status.json", func(r *http.Request) app.Response {
		return app.OK(map[string]string{"status": "good"})
	})

	s.Handle("/authors.json", func(r *http.Request) app.Response {
		return app.OK(authors.Authors())
	})

	http.ListenAndServe(s.Port, s.Mux)
}
