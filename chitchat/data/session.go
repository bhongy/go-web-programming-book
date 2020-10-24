package data

import "time"

// Session models a login session of a user
type Session struct {
	ID        int
	UUID      string
	Email     string
	UserID    int
	CreatedAt time.Time
}

// Check returns whether the session is valid in the database
func (s *Session) Check() (valid bool, err error) {
	var count int
	q := Db.QueryRow("SELECT count(id) FROM sessions WHERE uuid = $1", s.UUID)
	err = q.Scan(&count)
	if count > 0 {
		valid = true
	}
	return
}
