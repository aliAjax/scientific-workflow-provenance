package main

import (
	"log"
	"net/http"
	"os"
	"scientific-workflow-provenance/internal/transport"
)

func main() {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8088"
	}
	srv := transport.NewServer()
	log.Printf("workflow provenance listening on %s", addr)
	if err := http.ListenAndServe(addr, srv.Router()); err != nil {
		log.Fatal(err)
	}
}
