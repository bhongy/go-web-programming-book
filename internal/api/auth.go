package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bhongy/go-web-programming-book/internal/data"
)

// createAccount registers a new user account
// POST /account/create
func createAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	err := r.ParseForm()
	if err != nil {
		msg := fmt.Sprintf("Error parsing form data")
		redirectToErrorPage(w, r, msg)
		return
	}

	user := data.User{
		Name:  r.PostFormValue("name"),
		Email: r.PostFormValue("email"),
	}
	// plain text password
	password := r.PostFormValue("password")
	if err := user.Create(password); err != nil {
		log.Printf("create account: cannot create user - %v\n", err)
		redirectToErrorPage(w, r, "Error creating user")
		return
	}

	http.Redirect(w, r, "/login", 302)
}

// authenticate the user given the email and password
// POST /authenticate
func authenticate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	err := r.ParseForm()
	if err != nil {
		msg := fmt.Sprintf("Error parsing form data")
		redirectToErrorPage(w, r, msg)
		return
	}

	user, err := data.UserByEmail(r.PostFormValue("email"))
	// never tell the user if it was the username or password they got wrong
	// to prevent attackers from enumerating valid usernames without knowing their passwords
	// https://crackstation.net/hashing-security.htm
	// TODO: refactor the line below
	if err != nil || !user.CheckPassword(r.PostFormValue("password")) {
		msg := fmt.Sprintf("Invalid username or password")
		redirectToErrorPage(w, r, msg)
		return
	}

	session, err := user.CreateSession()
	if err != nil {
		log.Println("authenticate: cannot create session")
		redirectToErrorPage(w, r, "Cannot log in")
		return
	}

	cookie := http.Cookie{
		Name:     "_sess",
		Value:    session.UUID,
		HttpOnly: true,
		Secure:   true,
		// this must be set to "/" otherwise the browser will not set the cookie
		// even though `set-cookie` header is sent
		Path: "/",
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", 302)
}

func redirectToErrorPage(w http.ResponseWriter, r *http.Request, msg string) {
	url := fmt.Sprintf("/err?msg=%s", msg)
	http.Redirect(w, r, url, 302)
}
