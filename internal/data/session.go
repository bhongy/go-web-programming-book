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

// CheckSession
//   checks the request whether the session is still valid
//   returns the session if the session is still valid
func CheckSession(r *http.Request) (s Session, err error) {
	cookie, err := r.Cookie("_sess")
	if err != nil {
		return
	}
	uuid := cookie.Value
	query := `
		SELECT
			id, uuid, email, user_id, created_at
		FROM
			sessions
		WHERE
			uuid = $1
	`
	err = Db.
		QueryRow(query, uuid).
		Scan(&s.ID, &s.UUID, &s.Email, &s.UserID, &s.CreatedAt)
	if s.ID == 0 {
		err = errors.New("Invalid session")
	}
	return
}
