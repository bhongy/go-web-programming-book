package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bhongy/go-web-programming-book/data"
)

func main() {
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/authenticate", authenticate)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	log.Println("Server is running at: https://localhost:8080")
	if err := server.ListenAndServeTLS("localhost.crt", "localhost.key"); err != nil {
		log.Fatal(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if loggedin, _ := session(r); loggedin {
		generateHTML(w, nil, []string{"layout", "private.navbar", "index"})
	} else {
		generateHTML(w, nil, []string{"layout", "public.navbar", "index"})
	}
}

// GET /err?msg=
// shows the error message page
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

// GET /login
// Show the login page
func login(w http.ResponseWriter, r *http.Request) {
	if loggedin, _ := session(r); loggedin {
		generateHTML(w, nil, []string{"login.layout", "private.navbar", "login"})
	} else {
		generateHTML(w, nil, []string{"login.layout", "public.navbar", "login"})
	}
}

// GET /logout
// Logs the user out
func logout(w http.ResponseWriter, r *http.Request) {
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

// POST /authenticate
// Authenticate the user given the email and password
func authenticate(w http.ResponseWriter, r *http.Request) {
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
