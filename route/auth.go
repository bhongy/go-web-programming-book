package route

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bhongy/go-web-programming-book/data"
)

// Login shows the login page
// GET /login
func Login(w http.ResponseWriter, r *http.Request) {
	if loggedin, _ := session(r); loggedin {
		generateHTML(w, nil, []string{"login.layout", "private.navbar", "login"})
	} else {
		generateHTML(w, nil, []string{"login.layout", "public.navbar", "login"})
	}
}

// Logout shows the logout page
// GET /logout
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_sess")
	// if we have the session cookie, delete it
	if err != http.ErrNoCookie {
		session := data.Session{UUID: cookie.Value}
		err = session.Delete()
		if err != nil {
			log.Printf("logout: error deleting session - %v\n", err)
		}
		// remove cookie from the client
		cookie := http.Cookie{
			Name:   "_sess",
			MaxAge: -1,
		}
		http.SetCookie(w, &cookie)
	}
	http.Redirect(w, r, "/", 302)
}

// Authenticate the user given the email and password
// POST /authenticate
func Authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		msg := fmt.Sprintf("Cannot find user: %s", r.PostFormValue("email"))
		redirect := fmt.Sprintf("/err?msg=%s", msg)
		http.Redirect(w, r, redirect, 302)
		return
	}

	if user.Password != r.PostFormValue("password") {
		http.Redirect(w, r, "/login", 302)
		return
	}

	session, err := user.CreateSession()
	if err != nil {
		log.Println("authenticate: cannot create session")
		redirect := fmt.Sprintf("/err?msg=%s", "Cannot log in")
		http.Redirect(w, r, redirect, 302)
		return
	}

	cookie := http.Cookie{
		Name:     "_sess",
		Value:    session.UUID,
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", 302)
}
