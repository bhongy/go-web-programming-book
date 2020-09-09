package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func logRequest(req *http.Request) {
	log.Printf("request: %s\n", req.URL.Path)
}

func generateHTML(w http.ResponseWriter, data interface{}, filenames []string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}
