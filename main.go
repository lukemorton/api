package main

import (
	"fmt"
	"log"
	"net/http"
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
	fmt.Fprintln(w, "Hello world!!")
}
