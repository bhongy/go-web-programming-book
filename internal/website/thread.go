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

// readThread shows the detail of the thread for a given uuid
// including the posts and the form to write a post
// GET /thread/read?id={thread_uuid}
func readThread(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	uuid := r.URL.Query().Get("id")
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		RedirectToErrorPage(w, r, "Cannot read thread")
		return
	}

	if _, err := data.CheckSession(r); err == nil {
		// we need to pass by reference here because the methods are on the pointer receiver
		generateHTML(w, &thread, []string{"layout", "private.navbar", "private.thread"})
	} else {
		generateHTML(w, &thread, []string{"layout", "public.navbar", "public.thread"})
	}
}
