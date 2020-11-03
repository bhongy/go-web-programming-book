package route

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bhongy/go-web-programming-book/data"
)

// Signup shows the signup (account registration) page
// GET /signup
func Signup(w http.ResponseWriter, r *http.Request) {
	if loggedin, _ := session(r); loggedin {
		generateHTML(w, nil, []string{"login.layout", "private.navbar", "signup"})
	} else {
		generateHTML(w, nil, []string{"login.layout", "public.navbar", "signup"})
	}
}

// CreateAccount registers a new user account
// GET /account/create
func CreateAccount(w http.ResponseWriter, r *http.Request) {
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
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", 302)
}
