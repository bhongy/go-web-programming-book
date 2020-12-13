package data

import (
	"log"
	"time"
)

// Post models a reply in a discussion thread
type Post struct {
	ID        int
	UUID      string
	Body      string
	UserID    int
	ThreadID  int
	CreatedAt time.Time
}

// CreatedAtDate formats CreateAt data to display nicely on the screen
func (p *Post) CreatedAtDate() string {
	return p.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// Creator returns the User who made the post
func (p *Post) Creator() (u User) {
	query := `
		SELECT
			id, uuid, name, email, created_at
		FROM
			users
		WHERE
			id = $1
	`
	err := Db.
		QueryRow(query, p.UserID).
		Scan(&u.ID, &u.UUID, &u.Name, &u.Email, &u.CreatedAt)
	if err != nil {
		log.Printf("Post.Creator: cannot load user - %v\n", err)
	}
	return
}
