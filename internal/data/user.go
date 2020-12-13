package data

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User models a forum user
type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	CreatedAt time.Time

	password password
}

type password string

// newPassword creates a new password from a plain-text password
func newPassword(plaintext string) (password, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(plaintext), 12)
	return password(b), err
}

// compare checks whether two password values are equal
func (actual password) compare(given []byte) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(actual),
		[]byte(given),
	)
	if err != nil {
		if err != bcrypt.ErrMismatchedHashAndPassword {
			log.Printf(
				"password compare error: actual=%s, given=%s, err=%s",
				actual, given, err,
			)
		}
		return false
	}
	return true
}

// Create a new user in the database
func (u *User) Create(plaintextPassword string) (err error) {
	p, err := newPassword(plaintextPassword)
	if err != nil {
		return
	}

	query := `
		INSERT INTO users (uuid, name, email, password, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING
			id, uuid, created_at
	`
	err = Db.
		QueryRow(query, createUUID(), u.Name, u.Email, p, time.Now()).
		Scan(&u.ID, &u.UUID, &u.CreatedAt)
	return
}

// CheckPassword take a plain text password and compare it against
// the hashed user's password
func (u *User) CheckPassword(plaintextPassword string) bool {
	return u.password.compare([]byte(plaintextPassword))
}

// CreateSession creates a new login session for the user
func (u *User) CreateSession() (s Session, err error) {
	query := `
		INSERT INTO sessions (uuid, email, user_id, created_at)
		VALUES($1, $2, $3, $4)
		RETURNING
			id, uuid, email, user_id, created_at
	`
	err = Db.
		QueryRow(query, createUUID(), u.Email, u.ID, time.Now()).
		Scan(&s.ID, &s.UUID, &s.Email, &s.UserID, &s.CreatedAt)
	return
}

// CreateThread creates a new thread for the user
func (u *User) CreateThread(topic string) (t Thread, err error) {
	query := `
		INSERT INTO threads (uuid, topic, user_id, created_at)
		VALUES($1, $2, $3, $4)
		RETURNING
			id, uuid, topic, user_id, created_at
	`
	err = Db.
		QueryRow(query, createUUID(), topic, u.ID, time.Now()).
		Scan(&t.ID, &t.UUID, &t.Topic, &t.UserID, &t.CreatedAt)
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
		FROM
			users
		WHERE
			email = $1
	`
	err = Db.QueryRow(query, email).Scan(
		&u.ID,
		&u.UUID,
		&u.Name,
		&u.Email,
		&u.password,
		&u.CreatedAt,
	)
	return
}
