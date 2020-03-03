package main

import (
	"log"
	"net/http"

	"github.com/mdcurran/palatable/internal/pkg/server"
)

func main() {
	s := server.New()

	log.Println("Starting HTTP server")
	err := http.ListenAndServe(":8080", s.Router)
	if err != nil {
		log.Fatalf("Unable to initialise API server - %v", err)
	}
}
