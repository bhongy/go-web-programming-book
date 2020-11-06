package main

import (
	"log"
	"net/http"

	"github.com/bhongy/go-web-programming-book/internal/api"
	"github.com/bhongy/go-web-programming-book/internal/website"
)

func main() {
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", website.Index)
	mux.HandleFunc("/err", website.Err)

	mux.HandleFunc("/signup", website.Signup)
	mux.HandleFunc("/login", website.Login)
	mux.HandleFunc("/logout", website.Logout)

	mux.HandleFunc("/account/create", api.CreateAccount)
	mux.HandleFunc("/authenticate", api.Authenticate)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	log.Println("Server is running at: https://localhost:8080")
	if err := server.ListenAndServeTLS("localhost.crt", "localhost.key"); err != nil {
		log.Fatal(err)
	}
}
