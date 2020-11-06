package website

import (
	"net/http"
)

func NewServer() http.Handler {
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)

	return mux
}

// index shows the homepage
// GET /
func index(w http.ResponseWriter, r *http.Request) {
	if loggedin, _ := session(r); loggedin {
		generateHTML(w, nil, []string{"layout", "private.navbar", "index"})
	} else {
		generateHTML(w, nil, []string{"layout", "public.navbar", "index"})
	}
}

// err shows the error page given `msg` in the querystring
// GET /err?msg=
func err(w http.ResponseWriter, r *http.Request) {
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
