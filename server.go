package api

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

func NewServer(port string, mux *http.ServeMux) Server {
	log.Printf("Serving %s\n", port)
	return Server{port, mux}
}

type Server struct {
	Port string
	Mux  *http.ServeMux
}

type HandlerFunc func(r *http.Request) Response

func (s Server) Handle(path string, handler HandlerFunc) {
	s.Mux.Handle(path, handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := handler(r)
		w.WriteHeader(response.Status)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response.Body)
	})))
}
