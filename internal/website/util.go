package website

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/bhongy/go-web-programming-book/internal/data"
)

// Checks if the session is valid (or still valid)
func session(r *http.Request) (valid bool, err error) {
	cookie, err := r.Cookie("_sess")
	if err == nil {
		s := data.Session{UUID: cookie.Value}
		valid, err = s.Check()
	}
	return
}

func generateHTML(w http.ResponseWriter, data interface{}, filenames []string) {
	dir := "internal/website/templates"
	files := make([]string, len(filenames))
	for i, f := range filenames {
		files[i] = fmt.Sprintf("%s/%s.html", dir, f)
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}
