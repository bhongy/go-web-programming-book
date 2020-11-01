package main

import (
	"log"
	"net/http"

	"github.com/bhongy/go-web-programming-book/route"
)

func main() {
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", route.Index)
	mux.HandleFunc("/err", route.Err)

	mux.HandleFunc("/login", route.Login)
	mux.HandleFunc("/logout", route.Logout)
	mux.HandleFunc("/authenticate", route.Authenticate)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	log.Println("Server is running at: https://localhost:8080")
	if err := server.ListenAndServeTLS("localhost.crt", "localhost.key"); err != nil {
		log.Fatal(err)
	}
}
