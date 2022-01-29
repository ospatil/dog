package main

import (
	"log"

	"github.com/ospatil/dog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":7070")
	log.Fatal(srv.ListenAndServe())
}
