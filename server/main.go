package main

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/lukemorton/api/response"
)

const (
	port = ":3000"
)

func main() {
	log.Printf("Serving %s\n", port)
	http.HandleFunc("/", Handle)
	log.Fatal(http.ListenAndServe(port, nil))
}

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("200 /")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response.SuccessResponse())
}
