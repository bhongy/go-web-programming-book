package data

import (
	"errors"
	"net/http"
	"time"
)

// Session models a login session of a user
type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

// check returns whether the session is valid in the database
func (s *Session) check() (valid bool, err error) {
	var count int
	query := "SELECT count(id) FROM sessions WHERE uuid = $1"
	err = Db.QueryRow(query, s.UUID).Scan(&count)
	if err == nil && count > 0 {
		valid = true
	}
	return
}

// Delete the session with the UUID from the database
func (s *Session) Delete() (err error) {
	query := "DELETE FROM sessions WHERE uuid = $1"
	_, err = Db.Exec(query, s.UUID)
	return
}

// User returns the user from the session
func (s *Session) User() (u User, err error) {
	query := `
		SELECT
			id, uuid, name, email, created_at
		FROM
			users
		WHERE
			id = $1
	`
	err = Db.
		QueryRow(query, s.UserID).
		Scan(&u.ID, &u.UUID, &u.Name, &u.Email, &u.CreatedAt)
	return
}

// CheckSession checks the request whether the session is still valid
func CheckSession(r *http.Request) (s Session, err error) {
	cookie, err := r.Cookie("_sess")
	if err == nil {
		s.UUID = cookie.Value
		if ok, _ := s.check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}
