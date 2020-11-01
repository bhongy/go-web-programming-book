package route

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/bhongy/go-web-programming-book/data"
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
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

func redirectToErrorPage(w http.ResponseWriter, r *http.Request, msg string) {
	url := fmt.Sprintf("/err?msg=%s", msg)
	http.Redirect(w, r, url, 302)
}
