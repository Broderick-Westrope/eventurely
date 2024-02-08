package main

import (
	"log"
	"net/http"

	flag "github.com/spf13/pflag"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/event/create", eventCreate)
	mux.HandleFunc("/event/view", eventView)

	log.Printf("Starting server on %s", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
