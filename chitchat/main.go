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

	log.Println("Server is running at: https://localhost:8080")
	if err := server.ListenAndServeTLS("localhost.crt", "localhost.key"); err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	if loggedin, _ := session(r); loggedin {
		generateHTML(w, nil, []string{"layout", "private.navbar", "index"})
	} else {
		generateHTML(w, nil, []string{"layout", "public.navbar", "index"})
	}
}
