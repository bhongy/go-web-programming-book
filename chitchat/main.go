package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	log.Println("Server is running on port: 8080")
	server.ListenAndServe()
}

func index(w http.ResponseWriter, req *http.Request) {
	logRequest(req)
	generateHTML(w, nil, []string{"layout", "navbar", "index"})
}
