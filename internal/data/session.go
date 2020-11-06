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
func (s Session) Check() (valid bool, err error) {
	var count int
	query := "SELECT count(id) FROM sessions WHERE uuid = $1"
	err = Db.QueryRow(query, s.UUID).Scan(&count)
	if err == nil && count > 0 {
		valid = true
	}
	return
}

// Delete the session with the UUID from the database
func (s Session) Delete() (err error) {
	query := "DELETE FROM sessions WHERE uuid = $1"
	_, err = Db.Exec(query, s.UUID)
	return
}
