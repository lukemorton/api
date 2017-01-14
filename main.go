package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port = flag.String("port", ":3000", "Listen address")
)

func main() {
	flag.Parse()
  log.Printf("Serving %s\n", *port)
	http.HandleFunc("/", Handle)
	log.Fatal(http.ListenAndServe(*port, nil))
}

func Handle(w http.ResponseWriter, r *http.Request) {
  log.Println("200 /")
	fmt.Fprintln(w, "Hello world!!")
}
