package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/bhongy/go-web-programming-book/chitchat/data"
)

func logRequest(req *http.Request) {
	log.Printf("request: %s\n", req.URL.Path)
}

// Checks if the session is valid (or still valid)
func session(req *http.Request) (valid bool, err error) {
	cookie, err := req.Cookie("_cookie")
	if err == nil {
		s := data.Session{UUID: cookie.Value}
		valid, err = s.Check()
	}
	return
}

func generateHTML(w http.ResponseWriter, data interface{}, filenames []string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}
