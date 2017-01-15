package app

import (
	"os"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/handlers"
)

func NewServer(port string, mux *http.ServeMux) Server {
  log.Printf("Serving %s\n", port)
  return Server{port, mux}
}

type Server struct {
  Port string
  Mux *http.ServeMux
}

func (s Server) Handle(path string, h handlerFunc) {
	s.Mux.Handle(path, handler(func (w http.ResponseWriter, r *http.Request) {
		response := h(r)
		jsonResponse(w, response.Status, response.Body)
	}))
}

type handlerFunc func (r *http.Request) Response

func handler(h http.HandlerFunc) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, h)
}

func jsonResponse(w http.ResponseWriter, status int, body interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}
