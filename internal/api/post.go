package api

import (
	"fmt"
	"net/http"

	"github.com/bhongy/go-web-programming-book/internal/data"
	"github.com/bhongy/go-web-programming-book/internal/website"
)

func createPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
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
	body := r.PostFormValue("body")
	uuid := r.PostFormValue("thread_uuid")
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		website.RedirectToErrorPage(w, r, "Cannot read thread")
	}
	if _, err := user.CreatePost(thread, body); err != nil {
		website.RedirectToErrorPage(w, r, "Cannot create post")
	}
	url := fmt.Sprint("/thread/read?id=", uuid)
	http.Redirect(w, r, url, 302)
}
