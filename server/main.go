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

	Handle(mux, "/", func (w http.ResponseWriter, r *http.Request) {
		JSON(w, http.StatusBadRequest, response.Error("Bad request, check the docs. Docs coming soon... ;P"))
	})

	Handle(mux, "/status.json", func (w http.ResponseWriter, r *http.Request) {
		JSON(w, http.StatusOK, response.OK())
	})

	Handle(mux, "/authors.json", func (w http.ResponseWriter, r *http.Request) {
		JSON(w, http.StatusOK, authors.Authors())
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

func Handle(mux *http.ServeMux, path string, handler http.HandlerFunc) {
	mux.Handle(path, Handler(handler))
}
