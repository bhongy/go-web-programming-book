package api

import (
	"fmt"
	"net/http"

	"github.com/bhongy/go-web-programming-book/internal/data"
	"github.com/bhongy/go-web-programming-book/internal/website"
)

// createThread creates a new thread for the logged in user
// POST /thread/create
func createThread(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
	}
	sess, err := data.CheckSession(r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
		return
	}
	err = r.ParseForm()
	if err != nil {
		website.RedirectToErrorPage(w, r, "Cannot parse form")
		return
	}
	user, err := sess.User()
	if err != nil {
		website.RedirectToErrorPage(w, r, "Cannot get user from session")
		return
	}
	topic := r.PostFormValue("topic")
	if thread, err := user.CreateThread(topic); err != nil {
		website.RedirectToErrorPage(w, r, "Cannot create thread")
	} else {
		location := fmt.Sprintf("/thread/read?id=%s", thread.UUID)
		http.Redirect(w, r, location, 302)
	}
}
