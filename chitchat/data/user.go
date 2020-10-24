package data

import "time"

// User models a forum user
type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// UserByEmail retruns a user given the email
func UserByEmail(email string) (User, error) {
	u := User{}
	q := Db.QueryRow(`
		SELECT
			id,
			uuid,
			name,
			email,
			password,
			created_at
		FROM users
		WHERE email = $1
		`,
		email,
	)
	err := q.Scan(
		&u.ID,
		&u.UUID,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
	)
	return u, err
}
