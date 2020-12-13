package website

import (
	"fmt"
	"html/template"
	"net/http"
)

func generateHTML(w http.ResponseWriter, data interface{}, filenames []string) {
	dir := "internal/website/templates"
	files := make([]string, len(filenames))
	for i, f := range filenames {
		files[i] = fmt.Sprintf("%s/%s.html", dir, f)
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}
