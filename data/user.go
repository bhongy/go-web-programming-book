package data

import (
	"time"
)

// User models a forum user
type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// CreateSession ...
func (u User) CreateSession() (s Session, err error) {
	query := `
		INSERT INTO sessions (uuid, email, user_id, created_at)
		VALUES($1, $2, $3, $4)
		RETURNING
			id,
			uuid,
			email,
			user_id,
			created_at
	`
	err = Db.QueryRow(query, createUUID(), u.Email, u.ID, u.CreatedAt).Scan(
		&s.ID,
		&s.UUID,
		&s.Email,
		&s.UserID,
		&s.CreatedAt,
	)
	return
}

// UserByEmail returns a single user given the email
func UserByEmail(email string) (u User, err error) {
	query := `
		SELECT
			id,
			uuid,
			name,
			email,
			password,
			created_at
		FROM users
		WHERE email = $1
	`
	err = Db.QueryRow(query, email).Scan(
		&u.ID,
		&u.UUID,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
	)
	return
}
