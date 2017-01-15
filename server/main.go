package main

import (
	"os"
	"log"
	"net/http"
	"encoding/json"
	"github.com/lukemorton/api/response"
	"github.com/lukemorton/api/authors"
	"github.com/gorilla/handlers"
)

const (
	port = ":3000"
)

func main() {
	log.Printf("Serving %s\n", port)

	mux := http.NewServeMux()

	Handle(mux, "/", func (r *http.Request) response.Response {
		return response.Error(400, "Bad request, check the docs.")
	})

	Handle(mux, "/status.json", func (r *http.Request) response.Response {
		return response.DefaultOK()
	})

	Handle(mux, "/authors.json", func (r *http.Request) response.Response {
		return response.OK(authors.Authors())
	})

	http.ListenAndServe(port, mux)
}

func Handler(handler http.HandlerFunc) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, handler)
}

func JSON(w http.ResponseWriter, status int, body interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

type HandlerFunc func (r *http.Request) response.Response

func Handle(mux *http.ServeMux, path string, handler HandlerFunc) {
	mux.Handle(path, Handler(func (w http.ResponseWriter, r *http.Request) {
		response := handler(r)
		JSON(w, response.Status, response.Body)
	}))
}
