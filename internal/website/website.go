package website

import (
	"fmt"
	"net/http"

	"github.com/bhongy/go-web-programming-book/internal/data"
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

	mux.HandleFunc("/thread/new", newThread)

	return mux
}

// RedirectToErrorPage redirects the request to the error page
func RedirectToErrorPage(w http.ResponseWriter, r *http.Request, msg string) {
	url := fmt.Sprintf("/err?msg=%s", msg)
	http.Redirect(w, r, url, 302)
}

// index shows the homepage
// GET /
func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	threads, err := data.Threads()
	if err != nil {
		RedirectToErrorPage(w, r, "Cannot get threads")
		return
	}
	if _, err := data.CheckSession(r); err == nil {
		generateHTML(w, threads, []string{"layout", "private.navbar", "index"})
	} else {
		generateHTML(w, threads, []string{"layout", "public.navbar", "index"})
	}
}

// err shows the error page given `msg` in the querystring
// GET /err?msg=
func err(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	msg := r.URL.Query().Get("msg")
	if msg == "" {
		msg = "(no error message)"
	}
	if _, err := data.CheckSession(r); err == nil {
		generateHTML(w, msg, []string{"login.layout", "private.navbar", "error"})
	} else {
		generateHTML(w, msg, []string{"login.layout", "public.navbar", "error"})
	}
}
