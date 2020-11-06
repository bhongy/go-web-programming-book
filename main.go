package main

import (
	"log"
	"net/http"

	"github.com/bhongy/go-web-programming-book/internal/api"
	"github.com/bhongy/go-web-programming-book/internal/website"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", website.NewServer())
	mux.Handle("/api/", http.StripPrefix("/api", api.NewServer()))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	log.Println("Server is running at: https://localhost:8080")
	if err := server.ListenAndServeTLS("localhost.crt", "localhost.key"); err != nil {
		log.Fatal(err)
	}
}
