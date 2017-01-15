package main

import (
	"github.com/lukemorton/api"
	"net/http"
)

func main() {
	s := api.NewServer(":3000", http.NewServeMux())

	s.Handle("/", func(r *http.Request) api.Response {
		return api.Error(400, "Bad request, check the docs.")
	})

	s.Handle("/status.json", func(r *http.Request) api.Response {
		return api.OK(map[string]string{"status": "good"})
	})

	s.Handle("/authors.json", func(r *http.Request) api.Response {
		return api.OK(api.Authors())
	})

	http.ListenAndServe(s.Port, s.Mux)
}
