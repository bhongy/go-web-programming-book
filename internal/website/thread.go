package website

import (
	"net/http"

	"github.com/bhongy/go-web-programming-book/internal/data"
)

// newThread displays a page to create a new thread
// GET /thread/new
func newThread(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	if _, err := data.CheckSession(r); err == nil {
		generateHTML(w, nil, []string{"layout", "private.navbar", "new.thread"})
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
