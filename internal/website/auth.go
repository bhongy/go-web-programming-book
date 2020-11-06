package website

import (
	"log"
	"net/http"

	"github.com/bhongy/go-web-programming-book/internal/data"
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
