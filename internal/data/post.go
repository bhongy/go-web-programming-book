package data

import "time"

// Post models a reply in a discussion thread
type Post struct {
	ID        int
	UUID      string
	Body      string
	UserID    int
	ThreadID  int
	CreatedAt time.Time
}
