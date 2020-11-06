package website

import (
	"net/http"
)

// Index shows the homepage
// GET /
func Index(w http.ResponseWriter, r *http.Request) {
	if loggedin, _ := session(r); loggedin {
		generateHTML(w, nil, []string{"layout", "private.navbar", "index"})
	} else {
		generateHTML(w, nil, []string{"layout", "public.navbar", "index"})
	}
}

// Err shows the error page given `msg` in the querystring
// GET /err?msg=
func Err(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	if msg == "" {
		msg = "(no error message)"
	}
	if loggedin, _ := session(r); loggedin {
		generateHTML(w, msg, []string{"login.layout", "private.navbar", "error"})
	} else {
		generateHTML(w, msg, []string{"login.layout", "public.navbar", "error"})
	}
}
